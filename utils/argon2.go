package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

const (
	memory  = 64 * 1024 // 64 MB
	time    = 3         // number of iterations
	threads = 2         // number of threads
	keyLen  = 32        // output length
	saltLen = 16
)

func generateSalt() ([]byte, error) {
	b := make([]byte, saltLen)
	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func HashApiKey(key string) (string, error) {
	salt, err := generateSalt()
	if err != nil {
		return "", err
	}

	hashedKey := argon2.IDKey([]byte(key), salt, time, memory, threads, keyLen)

	base64Salt := base64.RawStdEncoding.EncodeToString(salt)
	base64Hash := base64.RawStdEncoding.EncodeToString(hashedKey)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, memory, time, threads, base64Salt, base64Hash)

	return encodedHash, nil
}

func VerifyApiKey(encodedHash, rawKey string) (bool, error) {
	var version int
	var memory, time, threads int
	var b64Salt, b64Hash string

	_, err := fmt.Sscanf(encodedHash, "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", &version, memory, time, threads, b64Salt, b64Hash)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(b64Salt)
	if err != nil {
		return false, err
	}

	hash, err := base64.RawStdEncoding.DecodeString(b64Hash)
	if err != nil {
		return false, err
	}

	computedHash := argon2.IDKey([]byte(rawKey), salt, uint32(time), uint32(memory), uint8(threads), uint32(len(hash)))
	if subtle.ConstantTimeCompare([]byte(encodedHash), computedHash) == 1 {
		return true, nil
	}

	return false, nil

}
