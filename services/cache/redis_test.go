package cache

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func TestRedis_All(t *testing.T) {
	assert := assert.New(t)
	server := miniredis.RunT(t)
	type fields struct {
		Addr       string
		Expiration time.Duration
		client     *redis.Client
	}
	tests := []struct {
		name   string
		fields fields
		key    string
		want   string
	}{
		{
			name: "Check Redis Cache - With Valid Address",
			fields: fields{
				Addr:       server.Host() + ":" + server.Port(),
				Expiration: 30,
				client:     &redis.Client{},
			},
			key:  "Key",
			want: "Test",
		},
		{
			name: "Check Redis Cache - With Invalid Address",
			fields: fields{
				Addr:       server.Host() + ":" + server.Port(),
				Expiration: 30,
				client:     &redis.Client{},
			},
			key:  "Key",
			want: "Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				Addr:       tt.fields.Addr,
				Expiration: tt.fields.Expiration,
				client:     tt.fields.client,
			}
			r.Connect()
			r.Set(tt.key, tt.want)
			got, found := r.Get(tt.key)
			assert.True(found, "Cache Set & Get Sequenced Failed fro key = %v", tt.key)
			assert.Equal(tt.want, got, "Get () = %v, want = %v", got, tt.want)
		})
	}
}
