package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/andrewbatallones/api/models"
	"github.com/andrewbatallones/api/utils"
)

func Sessions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var payload map[string]string
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot parse session: %s", err)

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"errors\": [\"Unable to log user in\"]}")
		return
	}

	conn, ok := utils.Connection()
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"errors\": [\"Unable to log user in\"]}")
		return
	}

	defer conn.Close()
	u, err := models.FindByUser(conn, map[string]string{"email": payload["email"]})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error finding user: %s", err)

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"errors\": [\"Unable to log user in\"]}")
		return
	}

	if u == nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "{\"errors\": [\"User email or password is incorrect\"]}")
		return
	}

	if !u.CheckPassword(payload["password"]) {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "{\"errors\": [\"User email or password is incorrect\"]}")
		return
	}

	user_json, err := json.Marshal(u)
	if err != nil {
		fmt.Fprintf(os.Stderr, "issue marshalling user: %s", err)

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"errors\": [\"Unable to log user in\"]}")
		return
	}

	// Generate Token
	token, err := u.GenerateJWT()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error finding user: %s", err)

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"errors\": [\"Unable to log user in\"]}")
		return
	}

	fmt.Fprintf(w, "{\"user\": %s, \"token\" %s}", user_json, token)
}
