package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

type word struct {
	key       string
	frequency int
}

func preliminaryCheck(words, last int) error {
	if words < last {
		return fmt.Errorf("Number of words ( %d ) in file is less than %d. Use -last flag when running the execution\n", words, last)
	}
	return nil
}

func cleanData(data string) []string {
	cleanedData := strings.Trim(data, "\n")
	// cleaning spaces in the beginning and end of the string
	cleanedData = strings.TrimPrefix(cleanedData, " ")
	cleanedData = strings.TrimSuffix(cleanedData, " ")
	// splitting the string into words
	dataInArray := strings.Split(cleanedData, " ")
	return dataInArray
}

func constructMap(data []string) map[string]int {
	// creating a map to store the words and their frequencies
	wordMap := make(map[string]int)
	for _, word := range data {
		wordMap[word]++
	}
	return wordMap
}

func main() {

	last := flag.Int("last", 10, "usage -last=10")
	fileName := flag.String("file", "./readme", "usage -file=./readme.md")
	flag.Parse()

	var words []word

	data, err := os.ReadFile(*fileName)
	if err != nil {
		panic(err)
	}

	dataInArray := cleanData(string(data))

	wordsMap := constructMap(dataInArray)

	for k, v := range wordsMap {
		words = append(words, word{k, v})
	}

	sort.Slice(words, func(i, j int) bool {
		return words[i].frequency < words[j].frequency
	})

	if err := preliminaryCheck(len(words), *last); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := len(words) - *last; i < len(words); i++ {
		fmt.Printf(" %d %s\n", words[i].frequency, words[i].key)
	}

}
