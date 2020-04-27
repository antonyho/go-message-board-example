package message

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

// Post is a message post on bulletin board
type Post struct {
	id           string
	name         string
	email        string
	text         string
	creationTime time.Time
	updateMutex  sync.Mutex
}

// NewPost instantiates a new Post pointer to struct
func NewPost(poster, email, content string) *Post {
	return &Post{
		id:           uuid.New().String(),
		name:         poster,
		email:        email,
		text:         content,
		creationTime: time.Now(),
		updateMutex:  sync.Mutex{},
	}
}

// Load instantiates a post pointer to struct from data source
func Load(id, poster, email, content string, creationTime time.Time) *Post {
	return &Post{
		id:           id,
		name:         poster,
		email:        email,
		text:         content,
		creationTime: creationTime,
		updateMutex:  sync.Mutex{},
	}
}

// Update a post only on its content
func (p *Post) Update(content string) {
	p.updateMutex.Lock()
	p.text = content
	p.updateMutex.Unlock()
}

// ID is the getter of Post id field
func (p *Post) ID() string {
	return p.id
}

// Name is the getter of Post name field
func (p *Post) Name() string {
	return p.name
}

// Email is the getter of Post email field
func (p *Post) Email() string {
	return p.email
}

// Text is the getter of Post text field
func (p *Post) Text() string {
	return p.text
}

// CreationTime is the getter of Post creationTime field
func (p *Post) CreationTime() time.Time {
	return p.creationTime
}
