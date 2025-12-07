package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Result part 1: %d\n", part1())
	fmt.Printf("Result part 2: %d\n", part2())
}

func part1() int {
	pwd, _ := os.Getwd()
	file, _ := os.Open(pwd + "/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	beamMap := map[int]bool{}
	splits := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			if line[i] == 'S' {
				beamMap[i] = true
				break
			}
			// beam hits splitter
			if line[i] == '^' && beamMap[i] {
				splits++
				beamMap[i] = false
				beamMap[i-1] = true
				beamMap[i+1] = true
			}
		}
	}
	return splits
}

func part2() int {
	pwd, _ := os.Getwd()
	file, _ := os.Open(pwd + "/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	// Map with line index as key, and number of possible timelines as value
	beamMap := map[int]int{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		for i := 0; i < len(line); i++ {
			if line[i] == 'S' {
				beamMap[i] = 1
				break
			}
			// beam hits splitter
			if _, exists := beamMap[i]; line[i] == '^' && exists {
				delete(beamMap, i)
				beamMap[i-1] = 1
				beamMap[i+1] = 1
			}
		}
	}

	result := 0
	// read bottom up and calculate number of possible timelines
	for i := len(lines) - 1; i >= 0; i-- {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == '^' {
				beamMap[j] = beamMap[j-1] + beamMap[j+1]
			}
			if lines[i][j] == 'S' {
				result = beamMap[j]
			}
		}
	}

	return result
}
