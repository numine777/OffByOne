package core

import (
	"fmt"
	"gorm.io/gorm"
)

type Master interface {
	PublishPost(PostId) (string, error)
	GetPost(PostId) (string, error)
	UpdatePost(PostId) (string, error)
	DeletePost(PostId) (string, error)
}

type Api struct{}

type Post struct {
	gorm.Model
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type PostId struct {
	Id string `json:"id"`
}

type Answer struct {
	gorm.Model
	QuestionId string `json:"question_id"`
	Post       Post   `json:"post"`
}

type Question struct {
	gorm.Model
	Post    Post     `json:"post"`
	Answers []string `json:"answer_ids"`
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
