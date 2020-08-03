package main

import (
	"flag"
	"fmt"
)

var camelCase string

func init() {
	flag.StringVar(&camelCase, "c", "", "camelcase")
}

func main() {
	flag.Parse()
	w := wordCount(camelCase)
	fmt.Printf("The number words in %s is %d\n", camelCase, w)
}

func wordCount(s string) int {
	w := 0
	if len(s) == 0 {
		return 0
	}
	for _, c := range s {
		ascii := int(c)
		if ascii <= 90 && ascii >= 65 {
			w++
		}
	}
	return w + 1
}
