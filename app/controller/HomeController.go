package controller

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"

	"go-web/app/config"

	"github.com/julienschmidt/httprouter"
)

type HomeController struct {
	BaseController
}

func (ctrl HomeController) Test(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte(`{"code":200,"message":"this is test"}`))
}

func (ctrl HomeController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func (ctrl HomeController) Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func (ctrl HomeController) Foo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data := map[string]interface{}{
		"code":    200,
		"message": "ok",
		"data":    nil,
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"code":500,"message":"server error!"}`))
		return
	}

	w.Write(jsonBytes)
}

func (ctrl HomeController) MockPanic(w http.ResponseWriter, r *http.Request) {
	panic("exec panic")
}

func (ctrl HomeController) Info(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"code":200,"message":"ok"}`))
}

//A very simple redis set data
func (ctrl HomeController) SetData(w http.ResponseWriter, r *http.Request) {
	redisObj, err := config.GetRedisObj("default")
	if err != nil {
		log.Println(err)
		w.Write([]byte(`{"code":500,"message":"redis connection error"}`))
		return
	}

	//用完就需要释放连接，防止过多的连接导致redis连接过多而陷入长久等待，从而redis崩溃
	defer redisObj.Close()

	_, err = redisObj.Do("set", "myname", "daheige")
	if err != nil {
		w.Write([]byte(`{"code":500,"message":"set redis data error"}`))
		return
	}

	w.Write([]byte(`{"code":200,"message":"ok"}`))
}
