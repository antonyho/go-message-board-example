package message

import (
	"errors"
	"sort"
)

// Errors for message board operations
var (
	ErrPostNotExist = errors.New("post does not exist")
)

// Board is for message bulletin
type Board struct {
	posts map[string]*Post // PostID:Post
}

// NewBoard instantiates a new Board pointer to struct
func NewBoard() *Board {
	return &Board{posts: make(map[string]*Post)}
}

// Load a slice of post into message board
// Assuming all posts in the slice has proper UUID(v4)
// Post with duplicated UUID will be overridden with the latter
func (b *Board) Load(posts []*Post) {
	for _, p := range posts {
		b.posts[p.ID] = p
	}
}

// LoadOne a single post into message board
// Assuming the post has proper UUID(v4)
// Post with duplicated UUID will be overridden with the latter
func (b *Board) LoadOne(post *Post) {
	b.posts[post.ID] = post
}

// List all posts on message board
func (b *Board) List() []*Post {
	resp := make([]*Post, 0, len(b.posts))
	for _, p := range b.posts {
		resp = append(resp, p)
	}

	sort.Slice(resp, func(i, j int) bool { return resp[i].CreationTime.After(resp[j].CreationTime) })

	return resp
}

// View provides the post for given ID
func (b *Board) View(id string) (*Post, error) {
	p, exist := b.posts[id]
	if !exist {
		return nil, ErrPostNotExist
	}

	return p, nil
}

// Paste a new message to the board
func (b *Board) Paste(message, poster, email string) {
	newPost := NewPost(poster, email, message)
	b.posts[newPost.ID] = newPost
}

// Update an existing post for given ID
// This function is not safe on concurrent operation
func (b *Board) Update(id, content string) error {
	p, exist := b.posts[id]
	if !exist {
		return ErrPostNotExist
	}

	p.Text = content

	return nil
}
