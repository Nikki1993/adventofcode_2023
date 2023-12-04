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

	var data [][]string
	mep := make(map[int]int)

	y := 0
	for s.Scan() {
		d := strings.Split(s.Text(), ":")
		scratches := strings.Split(d[1], "|")

		data = append(data, scratches)
		mep[y] = 1
		y++
	}

	for i, d := range data {
		winning, numbers := strings.Fields(d[0]), strings.Fields(d[1])

		m := make(map[int]struct{})
		inc := i + 1
		for _, num := range numbers {
			for _, win := range winning {
				if num != win {
					continue
				}

				_, ok := mep[inc]
				if !ok {
					break
				}

				m[inc] = struct{}{}
				mep[inc]++
				inc++
				break
			}
		}

		for z := 1; z < mep[i]; z++ {
			for k := range m {
				mep[k]++
			}
		}
	}

	total := 0
	for _, v := range mep {
		total += v
	}

	fmt.Println(total)
}
