package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/andrewbatallones/api/utils"
)

type HealthCheck struct {
	DbConnection    string `json:"db_connection"`
	RedisConnection string `json:"redis_connection"`
}

func NewHealthCheck() HealthCheck {
	return HealthCheck{"ok", "ok"}
}

func (hc *HealthCheck) TestConnection() {
	conn, ok := utils.Connection()
	if !ok {
		hc.DbConnection = "database is unable to connect"
		return
	}
	defer conn.Close()

	var test string
	err := conn.QueryRow(context.Background(), "SELECT 'Testing'").Scan(&test)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		hc.DbConnection = "database is unable to connect"
	}
}

func (hc *HealthCheck) TestRedisConnection() {
	rdb, ok := utils.RedisClient()
	if !ok {
		hc.RedisConnection = "redis is unable to connect"
		return
	}

	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Redis is unable to query: %v\n", err)
		hc.RedisConnection = "redis is unable to connect"
	}
}

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	hc := NewHealthCheck()

	hc.TestConnection()
	hc.TestRedisConnection()

	status, err := json.Marshal(hc)
	if err != nil {
		fmt.Printf("unable to create status JSON: %s", err)
		fmt.Fprint(w, "{\"status\": \"ALL ERROR\"}")
	}

	fmt.Fprintf(w, "{\"status\": %s}", status)
}
