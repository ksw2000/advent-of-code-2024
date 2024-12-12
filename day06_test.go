package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// const day6Part1Example = `....#.....
// .........#
// ..........
// ..#.......
// .......#..
// ..........
// .#..^.....
// ........#.
// #.........
// ......#...`

func TestDay6Part1(t *testing.T) {
	f, err := os.Open("data/day6.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// f := strings.NewReader(day6Part1Example)

	var u, v int
	scanner := bufio.NewScanner(f)
	table := [][]byte{}
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		table = append(table, []byte(line))
		for i := range line {
			if line[i] == '^' {
				u, v = row, i
			}
		}
		row++
	}

	rows := len(table)
	cols := len(table[0])
	for {
		for u >= 0 && table[u][v] != '#' {
			table[u][v] = 'x'
			u--
		}
		if u < 0 {
			break
		}
		u++
		for v < cols && table[u][v] != '#' {
			table[u][v] = 'x'
			v++
		}
		if v >= cols {
			break
		}
		v--
		for u < rows && table[u][v] != '#' {
			table[u][v] = 'x'
			u++
		}
		if u >= rows {
			break
		}
		u--
		for v >= 0 && table[u][v] != '#' {
			table[u][v] = 'x'
			v--
		}
		if v < 0 {
			break
		}
		v++
	}

	sum := 0
	for i := range table {
		for j := range table[i] {
			if table[i][j] == 'x' {
				sum++
			}
		}
	}

	fmt.Println(sum)
}

const (
	markU = uint8(1) << iota
	markR
	markD
	markL
	markO
)

func TestDay6Part2(t *testing.T) {
	f, err := os.Open("data/day6.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// f := strings.NewReader(day6Part1Example)

	start := time.Now()

	var u, v int
	scanner := bufio.NewScanner(f)
	mark := [][]uint8{}

	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		mark = append(mark, make([]uint8, len(line)))
		for i := range line {
			if line[i] == '^' {
				u, v = row, i
			} else if line[i] == '#' {
				mark[row][i] = markO
			}
		}
		row++
	}

	mark2 := make([][]uint8, len(mark))
	for i := range mark {
		mark2[i] = make([]uint8, len(mark[i]))
	}

	sum := 0
	for i := range mark {
		for j := range mark[i] {
			if mark[i][j] == 0 && !(i == u && j == v) {
				for i := range mark {
					copy(mark2[i], mark[i])
				}
				mark2[i][j] = markO
				if isCycle(mark2, u, v) {
					sum++
				}
			}
		}
	}

	fmt.Printf("%.3f ms\n", float64(time.Since(start).Microseconds())*0.001)
	fmt.Println(sum)
}

func TestDay6Part2Goroutine(t *testing.T) {
	f, err := os.Open("data/day6.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// f := strings.NewReader(day6Part1Example)

	start := time.Now()

	var u, v int
	scanner := bufio.NewScanner(f)
	mark := [][]uint8{}

	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		mark = append(mark, make([]uint8, len(line)))
		for i := range line {
			if line[i] == '^' {
				u, v = row, i
			} else if line[i] == '#' {
				mark[row][i] = markO
			}
		}
		row++
	}

	var wg sync.WaitGroup

	type pair struct {
		u int
		v int
	}

	sum := int64(0)
	cpu := runtime.NumCPU()
	pipes := make([]chan pair, cpu)
	for i := range pipes {
		pipes[i] = make(chan pair, 100)
	}

	for i := 0; i < cpu; i++ {
		go func(i int) {
			mark2 := make([][]uint8, len(mark))
			for k := range mark {
				mark2[k] = make([]uint8, len(mark[k]))
			}

			for o := range pipes[i] {
				for k := range mark {
					copy(mark2[k], mark[k])
				}
				mark2[o.u][o.v] = markO
				if isCycle(mark2, u, v) {
					atomic.AddInt64(&sum, 1)
				}
				wg.Done()
			}
		}(i)
	}

	worker := 0
	for i := range mark {
		for j := range mark[i] {
			if mark[i][j] == 0 && !(i == u && j == v) {
				wg.Add(1)
				pipes[worker%cpu] <- pair{i, j}
				worker++
			}
		}
	}

	wg.Wait()

	fmt.Printf("%.3f ms\n", float64(time.Since(start).Microseconds())*0.001)
	fmt.Println(sum)
}

func TestDay6Part2PruningWithGoroutine(t *testing.T) {
	f, err := os.Open("data/day6.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	start := time.Now()

	var u, v int
	scanner := bufio.NewScanner(f)
	mark := [][]uint8{}
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		mark = append(mark, make([]uint8, len(line)))
		for i := range line {
			if line[i] == '^' {
				u, v = row, i
			} else if line[i] == '#' {
				mark[row][i] = markO
			}
		}
		row++
	}

	type pair struct {
		u int
		v int
	}

	origin := pair{u, v}

	var wg sync.WaitGroup

	sum := int64(0)
	cpu := runtime.NumCPU()
	pipes := make([]chan pair, cpu)
	for i := range pipes {
		pipes[i] = make(chan pair, 100)
	}

	for i := 0; i < cpu; i++ {
		go func(i int) {
			mark2 := make([][]uint8, len(mark))
			for k := range mark {
				mark2[k] = make([]uint8, len(mark[k]))
			}

			for o := range pipes[i] {
				for k := range mark {
					copy(mark2[k], mark[k])
				}
				mark2[o.u][o.v] = markO
				if isCycle(mark2, origin.u, origin.v) {
					atomic.AddInt64(&sum, 1)
				}
				wg.Done()
			}
		}(i)
	}

	worker := 0
	rows := len(mark)
	cols := len(mark[0])
	set := map[pair]struct{}{}

	u--
	for {
		for u >= 0 && mark[u][v] != markO {
			set[pair{u, v}] = struct{}{}
			u--
		}
		if u < 0 {
			break
		}
		u++
		v++
		for v < cols && mark[u][v] != markO {
			set[pair{u, v}] = struct{}{}
			v++
		}
		if v >= cols {
			break
		}
		v--
		u++
		for u < rows && mark[u][v] != markO {
			set[pair{u, v}] = struct{}{}
			u++
		}
		if u >= rows {
			break
		}
		u--
		v--
		for v >= 0 && mark[u][v] != markO {
			set[pair{u, v}] = struct{}{}
			v--
		}
		if v < 0 {
			break
		}
		v++
		u--
	}

	delete(set, origin)

	for k := range set {
		wg.Add(1)
		pipes[worker%cpu] <- k
		worker++
	}
	wg.Wait()

	fmt.Printf("%.3f ms\n", float64(time.Since(start).Microseconds())*0.001)
	fmt.Println(sum)
}

func isCycle(mark [][]uint8, u, v int) bool {
	rows := len(mark)
	cols := len(mark[0])
	for {
		for u >= 0 && mark[u][v] != markO {
			if mark[u][v]&markU == markU {
				return true
			}
			mark[u][v] |= markU
			u--
		}
		if u < 0 {
			break
		}
		u++
		for v < cols && mark[u][v] != markO {
			if mark[u][v]&markR == markR {
				return true
			}
			mark[u][v] |= markR
			v++
		}
		if v >= cols {
			break
		}
		v--
		for u < rows && mark[u][v] != markO {
			if mark[u][v]&markD == markD {
				return true
			}
			mark[u][v] |= markD
			u++
		}
		if u >= rows {
			break
		}
		u--
		for v >= 0 && mark[u][v] != markO {
			if mark[u][v]&markL == markL {
				return true
			}
			mark[u][v] |= markL
			v--
		}
		if v < 0 {
			break
		}
		v++
	}
	return false
}
