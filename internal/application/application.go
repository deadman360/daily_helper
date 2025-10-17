package application

import (
	"github.com/deadman360/daily_helper/internal/application/task_application"
	"github.com/deadman360/daily_helper/internal/infrastructure"
)

type application struct {
	task_application.ITaskApplication
}

type IApplication interface {
	task_application.ITaskApplication
}

func NewRepository(repository infrastructure.IRepository) IApplication {
	task_application := task_application.NewTaskApplication(repository)

	return &application{
		ITaskApplication: task_application,
	}
}
