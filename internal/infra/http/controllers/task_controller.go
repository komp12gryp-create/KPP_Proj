package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
)

type TaskController struct {
	taskService app.TaskService
}

func NewTaskController(ts app.TaskService) TaskController {
	return TaskController{
		taskService: ts,
	}
}

func (c TaskController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)

		task, err := requests.Bind(r, requests.TaskRequest{}, domain.Task{})
		if err != nil {
			log.Printf("TaskController.Save(requests.Bind): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}

		task.UserId = user.Id
		task.Status = domain.NewTaskStatus
		task, err = c.taskService.Save(task)
		if err != nil {
			log.Printf("TaskController.Save(c.taskService.Save): %s", err)
			InternalServerError(w, err)
			return
		}

		taskDto := resources.TaskDto{}
		taskDto = taskDto.DomainToDto(task)

		Success(w, taskDto)
	}
}

func (c TaskController) FindList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)

		f, err := requests.ParseTaskFilter(r)
		if err != nil {
			log.Printf("TaskController.FindList(requests.ParseTaskFilter): %s", err)
			BadRequest(w, err)
			return
		}

		tasks, err := c.taskService.FindList(user.Id, f)
		if err != nil {
			log.Printf("TaskController.FindList(c.taskService.FindList): %s", err)
			InternalServerError(w, err)
			return
		}

		taskDto := resources.TaskDto{}
		tasksDto := taskDto.DomainToDtoCollection(tasks)

		Success(w, tasksDto)
	}
}

func (c TaskController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		task := r.Context().Value(TaskKey).(domain.Task)

		taskDto := resources.TaskDto{}
		taskDto = taskDto.DomainToDto(task)

		Success(w, taskDto)
	}
}

func (c TaskController) UpdateStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		task := r.Context().Value(TaskKey).(domain.Task)

		updTask, err := requests.Bind(r, requests.UpdateTaskStatusRequest{}, domain.Task{})
		if err != nil {
			log.Printf("TaskController.UpdateStatus(requests.Bind): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}

		task.Status = updTask.Status

		task, err = c.taskService.UpdateStatus(task)
		if err != nil {
			log.Printf("TaskController.UpdateStatus(c.taskService.UpdateStatus): %s", err)
			InternalServerError(w, err)
			return
		}

		taskDto := resources.TaskDto{}
		taskDto = taskDto.DomainToDto(task)

		Success(w, taskDto)
	}
}

func (c TaskController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		task := r.Context().Value(TaskKey).(domain.Task)

		updTask, err := requests.Bind(r, requests.TaskRequest{}, domain.Task{})
		if err != nil {
			log.Printf("TaskController.Save(requests.Bind): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}

		task.Title = updTask.Title
		task.Description = updTask.Description
		task.Deadline = updTask.Deadline

		task, err = c.taskService.Update(task)
		if err != nil {
			log.Printf("TaskController.Update(c.taskService.Update): %s", err)
			InternalServerError(w, err)
			return
		}

		taskDto := resources.TaskDto{}
		taskDto = taskDto.DomainToDto(task)

		Success(w, taskDto)
	}
}

func (c TaskController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		task := r.Context().Value(TaskKey).(domain.Task)

		err := c.taskService.Delete(task.Id)
		if err != nil {
			log.Printf("TaskController.Delete(c.taskService.Delete): %s", err)
			InternalServerError(w, err)
			return
		}

		noContent(w)
	}
}
