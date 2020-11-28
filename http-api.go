package main

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")

	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", patchPost).Methods("PATCH")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

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


func updatePost(w http.ResponseWriter, r *http.Request){
// GET id of post from route parameters 
var idParam string = mux.Vars(r)["id"]
id, err := strconv.Atoi(idParam)
if err != nil {
	w.WriteHeader(400)
	w.Write([]byte("ID could not be converted to integar"))
	return
}

if id >= len(posts){
	w.WriteHeader(400)
	w.Write([]byte("No post found with the specified Id"))
	return
}

// get Value from JSON body 
var updatedPost Post
json.NewDecoder(r.Body).Decode(&updatedPost)

posts[id] = updatedPost

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(updatedPost)
}

func getPost(w http.ResponseWriter, r *http.Request){
	// Get id of post from route parameter 
	var idParam string = mux.Vars(r)["id"]
// convert id from string to intergar  
// the function returns integar and an error 
id, err := strconv.Atoi(idParam)
if err !=nil {
	// there was an error  
	w.WriteHeader(400)
	// converting from string to slice of byte 
	w.Write([]byte("ID couldnt be converted to integar"))

	return
}

if id >= len(posts){
	w.WriteHeader(404)
	w.Write([]byte("No Post found with specified Id"))
	return
}

post:= posts[id]

w.Header().Set("Content-Type", "application/JSON")
json.NewEncoder(w).Encode(post)
}

func getAllPost(w http.ResponseWriter, r *http.Request){
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

func patchPost(w http.ResponseWriter, r *http.Request){
// GET id of post from route parameters 
var idParam string = mux.Vars(r)["id"]
id, err := strconv.Atoi(idParam)
if err != nil {
	w.WriteHeader(400)
	w.Write([]byte("ID could not be converted to integar"))
	return
}

if id >= len(posts){
	w.WriteHeader(400)
	w.Write([]byte("No post found with the specified Id"))
	return
}

// get the current value 
post := &posts[id]
json.NewDecoder(r.Body).Decode(post)


w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(post)
}

func deletePost(w http.ResponseWriter, r *http.Request){
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integar"))
		return
	}
	
	if id >= len(posts){
		w.WriteHeader(400)
		w.Write([]byte("No post found with the specified Id"))
		return
	}

	// deleting a post from a slice 
	posts = append(posts[:id], posts[id+1:]...)
	w.WriteHeader(200)
}

