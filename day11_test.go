package main

import (
	"fmt"
	"math"
	"strings"
	"sync"
	"testing"
	"time"
)

// const day10Part1Example = `125 17`
const day10Part1Test = `5 89749 6061 43 867 1965860 0 206250`

func TestDay11Part1(t *testing.T) {
	f := strings.NewReader(day10Part1Test)

	start := time.Now()

	var n int
	x, _ := fmt.Fscanf(f, "%d", &n)

	sum := 0
	for x > 0 {
		sum += testDay10Part1(n)
		x, _ = fmt.Fscanf(f, "%d", &n)
	}

	fmt.Printf("%.3f ms\n", float64(time.Since(start).Microseconds())*0.001)
	fmt.Println(sum)
}

func TestDay11Part2(t *testing.T) {
	f := strings.NewReader(day10Part1Test)

	start := time.Now()

	var n int
	x, _ := fmt.Fscanf(f, "%d", &n)

	sum := 0
	for x > 0 {
		sum += testDay10Part2(n)
		x, _ = fmt.Fscanf(f, "%d", &n)
	}

	fmt.Printf("%.3f ms\n", float64(time.Since(start).Microseconds())*0.001)
	fmt.Println(sum)
}

func TestDay11Part2Goroutine(t *testing.T) {
	f := strings.NewReader(day10Part1Test)

	start := time.Now()

	var n int
	x, _ := fmt.Fscanf(f, "%d", &n)

	sum := 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	for x > 0 {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			m := testDay10Part2(n)
			mu.Lock()
			sum += m
			mu.Unlock()
		}(n)
		x, _ = fmt.Fscanf(f, "%d", &n)
	}
	wg.Wait()
	fmt.Printf("%.3f ms\n", float64(time.Since(start).Microseconds())*0.001)
	fmt.Println(sum)
}

func testDay10Part2(n int) int {
	q1 := map[int]int{
		n: 1,
	}
	q2 := map[int]int{}

	for i := 0; i < 75; i++ {
		for q, k := range q1 {
			if a, b, two := blink(q); two {
				q2[a] += k
				q2[b] += k
			} else {
				q2[a] += k
			}
		}
		q1, q2 = q2, q1
		clear(q2)
	}

	sum := 0
	for _, k := range q1 {
		sum += k
	}
	return sum
}

func testDay10Part1(n int) int {
	q1 := []int{n}
	q2 := []int{}

	for i := 0; i < 25; i++ {
		for _, q := range q1 {
			if a, b, two := blink(q); two {
				q2 = append(q2, a, b)
			} else {
				q2 = append(q2, a)
			}
		}
		q1, q2 = q2, q1
		q2 = q2[:0]
	}
	return len(q1)
}

func blink(n int) (int, int, bool) {
	if n == 0 {
		return 1, 0, false
	}
	m := n
	d := 0
	for n > 0 {
		n = n / 10
		d++
	}
	if d%2 == 0 {
		mod := int(math.Pow10(d / 2))
		return m / mod, m % mod, true
	}
	return m * 2024, 0, false
}
