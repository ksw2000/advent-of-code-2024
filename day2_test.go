package main_test

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestDay2Part1(t *testing.T) {
	f, err := os.Open("data/day2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	safe := 0
	scanner := bufio.NewScanner(f)
next:
	for scanner.Scan() {
		line := scanner.Text()
		reader := strings.NewReader(line)
		increasing := false

		var pre, n int
		x, _ := fmt.Fscanf(reader, "%d %d", &pre, &n)
		if x != 2 {
			continue
		}
		delta := n - pre

		if delta == 0 {
			continue
		}

		increasing = delta > 0

		if !increasing {
			delta = -delta
		}

		if delta < 1 || delta > 3 {
			continue next
		}

		pre = n
		x, _ = fmt.Fscanf(reader, "%d", &n)
		for x > 0 {
			delta = n - pre
			if delta == 0 {
				continue next
			}
			if increasing && (delta < 0 || delta < 1 || delta > 3) {
				continue next
			}
			if !increasing && (delta > 0 || delta > -1 || delta < -3) {
				continue next
			}

			pre = n
			x, _ = fmt.Fscanf(reader, "%d", &n)
		}
		safe++
	}
	fmt.Println(safe)
}

func TestDay2Part2(t *testing.T) {
	f, err := os.Open("data/day2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	safe := 0
	scanner := bufio.NewScanner(f)
next:
	for scanner.Scan() {
		line := scanner.Text()
		reader := strings.NewReader(line)
		list := []int{}
		var n int
		x, _ := fmt.Fscanf(reader, "%d", &n)
		for x > 0 {
			list = append(list, n)
			x, _ = fmt.Fscanf(reader, "%d", &n)
		}

		if check(list) {
			safe++
			continue
		}

		list2 := make([]int, 0, len(list)-1)
		for i := 0; i < len(list); i++ {
			list2 = list2[:0]
			list2 = append(list2, list[:i]...)
			list2 = append(list2, list[i+1:]...)
			if check(list2) {
				safe++
				continue next
			}
		}
	}
	fmt.Println(safe)
}

func check(list []int) bool {
	if len(list) == 1 {
		panic("strange data")
	}

	delta := list[1] - list[0]

	if delta == 0 {
		return false
	}

	increasing := delta > 0

	if !increasing {
		delta = -delta
	}

	if delta < 1 || delta > 3 {
		return false
	}

	for i := 2; i < len(list); i++ {
		delta := list[i] - list[i-1]
		if delta == 0 {
			return false
		}
		if increasing && (delta < 0 || delta < 1 || delta > 3) {
			return false
		}
		if !increasing && (delta > 0 || delta > -1 || delta < -3) {
			return false
		}
	}
	return true
}
