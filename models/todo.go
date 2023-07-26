package models

type Todo struct {
	ID     string `json:"id" bson:"_id,omitempty"`
	Title  string `json:"title" bson:"title"`
	IsDone bool   `json:"isDone" bson:"isDone"`
}
