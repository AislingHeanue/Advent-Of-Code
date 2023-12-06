package day6

import "github.com/spf13/cobra"

func AddCommandsTo(root *cobra.Command) {
	root.AddCommand(aCommand())
	root.AddCommand(bCommand())
}
