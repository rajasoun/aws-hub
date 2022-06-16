package cache

import (
	"bytes"
	"log"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

const mockErr = "Mock Error"

func NewRedisClient(server *miniredis.Miniredis) *Redis {
	r := &Redis{
		Addr:       server.Host() + ":" + server.Port(),
		Expiration: 30,
		client:     &redis.Client{},
	}
	return r
}

func NewMockRedis(t *testing.T) (*miniredis.Miniredis, *Redis) {
	server := miniredis.RunT(t)
	client := NewRedisClient(server)
	client.Connect()
	return server, client
}

func TestConnectErr(t *testing.T) {
	assert := assert.New(t)
	server := miniredis.RunT(t)
	client := NewRedisClient(server)
	t.Run("Check Redis Connect Err", func(t *testing.T) {
		server.SetError(mockErr)
		client.Connect()
		err := client.Ping()
		assert.Error(err, "Err err = %v ", err)
	})
	defer server.Close()
	defer client.client.Close()
}

func TestRedisPing(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	// start Redis mock server
	server, client := NewMockRedis(t)
	defer server.Close()
	defer client.client.Close()

	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Check Ping",
			want:    "Successfully connected to Redis",
			wantErr: false,
		},
		{
			name:    "Check Ping Error",
			want:    "Cloudn't connect to Redis:",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				var outputBuffer bytes.Buffer
				log.SetOutput(&outputBuffer)
				server.SetError(mockErr)
				err := client.Ping()
				assert.Error(err, "Err err = %v ", err)
				gotLog := outputBuffer.String()
				assert.Contains(gotLog, tt.want, "Ping() got = %v , want = %v ", gotLog, tt.want)
				return
			}
			err := client.Ping()
			assert.NoError(err)
		})
	}
}

func TestRedisGetSet(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	// start Redis mock server
	server, client := NewMockRedis(t)
	defer server.Close()
	defer client.client.Close()
	tests := []struct {
		name     string
		key      string
		value    interface{}
		wantType string
		want     string
		wantErr  bool
	}{
		{
			name:     "Check Set Get",
			key:      "Key",
			value:    "Test",
			wantType: "RedisCache",
			want:     "Successfully connected to Redis",
			wantErr:  false,
		},
		{
			name:     "Check Set Get Error",
			key:      "Key",
			value:    map[string]interface{}{"foo": make(chan int)},
			wantType: "RedisCache",
			want:     mockErr,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				// With Wrong Value
				err := client.Set(tt.key, tt.value)
				assert.Error(err, "Error err = %v ", err)
				// With InValid Key
				got, foundKey := client.Get("Invalid")
				assert.False(foundKey, "Cache Set & Get Should Fail with Injected Err = %v", tt.key)
				assert.Empty(got, "Get () = %v", got)
				// With Redis Error
				server.SetError(mockErr)
				errMock := client.Set(tt.key, "dummy")
				assert.Error(errMock, "Simulated Err Failed err = %v ", errMock)
				return
			}
			err := client.Set(tt.key, tt.value)
			assert.NoError(err, "Error err = %v ", err)
			got, foundKey := client.Get(tt.key)
			assert.True(foundKey, "Cache Set & Get Sequenced Failed for key = %v", tt.key)
			assert.Equal(tt.value, got, "Get () = %v, want = %v", got, tt.value)
			assert.Equal(tt.wantType, client.Type(), "Get () = %v, want = %v", client.Type(), tt.wantType)
		})
	}
}
