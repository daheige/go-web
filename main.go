package main

import (
	"fmt"
	"log"
	"net/http"

	"flag"

	"github.com/julienschmidt/httprouter"
)

var port int

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func init() {
	flag.IntVar(&port, "port", 8080, "server port")
	flag.Parse()
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/index", Index)

	router.GET("/hello/:name", Hello)
	router.GET("/api/index", Index)

	//启动服务器
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	log.Println("server has run on: ", port)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal(err)
	}
}
