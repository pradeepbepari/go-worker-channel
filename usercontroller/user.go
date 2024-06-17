package usercontroller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"github.com/pradeepbepari/jsonplaceholder/helper"
	"github.com/pradeepbepari/jsonplaceholder/model"
)

const (
	users    = "https://jsonplaceholder.typicode.com/users"
	post     = "https://jsonplaceholder.typicode.com/posts"
	comments = "https://jsonplaceholder.typicode.com/comments"
	albums   = "https://jsonplaceholder.typicode.com/albums"
	photos   = "https://jsonplaceholder.typicode.com/photos"
	todo     = "https://jsonplaceholder.typicode.com/todos"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var user []model.User
	rsp, err := helper.LoadJson(users)
	if err != nil {
		log.Panic(err)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.Unmarshal(rsp, &user); err != nil {
		http.Error(w, "failed to unmarshal data", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)

}
func PostUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	resp := []model.CombaineStruct{
		{
			Post: &[]model.Post{},
		},
	}
	var posts []model.Post
	rsp, err := helper.LoadJson(post)
	if err != nil {
		log.Panic(err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.Unmarshal(rsp, &posts); err != nil {
		http.Error(w, "failed to unmarshal data", http.StatusInternalServerError)
		return
	}
	resp = append(resp, model.CombaineStruct{Post: &posts})
	json.NewEncoder(w).Encode(resp)

}
func Albums(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var album []model.Albums
	rsp, err := helper.LoadJson(albums)
	if err != nil {
		log.Panic(err)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.Unmarshal(rsp, &album); err != nil {
		http.Error(w, "failed to unmarshal data", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(album)
}

func Photos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var photo []model.Photos
	rsp, err := helper.LoadJson(photos)
	if err != nil {
		log.Panic(err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.Unmarshal(rsp, &photo); err != nil {
		http.Error(w, "failed to unmarshal data", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(photo)
}
func TODO(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var todos []model.Todo
	rsp, err := helper.LoadJson(todo)
	if err != nil {
		log.Panic(err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.Unmarshal(rsp, &todos); err != nil {
		http.Error(w, "failed to unmarshal data", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todos)
}
func Comments(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var comm []model.Comments
	rsp, err := helper.LoadJson(comments)
	if err != nil {
		log.Panic(err)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.Unmarshal(rsp, &comm); err != nil {
		http.Error(w, "failed to unmarshal data", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(comm)
}
func CommentsPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(r)
	id := vars["postid"]
	postid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return
	}
	channel := make(chan []model.Comments, postid)
	defer close(channel)
	var wg sync.WaitGroup
	jobChannel := make(chan int, postid)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(jobChannel, channel, &wg)
	}
	for i := 1; i <= postid; i++ {
		jobChannel <- i
	}
	close(jobChannel)
	wg.Wait()
	var allcomments []model.Comments
	for i := 1; i <= postid; i++ {
		comments := <-channel
		if comments != nil {
			allcomments = append(allcomments, comments...)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(allcomments); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
func worker(jobChannel <-chan int, ch chan<- []model.Comments, wg *sync.WaitGroup) {
	defer wg.Done()
	for id := range jobChannel {
		postComments(id, ch)
	}
}
func postComments(id int, ch chan<- []model.Comments) {

	url := "https://jsonplaceholder.typicode.com/comments?postId=" + strconv.Itoa(id)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to fetch comments for postId %d: %v", id, err)
		ch <- nil
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body for postId %d: %v", id, err)
		ch <- nil
		return
	}
	var comment []model.Comments
	if err := json.Unmarshal(body, &comment); err != nil {
		return
	}
	ch <- comment
}
