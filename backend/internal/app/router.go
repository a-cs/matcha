package app

import (
	"backend/internal/adpter/in/handler"
	"backend/internal/adpter/in/rest"
	goji "goji.io"
	"goji.io/pat"
	"net/http"
)

func initRouter(handler handler.Handler) *goji.Mux {
	mux := goji.NewMux()
	mux.Use(corsMiddleware)

	mapRoutes(mux, handler)

	return mux
}

func mapRoutes(mux *goji.Mux, handlers handler.Handler) {
	rest.Routes(mux, handlers)
	mux.HandleFunc(pat.Get("/ping"), ping)
}

func ping(serviceResponse http.ResponseWriter, _ *http.Request) {
	serviceResponse.WriteHeader(http.StatusOK)
	serviceResponse.Write([]byte("pong"))
}
