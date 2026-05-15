package domain

import "time"

type TaskFilter struct {
	Status       *TaskStatus
	DeadlineFrom *time.Time
	DeadlineTo   *time.Time
	CreatedFrom  *time.Time
	CreatedTo    *time.Time
	SortBy       TaskSortField
	SortOrder    SortOrder
	Pagination   *Pagination
}

type TaskSortField string

const (
	SortByDeadline    TaskSortField = "deadline"
	SortByCreatedDate TaskSortField = "created_date"
)

type SortOrder string

const (
	SortOrderAsc  SortOrder = "asc"
	SortOrderDesc SortOrder = "desc"
)
