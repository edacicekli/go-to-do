package main
import "testing"

func TestRemoveTodo(t *testing.T){ 
	removeTodo("0")
	if len(todos)!=3{
		t.Fatalf("TestRemoveTodo not passed")
	}
}