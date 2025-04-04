package cache_test

import (
	"context"
	"testing"

	"github.com/andrewbatallones/api/cache"
	"github.com/andrewbatallones/api/utils"
)

func deleteKey() {
	rdb, _ := utils.RedisClient()

	rdb.Del(context.Background(), "cache:/test")
}

func TestGetSetCach(t *testing.T) {
	deleteKey()

	body := "{\"test\": \"check\"}"
	setter := cache.Cache{
		CacheControl:  "",
		ContentType:   "application/json",
		ContentLength: "10",
		Body:          body,
	}

	err := setter.SetCache("/test")
	if err != nil {
		t.Errorf("failed to set cache: %s", err)
	}

	getter := cache.GetCache("/test")
	if getter == nil {
		t.Errorf("error retrieving cache")
	}

	if getter != nil && setter.Body != getter.Body {
		t.Errorf("the setter and getter did not match, expected %s, got %s", setter.Body, getter.Body)
	}
}

func TestNotFoundCache(t *testing.T) {
	deleteKey()

	getter := cache.GetCache("/test")
	if getter != nil {
		t.Errorf("cache found, expected nil, got %s", getter)
	}
}
