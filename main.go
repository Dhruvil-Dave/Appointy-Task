package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Post struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	username *username `json:"username"`
}

type username struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var Posts []Post

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Posts)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) 
	for _, item := range Posts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Post{})
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Post Post
	_ = json.NewDecoder(r.Body).Decode(&Post)
	Post.ID = strconv.Itoa(rand.Intn(100000000))
	Posts = append(Posts, Post)
	json.NewEncoder(w).Encode(Post)
}


func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Posts {
		if item.ID == params["id"] {
			Posts = append(Posts[:index], Posts[index+1:]...)
			var Post Post
			_ = json.NewDecoder(r.Body).Decode(&Post)
			Post.ID = params["id"]
			Posts = append(Posts, Post)
			json.NewEncoder(w).Encode(Post)
			return
		}
	}
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Posts {
		if item.ID == params["id"] {
			Posts = append(Posts[:index], Posts[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Posts)
}

func main() {
	
	r := mux.NewRouter()

	
	Posts = append(Posts, Post{ID: "1", Title: "Post One", username: &username{Firstname: "Clever", Lastname: "Harry"}})
	Posts = append(Posts, Post{ID: "2", Title: "Post Two", username: &username{Firstname: "Steve", Lastname: "Smith"}})

	// Route handles & endpoints
	r.HandleFunc("/Posts", getPosts).Methods("GET")
	r.HandleFunc("/Posts/{id}", getPost).Methods("GET")
	r.HandleFunc("/Posts", createPost).Methods("POST")
	r.HandleFunc("/Posts/{id}", updatePost).Methods("PUT")
	r.HandleFunc("/Posts/{id}", deletePost).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Request sample
// {
// 	"title":"Post Three",
// 	"username":{"firstname":"Harry","lastname":"Potter"}
// }
