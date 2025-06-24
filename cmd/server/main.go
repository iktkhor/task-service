package main

import (
    "fmt"
    "net/http"

    "github.com/iktkhor/task-service/internal/api"
    "github.com/iktkhor/task-service/internal/storage"
)

func main() {
    store := storage.NewTaskStore()
    http.HandleFunc("/tasks", api.CreateTaskHandler(store))
    http.HandleFunc("/tasks/", api.TaskByIDHandler(store))

    fmt.Println("Server running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}