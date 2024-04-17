package tests

import (
	"github.com/gofiber/fiber/v2"
	"semay.com/manager"
)

var TestApp *fiber.App

// initalaizing the app
func ReturnTestApp() {
	if TestApp == nil {
		app, _ := manager.MakeApp("test")
		TestApp = app
	}
}
