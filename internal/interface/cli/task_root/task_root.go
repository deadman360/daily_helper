package task_root

import (
	"github.com/deadman360/daily_helper/internal/application/task_application"
	"github.com/deadman360/daily_helper/internal/infrastructure"
	"github.com/deadman360/daily_helper/internal/interface/cli/task_root/task_commands"
	"github.com/spf13/cobra"
)

func CreateCommand(repository infrastructure.IRepository) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "task",
		Short: "Manage tasks",
	}

	rootCmd.Flags().String("name", "", "Task name (required)")
	rootCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(task_commands.CreateTaskCommand(task_application.NewTaskApplication(repository)))
	return rootCmd
}
