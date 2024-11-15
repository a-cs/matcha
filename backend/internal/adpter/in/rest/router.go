package rest

import (
	"backend/internal/adpter/in/handler"
	goji "goji.io"
	"goji.io/pat"
)

func Routes(mux *goji.Mux, handlers handler.Handler) {
	mux.HandleFunc(pat.Post("/signup"), handlers.CreateUser)
	mux.HandleFunc(pat.Post("/confirm"), handlers.ConfirmAccount)
	mux.HandleFunc(pat.Post("/login"), handlers.Login)
	mux.HandleFunc(pat.Get("/profile/:username"), handlers.GetProfile)
	mux.HandleFunc(pat.Put("/profile"), handlers.UpdateProfile)
}
