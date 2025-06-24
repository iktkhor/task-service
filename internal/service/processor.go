package service

import (
	"math/rand"
	"time"

	"github.com/iktkhor/task-service/internal/domain"
	"github.com/iktkhor/task-service/internal/storage"
)

func ProcessTask(t *domain.Task, store *storage.TaskStore) {
    t.Status = domain.StatusRunning
    t.StartedAt = time.Now()
    store.Set(t)

	// Генерация случайного времени выполнения от 180 до 300 секунд
	min := 180
	max := 300
	rs := rand.Intn(max-min+1) + min

    time.Sleep(time.Duration(rs) * time.Second)

    t.FinishedAt = time.Now()
    t.Duration = t.FinishedAt.Sub(t.StartedAt).String()
    t.Status = domain.StatusCompleted

    store.Set(t)
}