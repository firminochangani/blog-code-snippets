package app

import (
	"sync"
)

type App struct {
	lock *sync.RWMutex
	data map[string]*Task
}

func NewApp() *App {
	return &App{
		lock: &sync.RWMutex{},
		data: make(map[string]*Task),
	}
}

func (s *App) GetAllTasks() ([]Task, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	result := make([]Task, len(s.data))

	idx := 0
	for _, task := range s.data {
		result[idx] = *task
		idx++
	}

	return result, nil
}

func (s *App) CreateTask(name, ownerId string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	return nil
}
