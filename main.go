package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type Resp struct {
	Error bool   `json:"error"`
	Token string `json:"token"`
}

var access = Auth{Username: "c137@onecause.com", Password: "#th@nH@rm#y#r!$100%D0p#"}

func getToken() string {
	now := time.Now()
	hour := strconv.Itoa(now.Hour())
	min := strconv.Itoa(now.Minute())
	return "token-" + hour + min
}

func setCors(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	setCors(w)
	var auth Auth
	json.NewDecoder(r.Body).Decode(&auth)
	var resp = Resp{Error: true}
	if auth.Username == access.Username &&
		auth.Password == access.Password &&
		auth.Token == getToken() {
		resp.Error = false
	}
	json.NewEncoder(w).Encode(resp)
}

func corsHandle(w http.ResponseWriter, r *http.Request) {
	setCors(w)
}

func tokenHandler(w http.ResponseWriter, r *http.Request) {
	setCors(w)
	var resp = Resp{Error: false, Token: getToken()}
	json.NewEncoder(w).Encode(resp)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", tokenHandler).Methods("GET")
	r.HandleFunc("/", loginHandler).Methods("POST")
	r.HandleFunc("/", corsHandle).Methods("OPTIONS")
	log.Fatal(http.ListenAndServe(":3001", r))
}
