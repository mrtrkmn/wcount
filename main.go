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

func main() {

	last := flag.Int("last", 10, "usage -last=10")
	fileName := flag.String("file", "./readme", "usage -file=./readme.md")
	flag.Parse()
	wordsCount := make(map[string]int)
	var words []word

	data, err := os.ReadFile(*fileName)
	if err != nil {
		panic(err)
	}
	cleanedData := strings.Trim(string(data), "\n")
	dataInArray := strings.Split(cleanedData, " ")
	for _, s := range dataInArray {
		wordsCount[s]++
	}

	for k, v := range wordsCount {
		words = append(words, word{k, v})
	}
	sort.Slice(words, func(i, j int) bool {
		return words[i].frequency < words[j].frequency
	})

	for i := len(words) - *last; i < len(words); i++ {
		fmt.Printf(" %d %s\n", words[i].frequency, words[i].key)
	}
}
