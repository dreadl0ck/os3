package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"unicode"

	"github.com/mgutz/ansi"
)

var flagInput = flag.String("input", "crackme.txt", "supply an input file")

func main() {

	var subsitutionMap = map[rune]rune{
		'd': 'n',
		'h': 'l',
		'x': 'o',
		'p': 's',
		'c': 'h',
		'q': 't',
		'j': 'p',
		'v': 'e',
		'k': 'd',
		'n': 'a',
		'r': 'r',
		'z': 'u',
		'a': 'c',
		'o': 'f',
		'w': 'y',
		'e': 'b',
		't': 'w',
		'i': 'g',
		'm': 'v',
		'f': 'k',
		'g': 'i',
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

		if unicode.IsUpper(r) {
			if substitue, ok := subsitutionMap[unicode.ToLower(r)]; ok {
				fmt.Print(ansi.Red + string(unicode.ToUpper(substitue)) + ansi.Reset)
			}
		} else if unicode.IsLower(r) {
			if substitue, ok := subsitutionMap[unicode.ToLower(r)]; ok {
				fmt.Print(ansi.Red + string(unicode.ToLower(substitue)) + ansi.Reset)
			}
		} else {
			fmt.Print(string(r))
		}
	}

	os.Stdout.WriteString("\n")
}
