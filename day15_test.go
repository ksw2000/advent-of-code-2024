package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
	"time"
)

const day15example = `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`

func TestDay15Part1Recursion(t *testing.T) {
	f, err := os.Open("data/day15.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// f := strings.NewReader(day15example)

	table := [][]byte{}
	actions := []byte{}

	var startI, startJ int

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		for i := range line {
			if line[i] == '@' {
				startI = len(table)
				startJ = i
			}
		}
		table = append(table, []byte(line))
	}
	for scanner.Scan() {
		line := scanner.Text()
		actions = append(actions, []byte(line)...)
	}

	var move func(i, j, dI, dJ int) (int, int)
	move = func(i, j, dI, dJ int) (int, int) {
		switch table[i+dI][j+dJ] {
		case 'O':
			move(i+dI, j+dJ, dI, dJ)
			if table[i+dI][j+dJ] != '.' {
				break
			}
			fallthrough
		case '.':
			table[i][j], table[i+dI][j+dJ] = '.', table[i][j]
			return i + dI, j + dJ
		}
		return i, j
	}

	for i := range actions {
		switch actions[i] {
		case '^':
			startI, startJ = move(startI, startJ, -1, 0)
		case '>':
			startI, startJ = move(startI, startJ, 0, 1)
		case 'v':
			startI, startJ = move(startI, startJ, 1, 0)
		case '<':
			startI, startJ = move(startI, startJ, 0, -1)
		}
	}

	gps := 0
	for i := range table {
		for j := range table[i] {
			if table[i][j] == 'O' {
				gps += i*100 + j
			}
		}
	}

	fmt.Println("gps =", gps)
}

func TestDay15Part1Iteration(t *testing.T) {
	f, err := os.Open("data/day15.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// f := strings.NewReader(day15example)

	table := [][]byte{}
	actions := []byte{}

	var startI, startJ int

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		for i := range line {
			if line[i] == '@' {
				startI = len(table)
				startJ = i
			}
		}
		table = append(table, []byte(line))
	}
	for scanner.Scan() {
		line := scanner.Text()
		actions = append(actions, []byte(line)...)
	}

	move := func(i, j, dI, dJ int) (int, int) {
		k := 1
		for ; table[i+dI*k][j+dJ*k] == 'O'; k++ {
		}
		if table[i+dI*k][j+dJ*k] == '.' {
			for k >= 1 {
				table[i+dI*k][j+dJ*k] = table[i+dI*(k-1)][j+dJ*(k-1)]
				k--
			}
			table[i][j] = '.'
			return i + dI, j + dJ
		}
		return i, j
	}

	for i := range actions {
		switch actions[i] {
		case '^':
			startI, startJ = move(startI, startJ, -1, 0)
		case '>':
			startI, startJ = move(startI, startJ, 0, 1)
		case 'v':
			startI, startJ = move(startI, startJ, 1, 0)
		case '<':
			startI, startJ = move(startI, startJ, 0, -1)
		}
	}

	gps := 0
	for i := range table {
		for j := range table[i] {
			if table[i][j] == 'O' {
				gps += i*100 + j
			}
		}
	}

	fmt.Println("gps =", gps)
}

func TestDay15Part2(t *testing.T) {
	f, err := os.Open("data/day15.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// f := strings.NewReader(day15example)

	startTime := time.Now()

	table := [][]byte{}
	actions := []byte{}

	type pos struct {
		i, j int
	}

	start := pos{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			break
		}
		line2 := make([]byte, len(line)<<1)
		for i := range line {
			j := i << 1
			switch line[i] {
			case '#', '.':
				line2[j] = line[i]
				line2[j+1] = line[i]
			case 'O':
				line2[j] = '['
				line2[j+1] = ']'
			case '@':
				line2[j] = '@'
				line2[j+1] = '.'
				start = pos{
					i: len(table),
					j: j,
				}
			}
		}
		table = append(table, line2)
	}

	for scanner.Scan() {
		line := scanner.Text()
		actions = append(actions, []byte(line)...)
	}

	moveWide := func(p *pos, dI int) {
		q := []*pos{p}
		visited := map[pos]struct{}{}
		for i := 0; i < len(q); i++ {
			if _, v := visited[pos{q[i].i + dI, q[i].j - 1}]; !v && table[q[i].i+dI][q[i].j-1] == '[' {
				q = append(q, &pos{q[i].i + dI, q[i].j - 1})
				visited[pos{q[i].i + dI, q[i].j - 1}] = struct{}{}
			}
			if _, v := visited[pos{q[i].i + dI, q[i].j}]; !v && table[q[i].i+dI][q[i].j] == '[' {
				q = append(q, &pos{q[i].i + dI, q[i].j})
				visited[pos{q[i].i + dI, q[i].j}] = struct{}{}
			}
			if _, v := visited[pos{q[i].i + dI, q[i].j + 1}]; !v && table[q[i].i+dI][q[i].j+1] == '[' {
				q = append(q, &pos{q[i].i + dI, q[i].j + 1})
				visited[pos{q[i].i + dI, q[i].j + 1}] = struct{}{}
			}
			if table[q[i].i+dI][q[i].j] == '#' || table[q[i].i+dI][q[i].j+1] == '#' {
				return
			}
		}

		for i := len(q) - 1; i >= 0; i-- {
			table[q[i].i+dI][q[i].j] = '['
			table[q[i].i][q[i].j] = '.'
			table[q[i].i+dI][q[i].j+1] = ']'
			table[q[i].i][q[i].j+1] = '.'
		}
	}

	var move func(p *pos, dI, dJ int)
	move = func(p *pos, dI, dJ int) {
		if dI != 0 {
			switch table[p.i+dI][p.j+dJ] {
			case '[':
				moveWide(&pos{
					i: p.i + dI,
					j: p.j + dJ,
				}, dI)
			case ']':
				moveWide(&pos{
					i: p.i + dI,
					j: p.j + dJ - 1,
				}, dI)
			}
		} else {
			switch table[p.i+dI][p.j+dJ] {
			case '[', ']':
				move(&pos{
					i: p.i + dI,
					j: p.j + dJ,
				}, dI, dJ)
			}
		}
		if table[p.i+dI][p.j+dJ] == '.' {
			table[p.i][p.j], table[p.i+dI][p.j+dJ] = '.', table[p.i][p.j]
			p.i = p.i + dI
			p.j = p.j + dJ
		}
	}

	for i := range actions {
		switch actions[i] {
		case '^':
			move(&start, -1, 0)
		case '>':
			move(&start, 0, 1)
		case 'v':
			move(&start, 1, 0)
		case '<':
			move(&start, 0, -1)
		}
	}

	gps := 0
	for i := range table {
		for j := range table[i] {
			if table[i][j] == '[' {
				gps += i*100 + j
			}
		}
	}

	fmt.Printf("%.3f ms\n", float64(time.Since(startTime).Microseconds())*0.001)
	fmt.Println("gps =", gps)
}
