package cache

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
	"github.com/rajasoun/aws-hub/test"
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

func NewMockRedis(t *testing.T) (*miniredis.Miniredis, *Redis) {
	server := miniredis.RunT(t)
	client := NewRedisClient(server)
	client.Connect()
	return server, client
}

func pingSuccesPath(client *Redis, assert *assert.Assertions, tt struct {
	name    string
	want    string
	wantErr bool
}) {
	outputBuffer := test.SetLogOutputToBuffer()
	err := client.Ping()
	assert.NoError(err)
	gotLog := outputBuffer.String()
	assert.Contains(gotLog, tt.want, "Ping() got = %v , want = %v ", gotLog, tt.want)
}

func pingFailurePath(server *miniredis.Miniredis, client *Redis, assert *assert.Assertions, tt struct {
	name    string
	want    string
	wantErr bool
}) {
	outputBuffer := test.SetLogOutputToBuffer()
	server.SetError("Mock Error")
	err := client.Ping()
	assert.Error(err, "Err err = %v ", err)

	gotLog := outputBuffer.String()
	assert.Contains(gotLog, tt.want, "Ping() got = %v , want = %v ", gotLog, tt.want)
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
			switch {
			case tt.wantErr == false: //Happy Path
				pingSuccesPath(client, assert, tt)
			case tt.wantErr == true: //Edge Case
				pingFailurePath(server, client, assert, tt)
			}
		})
	}
}

func redisGetSetSucessPath(client *Redis, tt struct {
	name     string
	key      string
	value    interface{}
	wantType string
	wantMsg  string
	wantErr  bool
}, assert *assert.Assertions) {
	err := client.Set(tt.key, tt.value)
	assert.NoError(err, "Error err = %v ", err)
	got, foundKey := client.Get(tt.key)
	assert.True(foundKey, "Cache Set & Get Sequenced Failed for key = %v", tt.key)
	assert.Equal(tt.value, got, "Get () = %v, want = %v", got, tt.value)
	assert.Equal(tt.wantType, client.Type(), "Get () = %v, want = %v", client.Type(), tt.wantType)
}

func redisGetSetFailurePath(client *Redis, tt struct {
	name     string
	key      string
	value    interface{}
	wantType string
	wantMsg  string
	wantErr  bool
}, assert *assert.Assertions, server *miniredis.Miniredis) {
	//With Wrong Value
	err := client.Set(tt.key, tt.value)
	assert.Error(err, "Error err = %v ", err)
	//With InValid Key
	got, foundKey := client.Get("Invalid")
	assert.False(foundKey, "Cache Set & Get Should Fail with Injected Err = %v", tt.key)
	assert.Empty(got, "Get () = %v", got)
	//With Redis Error
	outputBuffer := test.SetLogOutputToBuffer()
	server.SetError("Mock Error")
	client.Set(tt.key, "dummy")
	gotErrLog := outputBuffer.String()
	assert.Contains(gotErrLog, tt.wantMsg, "Mock Set Failed")
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
		wantMsg  string
		wantErr  bool
	}{
		{
			name:     "Check Set Get",
			key:      "Key",
			value:    "Test",
			wantType: "RedisCache",
			wantMsg:  "Successfully connected to Redis",
			wantErr:  false,
		},
		{
			name:     "Check Set Get Error",
			key:      "Key",
			value:    map[string]interface{}{"foo": make(chan int)},
			wantType: "RedisCache",
			wantMsg:  "Mock Error",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch {
			case tt.wantErr == false: //Happy Path
				redisGetSetSucessPath(client, tt, assert)
			case tt.wantErr == true: //Edge Case
				redisGetSetFailurePath(client, tt, assert, server)
			}
		})
	}
}
