# paper-vault

``` 
paper-vault is a cli tool to backup/reveal secret. 
	
	Secret are encrypted with AES, then the cypher is encoding in Base64 and store to QR Code File. The QR File can be printed without having to hide it.

Usage:
  paper-vault [flags]
  paper-vault [command]

Available Commands:
  backup      Backup secret to qrcode
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  reveal      Reveal secret from qrcode or string
  version     Show the paper-vault version information

Flags:
  -h, --help   help for paper-vault

Use "paper-vault [command] --help" for more information about a command.
```