package data

import (
	"encoding/csv"
	"io"

	"github.com/antonyho/go-message-board-example/pkg/message"
	"github.com/gocarina/gocsv"
)

// Loader is a line by line CSV loader
type Loader struct {
	unmarshaller *gocsv.Unmarshaller
}

// NewLoader instantiates a new Loader pointer to struct
// A constructor function should not return error normally...
func NewLoader(r io.Reader) (*Loader, error) {
	umrshlr, err := gocsv.NewUnmarshaller(csv.NewReader(r), &message.Post{})
	if err != nil {
		return nil, err
	}
	return &Loader{unmarshaller: umrshlr}, nil
}

// ReadLine returns each next CSV row as Post
func (l *Loader) ReadLine() (*message.Post, error) {
	post, err := l.unmarshaller.Read()
	if err != nil {
		return nil, err
	}
	return post.(*message.Post), nil
}
