package task_commands

import (
	"github.com/deadman360/daily_helper/internal/infrastructure"
	"github.com/deadman360/daily_helper/internal/usecases/task_usecases"
	"github.com/spf13/cobra"
)

func CreateCommand(repository infrastructure.IRepository) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "task",
		Short: "Manage tasks",
	}

	rootCmd.Flags().String("name", "", "Task name (required)")
	rootCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(createTaskCommand(task_usecases.NewTaskInterface(repository)))
	return rootCmd
}
