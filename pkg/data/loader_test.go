package data

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/antonyho/go-message-board-example/pkg/message"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/testfile.csv")
	assert.NoError(t, err)
	posts, err := Load(data)
	assert.NoError(t, err)
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
