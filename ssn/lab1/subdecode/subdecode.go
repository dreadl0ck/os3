package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/mgutz/ansi"
)

var flagInput = flag.String("input", "crackme.txt", "supply an input file")

func main() {

	var m = map[rune]rune{
		'd': 'n',
		'h': 'l',
		'x': 'o',
		'p': 's',
		'c': 'h',
		'q': 't',
		'j': 'p',
		'v': 'e',
		'K': 'D',
		'k': 'd',
		'n': 'a',
		'r': 'r',
		'P': 'S',
		'z': 'u',
		'a': 'c',
		'o': 'f',
		'w': 'y',
		'e': 'b',
		'Q': 'T',
		'J': 'P',
		't': 'w',
		'i': 'g',
		'm': 'v',
		'N': 'A',
		'f': 'k',
		'g': 'i',
		'T': 'W',
		'b': 'j',
		's': 'm',
	}

	flag.Parse()

	c, err := ioutil.ReadFile(*flagInput)
	if err != nil {
		panic(err)
	}

	for _, char := range c {
		r := rune(char)

		if substitue, ok := m[r]; ok {
			fmt.Print(ansi.Red + string(substitue) + ansi.Reset)
		} else {
			fmt.Print(string(r))
		}
	}
}
