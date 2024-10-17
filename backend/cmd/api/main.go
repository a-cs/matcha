package main

import (
	"backend/internal/app"
	_ "github.com/lib/pq"
)

func main() {
	if err := app.Start(); err != nil {
		panic(err)
	}
}
