package cmd

import (
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear done tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		return service.ClearDone()
	},
}
