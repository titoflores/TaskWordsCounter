package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

func WordsCount(text string) map[string]int {
	lowerText := strings.ToLower(text)

	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		fmt.Print(err)
	}
	onlyAlphaText := reg.ReplaceAllString(lowerText, " ")
	words := strings.Fields(onlyAlphaText)

	wordsCountMap := make(map[string]int)
	for _, word := range words {
		wordsCountMap[word]++
	}
	return wordsCountMap
}

func SortWords(wordsCountMap map[string]int) []string {
	var words []string
	for word := range wordsCountMap {
		words = append(words, word)
	}
	sort.Strings(words)
	return words
}

func main() {
	dataBin, err := ioutil.ReadFile("words.txt")
	if err != nil {
		fmt.Print(err)
	}
	textFile := string(dataBin)
	countWord := WordsCount(textFile)
	wordOutput := SortWords(countWord)

	for _, word := range wordOutput {
		fmt.Println(word, countWord[word])
	}
}
