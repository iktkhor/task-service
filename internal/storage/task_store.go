package storage

import (
    "sync"

    "github.com/iktkhor/task-service/internal/domain"
)

type TaskStore struct {
    mu    sync.Mutex
    tasks map[string]domain.Task
}

func New() *TaskStore {
    return &TaskStore{
        tasks: make(map[string]domain.Task),
    }
}

func (s *TaskStore) Get(id string) (domain.Task, bool) {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    t, ok := s.tasks[id]
    return t, ok
}

func (s *TaskStore) Set(t domain.Task) {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.tasks[t.ID] = t
}

func (s *TaskStore) Delete(id string) {
    s.mu.Lock()
    defer s.mu.Unlock()
    delete(s.tasks, id)
}