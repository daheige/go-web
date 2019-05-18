package routers

import (
	"go-web/app/controller"
	"go-web/app/middleware"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func InitRouter(router *httprouter.Router) {
	//全局中间件处理
	//路由找不到的时候，响应方法,如果不指定就会走默认的http notfound抛出404页面
	reqWare := &middleware.RequestWare{}
	logWare := &middleware.LogWare{}

	router.NotFound = http.HandlerFunc(reqWare.NotFound) //http.Handler
	router.PanicHandler = reqWare.PanicHandler

	homeCtrl := &controller.HomeController{}
	router.GET("/", logWare.Access(homeCtrl.Index))
	router.GET("/test", homeCtrl.Test)

	router.GET("/index", homeCtrl.Index)

	// http://localhost:1339/hello/daheige
	router.GET("/hello/:name", homeCtrl.Hello)

	router.GET("/api/get-data", logWare.Access(homeCtrl.Foo)) //对httprouter添加访问日志功能

	//支持原生的http.handleFunc
	//添加访问日志功能
	router.HandlerFunc("GET", "/info", reqWare.Access(homeCtrl.Info))

	//模拟panic操作
	//http://localhost:1339/mock-panic
	router.HandlerFunc("GET", "/mock-panic", reqWare.Access(homeCtrl.MockPanic))

	//模拟redis设置
	router.HandlerFunc("GET", "/set-data", homeCtrl.SetData)
}
