package service

import (
	"errors"
	"sync"
	"time"
	"todo/internal/model"
)

var ErrNotFound = errors.New("task not found")

type TaskService struct {
	mu    sync.RWMutex
	tasks map[string]*model.Task
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks: make(map[string]*model.Task),
	}
}

func (s *TaskService) Create(task *model.Task) {
	s.mu.Lock()
	defer s.mu.Unlock()
	task.CreateAt = time.Now()
	s.tasks[task.Title] = task
}

func (s *TaskService) GetService(title string) (*model.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, ok := s.tasks[title]
	if !ok {
		return nil, ErrNotFound
	}
	return task, nil
}

func (s *TaskService) List(completed *bool) []*model.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []*model.Task
	for _, t := range s.tasks {
		if completed == nil || t.Completed == *completed {
			result = append(result, t)
		}
	}
	return result
}

func (s *TaskService) Complete(title string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, ok := s.tasks[title]
	if !ok {
		return ErrNotFound
	}
	complete := time.Now()
	task.Completed = true
	task.CompletedAt = &complete
	return nil
}

func (s *TaskService) DeleteAll() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tasks = make(map[string]*model.Task)
}
