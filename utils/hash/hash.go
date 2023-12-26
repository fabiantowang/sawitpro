package hash

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

// Define salt size
const saltSize = 16

// Define argon memory size
const argonMemory = 64 * 1024

// Define argon iterations
const argonIterations = 3

// Define argon parallelism
const argonParallelism = 2

// Define argon key length
const argonKeyLength = 32

// Generate 16 bytes randomly salt
func GenerateRandomSalt() ([]byte, error) {
	var salt = make([]byte, saltSize)

	rand.Read(salt)

	return salt, nil
}

// Hash password with salt using Argon2 algorithm
func HashPassword(password string, salt []byte) string {
	// Pass the plaintext password, salt and parameters to the argon2.IDKey
	// function. This will generate a hash of the password using the Argon2id
	// variant.
	return base64.RawStdEncoding.EncodeToString(argon2.IDKey([]byte(password), salt, argonIterations, argonMemory, argonParallelism, argonKeyLength))
}

// Check if two passwords match
func PasswordsMatch(hashedPassword, currPassword string, salt []byte) bool {
	var currPasswordHash = HashPassword(currPassword, salt)

	return hashedPassword == currPasswordHash
}
