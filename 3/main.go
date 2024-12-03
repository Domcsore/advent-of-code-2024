package main

import (
	"fmt"
	"io"
	"regexp"

	"github.com/domcsore/advent-of-code-2024/util"
)

func main() {
	inputReader := util.OpenFileFromArgs()

	input, err := io.ReadAll(inputReader)
	if err != nil {
		panic(err)
	}

	mulRe := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	matches := mulRe.FindAll(input, -1)

	if matches == nil {
		fmt.Println("No matches found")
		return
	}

	digRe := regexp.MustCompile(`\d{1,3}`)

	total := 0

	for i := 0; i < len(matches); i++ {
		multiplicants := digRe.FindAll(matches[i], -1)
		if len(multiplicants) != 2 {
			panic("Need exactly two multiplicants")
		}

		multiplicantOne, err := util.IntFromBytes(multiplicants[0])
		if err != nil {
			panic(err)
		}

		multiplicantTwo, err := util.IntFromBytes(multiplicants[1])
		if err != nil {
			panic(err)
		}

		total += multiplicantOne * multiplicantTwo
	}

	fmt.Printf("Total - %d\n", total)
}
