package handler

import (
	"Go-Handson/config"
	"Go-Handson/entity"
	"Go-Handson/service"
	"context"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . ListTasksService AddTaskService LoginService
type ListTasksService interface {
	ListTasks(ctx context.Context) (entity.Tasks, error)
}

type AddTaskService interface {
	AddTask(ctx context.Context, title string) (*entity.Task, error)
}

type RegisterUserService interface {
	RegisterUser(ctx context.Context, name, password, role string) (*entity.User, error)
}
type LoginService interface {
	Login(ctx context.Context, name, pw string) (string, error)
}

type RegistrationService interface {
	GenerateUserID() (string, error)
	CreateUserProfile(userID string, cfg *config.Config) (service.UserProfile, error)
}
