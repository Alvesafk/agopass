package storage

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
)

func HashMasterKey(mk string) []byte {
	hash := sha256.Sum256([]byte(mk))
	return hash[:]
}

func Encrypt(plain_text string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", nil
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return "", err
	}
	
	cipher_text := gcm.Seal(nonce, nonce, []byte(plain_text), nil)

	return base64.StdEncoding.EncodeToString(cipher_text), nil
}

func Decrypt(encoded string, key []byte) (string, error) {
	cipher_text, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce_size := gcm.NonceSize()
	if len(cipher_text) < nonce_size {
		return "", errors.New("Cipher text is too short.")
	}

	nonce, cipher_text := cipher_text[:nonce_size], cipher_text[nonce_size:]

	plain_text, err := gcm.Open(nil, nonce, cipher_text, nil)
	if err != nil {
		return "", nil
	}

	return string(plain_text), nil
}
