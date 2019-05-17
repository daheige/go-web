package routers

import (
	"go-web/app/controller"
	"go-web/app/middleware"

	"github.com/julienschmidt/httprouter"
)

func InitRouter(router *httprouter.Router) {
	homeCtrl := &controller.HomeController{}

	router.GET("/", homeCtrl.Index)
	router.GET("/test", homeCtrl.Test)

	router.GET("/index", homeCtrl.Index)

	// http://localhost:1339/hello/daheige
	router.GET("/hello/:name", homeCtrl.Hello)

	logWare := middleware.LogWare{}
	router.GET("/api/get-data", logWare.Access(homeCtrl.Foo))
}
