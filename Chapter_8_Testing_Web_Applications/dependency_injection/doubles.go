package main

type FakePost struct {
	Id      int
	Content string
	Author  string
}

func (post *FakePost) fetch(id int) (err error) {
	post.Id = id
	return
}

func (post *FakePost) create() (err error) {
	return
}

func (post *FakePost) update() (err error) {
	return
}

func (post *FakePost) delete() (err error) {
	return
}
