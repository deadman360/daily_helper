package task_application

import (
	"github.com/deadman360/daily_helper/internal/infrastructure"
	"github.com/spf13/cobra"
)

type task struct {
	repository infrastructure.IRepository
}

func NewTaskApplication(repository infrastructure.IRepository) ITaskApplication {
	return &task{
		repository: repository,
	}
}

type ITaskApplication interface {
	CreateTask(cmd *cobra.Command, args []string)
}
