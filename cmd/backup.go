package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"paper-vault/util"

	"github.com/spf13/cobra"
)

const (
	MIN_PASSWORD_LENGTH  = 8
	MIN_PASSWORD_ENTROPY = 50
	MAX_SECRET_SIZE      = 256
)

var (
	secretFile string

	backupCmd = &cobra.Command{
		Use:               "backup",
		Short:             DESC_BACKUP,
		Long:              DESC_BACKUP,
		RunE:              backupFunc,
		Example:           "paper-vault backup --input-file my_precious.txt ",
		DisableAutoGenTag: false,
	}
)

func backupFunc(cmd *cobra.Command, args []string) error {

	if secretFile != "" {
		password, err := util.PasswordPrompt(PASSWORD_LABEL)
		if err != nil {
			return err
		}
		confirmedPassword, err := util.PasswordPrompt(PASSWORD_LABEL_CONFIRM)
		if err != nil {
			return err
		}
		if confirmedPassword != password {
			return fmt.Errorf(ERROR_PASSWORD_NOT_MATCH)
		}
		if util.CheckPasswordRequierements(password) {
			if file, err := os.Stat(secretFile); err == nil {
				if file.Size() <= MAX_SECRET_SIZE {
					bf, err := ioutil.ReadFile(secretFile)
					if err != nil {
						return err
					} else {
						enc, err := util.Encrypt(password, bf)
						if err != nil {
							return err
						} else {
							err := util.GenerateQR(secretFile, enc)
							if err != nil {
								return err
							}
							fmt.Println(BACKUP_DONE)
						}
					}
				} else {
					return fmt.Errorf(ERROR_TOO_BIG_SECRET_SIZE)
				}
			} else {
				return err
			}
		} else {
			return fmt.Errorf(ERROR_PASSWORD_STRONG)
		}
	} else {
		return fmt.Errorf(ERROR_MISSING_INPUT_FILE)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(backupCmd)
	backupCmd.Flags().StringVarP(&secretFile, "input-file", "i", "", DESC_FLAG_INPUT_FILE)
}
