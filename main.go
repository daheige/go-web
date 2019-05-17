package main

import (
	"fmt"
	"go-web/app/routers"
	"log"
	"net/http"

	"flag"

	"github.com/julienschmidt/httprouter"
)

var port int

func init() {
	flag.IntVar(&port, "port", 1339, "server port")
	flag.Parse()
}

func main() {
	mux := httprouter.New()

	//初始化路由规则
	routers.InitRouter(mux)

	//启动服务器
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	log.Println("server has run on: ", port)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal(err)
	}
}
