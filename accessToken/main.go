package main

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strings"
)

var pool = newPool()
var client redis.Conn
var users = map[string]string{
	"lester": "123456",
	"claire": "654321",
}

func main(){
	client = pool.Get()
	defer client.Close()

	http.HandleFunc("/signin", SignIn)
	http.HandleFunc("/welcome", Welcome)
	log.Fatal(http.ListenAndServe(":8787", nil))
}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle: 80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}


type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Token struct {
	Token string `json:"token"`
	Expires string `json:"expires"`
}

type WelcomeMessage struct {
	Msg string `json:"message"`
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[creds.Username]
	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sessionToken := uuid.NewString()
	fmt.Printf("userName: %v \n", creds.Username)
	_, err = client.Do("SETEX", sessionToken, "120", creds.Username)
	if err != nil {
		fmt.Printf("err: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token := &Token{
		Token:   sessionToken,
		Expires: "120",
	}
	b, err := json.Marshal(token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	splitToken := strings.Split(auth, "Bearer ")

	if len(splitToken) < 2 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token := splitToken[1]

	res, err := client.Do("GET", token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if res == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	welcomeMsg := &WelcomeMessage{
		Msg: "Hello, there.",
	}
	msg, err := json.Marshal(welcomeMsg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/msg;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(msg)

}
