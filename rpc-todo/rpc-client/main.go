package main

import (
	"log"
	"net/rpc"
)

type Todo struct {
	Title  string
	Status string
}

func main() {

	var reply Todo

	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal("Connection Error: ", err)
	}

	finishApp := Todo{"Finish App", "Started"}
	makeDinner := Todo{"Make Dinner", "Not Started"}
	walkDog := Todo{"Walk the dog", "Not started"}

	client.Call("Task.MakeTodo", finishApp, &reply)
	client.Call("Task.MakeTodo", makeDinner, &reply)
	client.Call("Task.MakeTodo", walkDog, &reply)

	var allTodo []Todo
	client.Call("Task.GetAllTodo", "", &allTodo)

	log.Printf("All Todo's %+v\n", allTodo)

}
