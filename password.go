package password

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

const (
	digits       = "0123456789"
	symbols      = ".:;,/?!_-<>()[]*%=$&@#"
	lettersLower = "abcdefghijklmnopqrstuvwxyz"
	lettersUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// Config contains the parameters to generate a passwords' list
type Config struct {
	// Number of passwords to generate
	Number int
	// Length of the generated passwords
	Length int
	// Digits required
	Digits bool
	// Symbols required
	Symbols bool
	// Lower case characters required
	Lower bool
	// Upper case characters required
	Upper bool
}

// DefaultConfig stores default parameters for the generator
var DefaultConfig = Config{
	Number:  10,
	Length:  21,
	Digits:  true,
	Symbols: true,
	Lower:   true,
	Upper:   true,
}

// New returns a Config to pass to the generator
func New() *Config {
	return &DefaultConfig
}

// Generate returns a list of passwords based on the provided Config
func Generate(c *Config) []string {
	if c == nil {
		c = New()
	}
	chars, err := getCharacters(c)
	if err != nil {
		fmt.Println(err)
		return []string{}
	}
	list := make([]string, c.Number)
	rd := rand.Reader
	for i := 0; i < c.Number; i++ {
		var pwd string
		for j := 0; j < c.Length; j++ {
			n, _ := rand.Int(rd, big.NewInt(int64(len(chars))))
			pwd = pwd + string(chars[n.Int64()])
		}
		list[i] = pwd
	}
	return list
}

// getCharacters returns the string with authorized characters for the password
func getCharacters(c *Config) (string, error) {
	var chars string
	if c.Digits {
		chars += digits
	}
	if c.Symbols {
		chars += symbols
	}
	if c.Lower {
		chars += lettersLower
	}
	if c.Upper {
		chars += lettersUpper
	}
	if len(chars) == 0 {
		return chars, errors.New("cannot generate password from empty set")
	}
	return chars, nil
}
