package main

import (
	"fmt"
	"io"
	"os"

	"github.com/ksw2000/advent-of-code-2024/util"
)

type uf struct {
	url  string
	file string
}

func main() {
	list := []uf{
		{
			url:  "https://adventofcode.com/2024/day/1/input",
			file: "day1.txt",
		},
		{
			url:  "https://adventofcode.com/2024/day/2/input",
			file: "day2.txt",
		},
		{
			url:  "https://adventofcode.com/2024/day/3/input",
			file: "day3.txt",
		}, {
			url:  "https://adventofcode.com/2024/day/4/input",
			file: "day4.txt",
		}, {
			url:  "https://adventofcode.com/2024/day/5/input",
			file: "day5.txt",
		}, {
			url:  "https://adventofcode.com/2024/day/6/input",
			file: "day6.txt",
		}, {
			url:  "https://adventofcode.com/2024/day/7/input",
			file: "day7.txt",
		}, {
			url:  "https://adventofcode.com/2024/day/8/input",
			file: "day8.txt",
		}, {
			url:  "https://adventofcode.com/2024/day/9/input",
			file: "day9.txt",
		}, {
			url:  "https://adventofcode.com/2024/day/10/input",
			file: "day10.txt",
		}, {
			url:  "https://adventofcode.com/2024/day/11/input",
			file: "day11.txt",
		},
	}
	list = list[6:]
	for i := range list {
		fmt.Println(list[i])
		f, err := util.Fetch(list[i].url)
		if err != nil {
			panic(err)
		}

		g, err := os.Create(list[i].file)
		if err != nil {
			panic(err)
		}
		io.Copy(g, f)

		f.Close()
		g.Close()
	}
}
