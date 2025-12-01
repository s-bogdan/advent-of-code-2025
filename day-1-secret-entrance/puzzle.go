package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("Zeros 1: %d\n", part1())
	fmt.Printf("Zeros 2: %d\n", part2())
}

func part1() int {

	var (
		pos       = 50
		num       = 0
		direction = 1
		line      = ""
		zeros     = 0
	)

	dir, _ := os.Getwd()
	file, err := os.Open(dir + "/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line = scanner.Text()
		num, _ = strconv.Atoi(line[1:])

		if line[0] == 'R' {
			direction = 1
		} else {
			direction = -1
		}

		pos = (pos + direction*num) % 100

		if pos < 0 {
			pos = 100 + pos
		}

		if pos == 0 {
			zeros++
		}

	}

	return zeros
}

func part2() int {

	var (
		pos       = 50
		num       = 0
		direction = 1
		line      = ""
		zeros     = 0
		startZero = 0
	)

	dir, _ := os.Getwd()
	file, err := os.Open(dir + "/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line = scanner.Text()
		num, _ = strconv.Atoi(line[1:])

		if line[0] == 'R' {
			direction = 1
		} else {
			direction = -1
		}

		if pos == 0 {
			startZero = 0
		} else {
			startZero = 1
		}

		pos += direction * num

		if pos == 0 {
			zeros++
		} else if pos > 0 {
			zeros += pos / 100
		} else {
			zeros += startZero + direction*pos/100
		}

		pos = pos % 100

		if pos < 0 {
			pos = 100 + pos
		}

	}

	return zeros
}
