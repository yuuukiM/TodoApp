package models

import (
	"TodoApp/db"
	"time"
)

var pTodos *db.Todos

type Todo struct {
	ID         int
	Title      string
	Content    string
	CreateTime time.Time
	LimitTime  time.Time
}

func CreateTodo(Title string, Content string) {
	todo := Todo{}
	pTodos = append(pTodos, todo)
}
