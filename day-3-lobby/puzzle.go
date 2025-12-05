package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {

	dir, _ := os.Getwd()
	file, err := os.Open(dir + "/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		a, b, c := 0, 0, 0
		for i := 0; i < len(line); i++ {
			c, _ = strconv.Atoi(string(line[i]))
			if c > a && i < (len(line)-1) {
				a = c
				b, _ = strconv.Atoi(string(line[i+1]))
			} else if c > b {
				b = c
			}
		}
		sum += a*10 + b
	}

	return sum
}

func part2() int {
	const maxDigits = 12
	dir, _ := os.Getwd()
	file, err := os.Open(dir + "/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Printf("%s\n", line)
		var currentMax = make([]int, 12)
		for i := 0; i < maxDigits; i++ {
			currentMax[i], _ = strconv.Atoi(string(line[i]))
		}
		for i := maxDigits; i < len(line); i++ {
			n, _ := strconv.Atoi(string(line[i]))
			//fmt.Printf("%d -> %d ->", n, currentMax)

			for k := 0; k < maxDigits-1; k++ {
				if currentMax[k] < currentMax[k+1] {
					currentMax = append(append(currentMax[:k], currentMax[k+1:]...), n)
					//fmt.Printf("%d\n", currentMax)
					break
				}
			}
			if currentMax[maxDigits-1] < n {
				currentMax[maxDigits-1] = n
			}
		}

		temp := ""
		for i := 0; i < maxDigits; i++ {
			temp += strconv.Itoa(currentMax[i])
		}
		//fmt.Printf("%s -> %s\n", line, temp)
		sum += num(temp)
	}

	return sum
}

func num(str string) int {
	s, _ := strconv.Atoi(str)
	return s
}
