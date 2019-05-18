package middleware

import (
	"go-web/app/extension/Logger"
	"log"
	"net/http"
	"time"

	"github.com/daheige/thinkgo/common"

	"go-web/app/utils"

	"github.com/julienschmidt/httprouter"
)

type RequestWare struct {
}

//http.HandlerFunc实现了Handler的ServeHTTP方法，就实现了Handler接口
//http.HandlerFunc 可以把这样签名的函数转换为一个处理器
func (this *RequestWare) NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message":"404 - not found"}`))
}

//当请求发生了panic的时候处理函数
func (this *RequestWare) PanicHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	log.Println(err)
	//请求结束记录日志
	Logger.Emergency(r, "exec end", map[string]interface{}{
		"trace_error": err,
	})

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{"message":"server error!"}`))
}

//对于原生的http.HandleFunc加访问日志
func (this *RequestWare) Access(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		RequestAccess(h, w, r, nil)
	}
}

//请求访问日志
func RequestAccess(h interface{}, w http.ResponseWriter, r *http.Request, params interface{}) {
	t := time.Now()
	log.Println("request before")
	log.Println("request uri: ", r.RequestURI)

	reqId := r.Header.Get("x-request-id")
	if reqId == "" {
		reqId = common.RndUuidMd5()
	}

	log.Println("log_id: ", reqId)
	//将requestId 写入当前上下文中
	r = utils.ContextSet(r, "log_id", reqId)
	Logger.Info(r, "exec begin", nil)

	switch val := h.(type) {
	case httprouter.Handle: //httprouter处理器
		val(w, r, params.(httprouter.Params))
	case http.Handler: //http原生的处理器
		val.ServeHTTP(w, r)
	case http.HandlerFunc: //http原生的处理器函数
		val(w, r)
	default:
		panic("error handler")
	}

	log.Println("request end")
	//请求结束记录日志
	Logger.Info(r, "exec end", map[string]interface{}{
		"exec_time": time.Now().Sub(t).Seconds(),
	})
}
