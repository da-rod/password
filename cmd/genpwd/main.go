package main

import (
	"flag"
	"fmt"

	"github.com/da-rod/password"
)

const (
	n  = 25    // number of passwords
	l  = 25    // passwords' length
	d  = true  // include digits
	s  = true  // include symbols
	lw = true  // include lower case characters
	up = true  // include upper case characters
	i  = false // prepend passwords with their id
)

func main() {
	var (
		num     = flag.Int("n", n, "Number of passwords to generate")
		size    = flag.Int("l", l, "Length of the generated passwords")
		digits  = flag.Bool("d", d, "Include digits")
		symbols = flag.Bool("s", s, "Include symbols")
		lower   = flag.Bool("low", lw, "Include lower case characters")
		upper   = flag.Bool("up", up, "Include upper case characters")
		id      = flag.Bool("i", i, "Include password id in list")
	)
	flag.Parse()

	config := &password.Config{
		Number:  *num,
		Length:  *size,
		Digits:  *digits,
		Symbols: *symbols,
		Lower:   *lower,
		Upper:   *upper,
	}

	list := password.Generate(config)
	for k, pwd := range list {
		if *id {
			fmt.Printf("%d\t%s\n", k+1, pwd)
		} else {
			fmt.Println(pwd)
		}
	}
}
