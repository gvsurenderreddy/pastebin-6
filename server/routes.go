package server

import (
	"github.com/hoisie/mustache"
	"github.com/hoisie/web"
)

func routes() {
	web.Put("/", paste)  //Defined in paste.go
	web.Post("/", paste) //Defined in paste.go
	web.Get("/([a-zA-Z0-9]+)", getPaste)
	web.Get("/", index)
}

func index(ctx *web.Context) string {
	return mustache.RenderFile("templates/index.html")
}
