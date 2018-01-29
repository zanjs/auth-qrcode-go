package app

import (
	"city6/au/handler"
	"city6/au/middleware"
	"city6/au/models"
	"city6/au/utils"
	"fmt"

	"github.com/kataras/iris"

	"github.com/iris-contrib/middleware/cors"
	"github.com/jinzhu/configor"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

// InitApp is
func InitApp() {

	configor.Load(&models.Config, "app.yml")
	fmt.Printf("config port : %#v", models.Config)

	utils.Mkdir(utils.LogFIlePath)

	f := newLogFile()
	defer f.Close()

	app := iris.New()
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
	})

	app.Use(crs)

	app.Use(middleware.Before)

	// attach the file as logger, remember, iris' app logger is just an io.Writer.
	app.Logger().SetOutput(newLogFile())

	app.Get("/", handler.IndexHandler)
	app.Get("/create", handler.CreateAppSecret)
	app.Post("/create-link", handler.LinkAppSecret)
	app.Get("/get-link", handler.GetLinkUser)

	// test demo
	app.Get("/test/sign", handler.SingDemoT)

	// navigate to defafult config http://localhost:8080
	if err := app.Run(iris.Addr(":"+models.Config.APP.Port), iris.WithoutBanner); err != nil {
		if err != iris.ErrServerClosed {
			app.Logger().Warn("Shutdown with error: " + err.Error())
		}
	}
}
