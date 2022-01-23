package core

import (
	"fmt"
	"reflect"
	"runtime"
)

type Master interface {
	PublishPost(string) (string, error)
	GetPost(string) (string, error)
	UpdatePost(string) (string, error)
	DeletePost(string) (string, error)
}

type Api struct {}

type post struct {
	id      string
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type Answer struct {
	Question_id string
	Post        post
}

type Question struct {
	Post    post
	Answers []string
}

func GetFunctionName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func (a Api) UpdatePost(id string) (string, error) {
	return fmt.Sprint("Post updated"), nil
}

func (a Api) GetPost(id string) (string, error) {
	return fmt.Sprint("Post getd"), nil
}

func (a Api) PublishPost(id string) (string, error) {
	return fmt.Sprint("Post publishd"), nil
}

func (a Api) DeletePost(id string) (string, error) {
	return fmt.Sprint("Post deleted"), nil
}

func GetApi() Master {
    return Api{}
}

