package models

import "gopkg.in/mgo.v2/bson"

type Post struct {
	Id bson.ObjectId  `json:"id" bson:"_id"`
	Title string `json:"title" bson:"title"`
	Content string `json:"content" bson:"content"`
	Userid string `json:"userid" bson:"userid"`
}
