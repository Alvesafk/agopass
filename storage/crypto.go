/*
crypto.go file has the hash, encrypting and decrypting logic of the storage package.
*/

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

// HashMasterKey will hash the master key with sha256, it also add a salt to it.
func HashMasterKey(mk string) []byte {
	hash := sha256.Sum256([]byte(mk))
	return hash[:]
}

// Encrypt function accept plain_text and the master key hash, the later will be used as
// salt on the encryption, therefore the only key that can decrypt it is the one that 
// encrypted on the first place.
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

// Decrypt function receives a encrypted string and a master key, it will try to decrypt
// it, but if master key is not the same that was used on encryption, the decription will
// fail.
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

/*
Index:
func HashMasterKey(mk string) []byte
func Encrypt(plain_text string, key []byte) (string, error)
func Decrypt(encoded string, key []byte) (string, error)
*/
