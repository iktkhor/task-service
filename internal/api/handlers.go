package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/iktkhor/task-service/internal/domain"
	"github.com/iktkhor/task-service/internal/service"
	"github.com/iktkhor/task-service/internal/storage"
)



func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		postTask(w, r)
	case http.MethodGet:
		getTask(w, r)
	case http.MethodDelete:
		deleteTask(w, r)
	default:
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func postTask(w http.ResponseWriter, r *http.Request) {
		task := domain.Task{
			ID:        generateID(),
			Status:    http.StatusCreated,
			CreatedAt: time.Now(),
		}
		
		store.Set(task)
		go service.ProcessTask(task, store)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)

}

func getTask(w http.ResponseWriter, r *http.Request) {
	
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	
}

func TaskByIDHandler(store *storage.TaskStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/tasks/")
		switch r.Method {
		case http.MethodGet:
			if task, ok := store.Get(id); ok {
				resp := domain.TaskResponse{
					ID:        task.ID,
					Status:    string(task.Status),
					CreatedAt: task.CreatedAt.Format("02.01.2006 15:04:05"),
				}

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(resp)
				return
			}
			http.NotFound(w, r)
		case http.MethodDelete:
			if _, ok := store.Get(id); ok {
				store.Delete(id)
				w.WriteHeader(http.StatusNoContent)
				return
			}
			http.NotFound(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}
