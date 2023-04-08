package cmds

import (
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
	"net/http"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "HTTP API",
	Run: func(cmd *cobra.Command, args []string) {
		addr := ":8090"
		slog.Info("api listening", slog.String("addr", addr))
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			panic(err)
		}
	},
}
