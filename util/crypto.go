package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"io"
)

const (
	defaultIterationCount = 100
	keyLength             = 32
)

func Encrypt(password string, salt, data []byte) (string, error) {
	if salt == nil {
		salt = make([]byte, 8)
		rand.Read(salt)
	}
	key := computeKey(password)
	block, err := aes.NewCipher(key)

	if err != nil {
		return "", err
	}
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(password string, cryptoText string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)
	key := computeKey(password)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext)
}

func computeKey(password string) []byte {
	if len(password) == keyLength {
		return []byte(password)
	}
	h := sha1.New()
	h.Write([]byte(password))
	shaHash := hex.EncodeToString(h.Sum(nil))
	password += shaHash[:(keyLength - len(password))]
	return []byte(password)
}
