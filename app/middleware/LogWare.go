package middleware

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type LogWare struct{}

func (LogWare) Access(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log.Println("exec begin")
		log.Println("request uri: ", r.RequestURI)
		h(w, r, ps)

		log.Println("exec end")
	}
}
