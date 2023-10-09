package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/anti-duhring/goingcrazy/schema"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	Client *redis.Client
}

var (
	ErrNil = errors.New("no matching record found in redis database")
)

func InitializeCache() (*Cache, error) {
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))

	if err != nil {
		return nil, fmt.Errorf("Error parsing redis url: %s", err)
	}

	opt.PoolSize = 100
	client := redis.NewClient(opt)

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("Error pinging redis: %s", err)
	}

	return &Cache{
		Client: client,
	}, nil
}

func (c *Cache) GetPerson(ctx context.Context, key string) (schema.Person, error) {
	var person schema.Person
	val, err := c.Client.Get(ctx, key).Result()

	if err != nil && err != redis.Nil {
		logger.Errorf("Error getting person from redis: %s", err)
		return person, ErrRedisGetPerson
	}

	if val != "" {

		err := json.Unmarshal([]byte(val), &person)

		if err == nil {
			return person, nil
		}

		logger.Errorf("error unmarshalling person from redis: %s", err)

	}

	return person, ErrRedisGetPerson
}
