package domain

import (
	"net/http"
	"time"
)

type Task struct {
    ID         string     `json:"id"`
    Status     int        `json:"status"`
    CreatedAt  time.Time  `json:"created_at"`
    StartedAt  time.Time  `json:"started_at,omitempty"`
    FinishedAt time.Time  `json:"finished_at,omitempty"`
}

func (t * Task) Duration() string {
    var res string

    switch t.Status {
    case http.StatusCreated:
        res = "Task not started yet"
    case http.StatusProcessing:
        res = time.Since(t.StartedAt).String()
    case http.StatusOK:
        res = t.FinishedAt.Sub(t.StartedAt).String()
    default:
        res = "Duration isn't avaible"
    }

    return res
}