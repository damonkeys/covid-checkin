package business

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

const letters = "1234567890abcdefghijklmnopqrstuvwxyz"

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func randSeq(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//https://stackoverflow.com/questions/12321133/how-to-properly-seed-random-number-generator
func initSeed() error {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		return err
	}
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
	return nil
}

// generateCode generates a code of random characters and number from the constant letters with a length of 5
func generateCode() (string, error) {
	err := initSeed()
	if err != nil {
		return "", err
	}
	return randSeq(5), nil
}
