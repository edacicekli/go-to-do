package main

import (
	"testing"
	"time"
)

func TestRemoveTodo(t *testing.T) {

	newTodo := Todo{ID: "0", IsFinished: false, Title: "mocktodotitle", Body: "mocktodobody", DueDate: time.Now()}
	todos = append(todos, newTodo)

	if len(todos) != 1 {
		t.Fatalf("TestRemoveTodo not passed")
	}

	removeTodo("0")
	if len(todos) != 0 {
		t.Fatalf("TestRemoveTodo not passed")
	}

}
