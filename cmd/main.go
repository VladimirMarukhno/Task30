package main

import (
	user "GolandProjects/30.1"
	"GolandProjects/30.1/pkg"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	mux := chi.NewRouter()
	srv := pkg.Service{make(map[int]*user.User)}
	mux.Post("/create", srv.Create)
	mux.Get("/get_all", srv.GetAll)
	mux.Post("/make_friends", srv.MakeFriends)
	mux.Delete("/delete_user", srv.DeleteUser)
	mux.Get("/get_friend", srv.GetFriend)
	mux.Put("/update_age", srv.UpdateAge)
	http.ListenAndServe("localhost:8084", mux)
}
