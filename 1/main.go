package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	if len(os.Args) < 2 {
		panic("must provide input file")
	}

	inputPath := os.Args[1]

	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}

	fileReader := bufio.NewReader(file)

	heapOne := &IntHeap{}
	heap.Init(heapOne)
	heapTwo := &IntHeap{}
	heap.Init(heapTwo)

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

		heap.Push(heapOne, firstNum)

		secondNum, err := strconv.Atoi(string(secondBytes))
		if err != nil {
			panic("error reading number")
		}

		heap.Push(heapTwo, secondNum)
	}

	totalDiff := 0.0

	for heapOne.Len() > 0 {
		firstNum := heap.Pop(heapOne)
		secondNum := heap.Pop(heapTwo)

		absDiff := math.Abs(float64(secondNum.(int) - firstNum.(int)))
		totalDiff += absDiff
	}

	fmt.Printf("Total distance = %f\n", totalDiff)
}
