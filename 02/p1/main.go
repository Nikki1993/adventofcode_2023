package main

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var Reset = "\033[0m"
var Green = "\033[32m"
var Red = "\033[31m"

type cubes struct {
	red   int
	green int
	blue  int
}

func main() {
	fPath := utils.GetFilePath(os.Args)
	f, fClose := utils.MustOpenFile(fPath)
	defer fClose(f)
	c := cubes{
		red:   12,
		green: 13,
		blue:  14,
	}

	fScan := utils.ScanAndSplit(f)
	total := 0

	for fScan.Scan() {
		id, ok := parseGame(c, fScan.Text())
		if !ok {
			continue
		}

		total += id
	}

	fmt.Println()
	fmt.Println(" Sum of all IDs:", total)
}

func parseGame(cubes cubes, line string) (int, bool) {
	// separate game id
	v := strings.Split(line, ":")

	// get game ID as a string
	g := strings.Fields(v[0])
	gID, err := strconv.Atoi(g[1])
	if err != nil {
		log.Fatalf("Could not conver game %v to number %v", g, err)
	}

	// separate games string into separate items
	gtxt := v[1]
	gRounds := strings.Split(gtxt, ";")
	for _, r := range gRounds {
		cs := strings.Split(r, ",")
		for _, c := range cs {
			c := strings.Fields(c)

			cube := c[1]
			amount, err := strconv.Atoi(c[0])
			if err != nil {
				log.Fatalf("Error converting %v to an int %v", c[1], err)
			}

			switch cube {
			case "red":
				// returns false for positive numbers
				if math.Signbit(float64(cubes.red - amount)) {
					fmt.Println(Red, "ID", gID, Reset, "Remaining cubes", cubes, "Game Text", gtxt)
					return gID, false
				}
			case "green":
				if math.Signbit(float64(cubes.green - amount)) {
					fmt.Println(Red, "ID", gID, Reset, "Remaining cubes", cubes, "Game Text", gtxt)
					return gID, false
				}
			case "blue":
				if math.Signbit(float64(cubes.blue - amount)) {
					fmt.Println(Red, "ID", gID, Reset, "Remaining cubes", cubes, "Game Text", gtxt)
					return gID, false
				}
			default:
				log.Fatalf("Got unknown case %v", cube)
			}
		}
	}

	fmt.Println(Green, "ID", gID, Reset, "Remaining cubes", cubes, "Game Text", gtxt)
	return gID, true
}
