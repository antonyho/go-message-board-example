package data

import (
	"github.com/antonyho/go-message-board-example/pkg/message"
	"github.com/gocarina/gocsv"
)

// Load CSV file content into message Post slice
func Load(csvContent []byte) ([]*message.Post, error) {
	posts := new([]*message.Post)

	if err := gocsv.UnmarshalBytes(csvContent, posts); err != nil {
		return nil, err
	}

	return *posts, nil
}
