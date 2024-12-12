package main_test

import (
	"fmt"
	"os"
	"sort"
	"testing"
)

func TestDay1Part1(t *testing.T) {
	f, err := os.Open("data/day1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	left := []int{}
	right := []int{}
	var l, r int
	n, _ := fmt.Fscanf(f, "%d %d", &l, &r)
	for n > 0 {
		left = append(left, l)
		right = append(right, r)
		n, _ = fmt.Fscanf(f, "%d %d", &l, &r)
	}
	sort.Ints(left)
	sort.Ints(right)
	dis := 0
	for i := range left {
		dis += abs(left[i] - right[i])
	}
	fmt.Println(dis)
}

func TestDay1Part2(t *testing.T) {
	f, err := os.Open("data/day1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	leftMap := map[int]int{}
	rightMap := map[int]int{}
	var l, r int
	n, _ := fmt.Fscanf(f, "%d %d", &l, &r)
	for n > 0 {
		leftMap[l]++
		rightMap[r]++
		n, _ = fmt.Fscanf(f, "%d %d", &l, &r)
	}
	similarity := 0
	for k, v := range leftMap {
		similarity += k * v * rightMap[k]
	}
	fmt.Println(similarity)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
