package core

import (
	"fmt"
)

type Master interface {
	PublishPost(PostId) (string, error)
	GetPost(PostId) (string, error)
	UpdatePost(PostId) (string, error)
	DeletePost(PostId) (string, error)
}

type Api struct{}

type post struct {
	id      string
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type PostId struct {
    Id string `json:"id"`
}

type Answer struct {
	Question_id string
	Post        post
}

type Question struct {
	Post    post
	Answers []string
}

func (a Api) UpdatePost(id PostId) (string, error) {
	return fmt.Sprint("Post updated"), nil
}

func (a Api) GetPost(id PostId) (string, error) {
	return fmt.Sprint("Post getd"), nil
}

func (a Api) PublishPost(id PostId) (string, error) {
	return fmt.Sprint("Post publishd"), nil
}

func (a Api) DeletePost(id PostId) (string, error) {
	return fmt.Sprint("Post deleted"), nil
}

func GetApi() Master {
	return Api{}
}
