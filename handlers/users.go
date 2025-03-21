package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/andrewbatallones/api/models"
	"github.com/andrewbatallones/api/utils"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var payload map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot parse user: %s", err)

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"errors\": [\"Unable to create user\"]}")
		return
	}

	user_json, ok := payload["user"].(map[string]interface{})
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"errors\": [\"Please supply a user object\"]}")
		return
	}

	name, nameOk := user_json["name"].(string)
	email, emailOk := user_json["email"].(string)
	pass, passOk := user_json["password"].(string)

	if !nameOk || !emailOk || !passOk {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"errors\": [\"Missing name, email, or password in user\"]}")
		return
	}

	user := models.User{
		Name:  name,
		Email: email,
	}
	user.SetPassword(pass)

	conn, ok := utils.Connection()
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"errors\": [\"Unable to get product\"]}")
		return
	}
	defer conn.Close()

	err = user.Create(conn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot create user: %s", err)

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"errors\": [\"Unable to create user\"]}")
		return
	}

	created_user_json, err := json.Marshal(user)
	if err != nil {
		fmt.Fprintf(os.Stderr, "issue marshalling user: %s", err)

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"errors\": [\"User was created, but there was an issue returning the user\"]}")
		return
	}

	fmt.Fprintf(w, "{\"user\": %s}", created_user_json)
}
