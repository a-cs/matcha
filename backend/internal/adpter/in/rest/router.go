package rest

import (
	"backend/internal/adpter/in/handler"
	goji "goji.io"
	"goji.io/pat"
)

func Routes(mux *goji.Mux, handlers handler.Handler) {
	mux.HandleFunc(pat.Post("/signup"), handlers.CreateUser)
}