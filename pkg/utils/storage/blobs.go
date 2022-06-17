package storage

import "time"

//

type Blob struct {
	Key   string `json:"key"`
	Value string `json:"value"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BlobStorage interface {
	Get(key string) (Blob, bool, error)
	Put(key, value string) error
}

//

type NopBlobStorage struct{}

var _ BlobStorage = NopBlobStorage{}

func (n NopBlobStorage) Get(key string) (Blob, bool, error) {
	panic("unsupported")
}

func (n NopBlobStorage) Put(key, value string) error {
	panic("unsupported")
}

//

type HeapBlobStorage struct {
	m map[string]Blob
}

func NewHeapBlobStorage() BlobStorage {
	return &HeapBlobStorage{
		m: make(map[string]Blob),
	}
}

var _ BlobStorage = &HeapBlobStorage{}

func (bs *HeapBlobStorage) Get(key string) (Blob, bool, error) {
	b, ok := bs.m[key]
	return b, ok, nil
}

func (bs *HeapBlobStorage) Put(key, value string) error {
	now := time.Now()
	ca := now
	if b, ok := bs.m[key]; ok {
		ca = b.CreatedAt
	}
	bs.m[key] = Blob{
		Key:       key,
		Value:     value,
		CreatedAt: ca,
		UpdatedAt: now,
	}
	return nil
}
