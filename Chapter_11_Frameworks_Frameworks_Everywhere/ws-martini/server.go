package main

import (
	"github.com/sausheong/ws-martini/post"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/binding"

)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	
	
	m.Get("/post/:id", post.Retrieve, post.HandleGet)
	m.Post("/post", binding.Json(post.Post{}), post.HandlePost)
	m.Put("/post/:id", post.Retrieve, post.HandlePut)
	m.Delete("/post/:id", post.Retrieve, post.HandleDelete)
	m.Run()

}
