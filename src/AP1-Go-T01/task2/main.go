package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Words struct {
	list []string
	k    int
}

type wordCount struct {
	word  string
	count int
}

func getUserInput(words *Words) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input words:")
	scanner.Scan()
	input := scanner.Text()
	words.list = strings.Fields(input)

	fmt.Println("Input k (number of words):")
	for {
		var err error
		_, err = fmt.Scanln(&words.k)

		if err == nil {
			break
		} else {
			fmt.Println("Invalid input. Try again")
		}
	}
}

func listSort(words *Words) []wordCount {
	// мапа слово -> количество повторов
	freqMap := make(map[string]int)
	for _, wd := range words.list {
		freqMap[wd]++
	}

	var wordCounts []wordCount
	for w, c := range freqMap {
		wordCounts = append(wordCounts, wordCount{word: w, count: c})
	}

	// Сортировка: сначала по количеству по убыванию, затем лексикографически по возрастанию
	sort.Slice(wordCounts, func(i, j int) bool {
		if wordCounts[i].count == wordCounts[j].count {
			return wordCounts[i].word < wordCounts[j].word
		}
		return wordCounts[i].count > wordCounts[j].count
	})

	return wordCounts
}

func topKWords(wordCounts []wordCount, k int) []string {
	if k > len(wordCounts) {
		k = len(wordCounts)
	}

	var result []string
	for i := 0; i < k; i++ {
		result = append(result, wordCounts[i].word)
	}

	return result
}

func main() {
	var words Words

	getUserInput(&words)
	if len(words.list) == 0 {
		// Пустая строка на входе
		fmt.Println("")
		return
	}

	wordCounts := listSort(&words)
	topWords := topKWords(wordCounts, words.k)
	fmt.Println(strings.Join(topWords, " "))
}
