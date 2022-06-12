package cache

import (
	"log"
	"time"

	memoryCache "github.com/patrickmn/go-cache"
)

type Memory struct {
	Expiration time.Duration
	cache      *memoryCache.Cache
}

func (m *Memory) Type() string {
	return "InMemoryCache"
}

func (m *Memory) Connect() {
	defaultDuration := m.Expiration * time.Minute
	cleanupInterval := m.Expiration * time.Minute
	m.cache = memoryCache.New(defaultDuration, cleanupInterval)
	log.Println("Using in-memory cache")
}

func (m *Memory) Get(key string) (interface{}, bool) {
	return m.cache.Get(key)
}

func (m *Memory) Set(key string, value interface{}) error {
	duration := m.Expiration * time.Minute
	m.cache.Set(key, value, duration)
	return nil
}
