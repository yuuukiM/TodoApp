package db

import (
	"time"
)

type Todo struct {
	ID         int
	Title      string
	Content    string
	CreateTime time.Time
	LimitTime  time.Time
}

var Todos = make([]Todo, 1)
