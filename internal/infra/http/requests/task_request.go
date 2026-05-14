package requests

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type TaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Deadline    *int64 `json:"deadline"`
}

func (r TaskRequest) ToDomainModel() (interface{}, error) {
	var timeUnix int64
	if r.Deadline != nil {
		timeUnix = *r.Deadline
	}
	var deadline *time.Time
	if timeUnix != 0 {
		dl := time.Unix(timeUnix, 0)
		deadline = &dl
	}
	return domain.Task{
		Title:       r.Title,
		Description: r.Description,
		Deadline:    deadline,
	}, nil
}

type UpdateTaskStatusRequest struct {
	Status domain.TaskStatus `json:"status" validate:"required"`
}

func (r UpdateTaskStatusRequest) ToDomainModel() (interface{}, error) {
	return domain.Task{
		Status: r.Status,
	}, nil
}
