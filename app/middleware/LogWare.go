package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type LogWare struct{}

//和RequestWare中的LogAccess功能一样，为了兼容httpRouter路由设置
//对于httprouter.Handler加访问日志
func (LogWare) Access(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		RequestAccess(h, w, r, ps)
	}
}
