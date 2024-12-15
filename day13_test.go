package main

import (
	"fmt"
	"os"
	"testing"
)

// const day13Example = `Button A: X+94, Y+34
// Button B: X+22, Y+67
// Prize: X=8400, Y=5400

// Button A: X+26, Y+66
// Button B: X+67, Y+21
// Prize: X=12748, Y=12176

// Button A: X+17, Y+86
// Button B: X+84, Y+37
// Prize: X=7870, Y=6450

// Button A: X+69, Y+23
// Button B: X+27, Y+71
// Prize: X=18641, Y=10279`

func TestDay13Part1(t *testing.T) {
	f, err := os.Open("data/day13.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// f := strings.NewReader(day13Example)

	var a1, a2 int
	var b1, b2 int
	var c1, c2 int

	n, _ := fmt.Fscanf(f, "Button A: X+%d, Y+%d\nButton B: X+%d, Y+%d\nPrize: X=%d, Y=%d", &a1, &a2, &b1, &b2, &c1, &c2)

	sum := 0
	for n > 0 {
		// xa1 + yb1 = c1
		// xa2 + yb2 = c2

		delta := a1*b2 - a2*b1
		deltaX := c1*b2 - c2*b1
		deltaY := a1*c2 - a2*c1
		if delta != 0 {
			x, rx := deltaX/delta, deltaX%delta
			y, ry := deltaY/delta, deltaY%delta
			if rx == 0 && ry == 0 {
				sum += 3*x + y
			}
		} else if deltaX == 0 && deltaY == 0 {
			// xa1 + yb1 = c1
			for x := 0; ; x++ {
				y, ry := (c1-x*a1)/b1, (c1-x*a1)%b1
				if ry == 0 {
					sum += 3*x + y
					break
				}
				if y < 0 {
					break
				}
			}
		}

		fmt.Fscanln(f)
		n, _ = fmt.Fscanf(f, "Button A: X+%d, Y+%d\nButton B: X+%d, Y+%d\nPrize: X=%d, Y=%d", &a1, &a2, &b1, &b2, &c1, &c2)
	}
	fmt.Println("sum =", sum)
}

func TestDay13Part2(t *testing.T) {
	f, err := os.Open("data/day13.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// f := strings.NewReader(day13Example)

	var a1, a2 int
	var b1, b2 int
	var c1, c2 int

	n, _ := fmt.Fscanf(f, "Button A: X+%d, Y+%d\nButton B: X+%d, Y+%d\nPrize: X=%d, Y=%d", &a1, &a2, &b1, &b2, &c1, &c2)

	sum := 0
	for n > 0 {
		c1 += 10000000000000
		c2 += 10000000000000

		delta := a1*b2 - a2*b1
		deltaX := c1*b2 - c2*b1
		deltaY := a1*c2 - a2*c1
		if delta != 0 {
			x, rx := deltaX/delta, deltaX%delta
			y, ry := deltaY/delta, deltaY%delta
			if rx == 0 && ry == 0 {
				sum += 3*x + y
			}
		} else if deltaX == 0 && deltaY == 0 {
			// xa1 + yb1 = c1
			for x := 0; ; x++ {
				y, ry := (c1-x*a1)/b1, (c1-x*a1)%b1
				if ry == 0 {
					sum += 3*x + y
					break
				}
				if y < 0 {
					break
				}
			}
		}

		fmt.Fscanln(f)
		n, _ = fmt.Fscanf(f, "Button A: X+%d, Y+%d\nButton B: X+%d, Y+%d\nPrize: X=%d, Y=%d", &a1, &a2, &b1, &b2, &c1, &c2)
	}
	fmt.Println("sum =", sum)
}
