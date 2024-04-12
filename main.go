package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

var Users map[uuid.UUID]UserData = make(map[uuid.UUID]UserData)

type UserData struct {
	Name             string `json:"name"`
	Email            string `json:"email"`
	SubscriptionType int    `json:"subscriptionType"`
	IsAdmin          bool   `json:"isAdmin"`
	MetaData         struct {
		Bitrate  string `json:"bitrate"`
		Provider string `json:"provider"`
		Region   string `json:"region"`
	} `json:"metadata"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body := r.Body

	var ud UserData
    err := json.NewDecoder(body).Decode(&ud)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}
    
    uuid, err := uuid.NewRandom()
    if err != nil {
		fmt.Printf("%v", err.Error())
        return
    }

    Users[uuid] = ud

    w.Header().Set("id", uuid.String())
    w.WriteHeader(http.StatusCreated)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    userId := uuid.MustParse(r.Header.Get("id"))

    
    l1 := len(Users) 
    delete(Users, userId)
    l2 := len(Users)

    if l1 != l2 {
        w.WriteHeader(http.StatusAccepted)
    } else {
        w.WriteHeader(http.StatusInternalServerError)
    }
}

func GetUser(w http.ResponseWriter, r *http.Request) {
    userId := uuid.MustParse(r.Header.Get("id"))

    if ud, ok := Users[userId]; ok {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(ud)
    } else {
        w.WriteHeader(http.StatusNotFound)
    }
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	body := r.Body

    fmt.Printf("%+v\n", body)
	var ud UserData
    err := json.NewDecoder(body).Decode(&ud)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}
    
    userId := uuid.MustParse(r.Header.Get("id"))

    Users[userId] = ud

    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("id", userId.String())
    json.NewEncoder(w).Encode(ud)
}

func main() {
	http.HandleFunc("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetUser(w, r)
		case http.MethodPost:
			CreateUser(w, r)
		case http.MethodDelete:
			DeleteUser(w, r)
		case http.MethodPatch:
			UpdateUser(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("HTTP SERVER STARTED")

	log.Fatal(http.ListenAndServe(":9091", nil))
}
