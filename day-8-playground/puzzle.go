package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())

}

type JunctionBox struct {
	X int
	Y int
	Z int
}
type Distance struct {
	Distance float64
	FromIdx  int
	ToIdx    int
	From     JunctionBox
	To       JunctionBox
}

func part1() int {
	pwd, _ := os.Getwd()
	file, _ := os.Open(pwd + "/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	junctionsBoxes := make([]JunctionBox, 0)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, ",")
		junctionsBoxes = append(junctionsBoxes, JunctionBox{X: number(s[0]), Y: number(s[1]), Z: number(s[2])})
	}
	distances := make([]Distance, 0)
	for i := 0; i < len(junctionsBoxes); i++ {
		for j := i; j < len(junctionsBoxes); j++ {
			distances = append(distances, Distance{Distance: distance(junctionsBoxes[i], junctionsBoxes[j]), FromIdx: i, ToIdx: j, From: junctionsBoxes[i], To: junctionsBoxes[j]})
		}
	}
	distances = slices.DeleteFunc(distances, func(d Distance) bool { return d.Distance == 0 })
	slices.SortFunc(distances,
		func(a, b Distance) int {
			if a.Distance < b.Distance {
				return -1
			} else if a.Distance > b.Distance {
				return 1
			} else {
				return 0
			}
		})

	circuits := make([][]JunctionBox, 0)
	// map with JunctionBox as key, and index of the Circuit containing the JB in the circuits slice as value
	cMap := map[JunctionBox]int{}

	for i := 0; i < 1000; i++ {
		fromIdx, existsFrom := cMap[distances[i].From]
		toIdx, existsTo := cMap[distances[i].To]
		if existsTo && existsFrom {
			if fromIdx == toIdx {
				// both exist and are in the same circuit
				continue
			} else {
				// both exist but in separate circuits
				for _, jb := range circuits[fromIdx] {
					cMap[jb] = toIdx
				}
				circuits[toIdx] = append(circuits[toIdx], circuits[fromIdx]...)
				// empty but keep index, otherwise we need to reindex
				circuits[fromIdx] = make([]JunctionBox, 0)
				// point to the new circuit
				cMap[distances[i].From] = toIdx
			}
		} else if existsFrom {
			circuits[fromIdx] = append(circuits[fromIdx], distances[i].To)
			cMap[distances[i].To] = fromIdx
		} else if existsTo {
			circuits[toIdx] = append(circuits[toIdx], distances[i].From)
			cMap[distances[i].From] = toIdx
		} else {
			newCircuit := make([]JunctionBox, 0)
			newCircuit = append(newCircuit, distances[i].From)
			newCircuit = append(newCircuit, distances[i].To)
			circuits = append(circuits, newCircuit)
			cMap[distances[i].To] = len(circuits) - 1
			cMap[distances[i].From] = len(circuits) - 1
		}
	}

	slices.SortFunc(circuits, func(a, b []JunctionBox) int {
		if len(a) < len(b) {
			return 1
		} else if len(a) > len(b) {
			return -1
		} else {
			return 0
		}
	})

	return len(circuits[0]) * len(circuits[1]) * len(circuits[2])
}

func part2() int {
	pwd, _ := os.Getwd()
	file, _ := os.Open(pwd + "/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	junctionsBoxes := make([]JunctionBox, 0)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, ",")
		junctionsBoxes = append(junctionsBoxes, JunctionBox{X: number(s[0]), Y: number(s[1]), Z: number(s[2])})
	}
	maxJb := len(junctionsBoxes)
	distances := make([]Distance, 0)
	for i := 0; i < len(junctionsBoxes); i++ {
		for j := i; j < len(junctionsBoxes); j++ {
			distances = append(distances, Distance{Distance: distance(junctionsBoxes[i], junctionsBoxes[j]), FromIdx: i, ToIdx: j, From: junctionsBoxes[i], To: junctionsBoxes[j]})
		}
	}
	distances = slices.DeleteFunc(distances, func(d Distance) bool { return d.Distance == 0 })
	slices.SortFunc(distances,
		func(a, b Distance) int {
			if a.Distance < b.Distance {
				return -1
			} else if a.Distance > b.Distance {
				return 1
			} else {
				return 0
			}
		})

	circuits := make([][]JunctionBox, 0)
	// map with JunctionBox as key, and index of the Circuit containing the JB in the circuits slice as value
	cMap := map[JunctionBox]int{}

	for i := 0; i < len(distances); i++ {
		fromIdx, existsFrom := cMap[distances[i].From]
		toIdx, existsTo := cMap[distances[i].To]
		if existsTo && existsFrom {
			if fromIdx == toIdx {
				// both exist and are in the same circuit
				continue
			} else {
				// both exist but in separate circuits
				for _, jb := range circuits[fromIdx] {
					cMap[jb] = toIdx
				}
				circuits[toIdx] = append(circuits[toIdx], circuits[fromIdx]...)
				// empty but keep index, otherwise we need to reindex
				circuits[fromIdx] = make([]JunctionBox, 0)
				// point to the new circuit
				cMap[distances[i].From] = toIdx
			}
		} else if existsFrom {
			circuits[fromIdx] = append(circuits[fromIdx], distances[i].To)
			cMap[distances[i].To] = fromIdx
		} else if existsTo {
			circuits[toIdx] = append(circuits[toIdx], distances[i].From)
			cMap[distances[i].From] = toIdx
		} else {
			newCircuit := make([]JunctionBox, 0)
			newCircuit = append(newCircuit, distances[i].From)
			newCircuit = append(newCircuit, distances[i].To)
			circuits = append(circuits, newCircuit)
			cMap[distances[i].To] = len(circuits) - 1
			cMap[distances[i].From] = len(circuits) - 1
		}
		for j := range circuits {
			if len(circuits[j]) == maxJb {
				return distances[i].From.X * distances[i].To.X
			}
		}
	}

	return 0
}

func number(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
func distance(a JunctionBox, b JunctionBox) float64 {
	return math.Sqrt(float64((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y) + (a.Z-b.Z)*(a.Z-b.Z)))
}
