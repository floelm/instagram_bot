package cache

import (
	"github.com/go-redis/redis"
	"time"
)

type UserCache struct {
	redis *redis.Client
}

func NewUserCache() UserCache {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return UserCache{
		redis: client,
	}
}

func (c *UserCache) Set(url string) {
	expireAt := time.Now().Add(time.Second * 10)

	c.redis.Set(url, expireAt.Unix(), time.Hour*1000)
}

func (c *UserCache) Delete(url string) {
	c.redis.Del(url)
}

func (c *UserCache) Get(url string) int64 {
	cmd := c.redis.Get(url)

	expiresAt, err := cmd.Int64()
	if err != nil {
		panic(err)
	}

	return expiresAt
}

func (c *UserCache) GetAllKeys() []string {
	cmd := c.redis.Keys("*") //TODO: this is generally a bad idea :)
	keysSlice, err := cmd.Result()
	if err != nil {
		panic(err)
	}

	return keysSlice
}
