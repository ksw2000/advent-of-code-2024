package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
	"time"
)

// const day12Example = `RRRRIICCFF
// RRRRIICCCF
// VVRRRCCFFF
// VVRCCCJFFF
// VVVVCJJCFE
// VVIVCCJJEE
// VVIIICJJEE
// MIIIIIJJEE
// MIIISIJEEE
// MMMISSJEEE`

type uf struct {
	p []int
	r []int
	n int
}

func newUnionFind(n int) *uf {
	p := make([]int, n)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return &uf{
		p: p,
		r: r,
		n: n,
	}
}

func (u *uf) append(n int) {
	x := len(u.p)
	for i := x; i < x+n; i++ {
		u.p = append(u.p, i)
		u.r = append(u.r, 0)
	}
	u.n += n
}

func (u *uf) union(a, b int) {
	pa := u.find(a)
	pb := u.find(b)
	if pa != pb {
		if u.r[pa] < u.r[pb] {
			u.p[pa] = pb
		} else if u.r[pa] > u.r[pb] {
			u.p[pb] = pa
		} else {
			u.p[pb] = pa
			u.r[pa]++
		}
		u.n--
	}
}

func (u *uf) find(a int) int {
	for a != u.p[a] {
		u.p[a] = u.p[u.p[a]]
		a = u.p[a]
	}
	return a
}

func TestDay12Part1(t *testing.T) {
	f, err := os.Open("data/day12.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// f := strings.NewReader(day12Example)

	start := time.Now()

	perimeter := []int{}

	scanner := bufio.NewScanner(f)
	row := 0

	uf := newUnionFind(0)

	var previousLine string

	for scanner.Scan() {
		line := scanner.Text()

		uf.append(len(line))
		for i := range line {
			p := 4
			if i > 0 && line[i] == line[i-1] {
				p -= 2
				uf.union(row*len(line)+i, row*len(line)+i-1)
			}
			if row > 0 && line[i] == previousLine[i] {
				p -= 2
				uf.union(row*len(line)+i, (row-1)*len(line)+i)
			}
			perimeter = append(perimeter, p)
		}
		previousLine = line
		row++
	}

	areaOfComponent := make([]int, len(uf.p))
	perimeterOfComponent := make([]int, len(uf.p))
	for i := range areaOfComponent {
		p := uf.find(i)
		areaOfComponent[p]++
		perimeterOfComponent[p] += perimeter[i]
	}

	sum := 0
	for i := range areaOfComponent {
		if areaOfComponent[i] != 0 {
			sum += areaOfComponent[i] * perimeterOfComponent[i]
		}
	}

	fmt.Printf("%.3f ms\n", float64(time.Since(start).Microseconds())*0.001)
	fmt.Println("sum =", sum)
}

func TestDay12Part2(t *testing.T) {
	f, err := os.Open("data/day12.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	start := time.Now()

	// f := strings.NewReader(day12Example)

	// A B C
	// D *

	// same = 1
	// diff = 0

	// A B C D
	// 0 0 0 0  +4
	// 0 0 0 1   0
	// 0 0 1 0  +4
	// 0 0 1 1   0
	// 0 1 0 0   0
	// 0 1 0 1  -2
	// 0 1 1 0  +2
	// 0 1 1 1   0
	// 1 0 0 0  +4
	// 1 0 0 1  +2
	// 1 0 1 0  +4
	// 1 0 1 1  +2
	// 1 1 0 0  +2
	// 1 1 0 1  -2
	// 1 1 1 0  +4
	// 1 1 1 1   0

	const flagA = 8
	const flagB = 4
	const flagC = 2
	const flagD = 1
	do := [16]int{
		0:                     4,
		flagC:                 4,
		flagB | flagD:         -2,
		flagB | flagC:         2,
		flagA:                 4,
		flagA | flagD:         2,
		flagA | flagC:         4,
		flagA | flagC | flagD: 2,
		flagA | flagB:         2,
		flagA | flagB | flagD: -2,
		flagA | flagB | flagC: 4,
	}

	table := [][]byte{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		table = append(table, []byte(line))
	}

	rows := len(table)
	cols := len(table[0])

	uf := newUnionFind(rows * cols)
	perimeter := make([]int, rows*cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			flag := 0
			if j > 0 && table[i][j] == table[i][j-1] {
				flag = flag | flagD
				uf.union(i*cols+j, i*cols+j-1)
			}
			if i > 0 && table[i][j] == table[i-1][j] {
				flag = flag | flagB
				uf.union(i*cols+j, (i-1)*cols+j)
			}
			if i > 0 && j > 0 && table[i][j] == table[i-1][j-1] {
				flag = flag | flagA
			}
			if j+1 < cols && i > 0 && table[i][j] == table[i-1][j+1] {
				flag = flag | flagC
			}
			perimeter[i*cols+j] = do[flag]
		}
	}

	areaOfComponent := make(map[int]int, uf.n)
	perimeterOfComponent := make(map[int]int, uf.n)
	for i := range perimeter {
		p := uf.find(i)
		areaOfComponent[p]++
		perimeterOfComponent[p] += perimeter[i]
	}

	sum := 0
	for i := range areaOfComponent {
		// fmt.Println(areaOfComponent[i], "*", perimeterOfComponent[i])
		sum += areaOfComponent[i] * perimeterOfComponent[i]
	}

	fmt.Printf("%.3f ms\n", float64(time.Since(start).Microseconds())*0.001)
	fmt.Println("sum =", sum)
}
