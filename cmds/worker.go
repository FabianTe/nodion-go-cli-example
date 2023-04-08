package cmds

import (
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
	"time"
)

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "worker",
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("worker starting")
		time.Sleep(time.Second * 2)
		slog.Info("worker running")
		for {
			slog.Info("worker simulating work ...")
			time.Sleep(time.Second * 5)
		}
	},
}
