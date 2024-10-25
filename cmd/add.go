package cmd

import "github.com/spf13/cobra"

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task",
	Long:  `Adds a task`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Add task
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}