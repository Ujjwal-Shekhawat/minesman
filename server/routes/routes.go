package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Auth releated must be in auth.go
type Ts struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var resBody Ts

	if err := decoder.Decode(&resBody); err != nil {
		log.Fatal(err.Error())
	}

	resp := struct {
		Username string `json:"username"`
		Token    string `json:"token"`
	}{
		Username: "kamisama",
		Token:    "lmao_success_boi",
	}
	if resBody.Username == "kamisama" && resBody.Password == "kamisama" {
		json.NewEncoder(w).Encode(resp)
		return
	} else {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(struct{}{})
	}
}

func AuthConsole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Hitted Auth Endpoint")
	if r.Header.Get("auth-token") == "lmao_success_boi" {
		w.WriteHeader(200)
		resp := struct {
			Username string `json:"username"`
			Token    string `json:"token"`
			IsAuth   bool   `json:"isAuth"`
		}{
			Username: "kamisama",
			Token:    "lmao_success_boi",
			IsAuth:   true,
		}
		json.NewEncoder(w).Encode(resp)
		return
	} else {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(struct {
			IsAuth bool `json:"isAuth"`
		}{IsAuth: false})
	}
}
