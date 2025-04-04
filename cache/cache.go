package cache

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/andrewbatallones/api/utils"
)

// The cache object
type Cache struct {
	// Headers
	CacheControl  string `redis:"cache_control"`
	ContentType   string `redis:"content_type"`
	ContentLength string `redis:"content_length"`
	Body          string `redis:"body"`
}

func GetCache(path string) *Cache {
	rdb, ok := utils.RedisClient()
	if !ok {
		return nil
	}

	exists, err := rdb.Exists(context.Background(), fmt.Sprintf("cache:%s", path)).Result()
	if err != nil {
		log.Printf("unable to build cache: %s", err)
		return nil
	} else if exists == 0 {
		return nil
	}

	var cache Cache
	err = rdb.HGetAll(context.Background(), fmt.Sprintf("cache:%s", path)).Scan(&cache)
	if err != nil {
		log.Printf("unable to build cache: %s", err)
		return nil
	}

	log.Print("cache hit.")

	return &cache
}

func (c *Cache) SetCache(path string) error {
	rdb, ok := utils.RedisClient()
	if !ok {
		return errors.New("unable to connect to Redis")
	}

	c.CacheControl = "max-age=3600, public"

	err := rdb.HSet(context.Background(), fmt.Sprintf("cache:%s", path), c).Err()
	if err != nil {
		return err
	}

	log.Print("set cache")

	return rdb.Expire(context.Background(), fmt.Sprintf("cache:%s", path), 1*time.Hour).Err()
}
