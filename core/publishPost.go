package core

import (
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

func publishPost(post *Post, db *gorm.DB) (PostId, error) {
	db.Create(post)
	stringId := strconv.FormatUint(uint64(post.ID), 10)
    fmt.Printf("\nnew post with id %v created by user %v\n", stringId, post.Author)
	return PostId{Id: stringId}, nil
}
