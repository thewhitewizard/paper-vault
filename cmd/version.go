package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:               "version",
	Short:             DESC_VERSION,
	Long:              DESC_VERSION,
	Run:               versionFunc,
	Aliases:           []string{"v"},
	Example:           "paper-vault version",
	DisableAutoGenTag: true,
}

func versionFunc(cmd *cobra.Command, args []string) {
	fmt.Printf("paper-vault  version %v", Version)
}
