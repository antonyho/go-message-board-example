package message_test

import (
	"io"
	"os"
	"testing"

	"github.com/antonyho/go-message-board-example/pkg/data"
	"github.com/antonyho/go-message-board-example/pkg/message"
	"github.com/stretchr/testify/assert"
)

func TestBoard(t *testing.T) {
	file, err := os.Open("testdata/testfile.csv")
	assert.NoError(t, err)
	loader, err := data.NewLoader(file)
	assert.NoError(t, err)
	posts := make([]*message.Post, 0, 0)
	for post, err := loader.ReadLine(); err != io.EOF; post, err = loader.ReadLine() {
		if err != nil {
			assert.Fail(t, err.Error())
		}
		posts = append(posts, post)
	}

	board := message.NewBoard()
	board.Load(posts)
	listedPosts := board.List()
	assert.Len(t, listedPosts, 5)
	// Make sure the post are ordered anti-chronologically
	assert.Equal(t, listedPosts[0].ID, "AF380451-7EE5-3867-02C1-9D2ACC4259A2")

	_, err = board.View("post-does-not-exist")
	assert.Error(t, err)

	err = board.Update("post-does-not-exist", "doesn't matter")
	assert.Error(t, err)

	p, err := board.View("AF380451-7EE5-3867-02C1-9D2ACC4259A2")
	assert.NoError(t, err)
	assert.Equal(t, "est. Mauris eu turpis. Nulla aliquet. Proin velit.", p.Text)

	err = board.Update("AF380451-7EE5-3867-02C1-9D2ACC4259A2", "Updated content")
	assert.NoError(t, err)

	updatedPost, err := board.View("AF380451-7EE5-3867-02C1-9D2ACC4259A2")
	assert.NoError(t, err)
	assert.Equal(t, "Updated content", updatedPost.Text)

	board.Paste("Hello world!", "Antony", "do-not-reply@email.com")
	// Make sure my new post is there
	allNewPosts := board.List()
	assert.Equal(t, allNewPosts[0].Text, "Hello world!")
}
