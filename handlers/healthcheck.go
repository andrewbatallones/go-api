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
	DbConnection string `json:"db_connection"`
}

func NewHealthCheck() HealthCheck {
	return HealthCheck{"ok"}
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

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	hc := NewHealthCheck()

	hc.TestConnection()

	status, err := json.Marshal(hc)
	if err != nil {
		fmt.Printf("unable to create status JSON: %s", err)
		fmt.Fprint(w, "{\"status\": \"ALL ERROR\"}")
	}

	fmt.Fprintf(w, "{\"status\": \"%s\"}", status)
}
