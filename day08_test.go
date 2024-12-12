package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"
)

// const day8Example = `............
// ........0...
// .....0......
// .......0....
// ....0.......
// ......A.....
// ............
// ............
// ........A...
// .........A..
// ............
// ............`

func TestDay8Part1(t *testing.T) {
	// f := strings.NewReader(day8Part1Example)
	f, err := os.Open("data/day8.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	start := time.Now()

	scanner := bufio.NewScanner(f)

	table := [][]byte{}

	antennas := map[byte][]struct {
		row int
		col int
	}{}
	for scanner.Scan() {
		line := scanner.Text()
		for i := range line {
			if line[i] != '.' {
				antennas[line[i]] = append(antennas[line[i]], struct {
					row int
					col int
				}{
					row: len(table),
					col: i,
				})
			}
		}
		table = append(table, []byte(line))
	}
	set := map[struct {
		row int
		col int
	}]struct{}{}

	in := func(row, col int) bool {
		return row >= 0 && row < len(table) && col >= 0 && col < len(table[0])
	}

	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, positions := range antennas {
		if len(positions) > 0 {
			wg.Add(1)
			go func(positions []struct {
				row int
				col int
			}) {
				defer wg.Done()
				for i := 0; i < len(positions); i++ {
					for j := i + 1; j < len(positions); j++ {
						dRow := positions[i].row - positions[j].row
						dCol := positions[i].col - positions[j].col
						// antinode 1
						if in(positions[i].row+dRow, positions[i].col+dCol) {
							mu.Lock()
							set[struct {
								row int
								col int
							}{
								positions[i].row + dRow,
								positions[i].col + dCol,
							}] = struct{}{}
							mu.Unlock()
						}
						// antinode 2
						if in(positions[j].row-dRow, positions[j].col-dCol) {
							mu.Lock()
							set[struct {
								row int
								col int
							}{
								positions[j].row - dRow,
								positions[j].col - dCol,
							}] = struct{}{}
							mu.Unlock()
						}
					}
				}
			}(positions)
		}
	}
	wg.Wait()

	fmt.Printf("%.3f ms\n", float64(time.Since(start).Microseconds())*0.001)
	fmt.Println(len(set))
}

func TestDay8Part2(t *testing.T) {
	// f := strings.NewReader(day8Example)
	f, err := os.Open("data/day8.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	start := time.Now()

	scanner := bufio.NewScanner(f)

	table := [][]byte{}

	antennas := map[byte][]struct {
		row int
		col int
	}{}
	for scanner.Scan() {
		line := scanner.Text()
		for i := range line {
			if line[i] != '.' {
				antennas[line[i]] = append(antennas[line[i]], struct {
					row int
					col int
				}{
					row: len(table),
					col: i,
				})
			}
		}
		table = append(table, []byte(line))
	}

	in := func(row, col int) bool {
		return row >= 0 && row < len(table) && col >= 0 && col < len(table[0])
	}

	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, positions := range antennas {
		if len(positions) > 0 {
			wg.Add(1)
			go func(positions []struct {
				row int
				col int
			}) {
				defer wg.Done()
				for i := 0; i < len(positions); i++ {
					mu.Lock()
					table[positions[i].row][positions[i].col] = '#'
					mu.Unlock()
					for j := i + 1; j < len(positions); j++ {
						dRow := positions[i].row - positions[j].row
						dCol := positions[i].col - positions[j].col
						// antinode 1
						for r, c := dRow, dCol; in(positions[i].row+r, positions[i].col+c); {
							mu.Lock()
							table[positions[i].row+r][positions[i].col+c] = '#'
							mu.Unlock()
							r += dRow
							c += dCol
						}
						// antinode 2
						for r, c := dRow, dCol; in(positions[j].row-r, positions[j].col-c); {
							mu.Lock()
							table[positions[j].row-r][positions[j].col-c] = '#'
							mu.Unlock()
							r += dRow
							c += dCol
						}
					}
				}
			}(positions)
		}
	}
	wg.Wait()

	sum := 0
	for i := range table {
		for j := range table[i] {
			if table[i][j] == '#' {
				sum++
			}
		}
	}

	fmt.Printf("%.3f ms\n", float64(time.Since(start).Microseconds())*0.001)
	fmt.Println(sum)
}
