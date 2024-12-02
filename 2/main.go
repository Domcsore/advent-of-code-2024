package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func openFileFromArgs() *bufio.Reader {
	if len(os.Args) < 2 {
		panic("Must have input file as argument")
	}

	inputPath := os.Args[1]

	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}

	return bufio.NewReader(file)
}

func getLevelFromLine(levelLine []byte) []int {
	levelLine = levelLine[:len(levelLine)-1]

	levelBytes := bytes.Split(levelLine, []byte(" "))

	level := make([]int, 0)

	for _, v := range levelBytes {
		levelVal, err := strconv.Atoi(string(v))
		if err != nil {
			panic(err)
		}
		level = append(level, levelVal)
	}

	return level
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func isSafe(level []int) bool {
	isSafe := true
	isSigned := false
	for i := 1; i < len(level); i++ {
		diff := level[i-1] - level[i]
		absDiff := abs(diff)
		if absDiff > 3 || absDiff < 1 {
			isSafe = false
		}

		if i == 1 {
			isSigned = math.Signbit(float64(diff))
		} else if math.Signbit(float64(diff)) != isSigned {
			isSafe = false
		}
	}

	return isSafe
}

func main() {
	inputFileReader := openFileFromArgs()

	safeCount := 0

	for {
		levelLine, err := inputFileReader.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		level := getLevelFromLine(levelLine)
		levelIsSafe := isSafe(level)
		if levelIsSafe {
			safeCount++
		}
	}

	fmt.Printf("Safe levels = %d", safeCount)
}
