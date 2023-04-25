package storage

import (
	"time"

	"github.com/google/uuid"
)

//

type JournalEntry struct {
	Id uuid.UUID `json:"id"`

	Key     string `json:"key"`
	Content string `json:"content"`

	CreatedAt time.Time `json:"created_at"`
}

type Journal interface {
	Write(key, content string) error
}

//

type NopJournal struct{}

var _ Journal = NopJournal{}

func (j NopJournal) Write(key, content string) error {
	return nil
}

//

type HeapJournal struct {
	s []JournalEntry
}

func NewHeapJournal() Journal {
	return &HeapJournal{}
}

var _ Journal = &HeapJournal{}

func (j HeapJournal) Write(key, content string) error {
	j.s = append(j.s, JournalEntry{
		Id:        uuid.New(),
		Key:       key,
		Content:   content,
		CreatedAt: time.Now(),
	})
	return nil
}
