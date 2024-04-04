package flashstorage

import (
	"sync"
)

type MemoryStore struct {
	sync.RWMutex
	data map[string]string
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[string]string),
	}
}

func (ms *MemoryStore) set(key string, value string) {
	ms.data[key] = value
}

func (ms *MemoryStore) get(key string) string {
	if ms.count() > 0 {
		item := ms.data[key]
		return item
	}

	return ""
}

func (ms *MemoryStore) count() int {
	return len(ms.data)
}

func (ms *MemoryStore) Set(key string, value string) {
	ms.Lock()
	defer ms.Unlock()
	ms.set(key, value)
}

func (ms *MemoryStore) Get(key string) string {
	ms.RLock()
	defer ms.RUnlock()
	return ms.get(key)
}

func (ms *MemoryStore) Delete(key string) {
	ms.Lock()
	defer ms.Unlock()

	delete(ms.data, key)
}

func (ms *MemoryStore) Count() int {
	ms.RLock()
	defer ms.RUnlock()
	return ms.count()
}
