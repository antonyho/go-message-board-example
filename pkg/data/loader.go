package data

import (
	"encoding/csv"
	"io"
	"time"

	"github.com/antonyho/go-message-board-example/pkg/message"
	"github.com/gocarina/gocsv"
)

// PostData structure from data file
type PostData struct {
	ID           string    `csv:"id" json:"id"`
	Name         string    `csv:"name" json:"name"`
	Email        string    `csv:"email" json:"email"`
	Text         string    `csv:"text" json:"text"`
	CreationTime time.Time `csv:"creation_time" json:"creation_time"`
}

// Loader is a line by line CSV loader
type Loader struct {
	unmarshaller *gocsv.Unmarshaller
}

// NewLoader instantiates a new Loader pointer to struct
// A constructor function should not return error normally...
func NewLoader(r io.Reader) (*Loader, error) {
	umrshlr, err := gocsv.NewUnmarshaller(csv.NewReader(r), &PostData{})
	if err != nil {
		return nil, err
	}
	return &Loader{unmarshaller: umrshlr}, nil
}

// ReadLine returns each next CSV row as Post
func (l *Loader) ReadLine() (*message.Post, error) {
	postData, err := l.unmarshaller.Read()
	if err != nil {
		return nil, err
	}
	loadedPost := postData.(*PostData)
	post := message.Load(
		loadedPost.ID,
		loadedPost.Name,
		loadedPost.Email,
		loadedPost.Text,
		loadedPost.CreationTime,
	)
	return post, nil
}
