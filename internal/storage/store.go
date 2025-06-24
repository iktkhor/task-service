package storage

import (
    "sync"

    "github.com/iktkhor/task-service/internal/domain"
)

type TaskStore struct {
    mu    sync.RWMutex
    tasks map[string]*domain.Task
}

func NewTaskStore() *TaskStore {
    return &TaskStore{
        tasks: make(map[string]*domain.Task),
    }
}

func (s *TaskStore) Get(id string) (*domain.Task, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    t, ok := s.tasks[id]
    return t, ok
}

func (s *TaskStore) Set(t *domain.Task) {
    s.mu.Lock()
    s.tasks[t.ID] = t
    s.mu.Unlock()
}

func (s *TaskStore) Delete(id string) {
    s.mu.Lock()
    delete(s.tasks, id)
    s.mu.Unlock()
}