package domain

import "time"

type TaskStatus string

const (
    StatusPending   TaskStatus = "pending"
    StatusRunning   TaskStatus = "running"
    StatusCompleted TaskStatus = "completed"
)

type Task struct {
    ID         string     `json:"id"`
    Status     TaskStatus `json:"status"`
    CreatedAt  time.Time  `json:"created_at"`
    StartedAt  time.Time  `json:"started_at,omitempty"`
    FinishedAt time.Time  `json:"finished_at,omitempty"`
    Duration   string     `json:"duration,omitempty"`
}

type TaskResponse struct {
    ID          string `json:"id"`
    Status      string `json:"status"`
    CreatedAt   string `json:"created_at"`
    CurDuration string `json:"duration,omitempty"`
}