package app

import (
	"city6/au/controllers"
	"city6/au/middleware"

	"github.com/kataras/iris"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

// InitApp is
func InitApp() {
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

	app.Get("/", controllers.IndexHandler)
	app.Get("/create", controllers.CreateAppSecret)
	app.Post("/create-link", controllers.LinkAppSecret)

	// navigate to http://localhost:8080
	// and open the ./logs.txt file
	if err := app.Run(iris.Addr(":8080"), iris.WithoutBanner); err != nil {
		if err != iris.ErrServerClosed {
			app.Logger().Warn("Shutdown with error: " + err.Error())
		}
	}
}
