package cmd

import (
	"github.com/spf13/cobra"
)

func newCliCmd() *cobra.Command {
	return &cobra.Command{
		Use: "sq",
		Short: "squaaat-api application",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}
}
