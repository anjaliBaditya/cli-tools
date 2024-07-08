package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type PasswordGenerator struct{}

func NewPasswordGenerator() *PasswordGenerator {
	return &PasswordGenerator{}
}

func (p *PasswordGenerator) GeneratePassword(length int, useUpper bool, useLower bool, useNumbers bool, useSpecialChars bool) (string, error) {
	if length < 1 {
		return "", fmt.Errorf("length must be greater than 0")
	}

	charsets := []string{}
	if useUpper {
		charsets = append(charsets, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}
	if useLower {
		charsets = append(charsets, "abcdefghijklmnopqrstuvwxyz")
	}
	if useNumbers {
		charsets = append(charsets, "0123456789")
	}
	if useSpecialChars {
		charsets = append(charsets, "!@#$%^&*()_+-=")
	}

	if len(charsets) == 0 {
		return "", fmt.Errorf("must select at least one character set")
	}

	password := ""
	for i := 0; i < length; i++ {
		charset := charsets[rand.Intn(len(charsets))]
		password += string(charset[rand.Intn(len(charset))])
	}

	return password, nil
}

func main() {
	p := NewPasswordGenerator()
	password, err := p.GeneratePassword(12, true, true, true, true)
	if err!= nil {
		log.Fatal(err)
	}

	fmt.Println("Generated password:", password)
}
