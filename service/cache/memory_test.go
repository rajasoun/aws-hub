package cache

import (
	"testing"

	memoryCache "github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
)

func NewMemoryCacheClient() *Memory {
	m := &Memory{
		Expiration: 0,
		cache:      &memoryCache.Cache{},
	}
	return m
}

func TestMemoryCachePingGetSet(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	tests := []struct {
		name     string
		key      string
		value    string
		wantType string
	}{
		{"Check InMemory Cache ", "Key", "Test", "InMemoryCache"},
	}
	m := NewMemoryCacheClient()
	m.Connect()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.Set(tt.key, tt.value)
			got, found := m.Get(tt.key)
			assert.True(found, "Cache Set & Get Sequenced Failed for key = %v", tt.key)
			assert.Equal(tt.value, got, "Get () = %v, want = %v", got, tt.value)
			assert.Equal(tt.wantType, m.Type(), "Get () = %v, want = %v", m.Type(), tt.wantType)
		})
	}
}
