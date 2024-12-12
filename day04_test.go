package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func loadDay4Data() [][]byte {
	f, err := os.Open("data/day4.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var line []byte
	var lines [][]byte
	x, _ := fmt.Fscanf(f, "%s", &line)
	for x > 0 {
		lines = append(lines, line)
		x, _ = fmt.Fscanf(f, "%s", &line)
	}
	return lines
}

// func loadDay4Part1Example() [][]byte {
// 	return [][]byte{
// 		[]byte("MMMSXXMASM"),
// 		[]byte("MSAMXMSMSA"),
// 		[]byte("AMXSXMAAMM"),
// 		[]byte("MSAMASMSMX"),
// 		[]byte("XMASAMXAMM"),
// 		[]byte("XXAMMXXAMA"),
// 		[]byte("SMSMSASXSS"),
// 		[]byte("SAXAMASAAA"),
// 		[]byte("MAMMMXMMMM"),
// 		[]byte("MXMXAXMASX"),
// 	}
// }

func TestDay4Part1(t *testing.T) {
	data := loadDay4Data()
	start := time.Now()
	sum := 0
	for i := range data {
		for j := range data[i] {
			if data[i][j] == 'X' {
				if i-3 >= 0 &&
					j-3 >= 0 &&
					data[i-1][j-1] == 'M' &&
					data[i-2][j-2] == 'A' &&
					data[i-3][j-3] == 'S' {
					sum++
				}
				if i-3 >= 0 &&
					data[i-1][j] == 'M' &&
					data[i-2][j] == 'A' &&
					data[i-3][j] == 'S' {
					sum++
				}
				if i-3 >= 0 &&
					j+3 < len(data[i]) &&
					data[i-1][j+1] == 'M' &&
					data[i-2][j+2] == 'A' &&
					data[i-3][j+3] == 'S' {
					sum++
				}
				if j-3 >= 0 &&
					data[i][j-1] == 'M' &&
					data[i][j-2] == 'A' &&
					data[i][j-3] == 'S' {
					sum++
				}
				if j+3 < len(data[i]) &&
					data[i][j+1] == 'M' &&
					data[i][j+2] == 'A' &&
					data[i][j+3] == 'S' {
					sum++
				}
				if i+3 < len(data) &&
					j-3 >= 0 &&
					data[i+1][j-1] == 'M' &&
					data[i+2][j-2] == 'A' &&
					data[i+3][j-3] == 'S' {
					sum++
				}
				if i+3 < len(data) &&
					data[i+1][j] == 'M' &&
					data[i+2][j] == 'A' &&
					data[i+3][j] == 'S' {
					sum++
				}
				if i+3 < len(data) &&
					j+3 < len(data[i]) &&
					data[i+1][j+1] == 'M' &&
					data[i+2][j+2] == 'A' &&
					data[i+3][j+3] == 'S' {
					sum++
				}
			}
		}
	}
	fmt.Printf("%.3f ms\n", float64(time.Since(start).Microseconds())*0.001)
	fmt.Println(sum)
}

func TestDay4Part1Way2(t *testing.T) {
	data := loadDay4Data()

	start := time.Now()
	// data := loadDay4Data()
	matching := [][]struct {
		row int
		col int
	}{
		{{-1, -1}, {-2, -2}, {-3, -3}},
		{{-1, 0}, {-2, 0}, {-3, 0}},
		{{-1, 1}, {-2, 2}, {-3, 3}},
		{{0, -1}, {0, -2}, {0, -3}},
		{{0, 1}, {0, 2}, {0, 3}},
		{{1, -1}, {2, -2}, {3, -3}},
		{{1, 0}, {2, 0}, {3, 0}},
		{{1, 1}, {2, 2}, {3, 3}},
	}

	pattern := []byte("MAS")

	sum := 0
	for i := range data {
		for j := range data[i] {
			if data[i][j] == 'X' {
			find:
				for s := range matching {
					for t, v := range matching[s] {
						if i+v.row >= 0 && i+v.row < len(data) &&
							j+v.col >= 0 && j+v.col < len(data[i]) &&
							data[i+v.row][j+v.col] == pattern[t] {
						} else {
							continue find
						}
					}
					sum++
				}
			}
		}
	}
	fmt.Printf("%.3f ms\n", float64(time.Since(start).Microseconds())*0.001)
	fmt.Println(sum)
}

func loadDay4Part2Example() [][]byte {
	return [][]byte{
		[]byte(".M.S......"),
		[]byte("..A..MSMS."),
		[]byte(".M.S.MAA.."),
		[]byte("..A.ASMSM."),
		[]byte(".M.S.M...."),
		[]byte(".........."),
		[]byte("S.S.S.S.S."),
		[]byte(".A.A.A.A.."),
		[]byte("M.M.M.M.M."),
		[]byte(".........."),
	}
}

func TestDay4Part2(t *testing.T) {
	data := loadDay4Part2Example()
	fmt.Println("Example:", day4Part2(data))
	data = loadDay4Data()
	fmt.Println("Real:", day4Part2(data))
}

func day4Part2(data [][]byte) int {
	/*
		M.S
		.A.
		M.S
	*/
	/*
		M.M
		.A.
		S.S
	*/
	/*
		S.M
		.A.
		S.M
	*/
	/*
		S.S
		.A.
		M.M
	*/
	sum := 0
	for i := 2; i < len(data); i++ {
		for j := 2; j < len(data[i]); j++ {
			if data[i-1][j-1] != 'A' {
				continue
			}
			if data[i][j] == 'S' &&
				data[i][j-2] == 'M' &&
				data[i-2][j-2] == 'M' &&
				data[i-2][j] == 'S' {
				sum++
			}
			if data[i][j] == 'S' &&
				data[i][j-2] == 'S' &&
				data[i-2][j-2] == 'M' &&
				data[i-2][j] == 'M' {
				sum++
			}
			if data[i][j] == 'M' &&
				data[i][j-2] == 'S' &&
				data[i-2][j-2] == 'S' &&
				data[i-2][j] == 'M' {
				sum++
			}
			if data[i][j] == 'M' &&
				data[i][j-2] == 'M' &&
				data[i-2][j-2] == 'S' &&
				data[i-2][j] == 'S' {
				sum++
			}
		}
	}
	return sum
}
