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

type Api struct {
	Db *gorm.DB
}

type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type PostId struct {
	Id string `json:"id"`
}

type Answer struct {
	gorm.Model
	Question Question `gorm:"foreignKey:Question";json:"question_id"`
	Post     Post     `gorm:"embedded";json:"post"`
}

type Question struct {
	gorm.Model
	Post    Post     `gorm:"embedded";json:"post"`
	Answers []Answer `gorm:"foreignKey:Answer";json:"answer_ids"`
}

func (a Api) UpdatePost(id PostId) (string, error) {
	return fmt.Sprint("Post updated"), nil
}

func (a Api) GetPost(id PostId) (string, error) {
	return fmt.Sprint("Post getd"), nil
}

func (a Api) PublishPost(post Post) (PostId, error) {
	return publishPost(&post, a.Db)
}

func (a Api) DeletePost(id PostId) (string, error) {
	return fmt.Sprint("Post deleted"), nil
}

func GetApi(db *gorm.DB) Api {
	db.AutoMigrate(&Answer{}, &Question{}, &Post{})
	return Api{Db: db}
}
