package main

import (
	"flag"
	"fmt"
)

var input string
var rotate int

func init() {
	flag.StringVar(&input, "c", "", "encrypt string")
	flag.IntVar(&rotate, "r", 0, "rotate")
}

func main() {
	flag.Parse()
	enc := encrypt(input, rotate)
	fmt.Printf("Ceaser Cipher for %s is %s for rotation %d\n", input, enc, rotate)
}

func encrypt(s string, rotate int) string {
	enc := ""
	min, max := 97, 122
	minC, maxC := 65, 90

	if rotate == 26 {
		return s
	}
	rotate = rotate % 26

	for _, c := range s {
		ch := int(c)
		if ch >= min && ch <= max {
			ch = rotateCalc(ch, rotate, min, max)
		} else if ch >= minC && ch <= maxC {
			ch = rotateCalc(ch, rotate, minC, maxC)
		}
		enc = enc + string(ch)
	}
	return enc
}

func rotateCalc(ch, rotate, min, max int) int {
	ch = ch + rotate
	if ch > max {
		d := ch - max
		ch = min + d - 1
	}
	return ch
}
