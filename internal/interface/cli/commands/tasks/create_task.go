package interfaces

import (
	usecases "github.com/deadman360/daily_helper/internal/usecases/task"
	"github.com/spf13/cobra"
)

func createTaskCommand(task usecases.ITask) *cobra.Command {
	cmd := cobra.Command{
		Use:   "create",
		Short: "Create a new Task with timestamps",
		Run:   task.CreateTask,
	}
	cmd.Flags().StringP("name", "n", "", "Task name (required)")
	cmd.MarkFlagRequired("name")
	return &cmd
}
