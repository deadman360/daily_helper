package usecases

import (
	"github.com/deadman360/daily_helper/internal/infrastructure"
	"github.com/spf13/cobra"
)

type task struct {
	repository infrastructure.IRepository
}

func NewTaskInterface(repository infrastructure.IRepository) ITask {
	return &task{
		repository: repository,
	}
}

type ITask interface {
	CreateTask(cmd *cobra.Command, args []string)
}
