package cache

import (
	"bytes"
	"log"
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

func TestRedisPingGetSet(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	// start Redis mock server
	server := miniredis.RunT(t)
	defer server.Close()

	tests := []struct {
		name     string
		key      string
		value    string
		wantType string
	}{
		{
			name:     "Check Redis Cache Connect, Ping, Set & Get with valid JSON",
			key:      "Key",
			value:    "Test",
			wantType: "RedisCache",
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
			assert.Equal(tt.wantType, r.Type(), "Get () = %v, want = %v", r.Type(), tt.wantType)
		})
	}
}

func TestRedisGetWithErr(t *testing.T) {
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
			got, foundKey := r.Get(tt.key)
			assert.False(foundKey, "Cache Set & Get Should Fail with Injected Err = %v", tt.key)
			assert.Empty(got, "Get () = %v", got)
			//assert.NoError(r.Ping())
			server.SetError("Mock Error")
			var outputBuffer bytes.Buffer
			log.SetOutput(&outputBuffer)
			r.Set(tt.key, tt.value)
			gotLog := outputBuffer.String()
			assert.Contains(gotLog, "Mock Error", "Mock Set Failed")
		})
	}
}
