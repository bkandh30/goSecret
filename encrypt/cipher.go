package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
)

func Encrypt(key, plaintext string) (string, error) {
	block, err := newCipherBlock(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return hex.EncodeToString(cipherText), nil
}

func Decrypt(key, cipherHex string) (string, error) {
	block, err := newCipherBlock(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	cipherText, err := hex.DecodeString(cipherHex)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return "", errors.New("encrypt: cipherText too short")
	}

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func newCipherBlock(key string) (cipher.Block, error) {
	hash := sha256.Sum256([]byte(key))
	return aes.NewCipher(hash[:])
}
