package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	ID   int
	Name string
}

type TaskManager struct {
	Tasks   []Task
	counter int
}

func (tm *TaskManager) AddTask(name string) {
	task := Task{ID: tm.counter, Name: name}
	tm.Tasks = append(tm.Tasks, task)
	tm.counter++
	fmt.Println("Task added successfully!")
}

func (tm *TaskManager) PrintTasks() {
	if len(tm.Tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("Tasks:")
	for _, task := range tm.Tasks {
		fmt.Printf("ID: %d, Name: %s\n", task.ID, task.Name)
	}
}

func main() {
	taskManager := TaskManager{}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter task name (or 'quit' to exit): ")
		scanner.Scan()
		input := scanner.Text()

		if input == "quit" {
			break
		}

		taskManager.AddTask(input)
		taskManager.PrintTasks()
	}
}
