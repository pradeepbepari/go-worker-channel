package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pradeepbepari/jsonplaceholder/usercontroller"
)

func RegisteredRoutes(router *mux.Router) {
	router.HandleFunc("/users", usercontroller.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/posts", usercontroller.PostUser).Methods(http.MethodPost)
	router.HandleFunc("/albums", usercontroller.Albums).Methods(http.MethodGet)
	router.HandleFunc("/photos", usercontroller.Photos).Methods(http.MethodGet)
	router.HandleFunc("/todos", usercontroller.TODO).Methods(http.MethodGet)
	router.HandleFunc("/comments", usercontroller.Comments).Methods(http.MethodGet)
	router.HandleFunc("/comments/{postid}", usercontroller.CommentsPost).Methods(http.MethodGet)
}
