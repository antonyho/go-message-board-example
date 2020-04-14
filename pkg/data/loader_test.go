package data

import (
	"io"
	"os"
	"testing"
	"time"

	"github.com/antonyho/go-message-board-example/pkg/message"
	"github.com/stretchr/testify/assert"
)

func TestLoader(t *testing.T) {
	file, err := os.Open("testdata/testfile.csv")
	assert.NoError(t, err)
	loader, err := NewLoader(file)
	assert.NoError(t, err)
	posts := make([]*message.Post, 0, 0)
	for post, err := loader.ReadLine(); err != io.EOF; post, err = loader.ReadLine() {
		if err != nil {
			assert.Fail(t, err.Error())
		}
		posts = append(posts, post)
	}

	assert.Len(t, posts, 5)
	expectedFifthPostTime, _ := time.Parse(time.RFC3339, "2019-01-22T21:36:21-08:00")
	expectedFifthPost := &message.Post{
		ID:           "AF380451-7EE5-3867-02C1-9D2ACC4259A2",
		Name:         "Lawrence Alston",
		Email:        "rutrum.eu.ultrices@at.org",
		Text:         "est. Mauris eu turpis. Nulla aliquet. Proin velit.",
		CreationTime: expectedFifthPostTime,
	}
	assert.Equal(t, expectedFifthPost, posts[4])
}
