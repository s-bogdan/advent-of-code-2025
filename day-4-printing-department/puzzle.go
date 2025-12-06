package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Printf("Result part 1: %d", part1())
	fmt.Printf("Result part 2: %d", part2())
}
func part1() int {
	dir, _ := os.Getwd()
	file, err := os.Open(dir + "/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]int
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for j := 0; j < len(line); j++ {
			if line[j] == '@' {
				row[j] = 1
			} else {
				row[j] = 0
			}
		}
		row = append([]int{0}, row...)
		row = append(row, 0)
		grid = append(grid, row)
	}

	emptyRow := make([]int, len(grid[0]))
	for i := 0; i < len(grid[0]); i++ {
		emptyRow[i] = 0
	}

	grid = append([][]int{0: emptyRow}, grid...)
	grid = append(grid, emptyRow)

	printGrid(grid)
	fmt.Print("\n--------\n")
	result := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			if s := sum(grid[i+1][j], grid[i-1][j], grid[i][j+1], grid[i][j-1], grid[i+1][j+1], grid[i-1][j-1], grid[i-1][j+1], grid[i+1][j-1]); grid[i][j] == 1 && s <= 3 {
				result++
				mark9(&grid[i][j])
			}
		}
	}
	dumpToFile(grid)
	printGrid(grid)

	return result
}

func part2() int {
	dir, _ := os.Getwd()
	file, err := os.Open(dir + "/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]int
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for j := 0; j < len(line); j++ {
			if line[j] == '@' {
				row[j] = 1
			} else {
				row[j] = 0
			}
		}
		row = append([]int{0}, row...)
		row = append(row, 0)
		grid = append(grid, row)
	}

	emptyRow := make([]int, len(grid[0]))
	for i := 0; i < len(grid[0]); i++ {
		emptyRow[i] = 0
	}

	grid = append([][]int{0: emptyRow}, grid...)
	grid = append(grid, emptyRow)

	printGrid(grid)
	fmt.Print("\n--------\n")
	result := 0
	for removed := true; removed == true; {
		removed = false
		for i := 1; i < len(grid)-1; i++ {
			for j := 1; j < len(grid[0])-1; j++ {
				if s := sum(grid[i+1][j], grid[i-1][j], grid[i][j+1], grid[i][j-1], grid[i+1][j+1], grid[i-1][j-1], grid[i-1][j+1], grid[i+1][j-1]); grid[i][j] == 1 && s <= 3 {
					result++
					removed = true
					mark0(&grid[i][j])
				}
			}
		}
	}
	dumpToFile(grid)
	printGrid(grid)

	return result
}

func sum(nums ...int) int {
	s := 0
	for _, n := range nums {
		// 9 is already a roll marked taken
		if n == 9 {
			s++
		} else {
			s += n
		}

	}
	return s
}

// taken
func mark9(nums ...*int) {
	for _, n := range nums {
		if *n == 1 {
			*n = 9
		}
	}
}
func mark0(nums ...*int) {
	for _, n := range nums {
		if *n == 1 {
			*n = 0
		}
	}
}
func printGrid(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		fmt.Printf("[ ")
		for j := 0; j < len(grid[0]); j++ {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Printf("]\n")
	}
}

func dumpToFile(grid [][]int) {
	file, _ := os.Create("debug-1.txt")
	for i := 0; i < len(grid); i++ {
		file.WriteString("[ ")
		for j := 0; j < len(grid[0]); j++ {
			file.WriteString(fmt.Sprintf("%d ", grid[i][j]))
		}
		file.WriteString("]\n")
	}
	file.Close()
}
