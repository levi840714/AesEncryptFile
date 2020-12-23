package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"sync"
)

type AesKey struct {
	key       []byte
	syncMutex sync.Mutex
}

var commonKey = &AesKey{
	key: []byte(""), // either 16, 24, or 32 bytes to select
}

// encrypt plain text
func AesCFBEncrypt(plainText []byte) ([]byte, error) {
	block, err := aes.NewCipher(commonKey.key)
	if err != nil {
		return nil, err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	cipher.NewCFBEncrypter(block, iv).XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return cipherText, nil
}

// decrypt text
func AesCFBDecrypt(decryptText []byte) ([]byte, error) {
	cipherText := decryptText

	block, err := aes.NewCipher(commonKey.key)
	if err != nil {
		return nil, err
	}

	if len(cipherText) < aes.BlockSize {
		return nil, errors.New("cipher text too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	cipher.NewCFBDecrypter(block, iv).XORKeyStream(cipherText, cipherText)

	return cipherText, nil
}
