package cache

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func NewRedisClient(server *miniredis.Miniredis) *Redis {
	r := &Redis{
		Addr:       server.Host() + ":" + server.Port(),
		Expiration: 30,
		client:     &redis.Client{},
	}
	return r
}

func TestRedis_All(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	// start Redis mock server
	server := miniredis.RunT(t)
	defer server.Close()

	tests := []struct {
		name  string
		key   string
		value string
	}{
		{
			name:  "Check Redis Cache Connect, Ping, Set & Get with valid JSON",
			key:   "Key",
			value: "Test",
		},
	}
	r := NewRedisClient(server)
	r.Connect()
	defer r.client.Close()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NoError(r.Ping())
			r.Set(tt.key, tt.value)
			got, foundKey := r.Get(tt.key)
			assert.True(foundKey, "Cache Set & Get Sequenced Failed for key = %v", tt.key)
			assert.Equal(tt.value, got, "Get () = %v, want = %v", got, tt.value)
		})
	}
}
