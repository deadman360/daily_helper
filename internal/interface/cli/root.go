package cli

import (
	"fmt"

	"github.com/deadman360/daily_helper/internal/infrastructure"
	"github.com/deadman360/daily_helper/internal/interface/cli/task_root"
	"github.com/spf13/cobra"
)

func Execute(repository infrastructure.IRepository) {
	rootCmd := &cobra.Command{
		Use:   "dh",
		Short: "Daily Helper is a cli tool to help you to manage your daily tasks with notes.",
	}

	rootCmd.AddCommand(task_root.CreateCommand(repository))
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
