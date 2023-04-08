package cmds

import "github.com/spf13/cobra"

var RootCmd *cobra.Command

func init() {
	RootCmd = &cobra.Command{
		Use:   "nodion-example",
		Short: "nodion-example",
	}

	RootCmd.AddCommand(apiCmd)
	RootCmd.AddCommand(workerCmd)
}
