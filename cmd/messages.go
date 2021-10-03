package cmd

const (
	DESC_ROOT      = "paper-vault is a cli tool to backup/reveal secret "
	DESC_ROOT_LONG = `paper-vault is a cli tool to backup/reveal secret. 
	
	Secret are encrypted with AES, then the cypher is encoding in Base64 and store to QR Code File. The QR File can be printed without having to hide it.`

	DESC_BACKUP  = "Backup secret to qrcode"
	DESC_REVEAL  = "Reveal secret from qrcode or string"
	DESC_VERSION = "Show the paper-vault version information"

	BACKUP_DONE             = "backup done ! Remember to always remove secret file"
	DESC_FLAG_INPUT_FILE    = "File that contains the secret to backup"
	DESC_FLAG_INPUT_QR_FILE = "QR file that contains the secret to reveal"
	DESC_FLAG_INPUT_DATA    = "Encrypt secret to reveal"

	PASSWORD_LABEL                     = "enter your strong password for AES"
	PASSWORD_LABEL_CONFIRM             = "confirm password "
	ERROR_PASSWORD_STRONG              = "a strong password is required"
	ERROR_PASSWORD_NOT_MATCH           = "your password and confirmation password do not match"
	ERROR_TOO_BIG_SECRET_SIZE          = "secret is too big for paper-vault. Requiere < 256 Octets"
	ERROR_MISSING_INPUT_FILE           = "missing secretFile, see paper-vault backup --help"
	ERROR_MISSING_INPUT_FILE_OR_STRING = "missing qrcode file or string, see paper-vault reveal --help"
)
