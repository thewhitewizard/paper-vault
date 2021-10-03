# paper-vault


## 🤔 paper-vault ?

Petit utilitaire perso pour stocker des secrets dans un QR Code en ayant pris soin au préalable de chiffrer le secret avec AES et en encodant le chiffré obtenu en base64. L'utilitaire est capable de déchiffrer soir directement à partir du QR Code soit de son contenu. 
On peut normalement imprimer les QR Code sans devoir les cacher

Inspiration similaire : https://github.com/mtraver/qrbak


Usage
=====

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