package main

import (
	"bytes"
	"fmt"
	"github.com/domcsore/advent-of-code-2024/util"
	"io"
	"strconv"
)

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

func IsDistanceSafe(a int, b int) bool {
	distance := abs(a - b)
	if distance < 1 || distance > 3 {
		return false
	} else {
		return true
	}
}

func IsLevelSafe(level []int) bool {
	isAscending := true
	isDescending := true

	for i := 1; i < len(level); i++ {

		if !IsDistanceSafe(level[i-1], level[i]) {
			return false
		}

		if level[i-1] >= level[i] {
			isAscending = false
		}

		if level[i-1] <= level[i] {
			isDescending = false
		}
	}

	return isAscending || isDescending
}

func RemoveIndex(level []int, index int) []int {
	newLevel := make([]int, 0)
	newLevel = append(newLevel, level[:index]...)
	newLevel = append(newLevel, level[index+1:]...)

	return newLevel
}

func IsSafeWithRemovedLevel(level []int) bool {
	if IsLevelSafe(level) {
		return true
	}

	for i := 0; i < len(level); i++ {
		newLevel := RemoveIndex(level, i)
		if IsLevelSafe(newLevel) {
			return true
		}
	}

	return false
}

func main() {
	inputFileReader := util.OpenFileFromArgs()

	safeCount := 0
	safeCountWithDrop := 0

	for {
		levelLine, err := inputFileReader.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		level := getLevelFromLine(levelLine)
		levelIsSafe := IsLevelSafe(level)
		if levelIsSafe {
			safeCount++
		}

		levelIsSafeWithDrop := IsSafeWithRemovedLevel(level)
		if levelIsSafeWithDrop {
			safeCountWithDrop++
		}
	}

	fmt.Printf("Safe levels = %d\n", safeCount)
	fmt.Printf("Safe levels with drop = %d\n", safeCountWithDrop)
}
