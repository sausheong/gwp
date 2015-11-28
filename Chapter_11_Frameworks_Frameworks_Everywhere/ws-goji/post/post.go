package post

import (
	"github.com/zenazn/goji/web"
	"encoding/json"
	"net/http"
	"strconv"
)

// Retrieve a post
// GET /post/1
func HandleGet(c web.C, w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(c.URLParams["id"])

	post, err := retrieve(id)
	if err != nil {
		http.Error(w, "Cannot get post", 404)
		return
	}
	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		http.Error(w, "Cannot unmarshal JSON", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// Create a post
// POST /post/
func HandlePost(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var post Post
	json.Unmarshal(body, &post)
	err := post.create()
	if err != nil {
		http.Error(w, "Cannot create post", 500)
		return
	}
	w.WriteHeader(200)
	return
}

// Update a post
// PUT /post/1
func HandlePut(c web.C, w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(c.URLParams["id"])
	post, err := retrieve(id)
	if err != nil {
		http.Error(w, "Cannot get post", 404)
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &post)
	err = post.update()
	if err != nil {
		http.Error(w, "Cannot update post", 500)
		return
	}
	w.WriteHeader(200)
	return
}

// Delete a post
// DELETE /post/1
func HandleDelete(c web.C, w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(c.URLParams["id"])
	post, err := retrieve(id)
	if err != nil {
		http.Error(w, "Cannot get post", 404)
		return
	}
	err = post.delete()
	if err != nil {
		http.Error(w, "Cannot delete post", 500)
		return
	}
	w.WriteHeader(200)
	return
}
