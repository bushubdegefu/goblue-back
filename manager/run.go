package manager

// export PATH=$PATH:$(go env GOPATH)/bin

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strings"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/spf13/cobra"
	"gorm.io/gorm/clause"
	"semay.com/admin/database"
	"semay.com/admin/models"
	"semay.com/admin/responses"
	"semay.com/common"
	"semay.com/config"
	_ "semay.com/docs"
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
		regexp.MustCompile("^/api/v1/login"),
		regexp.MustCompile("^/lmetrics"),
		regexp.MustCompile("^/bluedocs"),
		regexp.MustCompile("^/docs"),
		regexp.MustCompile("^/metrics$"),
	}
	stop_flag  = 0
	route_name string
)

// API validation Key function to be run on middleware
// checks if token roles have required privilge to access the requested route
// func validateAPIKey(contx *fiber.Ctx, key string) (bool, error) {
// 	// Getting database session for app
// 	db := database.ReturnSession()

// 	// getting the name of the next function
// this is
// if stop_flag == 0 {
// 	stop_flag = 2
// 	contx.Next()
// 	route_name = contx.Route().Name
// 	contx.RestartRouting()

// }

// fmt.Println(stop_flag)
// fmt.Println(route_name)
// 	//  Getting list of roles required for the path
// 	roles := make([]string, 0, 20)
// 	var roles_fetch []models.Role
// 	var route models.EndPoints
// 	if res := db.Model(&models.EndPoints{}).Where("Name = ?", route_name).Find(&route); res.Error != nil {
// 		return false, contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
// 			Success: false,
// 			Message: "Token role fetching: " + res.Error.Error(),
// 			Data:    nil,
// 		})
// 	}
// 	db.Model(&route).Association("Roles").Find(&roles_fetch)

// 	for _, value := range roles_fetch {

// 		roles = append(roles, string(value.Name))
// 	}
// 	// parsing token
// 	token_roles, err := utils.ParseJWTToken(key)

// 	// validating token role against route
// 	tok_roles, _ := json.Marshal(token_roles)
// 	var tok_rol utils.UserClaim
// 	json.Unmarshal(tok_roles, &tok_rol)

// 	flag := true
// 	for _, route_priv := range roles {

// 		for _, tok_value := range tok_rol.Roles {
// 			if route_priv == tok_value {
// 				flag = true
// 				goto Exit // breaking out of loop if condition meet the requirement
// 			}
// 		}
// 	}
// 	// validating if value exists in available roles
// Exit:
// 	if err != nil {
// 		return flag, err
// 	}

// 	return flag, nil
// 	return true, nil
// }

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
func NextRoute(contx *fiber.Ctx, key string) (bool, error) {
	if stop_flag == 0 {
		stop_flag = 2
		contx.Next()
		route_name = contx.Route().Name
		contx.RestartRouting()
		// fmt.Println(route_name)
		// fmt.Println(contx.Route().Path)
	}
	return true, nil
}

func run() {
	// initalaizing the app
	app := fiber.New(fiber.Config{
		// Prefork:     true,
		// Network:     fiber.NetworkTCP,
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
		Format:     "\n${cyan}-[${time}]-[${ip}] -${white}${pid} ${red}${status} ${blue}[${method}] ${white}-${path}\n [${body}]\n[${error}]\n[${resBody}]\n[${reqHeaders}]\n[${queryParams}]\n",
		TimeFormat: "15:04:05",
		TimeZone:   "Local",
		Output:     file,
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
	// admin_app := app.Group("/api/v1")
	admin_app := app.Group("/api/v1", keyauth.New(keyauth.Config{
		Next:      authFilter,
		KeyLookup: "header:X-APP-TOKEN",
		Validator: NextRoute,
	}))

	setupRoutes(admin_app.(*fiber.Group))

	// updating response function route names to database
	// first parsing route names from app
	db := database.ReturnSession()
	// data, _ := json.MarshalIndent(app.Stack(), "", "  ")
	// fmt.Println(string(data))
	for _, routes := range app.Stack() {
		for _, route := range routes {

			if (route.Path != "") && (route.Path != "/") && (route.Name != "") && (route.Method != "HEAD") {
				response_meta := models.EndPoint{Name: route.Name + "_" + strings.ToLower(route.Method), RoutePaths: route.Path, Description: route.Name + "-" + route.Method, Method: route.Method}
				tx := db.Begin()
				if err := tx.Model(&models.EndPoint{}).Clauses(clause.OnConflict{DoNothing: true}).Create(&response_meta).Error; err != nil {
					tx.Rollback()
				}
				tx.Commit()

			}
		}
	}

	// recording available route name ends here
	port_1 := config.Config("PORT")
	// port_2 := config.Config("PORT_2")
	// starting on two provided ports

	// forking port 1
	go func() {
		log.Fatal(app.Listen(":" + port_1))
	}()
	// // forking port 2
	// go func() {
	// 	log.Fatal(app.Listen(":" + port_2))

	// }()

	c := make(chan os.Signal, 1)   // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt) // When an interrupt or termination signal is sent, notify the channel

	<-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	app.Shutdown()

	fmt.Println("Running cleanup tasks...")
	// Your cleanup tasks go here
	fmt.Println("Blue was successful shutdown.")

}

func setupRoutes(app *fiber.Group) {

	// adding endpoints
	app.Get("/checklogin", responses.CheckLogin).Name("check_login")
	app.Post("/login", responses.PostLogin).Name("login_route")

	app.Get("/roles", responses.GetRoles).Name("roles")
	app.Post("/roles", responses.PostRoles).Name("roles")
	app.Get("/roles/:id", responses.GetRolesID).Name("roles_single")
	app.Get("/droproles", responses.GetDropDownRoles).Name("drop_roles")
	app.Patch("/roles/:id", responses.PatchRoles).Name("roles_single")
	app.Delete("/roles/:id", responses.DeleteRoles).Name("roles_single")
	app.Put("/roles/:role_id", responses.ActivateDeactivateRoles).Name("activate_deactivate_role")
	app.Get("/role_endpoints", responses.GetRoleEndpointsID).Name("roles_endpoints")

	app.Get("/features", responses.GetFeatures).Name("features")
	app.Post("/features", responses.PostFeatures).Name("features")
	app.Get("/features/:id", responses.GetFeaturesID).Name("features_single")
	app.Get("/featuredrop", responses.GetDropFeatures).Name("drop_features")
	app.Patch("/features/:id", responses.PatchFeatures).Name("features_single")
	app.Delete("/features/:id", responses.DeleteFeatures).Name("features_single")
	app.Put("/features/:feature_id", responses.ActivateDeactivateFeature).Name("activate_deactivate_features")
	app.Patch("/featuresrole/:feature_id", responses.AddFeatureRole).Name("feature_role")
	app.Delete("/featuresrole/:feature_id", responses.DeleteFeatureRole).Name("feature_role")

	app.Get("/apps", responses.GetApps).Name("apps")
	app.Post("/apps", responses.PatchApps).Name("apps")
	app.Get("/apps/:id", responses.GetAppsID).Name("apps_single")
	app.Patch("/apps/:id", responses.PostApps).Name("apps_single")
	app.Delete("/apps/:id", responses.DeleteApps).Name("apps_single")

	app.Get("/users", responses.GetUsers).Name("users")
	app.Post("/users", responses.PostUsers).Name("users")
	app.Get("/users/:id", responses.GetUsersID).Name("user_single")
	app.Patch("/users/:id", responses.PatchUsers).Name("user_single")
	app.Delete("/users/:id", responses.DeleteUsers).Name("user_single")
	app.Put("/users/:user_id", responses.ActivateDeactivateUser).Name("activate_deactivate_user")
	app.Get("/userrole/:user_id", responses.GetUsersRolesByID).Name("get_user_roles")
	app.Post("/userrole/:user_id/:role_id", responses.AddUserRoles).Name("user_role")
	app.Delete("/userrole/:user_id/:role_id", responses.DeleteUserRoles).Name("user_role")

	app.Get("/pages", responses.GetPages).Name("pages")
	app.Post("/pages", responses.PostPage).Name("pages")
	app.Get("/pages/:id", responses.GetPageID).Name("page_single")
	app.Patch("/pages/:id", responses.PatchPage).Name("page_single")
	app.Delete("/pages/:id", responses.DeletePage).Name("page_single")
	app.Get("/pagesroles/:page_id", responses.GetPageRoles).Name("get_page_roles")
	app.Post("/pagerole/:page_id/:role_id", responses.AddPageRoles).Name("page_roles")
	app.Delete("/pagerole/:page_id/:role_id", responses.DeletePageRoles).Name("page_roles")

	app.Get("/endpoints", responses.GetEndPointResponse).Name("end_point")
	app.Post("/endpoints", responses.PostEndPoint).Name("end_point")
	app.Get("/endpoints/:id", responses.GetEndPointsID).Name("end_point_single")
	app.Get("/endpointdrop", responses.GetDropEndPoints).Name("drop_endpoints")
	app.Patch("/endpoints/:id", responses.PatchEndPoint).Name("end_point_single")
	app.Delete("/endpoints/:id", responses.DeleteEndPoint).Name("end_point_single")
	app.Patch("/feature_endpoint/:endpoint_id", responses.AddEndpointFeature).Name("feature_endpoint")
	app.Delete("/feature_endpoint/:endpoint_id", responses.DeleteEndpointFeature).Name("feature_endpoint")

	app.Post("/email", responses.SendEmail).Name("send_email")
}

func init() {
	goBlueCmd.AddCommand(runCmd)

}
