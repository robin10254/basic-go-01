package main

import (
	"fmt"
	"net/http"
)

var (
	task1 = "Learn new tech daily"
	task2 = "Increase my income"
	task3 = "Plan a Namibia tour"

	allTasks = []string{task1, task2, task3}
)

func main() {
	http.HandleFunc("/", welcomeTodos)
	http.HandleFunc("/show-tasks", showTasks)
	http.HandleFunc("/add-task", addTask)

	http.ListenAndServe(":9000", nil)
}

func welcomeTodos(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Welcome todos app!"
	fmt.Fprintln(writer, greeting)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Todos List: ")
	fmt.Fprintln(writer)
	for _, task := range allTasks {
		fmt.Fprintln(writer, task)
	}
}

func addTask(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Use POST Method", http.StatusMethodNotAllowed)
		return
	}

	task := request.FormValue("task")
	if task == "" {
		http.Error(writer, "Task is required", http.StatusBadRequest)
		return
	}

	allTasks = append(allTasks, task)

	showTasks(writer, request)
}
