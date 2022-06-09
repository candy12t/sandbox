package main

import (
	"github.com/candy12t/server/internal/web/router"
	"github.com/candy12t/server/internal/web/server"
)

func main() {
	r := router.Initialize()
	server.ListenAndServe(r)
}
