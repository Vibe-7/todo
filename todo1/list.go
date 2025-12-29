package todo1

import (
	"maps"
	"sync"
)

type List struct {
	tasks map[string]Task
	mtx   sync.RWMutex
}

func NewList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}

func (l *List) AddTaks(tash Task) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if _, ok := l.tasks[tash.Title]; ok {
		return ErrTaskAlreadyExists
	}
	l.tasks[tash.Title] = tash

	return nil

}

func (l *List) GetTask(tash string) (Task, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	task, ok := l.tasks[tash]
	if !ok {
		return Task{}, ErrTaskAlreadyExists
	}

	return task, nil
}

func (l *List) ListTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	tmp := make(map[string]Task, len(l.tasks))

	maps.Copy(tmp, l.tasks)
	return tmp
}

func (l *List) ListUncomletedTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	uncompletedTasks := make(map[string]Task)

	for title, task := range l.tasks {
		if !task.Completed {
			uncompletedTasks[title] = task
		}
	}
	return uncompletedTasks
}

func (l *List) CompleteTask(title string) (Task, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}
	task.Complete()

	l.tasks[title] = task

	return task, nil
}

func (l *List) UncompleteTasks(title string) (Task, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	tas, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}

	tas.Uncomplete()

	l.tasks[title] = tas

	return tas, nil

}

func (l *List) DeleteTask(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	_, ok := l.tasks[title]
	if !ok {
		return ErrTaskNotFound
	}

	delete(l.tasks, title)

	return nil
}
