package cmd

import (
	"github.com/alexleyoung/tasky/db"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes the database",
	Long:  `Initializes the database`,
	Run: func(cmd *cobra.Command, args []string) {
		db.InitDB()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}