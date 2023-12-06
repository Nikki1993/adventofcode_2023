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
	m := make(map[string][]int)
	for s.Scan() {
		td := strings.Split(s.Text(), ":")
		tag, v := td[0], td[1]
		values := strings.Fields(v)

		vi := make([]int, 0, len(values))
		for _, vs := range values {
			i, err := strconv.Atoi(vs)
			if err != nil {
				panic(err)
			}

			vi = append(vi, i)
		}

		m[tag] = vi
	}

	gNum := len(m["Time"])
	ch := make(chan int, gNum)

	for i := 0; i < gNum; i++ {
		i := i
		go processTimeAndDistance(ch, m["Time"][i], m["Distance"][i])
	}

	total := 0
	for i := 0; i < gNum; i++ {
		msg := <-ch
		if total == 0 {
			total = msg
			continue
		}

		total *= msg
	}

	close(ch)
	fmt.Println(total)
}

func processTimeAndDistance(ch chan<- int, t int, d int) {
	acc := 0
	for i := 0; i < t; i++ {
		nd := i * (t - i)
		if nd <= d {
			continue
		}

		acc += 1

	}

	ch <- acc
}
