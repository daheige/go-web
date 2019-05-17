package controller

import (
	"fmt"
	"net/http"

	"encoding/json"

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
