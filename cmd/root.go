package cmd

import (
	"log"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var (
	Version = "0.0.1"

	rootCmd = &cobra.Command{
		Use:   "paper-vault",
		Short: DESC_ROOT,
		Long:  DESC_ROOT_LONG,
		PreRun: func(rootCmd *cobra.Command, args []string) {
		},
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(aurora.Red(err))
		os.Exit(1)
	}
}
