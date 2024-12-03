package main

import (
	"bufio"
	"fmt"
	"github.com/domcsore/advent-of-code-2024/util"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
)

func getDistances(a []int, b []int) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("Lists are not the same length")
	}

	totalDiff := 0.0

	for i := 0; i < len(a); i++ {
		x := a[i]
		y := b[i]

		distance := math.Abs(float64(x) - float64(y))
		totalDiff += distance
	}

	return int(totalDiff), nil
}

func getSimilarity(a []int, b []int) int {
	freq := make(map[int]int)

	similarity := 0

	for i := 0; i < len(a); i++ {
		if _, ok := freq[a[i]]; !ok {
			for j := 0; j < len(b); j++ {
				if b[j] == a[i] {
					freq[a[i]] += 1
				}
			}
		}

		if v, ok := freq[a[i]]; ok {
			similarity += (a[i] * v)
		}
	}

	return similarity
}

func main() {

	fileReader := util.OpenFileFromArgs()

	listOne := make([]int, 0)
	listTwo := make([]int, 0)

	for {
		readBytes, readErr := fileReader.ReadBytes('\n')
		if readErr == io.EOF {
			break
		}
		if len(readBytes) != 14 {
			fmt.Printf("%d\n", len(readBytes))
			panic("malformed input")
		}

		firstBytes := readBytes[:5]
		secondBytes := readBytes[8:13]

		firstNum, err := strconv.Atoi(string(firstBytes))
		if err != nil {
			panic("error reading number")
		}

		listOne = append(listOne, firstNum)

		secondNum, err := strconv.Atoi(string(secondBytes))
		if err != nil {
			panic("error reading number")
		}

		listTwo = append(listTwo, secondNum)
	}

	sort.Slice(listOne, func(a int, b int) bool {
		return listOne[a] < listOne[b]
	})

	sort.Slice(listTwo, func(a int, b int) bool {
		return listTwo[a] < listTwo[b]
	})

	distance, err := getDistances(listOne, listTwo)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Total distance = %d\n", distance)

	fmt.Printf("Similarity = %d\n", getSimilarity(listOne, listTwo))
}
