package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func Connection() (*pgxpool.Pool, bool) {
	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, false
	}

	return conn, true
}

func RedisClient() (*redis.Client, bool) {
	url := GetEnv("REDIS_URL", "redis://user:password@localhost:6379/0")
	opts, err := redis.ParseURL(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse URL: %v\n", err)
		return nil, false
	}

	return redis.NewClient(opts), true
}
