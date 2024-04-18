package manager

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/earlydata"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/madflojo/tasks"

	"github.com/spf13/cobra"
	"gorm.io/gorm/clause"
	"semay.com/admin/database"
	"semay.com/admin/models"
	"semay.com/admin/responses"
	"semay.com/bluerabbit"
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
		regexp.MustCompile("^/api/v1/login"),
		regexp.MustCompile("^/api/v1/stream"),
		regexp.MustCompile("^/api/v1/pics"),
		regexp.MustCompile("^/lmetrics"),
		regexp.MustCompile("^/bluedocs"),
		regexp.MustCompile("^/docs"),
		regexp.MustCompile("^/metrics$"),
	}
)

// this is path filter which wavies token requirement for provided paths
func authFilter(c *fiber.Ctx) bool {
	originalURL := strings.ToLower(c.OriginalURL())

	for _, pattern := range protectedURLs {
		if pattern.MatchString(originalURL) {
			c.Request().Header.Add("X-APP-TOKEN", "allowed")
			return true
		}
	}
	return false
}

func testMiddleware(c *fiber.Ctx, key string) (bool, error) {
	return true, nil
}

func NextFunc(contx *fiber.Ctx) error {

	return nil
}

func NextRoute(contx *fiber.Ctx, key string) (bool, error) {
	contx.Next()
	route_name := contx.Route().Name + "_" + strings.ToLower(contx.Route().Method)

	if key == "anonymous" && models.Endpoints_JSON[route_name] == "Anonymous" {
		return true, nil
	}

	//  first validating the token
	claims, err := utils.ParseJWTToken(key)
	if err != nil {
		fmt.Println(err)
	}

	// check if the token have the desired role for the route
	role_test := utils.CheckValueExistsInSlice(claims.Roles, models.Endpoints_JSON[route_name])
	if role_test {
		return true, nil
	}
	return false, nil
}

func run() {
	scheduler := tasks.New()
	defer scheduler.Stop()

	app, log_file := MakeApp("dev")
	defer log_file.Close()
	// updating response function route names to database
	// first parsing route names from app
	db := database.ReturnSession()

	//  database logger file
	gormLoggeFile := database.GormLoggerFile()

	//  registering endpoints related to app
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

	// running background consumer
	go func() {
		bluerabbit.BlueConsumer()
	}()

	// recording available route name ends here
	port_1 := config.Config("PORT")
	// port_2 := config.Config("PORT_2")

	// starting on provided port
	go func(app *fiber.App) {
		log.Fatal(app.Listen(":" + port_1))
	}(app)

	// // Add a task to move to Logs Directory Every Interval, Interval to Be Provided From Configuration File
	if _, err := scheduler.Add(&tasks.Task{
		Interval: time.Duration(1 * time.Hour),
		TaskFunc: func() error {
			// currentTime := time.Now()
			// FileName := fmt.Sprintf("%v-%v-%v-%v-%v", currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute())
			// Command := fmt.Sprintf("cp goblue.log logs/blue-%v.log", FileName)
			// Command2 := fmt.Sprintf("cp gormblue.log logs/gorm-%v.log", FileName)
			// if _, err := exec.Command("bash", "-c", Command).Output(); err != nil {
			// 	fmt.Printf("error: %v\n", err)
			// }

			// if _, err := exec.Command("bash", "-c", Command2).Output(); err != nil {
			// 	fmt.Printf("error: %v\n", err)
			// }
			log_file.Truncate(0)
			gormLoggeFile.Truncate(0)
			return nil
		},
	}); err != nil {
		fmt.Println(err)

	}

	//  Salt Timer Tasks
	clear_run, _ := strconv.Atoi(config.Config("JWT_SALT_LIFE_TIME"))
	clear_run = int(clear_run)
	jwt_update_interval := time.Minute * time.Duration(clear_run)
	//  Task 2 for testing Make random heartbeat call

	if _, err := scheduler.Add(&tasks.Task{
		Interval: jwt_update_interval,
		TaskFunc: func() error {
			utils.JWTSaltUpdate()
			return nil
		},
	}); err != nil {
		fmt.Println(err)

	}

	c := make(chan os.Signal, 1)   // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt) // When an interrupt or termination signal is sent, notify the channel

	<-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	app.Shutdown()

	fmt.Println("Running cleanup tasks...")
	// Your cleanup tasks go here
	fmt.Println("Blue was successful shutdown.")

}

func MakeApp(appType string) (*fiber.App, *os.File) {

	// initalaizing the app
	app := fiber.New(fiber.Config{
		// Prefork: true,
		// Network:     fiber.NetworkTCP,
		// Immutable:   true,
		JSONEncoder:    json.Marshal,
		JSONDecoder:    json.Unmarshal,
		BodyLimit:      70 * 1024 * 1024,
		ReadBufferSize: 10 * 4096,
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
	models.GetAppFeatures("48015a9b-5a86-4a15-944b-94108aa78b4b")

	// adding prometheus instrumantation only if prod or dev
	if appType == "dev" || appType == "prod" {
		// prometheus middleware concrete instance
		prometheus := fiberprometheus.New("gobluefiber")
		prometheus.RegisterAt(app, "/metrics")

		// prometheus monitoring middleware
		app.Use(prometheus.Middleware)
	}
	// recover from panic attacks middlerware
	app.Use(recover.New())

	// allow cross origin request
	app.Use(cors.New())

	// Custom File Writer for logging
	file, err := os.OpenFile("goblue.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

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

	// setting up Applications from other moduels
	app.Get("/", func(contx *fiber.Ctx) error {
		return contx.JSON(common.ResponseHTTP{
			Success: true,
			Message: "Hello World!",
			Data:    nil,
		})

	}).Name("index_route")

	// adding group with authenthication middleware
	if appType == "dev" || appType == "prod" {
		// adding group with authenthication middleware
		admin_app := app.Group("/api/v1", keyauth.New(keyauth.Config{
			Next:      authFilter,
			KeyLookup: "header:X-APP-TOKEN",
			Validator: NextRoute,
		}))
		setupRoutes(admin_app.(*fiber.Group))
	} else {
		// adding group with authenthication middleware
		admin_app := app.Group("/api/v1", keyauth.New(keyauth.Config{
			Next:      authFilter,
			KeyLookup: "header:X-APP-TOKEN",
			Validator: testMiddleware,
		}))
		setupRoutes(admin_app.(*fiber.Group))
	}

	app.Use(idempotency.New(idempotency.Config{
		Lifetime: 10 * time.Second,
	}))

	app.Use(earlydata.New(earlydata.Config{
		Error: fiber.ErrTooEarly,
		// ...
	}))

	return app, file
}

func setupRoutes(app *fiber.Group) {

	// adding endpoints
	app.Get("/checklogin", NextFunc).Name("check_login").Get("/checklogin", responses.CheckLogin).Name("check_login")
	app.Post("/login", responses.PostLogin).Name("login_route")

	app.Get("/roles", NextFunc).Name("roles").Get("/roles", responses.GetRoles)
	app.Post("/roles", NextFunc).Name("roles").Post("/roles", responses.PostRoles)
	app.Get("/roles/:id", NextFunc).Name("roles_single").Get("/roles/:id", responses.GetRolesID)
	app.Get("/droproles", NextFunc).Name("drop_roles").Get("/droproles", responses.GetDropDownRoles)
	app.Patch("/roles/:id", NextFunc).Name("roles_single").Patch("/roles/:id", responses.PatchRoles)
	app.Patch("/approle/:role_id", NextFunc).Name("roles_app").Patch("/approle/:role_id", responses.UpdateRoleApp)
	app.Delete("/roles/:id", NextFunc).Name("roles_single").Delete("/roles/:id", responses.DeleteRoles)
	app.Put("/roles/:role_id", NextFunc).Name("activate_deactivate_role").Put("/roles/:role_id", responses.ActivateDeactivateRoles)
	app.Get("/role_endpoints", NextFunc).Name("roles_endpoints").Get("/role_endpoints", responses.GetRoleEndpointsID)

	app.Get("/features", NextFunc).Name("features").Get("/features", responses.GetFeatures)
	app.Post("/features", NextFunc).Name("features").Post("/features", responses.PostFeatures)
	app.Get("/features/:id", NextFunc).Name("features_single").Get("/features/:id", responses.GetFeaturesID)
	app.Get("/featuredrop", NextFunc).Name("drop_features").Get("/featuredrop", responses.GetDropFeatures)
	app.Patch("/features/:id", NextFunc).Name("features_single").Patch("/features/:id", responses.PatchFeatures)
	app.Delete("/features/:id", NextFunc).Name("features_single").Delete("/features/:id", responses.DeleteFeatures)
	app.Put("/features/:feature_id", NextFunc).Name("activate_deactivate_features").Put("/features/:feature_id", responses.ActivateDeactivateFeature)
	app.Patch("/featuresrole/:feature_id", NextFunc).Name("feature_role").Patch("/featuresrole/:feature_id", responses.AddFeatureRole)
	app.Delete("/featuresrole/:feature_id", NextFunc).Name("feature_role").Delete("/featuresrole/:feature_id", responses.DeleteFeatureRole)

	app.Get("/apps", NextFunc).Name("apps").Get("/apps", responses.GetApps)
	app.Post("/apps", NextFunc).Name("apps").Post("/apps", responses.PostApps)
	app.Get("/apps/:id", NextFunc).Name("apps_single").Get("/apps/:id", responses.GetAppsID)
	app.Get("/appsdrop", NextFunc).Name("drop_sppd").Get("/appsdrop", responses.GetDropApps)
	app.Get("/appsmatrix/:id", NextFunc).Name("apps_features").Get("/appsmatrix/:id", responses.GetAppMatrix)
	app.Patch("/apps/:id", NextFunc).Name("apps_single").Patch("/apps/:id", responses.PatchApps)
	app.Delete("/apps/:id", NextFunc).Name("apps_single").Delete("/apps/:id", responses.DeleteApps)

	app.Get("/users", NextFunc).Name("users").Get("/users", responses.GetUsers)
	app.Post("/users", NextFunc).Name("users").Post("/users", responses.PostUsers)
	app.Get("/users/:id", NextFunc).Name("user_single").Get("/users/:id", responses.GetUsersID)
	app.Patch("/users/:id", NextFunc).Name("user_single").Patch("/users/:id", responses.PatchUsers)
	app.Delete("/users/:id", NextFunc).Name("user_single").Delete("/users/:id", responses.DeleteUsers)
	app.Put("/users/:user_id", NextFunc).Name("activate_deactivate_user").Put("/users/:user_id", responses.ActivateDeactivateUser)
	app.Put("/users", NextFunc).Name("change_reset_password").Put("/users", responses.ChangePassword)
	app.Get("/userrole/:user_id", NextFunc).Name("get_user_roles").Get("/userrole/:user_id", responses.GetUsersRolesByID)
	app.Post("/userrole/:user_id/:role_id", NextFunc).Name("user_role").Post("/userrole/:user_id/:role_id", responses.AddUserRoles)
	app.Delete("/userrole/:user_id/:role_id", NextFunc).Name("user_role").Delete("/userrole/:user_id/:role_id", responses.DeleteUserRoles)

	app.Get("/pages", NextFunc).Name("pages").Get("/pages", responses.GetPages)
	app.Post("/pages", NextFunc).Name("pages").Post("/pages", responses.PostPage)
	app.Get("/pages/:id", NextFunc).Name("page_single").Get("/pages/:id", responses.GetPageID)
	app.Patch("/pages/:id", NextFunc).Name("page_single").Patch("/pages/:id", responses.PatchPage)
	app.Delete("/pages/:id", NextFunc).Name("page_single").Delete("/pages/:id", responses.DeletePage)
	app.Get("/pagesroles/:page_id", NextFunc).Name("get_page_roles").Get("/pagesroles/:page_id", responses.GetPageRoles)
	app.Post("/pagerole/:page_id/:role_id", NextFunc).Name("page_roles").Post("/pagerole/:page_id/:role_id", responses.AddPageRoles)
	app.Delete("/pagerole/:page_id/:role_id", NextFunc).Name("page_roles").Delete("/pagerole/:page_id/:role_id", responses.DeletePageRoles)

	app.Get("/endpoints", NextFunc).Name("end_point").Get("/endpoints", responses.GetEndPointResponse)
	app.Post("/endpoints", NextFunc).Name("end_point").Post("/endpoints", responses.PostEndPoint)
	app.Get("/endpoints/:id", NextFunc).Name("end_point_single").Get("/endpoints/:id", responses.GetEndPointsID)
	app.Get("/endpointdrop", NextFunc).Name("drop_endpoints").Get("/endpointdrop", responses.GetDropEndPoints)
	app.Patch("/endpoints/:id", NextFunc).Name("end_point_single").Patch("/endpoints/:id", responses.PatchEndPoint)
	app.Delete("/endpoints/:id", NextFunc).Name("end_point_single").Delete("/endpoints/:id", responses.DeleteEndPoint)
	app.Patch("/feature_endpoint/:endpoint_id", NextFunc).Name("feature_endpoint").Patch("/feature_endpoint/:endpoint_id", responses.AddEndpointFeature)
	app.Delete("/feature_endpoint/:endpoint_id", NextFunc).Name("feature_endpoint").Delete("/feature_endpoint/:endpoint_id", responses.DeleteEndpointFeature)

	app.Post("/email", NextFunc).Name("send_email").Post("/email", responses.SendEmail)

	app.Post("/blobpic", NextFunc).Name("blob_picture").Post("/blobpic", responses.UploadingPicture)
	app.Post("/blobvideo", NextFunc).Name("blob_video").Post("/blobvideo", responses.UploadingVideo)
	app.Get("/stream/:file_name", responses.StreamingVideo)
	app.Get("/pics", responses.StreamingPicture)

	app.Get("/dashboard", NextFunc).Name("dashboard").Get("/dashboard", responses.GetDashBoardGrouped)
	app.Get("/dashboardends", NextFunc).Name("dashboard").Get("/dashboardends", responses.GetAppEndpoitnsGroupedBy)
	app.Get("/dashboardfeat", NextFunc).Name("dashboard").Get("/dashboardfeat", responses.GetAppFeaturesGroupedBy)
	app.Get("/dashboardpages", NextFunc).Name("dashboard").Get("/dashboardpages", responses.GetAppPages)
	app.Get("/dashboardroles", NextFunc).Name("dashboard").Get("/dashboardroles", responses.GetAppRoles)
	app.Get("/dashboardrolespage", NextFunc).Name("dashboard").Get("/dashboardrolespage", responses.GetAppPagesInRoles)

}

func init() {
	goBlueCmd.AddCommand(runCmd)

}
