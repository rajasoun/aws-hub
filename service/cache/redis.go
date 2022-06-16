package cache

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type Redis struct {
	Addr       string
	Expiration time.Duration
	client     *redis.Client
}

func (r *Redis) Type() string {
	return "RedisCache"
}

func (r *Redis) Connect() {
	r.client = redis.NewClient(&redis.Options{
		Addr: r.Addr,
		DB:   0,
	})
	err := r.Ping()
	if err != nil {
		log.Printf("Err Ping() = %v", err)
	}
}

func (r *Redis) Ping() error {
	_, err := r.client.Ping().Result()
	if err != nil {
		log.Println("Cloudn't connect to Redis:", err)
		return err
	}
	log.Println("Successfully connected to Redis")
	return nil
}

func (r *Redis) Get(key string) (interface{}, bool) {
	val, err := r.client.Get(key).Result()
	if errors.Is(err, redis.Nil) {
		return val, false
	}
	// Set ensures no invalid value gets into cache
	// so ignoring err check for json.Unmarshal
	var data interface{}
	json.Unmarshal([]byte(val), &data)
	return data, true
}

func (r *Redis) Set(key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		log.Println("Error in Marshaling JSON ", err)
		return err
	}
	err = r.client.Set(key, data, r.Expiration*time.Minute).Err()
	if err != nil {
		log.Println("Error in Cache Set", err)
	}
	return err
}
