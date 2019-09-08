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
		'z': 't',
		'w': 'i',
		'm': 'g',
		'q': 'm',
		's': 'c',
		'r': 'h',
		'v': 'n',
		'c': 'e',
		'o': 'w',
		'u': 's',
		'i': 'a',
		'f': 'p',
		'p': 'r',
		'j': 'v',
		'x': 'd',
		'b': 'n',
		'y': 'y',
		'a': 'o',
		'h': 'f',
		'e': 'u',
		'l': 'l',
		'n': 'b',
		'd': 'z',
		'g': 'k',
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
			} else {
				fmt.Print(string(r))
			}
		} else if unicode.IsLower(r) {
			if substitue, ok := subsitutionMap[unicode.ToLower(r)]; ok {
				fmt.Print(ansi.Red + string(unicode.ToLower(substitue)) + ansi.Reset)
			} else {
				fmt.Print(string(r))
			}
		} else {
			fmt.Print(string(r))
		}
	}

	os.Stdout.WriteString("\n")
}
