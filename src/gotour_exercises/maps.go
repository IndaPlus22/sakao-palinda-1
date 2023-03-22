package main

import (
	"golang.org/x/tour/wc"; "strings"
)

func WordCount(s string) map[string]int {
	haha := strings.Fields(s)
	
	dict := map[string]int{}
    count := 0
    for _ , word := range haha {
        dict[word] = count+1
    }
	
	return dict
	
}

func main() {
	wc.Test(WordCount)
}
