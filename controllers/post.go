package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/vashuteotia123/instagram-backend-api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PostController struct {
	session *mgo.Session
}

func NewPostController(s *mgo.Session) *PostController {
	return &PostController{s}
}

func (post_controller PostController) GetPost(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}
	object_id := bson.ObjectIdHex(id)
	 post := models.Post{}

	 err := post_controller.session.DB("instagram-backend").C("posts").FindId(object_id).One(&post)
	 if err != nil {
		w.WriteHeader(404)
		return 
	 }
	 post_marshal, err := json.Marshal(post)
	 if err != nil { 
		fmt.Println(err)
	 }
	 w.Header().Set("Content-Type", "application/json")
	 w.WriteHeader(http.StatusOK)
	 fmt.Fprintf(w, "%s\n", post_marshal)
}

func (post_controller PostController) CreatePost(w http.ResponseWriter, r *http.Request , _ httprouter.Params){
	post := models.Post{}
	json.NewDecoder(r.Body).Decode(&post)
	post.Id = bson.NewObjectId()
	post.Time = time.Now()
	post_controller.session.DB("instagram-backend").C("posts").Insert(post)
	post_marshal, err := json.Marshal(post)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n",post_marshal)
}


func (post_controller PostController) GetUserPosts(w http.ResponseWriter,r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	posts := make([]models.Post, 0, 2)
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page == 0 {
		page = 1
	}
	err := post_controller.session.DB("instagram-backend").C("posts").Find(bson.M{"userid":id}).Skip((page-1)*2).Limit(2).All(&posts)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	post_marshal, err := json.Marshal(posts)
	 if err != nil {
		fmt.Println(err)
	 }
	 w.WriteHeader(http.StatusOK)
	 fmt.Fprintf(w, "%s\n", post_marshal)
}
