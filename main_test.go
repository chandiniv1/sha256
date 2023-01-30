package main

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSha256(t *testing.T) {
	message := []byte("hello world")

	hash := sha256.New()

	hash.Write(message)

	hashValue := hash.Sum(nil)

	want := fmt.Sprintf("%x", hashValue)

	got := "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"

	//go test
	assert.Equal(t, got, want)

}
