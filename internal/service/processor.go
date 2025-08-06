package service

import (
	"math/rand"
	"time"
	"net/http"
	

	"github.com/iktkhor/task-service/internal/domain"
	"github.com/iktkhor/task-service/internal/storage"
)

func GenRandSleep() int {
	// Генерация случайного времени выполнения от 180 до 300 секунд
	min := 180
	max := 300
	rs := rand.Intn(max-min+1) + min
	return rs
}

func ProcessTask(t domain.Task, ts *storage.TaskStore) {
    t.Status = http.StatusProcessing
    t.StartedAt = time.Now()
    ts.Set(t)

    time.Sleep(time.Duration(GenRandSleep()) * time.Second)

    t.FinishedAt = time.Now()
    t.Status = http.StatusOK

    ts.Set(t)
}