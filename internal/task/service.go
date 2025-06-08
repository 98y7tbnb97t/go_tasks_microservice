package task

import (
	"errors"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func validateTask(task *Task) error {
	if task.UserID == 0 {
		return errors.New("UserID must not be 0")
	}
	return nil
}

func (s *Service) GetTasks() ([]Task, error) {
	return s.repo.GetTasksFromBd()
}

func (s *Service) CreateTask(task *Task) error {
	if err := validateTask(task); err != nil {
		return err
	}
	return s.repo.CreateTaskFromBd(task)
}

func (s *Service) CreateTaskForUser(userID uint, task *Task) error {
	task.UserID = userID
	if err := validateTask(task); err != nil {
		return err
	}
	return s.repo.CreateTaskFromBd(task)
}

func (s *Service) UpdateTask(id string, task *Task) error {
	if err := validateTask(task); err != nil {
		return err
	}
	return s.repo.UpdateTaskFromBd(id, task)
}

func (s *Service) PatchTask(id string, task *Task) error {
	if err := validateTask(task); err != nil {
		return err
	}
	return s.repo.PatchTaskFromBd(id, task)
}

func (s *Service) DeleteTask(id string) error {
	return s.repo.DeleteTaskFromBd(id)
}

func (s *Service) GetTaskByID(id string, task *Task) error {
	return s.repo.GetTaskByIDFromBd(id, task)
}

func (s *Service) GetTasksForUser(userID uint) ([]Task, error) {
	return s.repo.GetTasksForUserFromBd(userID)
}
