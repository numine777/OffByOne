package core

import (
	"fmt"
	"github.com/satori/go.uuid"
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
	QId  uuid.UUID `gorm:"primary_key;type:char(36)";json:"question_id"`
	Post Post      `gorm:"embedded";json:"post"`
	// Quon Qon `gorm:"foreignKey:Question";json:"question_id"`
}

type Question struct {
	gorm.Model
	ID      uuid.UUID `gorm:"primary_key;type:char(36)"`
	Post    Post      `gorm:"embedded";json:"post"`
	Answers []Answer  `gorm:"foreignKey:QId";json:"answer_ids"`
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
