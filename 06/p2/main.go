package main

import (
	"adventofcode/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	p := utils.GetFilePath(os.Args)
	f, fClose := utils.MustOpenFile(p)
	defer fClose(f)

	s := utils.ScanAndSplit(f)
	m := make(map[string]int)
	for s.Scan() {
		td := strings.Split(s.Text(), ":")
		tag, v := td[0], td[1]
		values := strings.Fields(v)

		var sb strings.Builder
		for _, vs := range values {
			sb.WriteString(vs)
		}

		num, err := strconv.Atoi(sb.String())
		if err != nil {
			panic(err)
		}

		m[tag] = num
	}

	total := processTimeAndDistance(m["Time"], m["Distance"])
	fmt.Println(total)
}

func processTimeAndDistance(t int, d int) int {
	acc := 0
	for i := 0; i < t; i++ {
		nd := i * (t - i)
		if nd <= d {
			continue
		}

		acc += 1

	}

	return acc
}
