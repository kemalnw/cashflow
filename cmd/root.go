package cmd

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "cashflow-api",
	Short: "Cash Flow API",
	Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
}

func init() {
	// register command
	rootCmd.AddCommand(HttpCmd)

	// load environment variable
	if err := godotenv.Load(); err != nil {
		logrus.Fatalln("unable to load environment variable", err.Error())
	}
}

func Execute() error {
	cmd, _, err := rootCmd.Find(os.Args[1:])
	// Default run http command if args is not set
	if err == nil && cmd.Use == rootCmd.Use && cmd.Flags().Parse(os.Args[1:]) != pflag.ErrHelp {
		args := append([]string{"http"}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}

	return rootCmd.Execute()
}
