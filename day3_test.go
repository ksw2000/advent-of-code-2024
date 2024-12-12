package main_test

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"testing"
	"time"
)

const (
	d3Start = iota
	d3Mul
	d3Lp
	d3L
	d3Comma
	d3R
)

func TestDay3Part1(t *testing.T) {
	f, err := os.Open("data/day3.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	start := time.Now()

	var p, q, sum int
	var c rune
	var stack []rune
	state := d3Start
	x, _ := fmt.Fscanf(f, "%c", &c)
	for x > 0 {
		switch c {
		case 'l':
			if len(stack) > 1 && stack[len(stack)-2] == 'm' && stack[len(stack)-1] == 'u' {
				state = d3Mul
			} else {
				state = d3Start
			}
		case '(':
			if state == d3Mul {
				state = d3Lp
			} else {
				state = d3Start
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if state == d3Lp {
				state = d3L
				p = int(c - '0')
			} else if state == d3L {
				p = p*10 + int(c-'0')
			} else if state == d3Comma {
				state = d3R
				q = int(c - '0')
			} else if state == d3R {
				q = q*10 + int(c-'0')
			} else {
				state = d3Start
			}
		case ',':
			if state == d3L {
				state = d3Comma
			} else {
				state = d3Start
			}
		case ')':
			if state == d3R {
				state = d3Start
				sum += p * q
			} else {
				state = d3Start
			}
		default:
			state = d3Start
		}
		stack = append(stack, c)
		x, _ = fmt.Fscanf(f, "%c", &c)
	}

	fmt.Printf("%.3f ms\n", float64(time.Since(start).Microseconds())*0.001)
	fmt.Println(sum)
}

func TestDay3Part1Regexp(t *testing.T) {
	f, err := os.Open("data/day3.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	start := time.Now()

	input, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	r, err := regexp.Compile(`mul\(\d+,\d+\)`)
	if err != nil {
		panic(err)
	}
	sum := 0
	x := r.FindAllSubmatch(input, -1)
	for i := range x {
		for j := range x[i] {
			var p, q int
			fmt.Sscanf(string(x[i][j]), "mul(%d,%d)", &p, &q)
			sum += p * q
		}
	}

	fmt.Printf("%.3f ms\n", float64(time.Since(start).Microseconds())*0.001)
	fmt.Println(sum)
}

const DONOT = "don't("
const DO = "do("

func TestDay3Part2(t *testing.T) {
	f, err := os.Open("data/day3.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	start := time.Now()

	var p, q, sum int
	var c rune
	var stack []rune
	do := true
	state := d3Start
	x, _ := fmt.Fscanf(f, "%c", &c)
	for x > 0 {
		switch c {
		case 'l':
			if state == d3Start {
				if len(stack) > 1 && stack[len(stack)-2] == 'm' && stack[len(stack)-1] == 'u' {
					state = d3Mul
				} else {
					state = d3Start
				}
			} else {
				state = d3Start
			}
		case '(':
			if state == d3Mul {
				state = d3Lp
			} else {
				state = d3Start
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if state == d3Lp {
				state = d3L
				p = int(c - '0')
			} else if state == d3L {
				p = p*10 + int(c-'0')
			} else if state == d3Comma {
				state = d3R
				q = int(c - '0')
			} else if state == d3R {
				q = q*10 + int(c-'0')
			} else {
				state = d3Start
			}
		case ',':
			if state == d3L {
				state = d3Comma
			} else {
				state = d3Start
			}
		case ')':
			if len(stack) > len(DONOT)-1 && string(stack[len(stack)-len(DONOT):]) == DONOT {
				do = false
				state = d3Start
			} else if len(stack) > len(DO)-1 && string(stack[len(stack)-len(DO):]) == DO {
				do = true
				state = d3Start
			} else if state == d3R {
				state = d3Start
				if do {
					sum += p * q
				}
			} else {
				state = d3Start
			}
		default:
			state = d3Start
		}

		stack = append(stack, c)
		x, _ = fmt.Fscanf(f, "%c", &c)
	}

	fmt.Printf("%.3f ms\n", float64(time.Since(start).Microseconds())*0.001)
	fmt.Println(sum)
}

func TestDay3Part2Regexp(t *testing.T) {
	f, err := os.Open("data/day3.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	start := time.Now()

	input, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	r, err := regexp.Compile(`(?:mul\(\d+,\d+\))|do\(\)|don't\(\)`)
	if err != nil {
		panic(err)
	}
	sum := 0
	x := r.FindAllSubmatch(input, -1)
	do := true
	for i := range x {
		s := string(x[i][0])
		if s == "do()" {
			do = true
		} else if s == "don't()" {
			do = false
		} else if do {
			var p, q int
			fmt.Sscanf(s, "mul(%d,%d)", &p, &q)
			sum += p * q
		}
	}

	fmt.Printf("%.3f ms\n", float64(time.Since(start).Microseconds())*0.001)
	fmt.Println(sum)
}
