package app

import (
	"backend/internal/adpter/app"
	"backend/internal/adpter/in/handler"
	"backend/internal/application/port/in"
	"backend/internal/application/port/usecase"
	"fmt"
	"net/http"
	"os"
	"time"
)

func Start() error {
	app.Build()
	handlers := handler.NewHandler(CreateUser(), ConfirmAccount(), Login(), GetProfile())

	mux := initRouter(handlers)
	loggedMux := logRoutes(mux)

	return http.ListenAndServe("localhost:8000", loggedMux)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func CreateUser() in.CreateUser {
	return usecase.NewCreateUserUseCase()
}

func ConfirmAccount() in.ConfirmAccount {
	return usecase.NewConfirmAccountUseCase()
}

func Login() in.Login {
	return usecase.NewLoginUseCase()
}

func GetProfile() in.GetProfile {
	return usecase.NewGetProfileUseCase()
}

func logRoutes(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		fmt.Printf("Route: %s %s %s\n", r.Method, r.RequestURI, time.Since(start))
	})
}
