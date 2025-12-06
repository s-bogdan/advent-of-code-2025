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

type Database struct {
	FreshIngredientRange []Range
	Ingredients          []int
}
type Range struct {
	Start int
	End   int
}

func readDb() Database {
	dir, _ := os.Getwd()
	file, _ := os.Open(dir + "/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	db := Database{}
	db.FreshIngredientRange = make([]Range, 0)
	db.Ingredients = make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		data := strings.Split(line, "-")
		if len(data) == 2 {
			db.FreshIngredientRange = append(db.FreshIngredientRange, Range{Start: toNumber(data[0]), End: toNumber(data[1])})
		} else {
			db.Ingredients = append(db.Ingredients, toNumber(data[0]))
		}
	}

	return db
}
func toNumber(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
func part1() int {
	db := readDb()
	result := 0
	for _, i := range db.Ingredients {
		for _, r := range db.FreshIngredientRange {
			if i >= r.Start && i < r.End {
				result++
				break
			}
		}
	}
	return result
}

func part2() int {
	db := readDb()
	// within, outside, intersect
	vec := make([]Range, 0)
	for _, i := range db.FreshIngredientRange {
		for j := 0; j < len(vec); j++ {
			if vec[j].contains(i) {
				break
			}
			// new range contains smaller range
			if i.contains(vec[j]) {
				vec[j] = i
				break
			}
			// ranges overlaps
			if vec[j].overlaps(i) {
				vec[j].intersect(i)
				break
			}
		}
		// new range
		vec = append(vec, i)
	}

	fmt.Print("\n Unsorted ---------------\n")
	printSlice(vec)
	fmt.Print("\n---------------\n")

	slices.SortFunc(vec, func(i, j Range) int {
		return i.Start - j.Start
	})

	fmt.Print("\n Sorted ---------------\n")
	printSlice(vec)
	fmt.Print("\n---------------\n")

	// compact
	for compacted := true; compacted; {
		compacted = false
		for i := 0; i < len(vec)-1; i++ {
			if vec[i].contains(vec[i+1]) {
				fmt.Printf("%v contains %v\n", vec[i], vec[i+1])
				if i+1 == len(vec)-1 {
					vec = vec[:i+1]
				} else {
					vec = slices.Delete(vec, i+1, i+2)
				}
				compacted = true
			} else if vec[i+1].contains(vec[i]) {
				fmt.Printf("%v contains %v\n", vec[i+1], vec[i])
				if i == 0 {
					vec = vec[i+1:]
				} else {
					vec = slices.Delete(vec, i, i+1)
				}
				compacted = true
			} else if vec[i].overlaps(vec[i+1]) {
				fmt.Printf("%v overlaps %v -> %v\n", vec[i+1], vec[i], vec[i].intersect(vec[i+1]))
				if i == len(vec)-1-1 {
					vec[i] = vec[i].intersect(vec[i+1])
					// delete last element
					vec = vec[:i-1]
				} else {
					vec[i+1] = vec[i].intersect(vec[i+1])
					vec = slices.Delete(vec, i, i+1)
				}
				compacted = true
			}
			if compacted {
				break
			}
		}
	}

	fmt.Print("\n Compacted ---------------\n")
	printSlice(vec)
	fmt.Print("\n---------------\n")

	result := 0
	for _, i := range vec {
		result += i.End - i.Start + 1
	}

	return result
}

func (r *Range) contains(r1 Range) bool {
	return r1.Start >= r.Start && r1.End <= r.End
}
func (r *Range) overlaps(r1 Range) bool {
	return (r1.Start >= r.Start && r1.Start <= r.End && r1.End > r.End) || (r1.Start < r.Start && r1.End >= r.Start && r1.End < r.End)
}
func (r *Range) intersect(other Range) Range {
	return Range{Start: min(r.Start, other.Start), End: max(r.End, other.End)}
}

func printSlice(s []Range) {
	for _, i := range s {
		fmt.Printf("[ %d - %d ]\n", i.Start, i.End)
	}
}
