package main

import (
	"bufio"
	"container/list"
	"fmt"
	"maps"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Result part 1 %d\n", part1())
	//fmt.Printf("Result part 2 %d\n", part2())
}

type Machine struct {
	Lights     []bool
	Light      int
	Buttons    [][]bool
	ButtonList []int
	ButtonMap  map[int]ButtonList
	Counters   []int
}
type Sequence struct {
	ButtonList map[int]bool
	Xored      int
}
type ButtonList struct {
	Buttons []Button
}
type Button struct {
	OnPosition []bool
}

func (b *Button) isOn(pos int) bool {
	return b.OnPosition[pos]
}

type Possibility struct {
	Size    int
	Counter []int
	Valid   bool
	Found   bool
}

func parse() []Machine {
	pwd, _ := os.Getwd()
	file, _ := os.Open(pwd + "/input.txt")
	//file, _ := os.Open(pwd + "/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var machines []Machine
	for scanner.Scan() {
		line := scanner.Text()
		m := Machine{}
		l := strings.Split(line, " ")
		l0 := strings.Trim(l[0], "[]")
		size := len(l0)
		m.Lights = make([]bool, len(l0))
		for i := 0; i < size; i++ {
			if l0[i] == '#' {
				m.Lights[i] = true
				m.Light += pow2(size - 1 - i)
			}
		}
		for i := 1; i < len(l)-1; i++ {
			button := make([]bool, size)
			but := 0
			for _, s := range strings.Split(strings.Trim(l[i], "()"), ",") {
				n, _ := strconv.Atoi(s)
				button[n] = true
				but += pow2(size - 1 - n)
			}
			m.Buttons = append(m.Buttons, button)
			m.ButtonList = append(m.ButtonList, but)
		}

		for _, s := range strings.Split(strings.Trim(l[len(l)-1], "{}"), ",") {
			n, _ := strconv.Atoi(s)
			m.Counters = append(m.Counters, n)
		}

		machines = append(machines, m)
	}
	return machines
}
func part1() int {
	machines := parse()
	result := 0
	for i := 0; i < len(machines); i++ {
		queue := list.New()
		for j := 0; j < len(machines[i].ButtonList); j++ {
			queue.PushBack(Sequence{map[int]bool{machines[i].ButtonList[j]: true}, machines[i].ButtonList[j]})
		}
		for {
			front := queue.Front()
			if front == nil {
				break
			}
			seq := front.Value.(Sequence)

			if machines[i].Light == seq.Xored {
				fmt.Printf("%d,%d\n", i, len(seq.ButtonList))
				result += len(seq.ButtonList)
				break
			} else {
				for _, n := range machines[i].ButtonList {
					if _, ok := seq.ButtonList[n]; !ok {
						temp := maps.Clone(seq.ButtonList)
						temp[n] = true
						s := Sequence{ButtonList: temp, Xored: seq.Xored ^ n}
						queue.PushBack(s)
					}
				}
			}
			queue.Remove(front)
		}
	}

	return result
}
func part2() int {
	machines := parse()
	result := 0
	fmt.Printf("Start-------\n")
	for i := 0; i < len(machines); i++ {
		minButtons := math.MaxInt
		cCounter := slices.Clone(machines[i].Counters)
		//fmt.Printf("\n\nButtons: %v\n", machines[i].Buttons)
		fmt.Printf("Lights [%v]: %v -> ", i, machines[i].Counters)
		possibilityQueue := list.New()
		visited := make(map[string]bool)
		// start from first position
		for j := 0; j < len(machines[i].Buttons); j++ {
			cCounter = slices.Clone(machines[i].Counters)
			if _, ok := visited[identity(machines[i].Buttons[j], cCounter)]; !ok && isValid(machines[i].Buttons[j], cCounter) {
				//fmt.Printf("%v: %v - %v -> %v\n", 1, cCounter, prettyPrint(machines[i].Buttons[j]), diff(machines[i].Buttons[j], slices.Clone(cCounter)))
				possibilityQueue.PushBack(Possibility{Size: 1,
					Counter: diff(machines[i].Buttons[j], cCounter)})
			}
		}
		visited[identity2(cCounter)] = true

		found := false
		for found != true {
			front := possibilityQueue.Front()
			if front == nil {
				break
			}
			pos := front.Value.(Possibility)
			//fmt.Printf("%v\n", pos)
			for j := 0; j < len(machines[i].Buttons); j++ {
				cCounter = slices.Clone(pos.Counter)
				if _, ok := visited[identity(machines[i].Buttons[j], cCounter)]; !ok {
					p := diff2(machines[i].Buttons[j], pos, cCounter)
					if !p.Valid {
						continue
					} else if p.Found {
						found = true
						minButtons = min(minButtons, pos.Size)
						break
					} else {
						possibilityQueue.PushBack(p)
					}
				}
			}
			visited[identity2(cCounter)] = true

			possibilityQueue.Remove(front)
		}
		fmt.Printf("%v (from %v possibilities)\n", minButtons, len(visited))
		result += minButtons
	}
	//fmt.Println(machines)
	return result
}
func pow2(p int) int {
	return int(math.Pow(2, float64(p)))
}
func printMapValues(m map[int]bool) string {
	s := strings.Builder{}
	s.WriteString("[")
	for k, _ := range m {
		s.WriteString(fmt.Sprintf("%d,", k))
	}
	s.WriteString("]")
	return s.String()
}

//	func combine(b ...[]bool) []bool {
//		result := make([]bool, len(b))
//		for i := 0; i < len(b); i++ {
//
//		}
//	}
func isValidNum(a int, nums ...int) bool {
	if len(nums) == 1 {
		return a == nums[0]
	} else {
		xored := 0
		for i := 0; i < len(nums); i++ {
			xored = xored ^ nums[i]
		}
		return a == xored
	}
}
func identity2(counter []int) string {
	sb := strings.Builder{}
	for i := 0; i < len(counter); i++ {
		sb.WriteString(strconv.Itoa(counter[i]) + "-")
	}
	return sb.String()
}
func identity(buttons []bool, counter []int) string {
	sb := strings.Builder{}
	for i := 0; i < len(counter); i++ {
		sb.WriteString(strconv.Itoa(counter[i]))
	}
	for _, b := range buttons {
		if b {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
	}
	return sb.String()
}
func prettyPrint(b []bool) string {
	sb := strings.Builder{}
	sb.WriteString("[")
	for i := 0; i < len(b); i++ {
		if b[i] {
			sb.WriteString("1 ")
		} else {
			sb.WriteString("0 ")
		}
	}
	return strings.TrimRight(sb.String(), " ") + "]"
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func notZeroAt(a []int, pos int) bool {
	return a[pos] != 0
}
func allZero(a []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != 0 {
			return false
		}
	}
	return true
}
func isValid(b []bool, counter []int) bool {
	for i := 0; i < len(counter); i++ {
		if b[i] && counter[i]-1 < 0 {
			return false
		}
	}
	return true
}
func diff(b []bool, counter []int) []int {
	for i := 0; i < len(counter); i++ {
		if b[i] {
			counter[i]--
		}
	}
	return counter
}
func diff2(b []bool, p Possibility, counters []int) Possibility {
	zeros := true
	found := false
	valid := true
	for i := 0; i < len(p.Counter); i++ {
		if b[i] {
			p.Counter[i]--
		}
		if p.Counter[i] < 0 {
			valid = false
			break
		}
		if !zeros || p.Counter[i] > 0 {
			zeros = false
		}
	}
	if p.Valid && zeros {
		found = true
	}
	return Possibility{Size: p.Size + 1, Counter: counters, Valid: valid, Found: found}
}
func sumInts(a [][]bool) string {
	res := make([]int, len(a[0]))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] {
				res[j]++
			}
		}
	}
	return fmt.Sprintf("%v", res)
}
func dumpToFile(m []Machine) {
	file, _ := os.Create("debug.txt")
	sb := strings.Builder{}
	for i := 0; i < len(m); i++ {
		sb.WriteString(fmt.Sprintf("%v - %v\n", m[i].Counters, sumInts(m[i].Buttons)))
	}
	file.WriteString(sb.String())
	file.Close()
}
