package controllers

import (
	"strconv"
	"fmt"
	"github.com/sausheong/ws-b/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Post
type PostController struct {
	beego.Controller
}

// @Title create
// @Description create post
// @Param	body		body 	models.Post	true		The Post data
// @Success 200 {string} models.Posts.Id
// @Failure 403 body is empty
// @router / [post]
func (p *PostController) Post() {
	var post models.Posts
	json.Unmarshal(p.Ctx.Input.RequestBody, &post)
	id, err := models.AddPost(&post)
	if err != nil {
		p.Data["json"] = err
	} else {
		p.Data["json"] = id
	}
	p.ServeJson()
}

// @Title Get
// @Description find post by id
// @Param	id		path 	string	true		the id you want to get
// @Success 200 {post} models.Posts
// @Failure 403 :id is empty
// @router /:id [get]
func (p *PostController) Get() {
	id, _ := strconv.ParseInt(p.Ctx.Input.Params[":id"], 10, 64)
	post := &models.Posts{Id: id}
	if id != 0 {
		err := models.GetPost(post)
		if err != nil {
			p.Data["json"] = err
		} else {
			p.Data["json"] = post
		}
	}

	p.ServeJson()
}

// @Title update
// @Description update the Post
// @Param	id		path 	string	true		The id you want to update
// @Param	body		body 	models.Posts	true		The Post data
// @Success 200 {post} models.Posts
// @Failure 403 :id is empty
// @router /:id [put]
func (p *PostController) Put() {
	id, _ := strconv.ParseInt(p.Ctx.Input.Params[":id"], 10, 64)
	post := &models.Posts{Id: id}
	json.Unmarshal(p.Ctx.Input.RequestBody, post)
	
	err := models.UpdatePost(post)
	if err != nil {
		p.Data["json"] = err
	} else {
		p.Data["json"] = "update success!"
	}
	p.ServeJson()
}

// @Title delete
// @Description delete the object
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (p *PostController) Delete() {
	id, _ := strconv.ParseInt(p.Ctx.Input.Params[":id"], 10, 64)
	post := &models.Posts{Id: id}
	json.Unmarshal(p.Ctx.Input.RequestBody, post)	
	models.DeletePost(post)
	p.Data["json"] = "delete success!"
	p.ServeJson()
}

