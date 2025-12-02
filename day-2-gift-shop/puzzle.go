package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	a, b int
	s    string
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	dir, _ := os.Getwd()
	input, err := os.ReadFile(dir + "/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	ranges := strings.Split(string(input), ",")

	sum := 0
	for _, r := range ranges {
		p := strings.Split(r, "-")
		a, _ = strconv.Atoi(p[0])
		b, _ = strconv.Atoi(p[1])
		for i := a; i <= b; i++ {
			s = strconv.Itoa(i)
			if len(s)%2 != 0 {
				continue
			} else if s[:len(s)/2] == s[len(s)/2:] {
				sum += i
			}
		}
	}
	return sum
}

func part2() int {
	dir, _ := os.Getwd()
	input, err := os.ReadFile(dir + "/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	ranges := strings.Split(string(input), ",")

	sum := 0
	invalid := true
	for _, r := range ranges {
		p := strings.Split(r, "-")
		a, _ = strconv.Atoi(p[0])
		b, _ = strconv.Atoi(p[1])
		for i := a; i <= b; i++ {
			s = strconv.Itoa(i)
			//fmt.Printf("%s\n", s)
			invalid = true
			for step := 1; step <= len(s)/2; step++ {
				if len(s)%step != 0 {
					continue
				}
				for k := 0; k < len(s)/step-1; k++ {
					if s[k*step:k*step+step] != s[k*step+step:k*step+step+step] {
						invalid = false
						break
					} else {
						invalid = true
					}
				}
				if invalid {
					sum += i
					break
				}
			}
		}
	}
	return sum
}
