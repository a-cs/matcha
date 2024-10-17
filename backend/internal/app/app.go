package app

import (
	"backend/internal/adpter/in/handler"
	"backend/internal/application/port/in"
	"backend/internal/application/port/usecase"
	"fmt"
	"net/http"
	"time"
)

func Start() error {
	handlers := handler.NewHandler(CreateUser())

	mux := initRouter(handlers)
	loggedMux := logRoutes(mux)

	return http.ListenAndServe("localhost:8000", loggedMux)
}

func CreateUser() in.CreateUser {
	return usecase.NewCreateUserUseCase()
}

func logRoutes(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		fmt.Printf("Route: %s %s %s\n", r.Method, r.RequestURI, time.Since(start))
	})
}
