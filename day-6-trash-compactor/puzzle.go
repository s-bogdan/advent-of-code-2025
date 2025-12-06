package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Result part 1: %d\n", part1())
	fmt.Printf("Result part 2: %d\n", part2())
}
func part1() int {
	problems := loadProblemList()
	result := 0
	for _, i := range problems {
		result += i.solveProblem()
	}
	return result
}
func part2() int {
	dir, _ := os.Getwd()
	file, _ := os.Open(dir + "/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}
	result := 0
	currentSolution := 0
	currentOperator := ""
	for i := 0; i < len(lines[0]); i++ {
		s := string(lines[len(lines)-1][i])
		if s == "+" || s == "*" {
			fmt.Printf(" = %d\n", currentSolution)
			result += currentSolution
			currentOperator = s
			if s == "*" {
				currentSolution = 1
			} else {
				currentSolution = 0
			}
		}
		n := getRtlNum(lines, i)
		if n == 0 {
			// empty row
			continue
		}
		if currentOperator == "*" {
			fmt.Printf(" * %d", n)
			currentSolution *= n
		} else {
			currentSolution += n
			fmt.Printf(" + %d", n)
		}
	}
	result += currentSolution
	return result
}

// + 1 2 3 4 => 4321
func getRtlNum(lines []string, index int) int {
	s := ""
	for i := 0; i < len(lines)-1; i++ {
		s += string(lines[i][index])
	}
	n, _ := strconv.Atoi(strings.TrimSpace(s))
	return n
}

func loadProblemList() []Problem {
	dir, _ := os.Getwd()
	file, _ := os.Open(dir + "/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	grid := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, " ")
		row = slices.DeleteFunc(row, func(s string) bool { return strings.Trim(s, " ") == "" })
		grid = append(grid, row)
	}
	printGrid(grid)
	problems := make([]Problem, 0)
	for i := 0; i < len(grid[0]); i++ {
		p := Problem{}
		p.Operands = make([]int, 0)
		for j := 0; j < len(grid)-1; j++ {
			op, _ := strconv.Atoi(grid[j][i])
			p.Operands = append(p.Operands, op)
		}
		p.Operator = grid[len(grid)-1][i]

		fmt.Printf("%v\n", p)
		problems = append(problems, p)
	}
	return problems
}

type Problem struct {
	Operands []int
	Operator string
}

func (p *Problem) solveProblem() int {
	result := 0
	if p.Operator == "*" {
		result = 1
	}
	for i := 0; i < len(p.Operands); i++ {
		if p.Operator == "+" {
			result += p.Operands[i]
		} else {
			result *= p.Operands[i]
		}
	}
	fmt.Printf("%v %v -> %v\n", p.Operands, p.Operator, result)
	return result
}

func printGrid(grid [][]string) {
	for i := 0; i < len(grid); i++ {
		fmt.Printf("[ ")
		for j := 0; j < len(grid[0]); j++ {
			fmt.Printf("%v ", grid[i][j])
		}
		fmt.Printf("]\n")
	}
}
