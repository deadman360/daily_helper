package task_application

import (
	"fmt"

	"github.com/deadman360/daily_helper/internal/domain"
	"github.com/spf13/cobra"
)

func (task *task_application) CreateTask(cmd *cobra.Command, args []string) {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		fmt.Println("Error reading flag:", err)
		return
	}
	entity := &domain.Task{
		Name: name,
	}
	task.repository.CreateTask(entity)
}
