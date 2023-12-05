package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

type Argon2Parameters struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

// HashPassword takes the password to be hashed and the Argon2 set of parameters.
// Return the encoded hash.
func HashPassword(password string, params *Argon2Parameters) (encodedHash string) {
	salt, err := generateSalt(params.SaltLength)
	if err != nil {
		return ""
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		params.Iterations,
		params.Memory,
		params.Parallelism,
		params.KeyLength,
	)

	bash64Salt := base64.RawStdEncoding.EncodeToString(salt)
	bash64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Create the standard encode hash representation of the password.

	encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		params.Memory,
		params.Iterations,
		params.Parallelism,
		bash64Salt,
		bash64Hash)

	return encodedHash
}

// generateSalt returns a random salt that is `num` length
func generateSalt(num uint32) ([]byte, error) {
	salt := make([]byte, num)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}

	return salt, nil
}

// CheckPasswordMatch compare two strings to see if their hashes match.
// Return true if the hashes match, otherwise return false.
func CheckPasswordMatch(password, encodedHash string) (match bool) {
	pass, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false
	}

	// Get the key from the plaintext password
	newHash := argon2.IDKey(
		[]byte(password),
		salt,
		pass.Iterations,
		pass.Memory,
		pass.Parallelism,
		pass.KeyLength,
	)

	if subtle.ConstantTimeCompare(hash, newHash) == 1 {
		return true
	}

	return false
}

// decodeHash decodes the Argon2 hash and returns the information from the hash
// so that comparisons can be made.
func decodeHash(encodedHash string) (p *Argon2Parameters, salt, hash []byte, err error) {
	values := strings.Split(encodedHash, "$")
	if len(values) != 6 {
		return nil, nil, nil, errors.New("the encoded hash is not in the correct format")
	}

	var version int
	_, err = fmt.Sscanf(values[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, errors.New("incompatible version")
	}

	params := &Argon2Parameters{}
	_, err = fmt.Sscanf(values[3], "m=%d,t=%d,p=%d",
		&params.Memory,
		&params.Iterations,
		&params.Parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(values[4])
	if err != nil {
		return nil, nil, nil, err
	}
	params.SaltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(values[5])
	if err != nil {
		return nil, nil, nil, err
	}
	params.KeyLength = uint32(len(hash))

	return params, salt, hash, nil
}
