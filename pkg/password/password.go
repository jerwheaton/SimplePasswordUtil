package password

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/jerwheaton/SimplePasswordUtil/data"
	"github.com/willf/bloom"
)

const (
	minLen = 8
	maxLen = 72

	maxDesiredEntropy = 100.0
)

var badPasswords = []string{}

func Check(listPath, secret string, useBloom bool) (bool, error) {
	if listPath == "" {
		fmt.Println("Would like to import password set...", len(data.TotalSet))
		return false, nil
	}

	file, err := os.Open(listPath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		badPasswords = append(badPasswords, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return false, err
	}

	if useBloom {
		filter := bloom.New(10*uint(len(badPasswords)), 5)
		for i := range badPasswords {
			filter.Add([]byte(badPasswords[i]))
		}
		m := filter.Test([]byte(secret))
		fp := filter.EstimateFalsePositiveRate(uint(len(badPasswords)))
		fmt.Println("False positive rate: ", fp)
		return m, nil
	}

	sort.Strings(badPasswords)
	i := sort.SearchStrings(badPasswords, secret)

	return badPasswords[i] == secret, nil
}

var charSets = []string{
	"abcdefghijklmnopqrstuvwxyz",
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"0123456789",
	"~`!@#$%^&*()",
	"-_=+[]{}\\|;:\"'<>.,?/",
}

func totalSet(secret string) (int, []rune) {
	collections := make([]int, len(charSets))
	keys := make(map[rune]bool)
	unique := []rune{}
	unicodeChars := 0

	// String is short, may as well iterate over a couple times
	// Reduces work for second iteration, which isn't super optimized
	// and allows us to pass back slice of unique runes.
	for _, c := range secret {
		if _, v := keys[c]; !v {
			keys[c] = true
			unique = append(unique, c)
		}
	}

	for _, r := range unique {
		for i, set := range charSets {
			if strings.ContainsRune(set, r) {
				collections[i] = len(set)
				break
			} else if r >= 127 {
				// Not going to characterize unicode char sets
				unicodeChars++
			}
		}
	}

	set := unicodeChars
	for _, count := range collections {
		set += count
	}

	return set, unique
}

// Rate is a simple function to output the approximate entropy of a string.
func Rate(secret string) float64 {
	set, unique := totalSet(secret)
	entropy := float64(len(secret)) * math.Log2(float64(set))

	// Use square root to reduce weight of uniqueness
	uniqueness := math.Sqrt(float64(len(unique)) / float64(len(secret)))
	weightedEntropy := entropy * uniqueness

	rating := weightedEntropy / maxDesiredEntropy

	return math.Round(rating*1000) / 1000 // Round to 3 decimal places

}
