/*
 * Created for OS3 - Security and Network Engineering
 * Copyright (c) 2019 Philipp Mieden <dreadl0ck [at] protonmail [dot] ch>
 */

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strconv"
	"unicode"

	"github.com/evilsocket/islazy/tui"
)

var flagInput = flag.String("input", "unknown.txt", "supply an input file")

func main() {

	flag.Parse()

	c, err := ioutil.ReadFile(*flagInput)
	if err != nil {
		panic(err)
	}

	fmt.Println(analyze(c, *flagInput))
}

// CipherTextInfo

type cipherTextInfo struct {
	Filename       string
	HasDigits      bool
	HasSpacing     bool
	HasPunctuation bool
	HasSymbols     bool
	LetterFreqs    map[rune]int
	DiphoFreqs     map[string]int
	TotalBytes     int
}

func (c *cipherTextInfo) String() string {

	var b bytes.Buffer
	b.WriteString("Filename: " + filepath.Base(c.Filename) + "\n")

	cols := [][]string{}
	cols = append(cols, []string{"TotalBytes", strconv.Itoa(c.TotalBytes)})
	cols = append(cols, []string{"HasDigits", strconv.FormatBool(c.HasDigits)})
	cols = append(cols, []string{"HasSpacing", strconv.FormatBool(c.HasSpacing)})
	cols = append(cols, []string{"HasPunctuation", strconv.FormatBool(c.HasPunctuation)})
	cols = append(cols, []string{"HasSymbols", strconv.FormatBool(c.HasSymbols)})

	tui.Table(&b, []string{"Attribute", "Value"}, cols)
	cols = [][]string{}

	var letters []int
	for _, count := range c.LetterFreqs {
		letters = append(letters, count)
	}
	sort.Ints(letters)

	b.WriteString("\nLetterFreqs:\n")
	for _, num := range letters[len(letters)-5:] {
		for r, count := range c.LetterFreqs {
			if count == num {
				cols = append(cols, []string{string(r), strconv.Itoa(count), strconv.FormatFloat(float64(count)/float64(len(c.LetterFreqs)), 'f', 2, 64) + "%"})
			}
		}
	}

	tui.Table(&b, []string{"Letter", "Count", "Share"}, cols)
	cols = [][]string{}

	var diphors []int
	for _, count := range c.DiphoFreqs {
		diphors = append(diphors, count)
	}
	sort.Ints(diphors)

	b.WriteString("\nDiphorFreqs:\n")
	for _, num := range diphors[len(diphors)-5:] {
		for d, count := range c.DiphoFreqs {
			if count == num {
				cols = append(cols, []string{d, strconv.Itoa(count), strconv.FormatFloat(float64(count)/float64(len(c.DiphoFreqs)), 'f', 2, 64) + "%"})
			}
		}
	}

	tui.Table(&b, []string{"Diphor", "Count", "Share"}, cols)

	return b.String()
}

// Utils

func analyze(buf []byte, path string) (info *cipherTextInfo) {

	info = &cipherTextInfo{
		LetterFreqs: map[rune]int{},
		DiphoFreqs:  map[string]int{},
	}
	info.Filename = path

	var diphor = ""

	for _, b := range buf {

		r := rune(b)
		info.LetterFreqs[r]++

		if unicode.IsSpace(r) {
			info.HasSpacing = true
			continue
		}
		diphor += string(b)

		if len(diphor) == 2 {

			// skip if dihpor has already been counted
			if _, ok := info.DiphoFreqs[diphor]; ok {
				continue
			}

			cmp := ""
			for _, b := range buf {

				// skip space, numbers and punctuation
				if !unicode.IsLetter(rune(b)) {
					continue
				}

				cmp += string(b)
				if len(cmp) == 2 {

					// fmt.Println("cmp", diphor, cmp)
					if cmp == diphor {
						info.DiphoFreqs[diphor]++
					}

					// remove first letter
					cmp = cmp[1:]
				}
			}
			// remove first letter
			diphor = diphor[1:]
		}

		switch {
		case unicode.IsDigit(r):
			info.HasDigits = true
		case unicode.IsPunct(r):
			info.HasPunctuation = true
		case unicode.IsSymbol(r):
			info.HasSymbols = true
		}
	}

	info.TotalBytes = len(buf)

	return
}
