package models

type Todo struct {
	ID     string `json:"id" bson:"_id"`
	Title  string `json:"title" bson:"title"`
	IsDone string `json:"isDone" bson:"isDone"`
}
