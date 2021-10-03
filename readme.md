# paper-vault


## 🤔 paper-vault à quoi ça sert ?

Petit utilitaire perso pour stocker des secrets dans un QR Code en ayant pris soin au préalable de chiffrer le secret avec AES et en encodant le chiffré obtenu en base64. L'utilitaire est capable de déchiffrer soir directement à partir du QR Code soit de son contenu. 

L'utilitaire peut etre utile notamment pour stocker les mots des wallets crypto si on ne souhaite pas passer par un soft.
On peut normalement imprimer les QR Code sans devoir les cacher.


## AVERTISSEMENT / WARNING
⚠️ Utilisation à vos risques et périls. 

⚠️ Use at your own risk



## Usage


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

``` 
 ╭─frederic@frederic in repo: paper-vault on  master [!?] via  v1.17.1 took 7s
 ╰─λ echo  "mon super secret" > secret.txt

 ╭─frederic@frederic in repo: paper-vault on  master [!?] via  v1.17.1 took 2ms
 ╰─λ ./paper-vault backup -i secret.txt
enter your strong password for AES: *************
confirm password : *************
backup done ! Remember to always remove secret file

 ╭─frederic@frederic in repo: paper-vault on  master [!?] via  v1.17.1 took 13s
 ╰─λ ./paper-vault reveal -i secret.txt.jpeg
enter your strong password for AES: *************
mon super secret

```