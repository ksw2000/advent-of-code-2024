package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

// const day10Part1Example = `89010123
// 78121874
// 87430965
// 96549874
// 45678903
// 32019012
// 01329801
// 10456732`

func TestDay10Part1(t *testing.T) {
	// f := strings.NewReader(day10Part1Example)

	f, err := os.Open("data/day10.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	table := [][]int{}
	row := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		table = append(table, []int{})
		for i := range line {
			table[row] = append(table[row], int(line[i]-'0'))
		}
		row++
	}

	sum := 0
	for i := range table {
		for j := range table[i] {
			if table[i][j] == 0 {
				q1 := map[struct {
					row int
					col int
				}]struct{}{}

				q1[struct {
					row int
					col int
				}{
					i, j,
				}] = struct{}{}

				q2 := map[struct {
					row int
					col int
				}]struct{}{}

				depth := 1
				for len(q1) > 0 && depth < 10 {
					for q := range q1 {
						if q.row-1 >= 0 && table[q.row-1][q.col] == depth {
							q2[struct {
								row int
								col int
							}{
								q.row - 1,
								q.col,
							}] = struct{}{}
						}
						if q.row+1 < len(table) && table[q.row+1][q.col] == depth {
							q2[struct {
								row int
								col int
							}{
								q.row + 1,
								q.col,
							}] = struct{}{}
						}
						if q.col-1 >= 0 && table[q.row][q.col-1] == depth {
							q2[struct {
								row int
								col int
							}{
								q.row,
								q.col - 1,
							}] = struct{}{}
						}
						if q.col+1 < len(table[0]) && table[q.row][q.col+1] == depth {
							q2[struct {
								row int
								col int
							}{
								q.row,
								q.col + 1,
							}] = struct{}{}
						}
					}
					depth++
					q1, q2 = q2, q1
					clear(q2)
				}
				if depth == 10 {
					sum += len(q1)
				}
			}
		}
	}

	fmt.Println(sum)
}

func TestDay10Part2(t *testing.T) {
	// f := strings.NewReader(day10Part1Example)

	f, err := os.Open("data/day10.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	table := [][]int{}
	row := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		table = append(table, []int{})
		for i := range line {
			table[row] = append(table[row], int(line[i]-'0'))
		}
		row++
	}

	score := make([][]int, len(table))
	for i := range score {
		score[i] = make([]int, len(table[0]))
	}

	for i := range table {
		for j := range table[i] {
			if table[i][j] == 9 {
				score[i][j] = 1
			}
		}
	}
	sum := 0
	for k := 8; k >= 0; k-- {
		for i := range table {
			for j := range table[i] {
				if table[i][j] == k {
					s := 0
					if i-1 >= 0 && table[i-1][j] == k+1 {
						s += score[i-1][j]
					}
					if i+1 < len(table) && table[i+1][j] == k+1 {
						s += score[i+1][j]
					}
					if j-1 >= 0 && table[i][j-1] == k+1 {
						s += score[i][j-1]
					}
					if j+1 < len(table[0]) && table[i][j+1] == k+1 {
						s += score[i][j+1]
					}
					score[i][j] = s
					if k == 0 {
						sum += s
					}
				}
			}
		}
	}

	fmt.Println(sum)
}
