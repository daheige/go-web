package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type BaseController struct{}

func (b BaseController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte(`{"code":200,"message":"ok"}`))
}
