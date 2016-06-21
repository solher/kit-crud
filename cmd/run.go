package cmd

import (
	"net/http"
	"os"
	"strconv"

	"github.com/go-kit/kit/log"
	"github.com/solher/kit-crud/app"
	"github.com/spf13/cobra"
)

const (
	Port = "app.port"
)

var (
	logger = log.NewLogfmtLogger(os.Stdout)
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the service",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger.Log("msg", "building")
		appHandler, err := app.Builder(vip, logger)
		if err != nil {
			return err
		}

		port := strconv.Itoa(vip.GetInt(Port))

		logger.Log("port", port, "msg", "listening")

		http.Handle("/", appHandler)
		return http.ListenAndServe(":"+port, appHandler)
	},
}

func init() {
	logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)

	rootCmd.AddCommand(runCmd)

	runCmd.PersistentFlags().IntP("port", "p", 3000, "listening port")
	vip.BindPFlag(Port, runCmd.PersistentFlags().Lookup("port"))
}
