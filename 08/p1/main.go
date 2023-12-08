package main

import (
	"adventofcode/utils"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	p := utils.GetFilePath(os.Args)
	f, fClose := utils.MustOpenFile(p)
	re := regexp.MustCompile(`\(.*?\)`)
	defer fClose(f)

	s := utils.ScanAndSplit(f)
	m := make(map[string][2]string)
	var txtSlice []string
	for s.Scan() {
		txtSlice = append(txtSlice, s.Text())
	}

	var pattern []int
	for _, r := range txtSlice[0] {
		d := 0
		if r == 'R' {
			d = 1
		}

		pattern = append(pattern, d)
	}

	length := len(txtSlice)
	current, end := "AAA", "ZZZ"
	// skipping 1 as it's a blank line
	for i := 2; i < length; i++ {
		data := strings.Split(txtSlice[i], "=")
		step, direction := strings.TrimSpace(data[0]), strings.TrimSpace(data[1])
		directionString := re.FindString(direction)
		directionString = directionString[1 : len(directionString)-1]
		paths := strings.Split(directionString, ", ")

		m[step] = [2]string{paths[0], paths[1]}
	}

	steps := 0
	match := false
	for !match {
		fmt.Println(pattern)
		for _, v := range pattern {
			fmt.Println("CURRENT", m[current])
			if v == 0 {
				fmt.Println("TURNING LEFT", m[current][v])
			} else {
				fmt.Println("TURNING RIGHT", m[current][v])
			}
			steps++
			current = m[current][v]

			if current != end {
				continue
			}

			match = true
			break
		}

	}

	fmt.Println("CURRENT", current, "END", end)
	fmt.Println("Number of steps taken", steps)
}
