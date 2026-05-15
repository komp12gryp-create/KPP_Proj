package requests

import (
	"errors"
	"net/http"
	"strconv"
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

func ParseTaskFilter(r *http.Request) (domain.TaskFilter, error) {
	var f domain.TaskFilter

	if status := r.URL.Query().Get("status"); status != "" {
		s := domain.TaskStatus(status)
		f.Status = &s
	}

	if dl := r.URL.Query().Get("deadline"); dl != "" {
		if dl == "today" {
			now := time.Now()
			from := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
			to := from.Add(24*time.Hour - time.Second)
			f.DeadlineFrom = &from
			f.DeadlineTo = &to
		} else {
			ts, err := strconv.ParseInt(dl, 10, 64)
			if err != nil {
				return f, errors.New("invalid deadline parameter")
			}
			t := time.Unix(ts, 0)
			from := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
			to := from.Add(24*time.Hour - time.Second)
			f.DeadlineFrom = &from
			f.DeadlineTo = &to
		}
	}

	if v := r.URL.Query().Get("deadline_from"); v != "" {
		ts, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return f, errors.New("invalid deadline_from parameter")
		}
		t := time.Unix(ts, 0)
		f.DeadlineFrom = &t
	}
	if v := r.URL.Query().Get("deadline_to"); v != "" {
		ts, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return f, errors.New("invalid deadline_to parameter")
		}
		t := time.Unix(ts, 0)
		f.DeadlineTo = &t
	}

	if cd := r.URL.Query().Get("created_date"); cd != "" {
		if cd == "today" {
			now := time.Now()
			from := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
			to := from.Add(24*time.Hour - time.Second)
			f.CreatedFrom = &from
			f.CreatedTo = &to
		} else {
			ts, err := strconv.ParseInt(cd, 10, 64)
			if err != nil {
				return f, errors.New("invalid created_date parameter")
			}
			t := time.Unix(ts, 0)
			from := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
			to := from.Add(24*time.Hour - time.Second)
			f.CreatedFrom = &from
			f.CreatedTo = &to
		}
	}

	if v := r.URL.Query().Get("created_from"); v != "" {
		ts, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return f, errors.New("invalid created_from parameter")
		}
		t := time.Unix(ts, 0)
		f.CreatedFrom = &t
	}
	if v := r.URL.Query().Get("created_to"); v != "" {
		ts, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return f, errors.New("invalid created_to parameter")
		}
		t := time.Unix(ts, 0)
		f.CreatedTo = &t
	}

	if sortBy := r.URL.Query().Get("sort_by"); sortBy != "" {
		f.SortBy = domain.TaskSortField(sortBy)
	}
	if sortOrder := r.URL.Query().Get("sort_order"); sortOrder != "" {
		f.SortOrder = domain.SortOrder(sortOrder)
	}

	if page := r.URL.Query().Get("page"); page != "" {
		p, err := strconv.ParseUint(page, 10, 64)
		if err != nil {
			return f, errors.New("invalid page parameter")
		}
		countPerPage := uint64(10)
		if cpp := r.URL.Query().Get("count_per_page"); cpp != "" {
			countPerPage, err = strconv.ParseUint(cpp, 10, 64)
			if err != nil {
				return f, errors.New("invalid count_per_page parameter")
			}
		}
		f.Pagination = &domain.Pagination{
			Page:         p,
			CountPerPage: countPerPage,
		}
	}

	return f, nil
}
