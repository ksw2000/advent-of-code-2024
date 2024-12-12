package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
	"testing"
)

const day5part1Example = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

// Since the section 1 dependencies contains cycle, i.e. is not a DAG
// Thus, we cannot use topological sort to decide its order
func TestDay5Part1WRONG(t *testing.T) {
	f := strings.NewReader(day5part1Example)

	adjList := map[int][]int{}
	inDegree := map[int]int{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		var p, q int
		fmt.Sscanf(line, "%d|%d", &p, &q)
		adjList[p] = append(adjList[p], q)

		inDegree[q]++
		if _, ok := inDegree[p]; !ok {
			inDegree[p] = 0
		}
	}

	// topological sort

	labels := map[int]int{}
	label := 0

	q1 := []int{}
	for k := range inDegree {
		if inDegree[k] == 0 {
			q1 = append(q1, k)
			labels[k] = label
		}
	}

	q2 := []int{}
	for len(q1) > 0 {
		label++
		for i := range q1 {
			for _, n := range adjList[q1[i]] {
				inDegree[n]--
				if inDegree[n] == 0 {
					labels[n] = label
					q2 = append(q2, n)
				}
			}
		}
		q2, q1 = q1, q2
		q2 = q2[:0]
	}

	fmt.Println(labels)

	sum := 0
scan:
	for scanner.Scan() {
		line := scanner.Bytes()
		reader := bytes.NewReader(line)

		var n int
		list := []int{}
		x, _ := fmt.Fscanf(reader, "%d", &n)
		pre := -1
		for x > 0 {
			if labels[n] < pre {
				continue scan
			}
			pre = labels[n]
			list = append(list, n)
			x, _ = fmt.Fscanf(reader, ",%d", &n)
		}
		sum += list[len(list)/2]
	}
	fmt.Println(sum)
}

func TestDay5Part1(t *testing.T) {
	f, err := os.Open("data/day5.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	adjList := map[int][]int{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		var p, q int
		fmt.Sscanf(line, "%d|%d", &p, &q)
		adjList[p] = append(adjList[p], q)
	}

	sum := 0
scan:
	for scanner.Scan() {
		line := scanner.Bytes()
		reader := bytes.NewReader(line)

		var n int
		list := []int{}
		x, _ := fmt.Fscanf(reader, "%d", &n)
		for x > 0 {
			list = append(list, n)
			x, _ = fmt.Fscanf(reader, ",%d", &n)
		}

		for i := 1; i < len(list); i++ {
			// check list[i-1] -> list[i]
			if ok := slices.Contains(adjList[list[i-1]], list[i]); !ok {
				continue scan
			}
		}

		sum += list[len(list)/2]
	}
	fmt.Println(sum)
}

func TestDay5Part2(t *testing.T) {
	f, err := os.Open("data/day5.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	type pair struct {
		u int
		v int
	}

	dependencies := map[pair]bool{}

	// set true if u before v

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		var p, q int
		fmt.Sscanf(line, "%d|%d", &p, &q)
		dependencies[pair{
			p, q,
		}] = true
	}

	sum := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		reader := bytes.NewReader(line)

		var n int
		list := []int{}
		x, _ := fmt.Fscanf(reader, "%d", &n)
		for x > 0 {
			list = append(list, n)
			x, _ = fmt.Fscanf(reader, ",%d", &n)
		}

		needSorting := false
		for i := 1; i < len(list); i++ {
			if !dependencies[pair{
				list[i-1], list[i],
			}] {
				needSorting = true
				break
			}
		}

		if needSorting {
			sort.Slice(list, func(i, j int) bool {
				return dependencies[pair{
					list[i], list[j],
				}]
			})
			sum += list[len(list)/2]
		}
	}
	fmt.Println(sum)
}
