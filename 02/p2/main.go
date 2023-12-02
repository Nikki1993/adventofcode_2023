package main

import (
	"bufio"
	"fmt"
	"log"
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
	fPath := "input.txt"
	if len(os.Args) > 1 {
		fPath = os.Args[1]
	}

	c := cubes{
		red:   0,
		green: 0,
		blue:  0,
	}

	f, err := os.Open(fPath)
	if err != nil {
		log.Fatalf("Error opening a fPath %v", err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Printf("Error closing the fPath %v", err)
		}
	}(f)

	fScan := bufio.NewScanner(f)
	fScan.Split(bufio.ScanLines)
	total := 0

	for fScan.Scan() {
		cubes := parseGame(c, fScan.Text())
		total += cubes.red * cubes.green * cubes.blue
	}

	fmt.Println(" Sum of all cube powers:", total)
}

func parseGame(cubes cubes, line string) cubes {
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
				if cubes.red > amount {
					continue
				}

				cubes.red = amount
			case "green":
				if cubes.green > amount {
					continue
				}

				cubes.green = amount
			case "blue":
				if cubes.blue > amount {
					continue
				}

				cubes.blue = amount
			default:
				log.Fatalf("Got unknown case %v", cube)
			}
		}
	}

	fmt.Println("ID", gID, "Least amount needed", cubes, "Game Text", gtxt)
	return cubes
}
