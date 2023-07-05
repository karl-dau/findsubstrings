package main

import (
	"flag"
	"fmt"

	"github.com/karl-dau/findsubstrings/substrings"
)

func main() {
	var textToSearch, subtext string
	// read in using flags package
	flag.StringVar(&textToSearch, "textToSearch", "", "text to search")
	flag.StringVar(&subtext, "subText", "", "subText to search for")
	flag.Parse()

	if textToSearch == "" || subtext == "" {
		fmt.Println("Please provide both textToSearch and subText")
		// nothing to do, early exit
		return
	}

	indices := substrings.FindIndices(textToSearch, subtext)
	if indices == nil {
		// no matches found, don't print anything
		return
	}
	for _, index := range indices {
		fmt.Print(index, " ")
	}
	fmt.Println()

}
