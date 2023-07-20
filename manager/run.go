package manager

// export PATH=$PATH:$(go env GOPATH)/bin

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/ansrivas/fiberprometheus/v2"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/cobra"
	"gorm.io/gorm/clause"
	"semay.com/admin/database"
	"semay.com/admin/models"
	"semay.com/admin/responses"
	"semay.com/common"
	"semay.com/config"
	_ "semay.com/docs"
	"semay.com/utils"
)

var (
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run Run Development server ",
		Long:  `Run Fiber development server Run Development server `,
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
	protectedURLs = []*regexp.Regexp{
		regexp.MustCompile("^/api/login"),
		regexp.MustCompile("^/lmetrics"),
		regexp.MustCompile("^/bluedocs"),
		regexp.MustCompile("^/docs"),
		regexp.MustCompile("^/metrics$"),
	}
	stop_flag = 0
)

// API validation Key function to be run on middleware
// checks if token roles have required privilge to access the requested route
func validateAPIKey(contx *fiber.Ctx, key string) (bool, error) {
	// Getting database session for app
	db := database.ReturnSession()

	// getting the name of the next function

	contx.Next()
	route_name := contx.Route().Name
	if stop_flag == 0 {
		stop_flag++
		contx.RestartRouting()
	}
	//  Getting list of roles required for the path
	roles := make([]string, 0, 20)
	var roles_fetch []models.Role
	var route models.RouteResponse
	if res := db.Model(&models.RouteResponse{}).Where("Name = ?", route_name).Find(&route); res.Error != nil {
		return false, contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Token role fetching: " + res.Error.Error(),
			Data:    nil,
		})
	}
	db.Model(&route).Association("Roles").Find(&roles_fetch)

	for _, value := range roles_fetch {

		roles = append(roles, string(value.Name))
	}
	// parsing token
	token_roles, err := utils.ParseJWTToken(key)

	// validating token role against route
	tok_roles, _ := json.Marshal(token_roles)
	var tok_rol utils.UserClaim
	json.Unmarshal(tok_roles, &tok_rol)

	flag := false
	for _, route_priv := range roles {

		for _, tok_value := range tok_rol.Roles {
			if route_priv == tok_value {
				flag = true
				goto Exit // breaking out of loop if condition meet the requirement
			}
		}
	}
	// validating if value exists in available roles
Exit:
	if err != nil {
		return flag, err
	}

	return flag, nil
}

// this is path filter which wavies token requirement for provided paths
func authFilter(c *fiber.Ctx) bool {
	originalURL := strings.ToLower(c.OriginalURL())

	for _, pattern := range protectedURLs {
		if pattern.MatchString(originalURL) {

			return true
		}
	}
	return false
}

func run() {
	// initalaizing the app
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			// Send custom error page
			err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
			if err != nil {
				// In case the SendFile fails
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}

			// Return from handler
			return nil
		},
	})

	// prometheus middleware concrete instance
	prometheus := fiberprometheus.New("gobluefiber")
	prometheus.RegisterAt(app, "/metrics")

	// prometheus monitoring middleware
	app.Use(prometheus.Middleware)

	// recover from panic attacks middlerware
	app.Use(recover.New())

	// allow cross origin request
	app.Use(cors.New())

	// Custom File Writer for logging
	file, err := os.OpenFile("goblue.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	// logger middle ware with the custom file writer object
	app.Use(logger.New(logger.Config{
		Format:     "${cyan}-[${time}]-[${ip}] -${white}${pid} ${red}${status} ${blue}[${method}] ${white}-${path}\n [${body}]\n",
		TimeFormat: "15:04:05",
		TimeZone:   "Local",

		Output: file,
	}))

	// mouting static files
	app.Static("/bluedocs", "./documents").Name("static_file_routes")

	// swagger documentation endpoints
	app.Get("/docs/*", swagger.HandlerDefault)
	app.Get("/docs/*", swagger.New()).Name("swagger_routes")

	// fiber native monitoring metrics endpoint
	app.Get("/lmetrics", monitor.New(monitor.Config{Title: "goBlue Metrics Page"})).Name("custom_metrics_route")

	// setting up applications from other moduels

	app.Get("/", func(contx *fiber.Ctx) error {

		return contx.JSON(common.ResponseHTTP{
			Success: true,
			Message: "Hello World!",
			Data:    nil,
		})

	}).Name("index_route")

	// adding group with authenthication middleware
	admin_app := app.Group("/api/", keyauth.New(keyauth.Config{
		Next:      authFilter,
		KeyLookup: "header:X-APP-TOKEN",
		Validator: validateAPIKey,
	}))
	setupRoutes(admin_app.(*fiber.Group))

	// updating response function route names to database
	// first parsing route names from app
	response_names := make([]models.RouteResponse, 0)

	for _, routes := range app.Stack() {
		for _, route := range routes {
			if route.Name != "" {
				response_meta := models.RouteResponse{Name: route.Name}
				response_names = append(response_names, response_meta)

			}
		}
	}
	// building superuser list for routeroles for routes above
	super_route_roles := make([]models.RouteRoles, 0)
	for i := range response_names {
		route_id := i + 1
		super_route_roles = append(super_route_roles, models.RouteRoles{RoleID: 1, RouteResponseID: uint(route_id)})
	}

	// getting database session to update
	db := database.ReturnSession()
	tx := db.Begin()
	db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(super_route_roles, len(super_route_roles))
	db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(response_names, len(response_names))

	tx.Commit()

	// recording available route name ends here
	port := config.Config("PORT")
	log.Fatal(app.Listen("0.0.0.0:" + port))
}

func setupRoutes(app *fiber.Group) {

	// adding middleware
	app.Post("/login", responses.PostLogin).Name("login_route")

	app.Get("/roles", responses.GetRoles).Name("get_roles")
	app.Post("/roles", responses.PostRoles).Name("add_roles")
	app.Get("/roles/:id", responses.GetRolesID).Name("get_one_role")
	app.Patch("/roles/:id", responses.PatchRoles).Name("patch_role")
	app.Delete("/roles/:id", responses.DeleteRoles).Name("delete_role")

	app.Get("/users", responses.GetUsers).Name("get_users")
	app.Get("users/:id", responses.GetUsersID).Name("get_one_user")
	app.Get("/userrole/:user_id", responses.GetUsersRolesByID).Name("get_user_roles")
	app.Post("/users", responses.PostUsers).Name("add_user")
	app.Patch("/users/:id", responses.PatchUsers).Name("patch_user")
	app.Delete("/users/:id", responses.DeleteUsers).Name("delete_user")
	app.Post("/userrole/:user_id/:role_id", responses.AddUserRoles).Name("add_user_role")
	app.Delete("/userrole/:user_id/:role_id", responses.DeleteUserRoles).Name("delete_user_role")

	app.Get("/pages", responses.GetPages).Name("get_pages")
	app.Get("/pages/:id", responses.GetPageID).Name("get_one_page")
	app.Get("/pageroute/:page_id", responses.GetPageRoutes).Name("get_page_routes")
	app.Post("/pages", responses.PostPage).Name("add_page")
	app.Patch("/pages/:id", responses.PatchPage).Name("patch_page")
	app.Delete("/pages/:id", responses.DeletePage).Name("delete_page")
	app.Post("/pageroute/:page_id/:route_id", responses.AddPageRoutes).Name("add_page_route")
	app.Delete("/pageroute/:page_id/:route_id", responses.DeletePageRoutes).Name("delete_page_route")

	app.Get("/routes", responses.GetRouteResponse).Name("get_routes")
	app.Get("/routes/:id", responses.GetRoutesID).Name("get_one_route")
	app.Get("/routerole/:route_id", responses.GetRouteRoles).Name("get_route_roles")
	app.Post("/routes", responses.PostRoute).Name("add_route")
	app.Patch("/routes/:id", responses.PatchRoute).Name("patch_route")
	app.Delete("/routes/:id", responses.DeleteRoute).Name("delete_route")
	app.Post("/routerole/:route_id/:role_id", responses.AddRouteRoles).Name("add_route_role")
	app.Delete("/routerole/:route_id/:role_id", responses.DeleteRouteRoles).Name("delete_route_role")

	app.Post("/email", responses.SendEmail).Name("send_email")
}

func init() {
	goBlueCmd.AddCommand(runCmd)

}
