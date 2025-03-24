package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/andrewbatallones/api/auth"
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

	userJson, ok := payload["user"].(map[string]interface{})
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"errors\": [\"Please supply a user object\"]}")
		return
	}

	name, nameOk := userJson["name"].(string)
	email, emailOk := userJson["email"].(string)
	pass, passOk := userJson["password"].(string)

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

	createdUserJson, err := json.Marshal(user)
	if err != nil {
		fmt.Fprintf(os.Stderr, "issue marshalling user: %s", err)

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"errors\": [\"User was created, but there was an issue returning the user\"]}")
		return
	}

	fmt.Fprintf(w, "{\"user\": %s}", createdUserJson)
}

func UserShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "{\"errors\": [\"Unable to authenticate user\"]}")
		return
	}

	conn, ok := utils.Connection()
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "{\"errors\": [\"Unable to authenticate user\"]}")
		return
	}
	defer conn.Close()

	u, err := auth.ValidateUserJWT(conn, strings.Split(authHeader, " ")[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to authenticate user: %s", err)

		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "{\"errors\": [\"Unable to authenticate user\"]}")
		return
	}

	userId, err := strconv.Atoi(r.PathValue("user_id"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get user_id: %s", err)

		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "{\"errors\": [\"Unable to authenticate user\"]}")
		return
	}

	if *u.Id != int(userId) {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "{\"errors\": [\"Unable to authenticate user\"]}")
		return
	}

	userJson, err := json.Marshal(u)
	if err != nil {
		fmt.Fprintf(os.Stderr, "issue marshalling user: %s", err)

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"errors\": [\"Cannot get user\"]}")
		return
	}

	fmt.Fprintf(w, "{\"user\": %s}", userJson)
}
