package cmd

import (
	"fmt"
	"os"
	"paper-vault/util"

	"github.com/spf13/cobra"
)

var (
	qrFile      string
	encryptData string

	revealCmd = &cobra.Command{
		Use:               "reveal",
		Short:             DESC_REVEAL,
		Long:              DESC_REVEAL,
		RunE:              revealFunc,
		Example:           "paper-vault reveal -i qr.png",
		DisableAutoGenTag: false,
	}
)

func revealFunc(cmd *cobra.Command, args []string) error {

	if qrFile != "" || encryptData != "" {
		password, err := util.PasswordPrompt(PASSWORD_LABEL)
		if err != nil {
			return err
		}
		if encryptData != "" {
			fmt.Println(util.Decrypt(password, encryptData))
		} else {
			if _, err := os.Stat(qrFile); err == nil {
				encryptData, err = util.ReadDataFromQRFile(qrFile)
				if err != nil {
					return err
				}
				fmt.Println(util.Decrypt(password, encryptData))
			} else {
				return err
			}
		}
	} else {
		return fmt.Errorf(ERROR_MISSING_INPUT_FILE_OR_STRING)

	}
	return nil
}

func init() {
	rootCmd.AddCommand(revealCmd)
	revealCmd.Flags().StringVarP(&qrFile, "input-file", "i", "", DESC_FLAG_INPUT_QR_FILE)
	revealCmd.Flags().StringVarP(&encryptData, "input-data", "s", "", DESC_FLAG_INPUT_DATA)
}
