package cmd

import (
	"fmt"
	"github.com/kemalnw/cashflow/app/api"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var HttpCmd = &cobra.Command{
	Use:   "http",
	Short: "Run Http Server",
	Long:  "Run Http Server",
	RunE: func(cmd *cobra.Command, args []string) error {
		initHTTP()

		app := api.New(baseHandler)
		Chan := make(chan error)
		go func() {
			Chan <- app.Run()
		}()
		term := make(chan os.Signal, 1)
		signal.Notify(term, os.Interrupt, syscall.SIGTERM)

		select {
		case <-term:
			logrus.Infoln("signal terminated detected")
			return nil
		case err := <-Chan:
			return fmt.Errorf("service runtime error: %w", err)
		}
	},
}
