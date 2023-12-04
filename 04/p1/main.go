package main

import (
	"adventofcode/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	p := utils.GetFilePath(os.Args)
	f, fClose := utils.MustOpenFile(p)
	defer fClose(f)

	s := utils.ScanAndSplit(f)

	total := 0
	for s.Scan() {
		d := strings.Split(s.Text(), ":")
		scratches := strings.Split(d[1], "|")
		winning, numbers := strings.Fields(scratches[0]), strings.Fields(scratches[1])

		acc := 0
		fmt.Println(winning, numbers, acc)
		for _, num := range numbers {
			for _, win := range winning {
				if num != win {
					continue
				}

				fmt.Println("Match found", num, "Winning", win)

				if acc == 0 {
					acc = 1
					break
				}

				acc = acc * 2
				break
			}
		}

		total += acc
	}

	fmt.Println(total)
}
