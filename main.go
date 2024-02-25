package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input file> <output file>")
		os.Exit(1)
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	inputData, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}
	text := string(inputData)
	formattedText := FormatText(text)
	err = os.WriteFile(outputFile, []byte(formattedText), 0o644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		os.Exit(1)
	}
	fmt.Println("Text conversion successful")
}
func Capitalize(s string) string {
	c := true
	sb := []rune(s)
	for i, b := range sb {
		if (b >= 'a' && b <= 'z') && c {
			sb[i] = b - 32
		} else if (b >= 'A' && b <= 'Z') && !c {
			sb[i] = b + 32
		}
		if (b < 'A' || b > 'Z') && (b < 'a' || b > 'z') && (b < '0' || b > '9') {
			c = true
		} else {
			c = false
		}
	}
	return string(sb)
}
func FormatText(text string) string {
	// Define regex patterns
	hexRegex := regexp.MustCompile(`\b([0-9A-Fa-f]+)\s+\(hex\)`)
	binRegex := regexp.MustCompile(`\b([01]+)\s+\(bin\)`)
	upRegex := regexp.MustCompile(`\b(\w+)\s+\(up\)`)
	lowRegex := regexp.MustCompile(`\b(\w+)\s+\(low\)`)
	capRegex := regexp.MustCompile(`\b(\w+)\s+\(cap\)`)
	upNumRegex := regexp.MustCompile(`\b((?:\w+\s+){0,}\w+)\s+\(up,\s*(\d+)\)`)
	lowNumRegex := regexp.MustCompile(`\b((?:\w+\s+){0,}\w+)\s+\(low,\s*(\d+)\)`)
	capNumRegex := regexp.MustCompile(`\b((?:\w+\s+){0,}\w+)\s+\(cap,\s*(\d+)\)`)
	puncRegex := regexp.MustCompile(`\s+([.,!?:;])`)
	puncGroupRegex := regexp.MustCompile(`([.,!?:;])\s+([.,!?:;])`)
	quoteRegex := regexp.MustCompile(`'\s+(.*?)\s+'`)
	aAnRegex := regexp.MustCompile(`\ba\s+([aeiouhAEIOUH])`)
	text = regexp.MustCompile(`,(\S)`).ReplaceAllString(text, ", $1")
	//
	// Replace hex and bin
	text = hexRegex.ReplaceAllStringFunc(text, func(s string) string {
		hexNum := hexRegex.FindStringSubmatch(s)[1]
		num, _ := strconv.ParseInt(hexNum, 16, 64)
		return fmt.Sprintf("%d", num)
	})
	text = binRegex.ReplaceAllStringFunc(text, func(s string) string {
		binNum := binRegex.FindStringSubmatch(s)[1]
		num, _ := strconv.ParseInt(binNum, 2, 64)
		return fmt.Sprintf("%d", num)
	})
	// Replace up, low and cap
	text = upRegex.ReplaceAllStringFunc(text, func(s string) string {
		word := upRegex.FindStringSubmatch(s)[1]
		return strings.ToUpper(word)
	})
	text = lowRegex.ReplaceAllStringFunc(text, func(s string) string {
		word := lowRegex.FindStringSubmatch(s)[1]
		return strings.ToLower(word)
	})
	text = capRegex.ReplaceAllStringFunc(text, func(s string) string {
		word := capRegex.FindStringSubmatch(s)[1]
		return Capitalize(word)
	})
	// Replace upNum, lowNum and capNum
	text = upNumRegex.ReplaceAllStringFunc(text, func(s string) string {
		match := upNumRegex.FindStringSubmatch(s)
		numWordsStr := match[1]
		numWordsSlice := strings.Fields(numWordsStr)
		numWordsToChangeStr := match[2]
		numWordsToChangeInt, _ := strconv.Atoi(numWordsToChangeStr)
		if numWordsToChangeInt > len(numWordsSlice) {
			numWordsToChangeInt = len(numWordsSlice)
		}
		for i := len(numWordsSlice) - numWordsToChangeInt; i < len(numWordsSlice); i++ {
			numWordsSlice[i] = strings.ToUpper(numWordsSlice[i])
		}
		return strings.Join(numWordsSlice, " ")
	})
	text = lowNumRegex.ReplaceAllStringFunc(text, func(s string) string {
		match := lowNumRegex.FindStringSubmatch(s)
		numWordsStr := match[1]
		numWordsSlice := strings.Fields(numWordsStr)
		numWordsToChangeStr := match[2]
		numWordsToChangeInt, _ := strconv.Atoi(numWordsToChangeStr)
		if numWordsToChangeInt > len(numWordsSlice) {
			numWordsToChangeInt = len(numWordsSlice)
		}
		for i := len(numWordsSlice) - numWordsToChangeInt; i < len(numWordsSlice); i++ {
			numWordsSlice[i] = strings.ToLower(numWordsSlice[i])
		}
		return strings.Join(numWordsSlice, " ")
	})
	text = capNumRegex.ReplaceAllStringFunc(text, func(s string) string {
		match := capNumRegex.FindStringSubmatch(s)
		numWordsStr := match[1]
		numWordsSlice := strings.Fields(numWordsStr)
		numWordsToChangeStr := match[2]
		numWordsToChangeInt, _ := strconv.Atoi(numWordsToChangeStr)
		if numWordsToChangeInt > len(numWordsSlice) {
			numWordsToChangeInt = len(numWordsSlice)
		}
		for i := len(numWordsSlice) - numWordsToChangeInt; i < len(numWordsSlice); i++ {
			numWordsSlice[i] = Capitalize(numWordsSlice[i])
		}
		return strings.Join(numWordsSlice, " ")
	})
	// Replace punctuation
	text = puncRegex.ReplaceAllString(text, "$1")
	text = puncGroupRegex.ReplaceAllString(text, "$1$2")
	text = quoteRegex.ReplaceAllString(text, "'$1'")
	// Replace a/an
	text = aAnRegex.ReplaceAllStringFunc(text, func(s string) string {
		return "an " + aAnRegex.FindStringSubmatch(s)[1]
	})
	return text
}
