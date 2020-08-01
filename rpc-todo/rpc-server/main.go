package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Todo struct {
	Title  string
	Status string
}

type Task int

var todoSlice []Todo

func (t *Task) MakeTodo(todo Todo, reply *Todo) error {
	todoSlice = append(todoSlice, todo)
	*reply = todo
	return nil
}

func (t *Task) GetTodo(title string, reply *Todo) error {
	var found Todo

	for _, v := range todoSlice {
		if v.Title == title {
			found = v
		}
	}

	*reply = found
	return nil
}

func (t *Task) DeleteTodo(todo Todo, reply *Todo) error {
	var deleted Todo
	for i, v := range todoSlice {
		if v.Title == todo.Title && v.Status == todo.Status {
			todoSlice = append(todoSlice[:i], todoSlice[i+1:]...)
			deleted = todo
			break
		}
	}
	*reply = deleted
	return nil
}

func (t *Task) GetAllTodo(todo string, reply *[]Todo) error {
	*reply = todoSlice
	return nil
}

func main() {
	task := new(Task)
	err := rpc.Register(task)

	if err != nil {
		log.Fatal("Format for service Task is not correct", err)
	}
	rpc.HandleHTTP()

	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("Listen err: ", err)
	}
	log.Printf("Serving RPC server on port: %d\n", 1234)

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving: ", err)
	}

}
