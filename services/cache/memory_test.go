package cache

import (
	"testing"
	"time"

	memoryCache "github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
)

func TestMemory_All(t *testing.T) {
	assert := assert.New(t)
	type fields struct {
		Expiration time.Duration
		cache      *memoryCache.Cache
	}
	tests := []struct {
		name   string
		fields fields
		key    string
		want   string
	}{
		{"Check InMemory Cache ", fields{
			Expiration: 0,
			cache:      &memoryCache.Cache{},
		}, "Key", "Test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Memory{
				Expiration: tt.fields.Expiration,
				cache:      tt.fields.cache,
			}
			m.Connect()
			m.Set(tt.key, tt.want)
			got, found := m.Get(tt.key)
			assert.True(found, "Cache Set & Get Sequenced Failed fro key = %v", tt.key)
			assert.Equal(tt.want, got, "Get () = %v, want = %v", got, tt.want)
		})
	}
}
