package post

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"strconv"
)

func Retrieve(c martini.Context, params martini.Params, r render.Render) {
	id, _ := strconv.Atoi(params["id"])
	post, err := retrieve(id)
	if err != nil {
		r.Error(404)
		return
	}
	c.Map(post)
}

// Retrieve a post
// GET /post/1
func HandleGet(post Post, r render.Render) {
	r.JSON(200, post)
}

// Create a post
// POST /post/
func HandlePost(post Post, r render.Render) {
	err := post.create()
	if err != nil {
		r.Error(500)
		return
	}
	r.Status(200)
}

// Update a post
// PUT /post/1
func HandlePut(post Post, r render.Render) {
	err := post.update()
	if err != nil {
		r.Error(500)
		return
	}
	r.Status(200)
}

// Delete a post
// DELETE /post/1
func HandleDelete(post Post, r render.Render) {
	err := post.delete()
	if err != nil {
		r.Error(500)
		return
	}
	r.Status(200)
}
