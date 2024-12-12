package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
)

const day7Part1Example = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func TestDay7Part1(t *testing.T) {
	f := strings.NewReader(day7Part1Example)
	fmt.Println("Example:", testDay7Part1(f))

	g, err := os.Open("data/day7.txt")
	if err != nil {
		panic(err)
	}
	defer g.Close()
	fmt.Println("Real:", testDay7Part1(g))
}

func testDay7Part1(f io.Reader) int64 {
	scanner := bufio.NewScanner(f)
	sum := int64(0)
	var wg sync.WaitGroup
	for scanner.Scan() {
		line := scanner.Text()
		wg.Add(1)
		go func(reader *strings.Reader) {
			defer wg.Done()

			var target, n int
			fmt.Fscanf(reader, "%d: %d", &target, &n)

			q1 := []int{n}
			q2 := []int{}

			x, _ := fmt.Fscanf(reader, "%d", &n)
			for x > 0 {
				for _, m := range q1 {
					if m+n <= target {
						q2 = append(q2, m+n)
					}
					if m*n <= target {
						q2 = append(q2, m*n)
					}
				}
				q1, q2 = q2, q1
				q2 = q2[:0]
				x, _ = fmt.Fscanf(reader, "%d", &n)
			}
			if slices.Contains(q1, target) {
				atomic.AddInt64(&sum, int64(target))
			}
		}(strings.NewReader(line))
	}
	wg.Wait()
	return sum
}

func TestDay7Part2(t *testing.T) {
	f := strings.NewReader(day7Part1Example)
	fmt.Println("Example:", testDay7Part2(f))

	g, err := os.Open("data/day7.txt")
	if err != nil {
		panic(err)
	}
	defer g.Close()
	fmt.Println("Real:", testDay7Part2(g))
}

func concat(a, b int) int {
	c := b
	for b > 0 {
		a = a * 10
		b = b / 10
	}
	return a + c
}

func testDay7Part2(f io.Reader) int64 {
	scanner := bufio.NewScanner(f)
	sum := int64(0)
	var wg sync.WaitGroup
	for scanner.Scan() {
		line := scanner.Text()
		wg.Add(1)
		go func(reader *strings.Reader) {
			defer wg.Done()

			var target, n int
			fmt.Fscanf(reader, "%d: %d", &target, &n)

			q1 := []int{n}
			q2 := []int{}

			x, _ := fmt.Fscanf(reader, "%d", &n)
			for x > 0 {
				for _, m := range q1 {
					if d := m + n; d <= target {
						q2 = append(q2, d)
					}
					if d := m * n; d <= target {
						q2 = append(q2, d)
					}
					if d := concat(m, n); d <= target {
						q2 = append(q2, d)
					}
				}
				q1, q2 = q2, q1
				q2 = q2[:0]
				x, _ = fmt.Fscanf(reader, "%d", &n)
			}
			if slices.Contains(q1, target) {
				atomic.AddInt64(&sum, int64(target))
			}
		}(strings.NewReader(line))
	}
	wg.Wait()
	return sum
}
