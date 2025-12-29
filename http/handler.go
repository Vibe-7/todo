package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
	"todo/todo1"

	"github.com/gorilla/mux"
)

type HttpHandlers struct {
	todoList *todo1.List
}

func NewHttpHandlers(todoList *todo1.List) *HttpHandlers {
	return &HttpHandlers{
		todoList: todoList,
	}
}

func NewErrorDTO(err error) ErrorDTO {
	return ErrorDTO{
		Message: err.Error(),
		Time:    time.Now(),
	}
}

func (h *HttpHandlers) HandlersCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO TaskDTO
	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err := taskDTO.ValidateForCreate(); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
	}

	todoTask := todo1.NewTasks(taskDTO.Title, taskDTO.Description)
	if err := h.todoList.AddTaks(todoTask); err != nil {
		errDTO := NewErrorDTO(err)

		if errors.Is(err, todo1.ErrTaskAlreadyExists) {
			http.Error(w, errDTO.ToString(), http.StatusConflict)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}

	b, err := json.MarshalIndent(todoTask, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
}

func (h *HttpHandlers) HandlersGetTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	task, err := h.todoList.GetTask(title)
	if err != nil {
		errDto := NewErrorDTO(err)
		if errors.Is(err, todo1.ErrTaskNotFound) {
			http.Error(w, errDto.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDto.ToString(), http.StatusInternalServerError)
		}
		return
	}

	b, err := json.MarshalIndent(task, "", "   ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}

}

func (h *HttpHandlers) HandlersGetAllTask(w http.ResponseWriter, r *http.Request) {
	tasks := h.todoList.ListTasks()

	b, err := json.MarshalIndent(tasks, "", "		")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}

}

func (h *HttpHandlers) HandlersGetAllUncompletedTask(w http.ResponseWriter, r *http.Request) {
	uncomplitedTasks := h.todoList.ListUncomletedTasks()

	b, err := json.MarshalIndent(uncomplitedTasks, "", "		")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}

}

func (h *HttpHandlers) HandlerCompletedTask(w http.ResponseWriter, r *http.Request) {
	var completeDTO CompleteTaskDTO
	if err := json.NewDecoder(r.Body).Decode(&completeDTO); err != nil {
		errDTO := NewErrorDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	title := mux.Vars(r)["title"]

	var (
		chTask todo1.Task
		err    error
	)

	if completeDTO.Complete {
		chTask, err = h.todoList.CompleteTask(title)
	} else {
		chTask, err = h.todoList.UncompleteTasks(title)
	}

	if err != nil {
		errDTO := NewErrorDTO(err)
		if errors.Is(err, todo1.ErrTaskAlreadyExists) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}

	b, err := json.MarshalIndent(chTask, "", "		")
	if err != nil {
		log.Fatal("Ошибка фатальная")
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("failed write to http response")
		return
	}

}

func (h *HttpHandlers) HandlerDeleteTask(w http.ResponseWriter, r *http.Request) {

	title := mux.Vars(r)["title"]

	if err := h.todoList.DeleteTask(title); err != nil {
		errDTO := NewErrorDTO(err)
		if errors.Is(err, todo1.ErrTaskAlreadyExists) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}
}
