package main

import "time"

//Todo struct
type Todo struct {
	Name     string
	Complete bool
	Due      time.Time
}

//Todos struct
type Todos []Todo
