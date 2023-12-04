package main

import (
	"adventofcode/utils"
	"fmt"
	"os"
	"regexp"
)

func main() {
	reS := regexp.MustCompile(`[^\w.\s]`)
	reN := regexp.MustCompile(`\d+`)
	fPath := utils.GetFilePath(os.Args)
	f, fClose := utils.MustOpenFile(fPath)
	defer fClose(f)

	s := utils.ScanAndSplit(f)

	var txtSlice []string

	for s.Scan() {
		txtSlice = append(txtSlice, s.Text())
	}

	txtSliceLen := len(txtSlice)
	total := 0
	for i := 1; i < txtSliceLen; i++ {
		prev := txtSlice[i-1]
		curr := txtSlice[i]

		nums := reN.FindAllStringIndex(prev, -1)
		syms := reS.FindAllStringIndex(prev, -1)
		var prevSyms []int
		for _, v := range syms {
			prevSyms = append(prevSyms, v[0])
		}

		var currSyms []int
		syms = reS.FindAllStringIndex(curr, -1)
		for _, v := range syms {
			currSyms = append(currSyms, v[0])
		}

		fmt.Println("NUMBERS", nums, "SYMBOLS PREV", prevSyms, "SYMBOLS CURRENT", currSyms, prev)

	}

	fmt.Println("TOTAL", total)
}
