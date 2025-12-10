package main

import (
	"bufio"
	"container/list"
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

type Coordinate struct {
	X int
	Y int
}

func part1() int {
	pwd, _ := os.Getwd()
	file, _ := os.Open(pwd + "/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	coordinates := make([]Coordinate, 0)
	for scanner.Scan() {
		line := scanner.Text()
		c := getCoordinate(line)
		coordinates = append(coordinates, c)
	}

	m := 0
	for i := 0; i < len(coordinates); i++ {
		for j := i; j < len(coordinates); j++ {
			m = max(m, int((math.Abs(float64(coordinates[i].X-coordinates[j].X))+1)*(math.Abs(float64(coordinates[i].Y-coordinates[j].Y))+1)))
		}
	}

	return m
}

var (
	grid     [][]bool
	xAxisMap map[int][]int
	yAxisMap map[int][]int
)

func part2() int {
	pwd, _ := os.Getwd()
	file, _ := os.Open(pwd + "/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	coord := make([]Coordinate, 0)

	maxX, maxY := 0, 0
	minX, minY := 100000, 1000000
	for scanner.Scan() {
		line := scanner.Text()
		c1 := getCoordinate(line)
		maxX = max(maxX, c1.X)
		maxY = max(maxY, c1.Y)
		minX = min(minX, c1.X)
		minY = min(minY, c1.Y)
		coord = append(coord, c1)
	}
	// add again so we can close the loop when we mark the lines
	coord = append(coord, coord[0])

	sqSize := max(maxX+1, maxY+1)
	for i := 0; i < sqSize; i++ {
		grid = append(grid, make([]bool, sqSize))
	}

	queue := list.New()
	for i := 0; i < len(coord)-1; i++ {
		grid[coord[i].X][coord[i].Y] = true
		queue.PushBack(Coordinate{coord[i].X, coord[i].Y})
		grid[coord[i+1].X][coord[i+1].Y] = true
		queue.PushBack(Coordinate{coord[i+1].X, coord[i+1].Y})

		if coord[i].X == coord[i+1].X {
			miy := min(coord[i].Y, coord[i+1].Y)
			mxy := max(coord[i].Y, coord[i+1].Y)
			for j := miy; j < mxy; j++ {
				grid[coord[i].X][j] = true
				queue.PushBack(Coordinate{coord[i+1].X, j})
			}
		} else {
			mix := min(coord[i].X, coord[i+1].X)
			mxx := max(coord[i].X, coord[i+1].X)
			for j := mix; j < mxx; j++ {
				grid[j][coord[i].Y] = true
				queue.PushBack(Coordinate{j, coord[i].Y})
			}
		}
	}

	xAxisMap = make(map[int][]int)
	yAxisMap = make(map[int][]int)
	for {
		front := queue.Front()
		if front == nil {
			break
		}
		c := front.Value.(Coordinate)
		xAxisMap[c.X] = append(xAxisMap[c.X], c.Y)
		yAxisMap[c.Y] = append(yAxisMap[c.Y], c.X)
		queue.Remove(front)
	}
	for _, v := range xAxisMap {
		slices.Sort(v)
	}
	for _, v := range yAxisMap {
		slices.Sort(v)
	}

	//printGrid(grid, coord)
	m := 0
	for i := 0; i < len(coord); i++ {
		for j := i; j < len(coord); j++ {
			curr := int((math.Abs(float64(coord[i].X-coord[j].X)) + 1) * (math.Abs(float64(coord[i].Y-coord[j].Y)) + 1))
			if m < curr && isValidRectangle(coord[i], coord[j]) {
				m = curr
			}
		}
	}

	return m
}

func isValidRectangle(a, b Coordinate) bool {
	mix := min(a.X, b.X)
	mxx := max(a.X, b.X)
	miy := min(a.Y, b.Y)
	mxy := max(a.Y, b.Y)

	// point is either on a line, or between 2 points on both the X and Y axis
	for i := mix; i < mxx; i++ {
		if !((grid[i][a.Y] && grid[i][b.Y]) ||
			(pointInside(Coordinate{X: i, Y: a.Y}) && pointInside(Coordinate{i, b.Y}))) {
			return false
		}
	}
	for i := miy; i < mxy; i++ {
		if !((grid[a.X][i] && grid[b.X][i]) ||
			(pointInside(Coordinate{a.X, i}) && pointInside(Coordinate{b.X, i}))) {
			return false
		}
	}
	return true
}
func pointInside(a Coordinate) bool {
	xAxis := xAxisMap[a.X]
	yAxis := yAxisMap[a.Y]
	return xAxis[0] <= a.Y && a.Y <= xAxis[len(xAxis)-1] && yAxis[0] <= a.X && a.X <= yAxis[len(yAxis)-1]
}

func getCoordinate(s string) Coordinate {
	c := strings.Split(s, ",")
	x, _ := strconv.Atoi(c[0])
	y, _ := strconv.Atoi(c[1])
	return Coordinate{x, y}
}
func printGrid(grid [][]bool, c []Coordinate) {
	mp := map[Coordinate]bool{}
	for i := 0; i < len(c); i++ {
		mp[c[i]] = true
	}
	sb := strings.Builder{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if _, ok := mp[Coordinate{j, i}]; ok {
				sb.WriteRune('#')
			} else if grid[j][i] {
				sb.WriteRune('X')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteString("\n")
	}
	file, _ := os.Create("debug-1.txt")
	defer file.Close()
	_, _ = file.WriteString(sb.String())
}
