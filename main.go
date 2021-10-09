package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/vashuteotia123/instagram-backend-api/controllers"
	"gopkg.in/mgo.v2"
)

/*
Submitted in last 5 minutes as pushing my code to public repo sooner than this might result
in code duplication.
Made by: Vishal Teotia
*/
func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	pc := controllers.NewPostController(getSession())
	r.GET("/users/:id", uc.GetUser)
	r.POST("/users", uc.CreateUser)
	r.GET("/post/:id",pc.GetPost)
	r.POST("/post",pc.CreatePost)
	r.GET("/posts/users/:id", pc.GetUserPosts)
	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
