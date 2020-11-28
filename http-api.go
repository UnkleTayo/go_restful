package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// // Item is a struct that groups all necessaty field into a single unit
// type Item struct {
// 	// This will not be exported if it is in lowercase
// 	// json tag converts the json to a lowecase key
// 	Data 			string `json:"data"`
// 	OtherData int 		`json"otherData"`
// }
// type Item struct {
// 	// This will not be exported if it is in lowercase
// 	// json tag converts the json to a lowecase key
// 	Data 			string `json:"data"`
// 	OtherData int 		`json"otherData"`
// }

// User is a struct that represnt a user in our App
type User struct {
	Fullname  string 	`json:"fullName"`
	Username  string 	`json:"userName"`
	Email 	  string 	`json:"email"`
}


// Post is a struct representing a single post
type Post struct {
	Title  string 	`json:"title"`
	Body	 string 	`json:"body"`
	Author User 	`json:"author"`
}

var posts []Post = []Post{}


func main() {
	// Using the Gorilla mux package 
	router := mux.NewRouter()
	// router.HandleFunc("/test", test)
	router.HandleFunc("/posts", addNewPost).Methods("POST")
	router.HandleFunc("/posts", getAllPost).Methods("GET")

	// creating a route
	http.ListenAndServe(":5000", router)
}

// func test(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type", "application/json")

// 	// New encoder takes in a io.writer
// 	json.NewEncoder(w).Encode(struct {
// 		ID string
// 	}{"12434"})
// }

func getAllPost(w http.ResponseWriter, r * http.Request){
	w.Header().Set("Content-Type", "application/JSON")
	json.NewEncoder(w).Encode(posts)
}

// Route handler
func addNewPost(w http.ResponseWriter, r *http.Request){
	// Gorilla mux vars returns a map with key of string and value of string  
	// routeVariable := mux.Vars(r)["item"] //used to get var from url
	// get Item value from json body 

	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)
	
	posts = append(posts, newPost)


	w.Header().Set("Content-Type", "application/JSON")
	json.NewEncoder(w).Encode(posts)
}

