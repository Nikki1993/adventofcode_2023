package main

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	re := regexp.MustCompile(`(\d)`)
	words := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

	fPath := utils.GetFile(os.Args)
	f, fClose := utils.MustOpenFile(fPath)
	defer fClose(f)

	fScan := utils.ScanAndSplit(f)
	total := 0

	for fScan.Scan() {
		txt := fScan.Text()

		// find all numbers
		d := re.FindAllIndex([]byte(txt), -1)
		for k := range words {
			if !strings.Contains(txt, k) {
				continue
			}

			firstIndex := strings.Index(txt, k)
			lastIndex := strings.LastIndex(txt, k)
			length := len(k)

			// could optimize by caching length of words
			end := firstIndex + length
			d = append(d, []int{firstIndex, end})

			end = lastIndex + length
			d = append(d, []int{lastIndex, end})
		}

		sort.SliceStable(d, func(i, j int) bool {
			return d[i][0] < d[j][0]
		})

		start := txt[d[0][0]:d[0][1]]
		ms, ok := words[start]
		if ok {
			start = ms
		}

		l := len(d) - 1
		end := txt[d[l][0]:d[l][1]]
		me, ok := words[end]
		if ok {
			end = me
		}

		// convert to numbers
		str := start + end
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalf("Error converting string %v to a number %v", str, err)
		}

		total += num
	}

	fmt.Println(total)
}
