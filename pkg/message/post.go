package message

import (
	"time"

	"github.com/google/uuid"
)

// Post is a message post on bulletin board
type Post struct {
	ID           string    `csv:"id" json:"id"`
	Name         string    `csv:"name" json:"name"`
	Email        string    `csv:"email" json:"email"`
	Text         string    `csv:"text" json:"text"`
	CreationTime time.Time `csv:"creation_time" json:"creation_time"`
}

// NewPost instantiates a new Post pointer to struct
func NewPost(poster, email, content string) *Post {
	return &Post{
		ID:           uuid.New().String(),
		Name:         poster,
		Email:        email,
		Text:         content,
		CreationTime: time.Now(),
	}
}
