package api

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/iktkhor/task-service/internal/domain"
	"github.com/iktkhor/task-service/internal/service"
	"github.com/iktkhor/task-service/internal/storage"
)

func generateID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

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
	id := generateID()

}

func getTask(w http.ResponseWriter, r *http.Request) {
	
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	
}

func CreateTaskHandler(store *storage.TaskStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			return
		}
		task := &domain.Task{
			ID:        id,
			Status:    http.StatusCreated,
			CreatedAt: time.Now(),
		}
		store.Set(task)
		go service.ProcessTask(task, store)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}
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
