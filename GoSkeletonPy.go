package GoSkeletonPy

import (
	"os"
	"strings"
)

// DefaultSkeletonFileExtension is the default file extension for the skeleton file
const DefaultSkeletonFileExtension = "skeletonpy"

// FileToSkeleton takes a Python file, encrypts it using AES, and stores it in a file.
// It then opens the original file and replaces all the function bodies with ...
// If a password is provided, it is used for the encryption.
// The function returns true if the operation was successful, false otherwise.
func FileToSkeleton(path, password, ff string) (bool, error) {
	// If no file format is provided, use the default
	if ff == "" {
		ff = DefaultSkeletonFileExtension
	}

	// Append .py to the path if it doesn't already end with .py
	if !strings.HasSuffix(path, ".py") {
		path += ".py"
	}

	// Check if the file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, err
	}

	// Read the file
	data, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}

	// If a password is provided, encrypt the data
	encryptedData := data
	if password != "" {
		encryptedData, err = Encrypt(data, password)
	}

	if err != nil {
		return false, err
	}

	// Get the file path without the extension and append the file format
	encryptedPath, _ := strings.CutSuffix(path, ".py")
	encryptedPath += "." + ff

	// Write the encrypted data to the new file
	err = os.WriteFile(encryptedPath, encryptedData, 0644)
	if err != nil {
		return false, err
	}

	// Delete the original file
	err = os.Remove(path)
	if err != nil {
		panic(err)
	}

	// Write the mock file
	err = os.WriteFile(path, []byte(strings.Join(GeneratePythonSkeleton(strings.Split(string(data), "\n")), "\n")), 0644)
	if err != nil {
		return false, err
	}

	return true, nil
}

// RestoreSkeletonFile takes a skeleton file and restores the original Python file.
// If a password is provided, it is used for the decryption.
// The function returns true if the operation was successful, false otherwise.
func RestoreSkeletonFile(path string, password string, ff string) (bool, error) {
	// If no file format is provided, use the default
	if ff == "" {
		ff = DefaultSkeletonFileExtension
	}

	// Adjust the path if necessary
	if !strings.HasSuffix(path, "."+ff) && strings.HasSuffix(path, ".py") {
		path, _ = strings.CutSuffix(path, ".py")
		path += "." + ff
	}

	// Check if the file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, err
	}

	// Delete the mock file
	mockPath, _ := strings.CutSuffix(path, "."+ff)
	mockPath += ".py"

	// Check if the mock file exists
	if _, err := os.Stat(mockPath); os.IsNotExist(err) {
		return false, err
	}

	// Remove the mock file
	err := os.Remove(mockPath)
	if err != nil {
		return false, err
	}

	// Read the encrypted data
	data, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}

	// If a password is provided, decrypt the data
	decryptedData := data
	if password != "" {
		decryptedData, err = Decrypt(data, password)
	}

	if err != nil {
		return false, err
	}

	// Write the decrypted data to the original file
	err = os.WriteFile(mockPath, decryptedData, 0644)
	if err != nil {
		return false, err
	}

	// Delete the encrypted file
	err = os.Remove(path)
	if err != nil {
		return false, err
	}

	return true, nil
}
