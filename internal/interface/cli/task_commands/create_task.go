package task_commands

import (
	"github.com/deadman360/daily_helper/internal/application/task_application"
	"github.com/spf13/cobra"
)

func createTaskCommand(task task_application.ITaskApplication) *cobra.Command {
	cmd := cobra.Command{
		Use:   "create",
		Short: "Create a new Task with timestamps",
		Run:   task.CreateTask,
	}
	cmd.Flags().StringP("name", "n", "", "Task name (required)")
	cmd.MarkFlagRequired("name")
	return &cmd
}
