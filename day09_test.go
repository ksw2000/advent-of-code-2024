package main

import (
	"container/list"
	"fmt"
	"os"
	"testing"
	"time"
)

// const day9Example = `2333133121414131402`

func TestDay9Part1(t *testing.T) {
	f, err := os.Open("data/day9.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// f := strings.NewReader(day9Example)

	start := time.Now()

	disk := []int{}
	free := []int{}
	var n byte
	var round int
	for x, _ := fmt.Fscanf(f, "%c", &n); x > 0; x, _ = fmt.Fscanf(f, "%c", &n) {
		if round&1 == 0 {
			disk = append(disk, int(n-'0'))
		} else {
			free = append(free, int(n-'0'))
		}
		round++
	}

	checksum := 0
	offset := 0
	for i := 0; i < len(disk); i++ {
		// consume disk
		for j := 0; j < disk[i]; j++ {
			checksum += i * offset
			offset++
		}
		// process free space
		if i < len(free) {
			f := free[i]
			for j := len(disk) - 1; j > i; j-- {
				d := disk[j]
				for k := 0; k < d && f > 0; k++ {
					checksum += j * offset
					offset++
					f--
					disk[j]--
				}
				if disk[j] == 0 {
					disk = disk[:len(disk)-1]
				}
			}
		}
	}

	fmt.Printf("%.3f ms\n", float64(time.Since(start).Microseconds())*0.001)
	fmt.Println(checksum)
}

func TestDay9Part2LinkedList(t *testing.T) {
	f, err := os.Open("data/day9.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// f := strings.NewReader(day9Example)

	start := time.Now()

	type entry struct {
		length int
		label  int
		isFree bool
	}
	var n byte
	var round int
	label := 0
	l := list.New()
	looked := -1
	for x, _ := fmt.Fscanf(f, "%c", &n); x > 0; x, _ = fmt.Fscanf(f, "%c", &n) {
		l.PushBack(&entry{
			length: int(n - '0'),
			label:  label,
			isFree: round&1 == 1,
		})
		if round&1 == 0 {
			looked = label
		} else {
			label++
		}
		round++
	}

	// DEBUG PRINT:
	// for i, cur := 0, l.Front(); i < l.Len(); i, cur = i+1, cur.Next() {
	// 	e := cur.Value.(*entry)
	// 	for j := 0; j < e.length; j++ {
	// 		if e.isFree {
	// 			fmt.Print(".")
	// 		} else {
	// 			fmt.Print(e.label)
	// 		}
	// 	}
	// }
	// fmt.Println()

	for cur0 := l.Back(); cur0 != l.Front(); {
		pre := cur0.Prev()
		entry0 := cur0.Value.(*entry)
		if entry0.isFree || entry0.label > looked {
			cur0 = pre
			continue
		}

		// try to move this block to the free space
		for cur1 := l.Front(); cur1 != cur0; cur1 = cur1.Next() {
			if entry1 := cur1.Value.(*entry); entry1.isFree && entry1.length >= entry0.length {
				// fmt.Printf("%d find the free space, move it!\n", entry0.label)
				entry0.isFree = true
				l.InsertBefore(&entry{
					length: entry0.length,
					label:  entry0.label,
					isFree: false,
				}, cur1)
				entry1.length -= entry0.length
				break
			}
		}
		looked = entry0.label - 1
		cur0 = pre
	}

	checksum := 0
	offset := 0
	for i, cur := 0, l.Front(); i < l.Len(); i, cur = i+1, cur.Next() {
		e := cur.Value.(*entry)
		if !e.isFree {
			checksum += e.label * (offset + (offset + e.length - 1)) * e.length / 2
		}
		offset += e.length
	}

	fmt.Printf("%.3f ms\n", float64(time.Since(start).Microseconds())*0.001)
	fmt.Println("checksum =", checksum)
}

func TestDay9Part2(t *testing.T) {
	f, err := os.Open("data/day9.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// f := strings.NewReader(day9Example)

	start := time.Now()

	type entry struct {
		offset int
		length int
	}

	var n byte
	var round bool
	data := []entry{}
	free := []entry{}
	offset := 0
	for x, _ := fmt.Fscanf(f, "%c", &n); x > 0; x, _ = fmt.Fscanf(f, "%c", &n) {
		e := entry{
			offset: offset,
			length: int(n - '0'),
		}

		if !round {
			data = append(data, e)
		} else {
			free = append(free, e)
		}

		offset += e.length
		round = !round
	}

	checksum := 0

	for i := len(data) - 1; i >= 0; i-- {
		for j := 0; j < len(free) && free[j].offset < data[i].offset; j++ {
			if free[j].length >= data[i].length {
				// move data
				data[i].offset = free[j].offset
				free[j].length -= data[i].length
				free[j].offset += data[i].length
				break
			}
		}
		checksum += i * (data[i].offset*2 + data[i].length - 1) * data[i].length / 2
	}

	fmt.Printf("%.3f ms\n", float64(time.Since(start).Microseconds())*0.001)
	fmt.Println("checksum =", checksum)
}
