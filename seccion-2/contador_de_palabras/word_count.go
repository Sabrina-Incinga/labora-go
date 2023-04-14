package main

import (
	"fmt"
	"strings"
	//"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	var countMap = make(map[string]int)

	for _, word := range words {
		count, ok := countMap[word]
		if ok {
			countMap[word] = count + 1
		} else {
			countMap[word] = 1
		}
	}

	return countMap
}
func main() {
	//wc.Test(WordCount)
	fmt.Println(WordCount("I ate a donut. Then I ate another donut."))
}
