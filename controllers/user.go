package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/julienschmidt/httprouter"
	"github.com/vashuteotia123/instagram-backend-api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (user_controller UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectIdHex(id)
	 user := models.User{}

	 err := user_controller.session.DB("instagram-backend").C("users").FindId(oid).One(&user)
	 if err != nil {
		w.WriteHeader(404)
		return 
	 }
	 user_marshal, err := json.Marshal(user)
	 if err != nil { 
		fmt.Println(err)
	 }
	 w.Header().Set("Content-Type", "application/json")
	 w.WriteHeader(http.StatusOK)
	 fmt.Fprintf(w, "%s\n", user_marshal)
}


func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}
// HashCompare function can also be added here to create a login session for a user with email and password


func (user_controller UserController) CreateUser(w http.ResponseWriter, r *http.Request , _ httprouter.Params){
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	user.Id = bson.NewObjectId()
	user.Password, _ = HashPassword(user.Password)
	println(user.Password)
	user_controller.session.DB("instagram-backend").C("users").Insert(user)
	user_marshal, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n",user_marshal)
}


