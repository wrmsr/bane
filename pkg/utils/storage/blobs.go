package storage

import "time"

type Blob struct {
	Key   string `json:"key"`
	Value string `json:"value"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BlobStorage interface {
	Get(key string) (Blob, error)
	Put(key string, value string) error
}
