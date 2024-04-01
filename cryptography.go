package GoSkeletonPy

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"golang.org/x/crypto/argon2"
	"io"
)

// generateRandomSalt generates a random salt of 16 bytes.
// It panics if it encounters an error during the generation.
func generateRandomSalt() []byte {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}
	return salt
}

// Encrypt takes a byte slice of data and a password string as input.
// It generates a random salt and uses it along with the password to derive an AES key using Argon2.
// It then creates a new GCM cipher with the derived key and generates a random nonce.
// The data is encrypted using the GCM cipher and the nonce, and the resulting ciphertext is returned along with the salt.
// If an error occurs during any of these steps, it is returned.
func Encrypt(data []byte, password string) ([]byte, error) {
	salt := generateRandomSalt()

	block, err := aes.NewCipher(argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return append(salt, ciphertext...), nil
}

// Decrypt takes a byte slice of data and a password string as input.
// It extracts the salt from the data and uses it along with the password to derive an AES key using Argon2.
// It then creates a new GCM cipher with the derived key.
// The nonce and the ciphertext are extracted from the data, and the ciphertext is decrypted using the GCM cipher and the nonce.
// The resulting plaintext is returned.
// If an error occurs during any of these steps, it is returned.
func Decrypt(data []byte, password string) ([]byte, error) {
	salt := data[:len(generateRandomSalt())]
	data = data[len(generateRandomSalt()):]

	block, err := aes.NewCipher(argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, err
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
