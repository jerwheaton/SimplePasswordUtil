package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func writeLine(f *os.File, str string) {
	byts, err := f.WriteString(fmt.Sprintf("%s\n", str))
	if err != nil {
		panic(err)
	}
	if byts == 0 {
		panic(fmt.Errorf("0 bytes written"))
	}
}

func main() {
	badPasswords := []string{}

	file, err := os.Open("./passwords.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		badPasswords = append(badPasswords, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	f, err := os.Create("../data/data.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	writeLine(f, "package data")
	writeLine(f, "")
	writeLine(f, "var TotalSet = map[string]bool{")

	// foreach entry in file, add to set
	for i := range badPasswords {
		quoted := []byte("")
		quoted = strconv.AppendQuote(quoted, badPasswords[i])
		writeLine(f, fmt.Sprintf("\t%s: true,", string(quoted)))
	}

	writeLine(f, "}")

	f.Sync()
}
