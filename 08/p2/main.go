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
	multiCurrent := make(map[string]int)
	// skipping 1 as it's a blank line
	for i := 2; i < length; i++ {
		data := strings.Split(txtSlice[i], "=")
		step, direction := strings.TrimSpace(data[0]), strings.TrimSpace(data[1])
		directionString := re.FindString(direction)
		directionString = directionString[1 : len(directionString)-1]
		paths := strings.Split(directionString, ", ")

		if step[2] == 'A' {
			multiCurrent[step] = 0
		}

		m[step] = [2]string{paths[0], paths[1]}
	}

	ch := make(chan int)

	for k := range multiCurrent {
		go getSteps(pattern, k, m, ch)
	}

	num := 0
	for i := 0; i < len(multiCurrent); i++ {
		if num == 0 {
			num = <-ch
			continue
		}

		num = lcm(num, <-ch)
	}

	close(ch)
	fmt.Println(num)
}

func getSteps(pattern []int, k string, m map[string][2]string, ch chan<- int) {
	steps := 0
	match := false
	for !match {
		for _, v := range pattern {
			steps++
			k = m[k][v]

			if rune(k[2]) != 'Z' {
				continue
			}

			match = true
			break
		}
	}

	ch <- steps
}

// study up on this shit. It's still going over my head
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func lcm(a, b int) int {
	fmt.Println(a, b, gcd(a, b))
	return a * b / gcd(a, b)
}
