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

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PATCH", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Authorization", "Content-Type", "Accept"},
		AllowCredentials: true, // allows everything, use that to change the hosts.
	})

	app.Use(crs)

	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())
	// app.UseGlobal(func(ctx iris.Context) {
	// 	ctx.Header("Access-Control-Allow-Origin", "*")
	// 	ctx.Next()
	// })

	app.Use(middleware.Before)

	// attach the file as logger, remember, iris' app logger is just an io.Writer.
	app.Logger().SetOutput(newLogFile())

	v1 := app.Party("/api/v1")
	{
		v1.Use(cors.Default())
		v1.Get("/", handler.IndexHandler)
		v1.Get("/create", handler.CreateAppSecret)
		v1.Options("/create", handler.CreateAppSecret)
		v1.Post("/create-link", handler.LinkAppSecret)
		v1.Options("/create-link", handler.LinkAppSecret)
		v1.Get("/get-link", handler.GetLinkUser)
		v1.Options("/get-link", handler.GetLinkUser)
		// test demo
		v1.Get("/test/sign", handler.SingDemoT)
	}

	// app.WrapRouter(cors.WrapNext(cors.Options{
	// 	AllowedOrigins:   []string{"*"},
	// 	AllowCredentials: true,
	// }))

	// navigate to defafult config http://localhost:8080
	if err := app.Run(iris.Addr(":"+models.Config.APP.Port), iris.WithoutBanner); err != nil {
		if err != iris.ErrServerClosed {
			app.Logger().Warn("Shutdown with error: " + err.Error())
		}
	}
}
