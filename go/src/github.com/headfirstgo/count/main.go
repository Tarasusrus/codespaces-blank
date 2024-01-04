// count подсчитывает количество вхождений
// каждой строки в файле.

package main

import (
	"fmt"
	"log"
	"test/go/src/github.com/headfirstgo/datafile"
)
func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	for name, count := range counts {
		fmt.Println(name, count)
	}
}