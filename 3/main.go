package main

import (
	"fmt"
	"io"
	"regexp"

	"github.com/domcsore/advent-of-code-2024/util"
)

func GetMulCommands(b []byte) [][]byte {
	mulRe := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	return mulRe.FindAll(b, -1)
}

func GetMulCommandMultiplicants(b []byte) (int, int, error) {
	digRe := regexp.MustCompile(`\d{1,3}`)
	multiplicants := digRe.FindAll(b, -1)

	if len(multiplicants) != 2 {
		return 0, 0, fmt.Errorf("mul commands must contain 2 numbers")
	}

	multiplicantOne, err := util.IntFromBytes(multiplicants[0])
	if err != nil {
		return 0, 0, err
	}

	multiplicantTwo, err := util.IntFromBytes(multiplicants[1])
	if err != nil {
		return 0, 0, err
	}

	return multiplicantOne, multiplicantTwo, nil
}

func CalculateAllMulInstructions(b []byte) int {
	matches := GetMulCommands(b)
	if matches == nil {
		return 0
	}

	total := 0

	for i := 0; i < len(matches); i++ {
		multiplicantOne, multiplicantTwo, err := GetMulCommandMultiplicants(matches[i])
		if err != nil {
			panic(err)
		}

		total += multiplicantOne * multiplicantTwo
	}

	return total
}

func GetDoBlocks(b []byte) [][]byte {
	doBlockRe := regexp.MustCompile(`(\A|do\(\))(.|\n)*?(don\'t\(\)|\z)`)

	return doBlockRe.FindAll(b, -1)
}

func CalculateDoBlockInstructions(b []byte) int {
	doBlocks := GetDoBlocks(b)

	total := 0

	for i := 0; i < len(doBlocks); i++ {
		// fmt.Printf("%s\f", doBlocks[i])
		blockTotal := CalculateAllMulInstructions(doBlocks[i])
		total += blockTotal
	}

	return total
}

func main() {
	inputReader := util.OpenFileFromArgs()

	input, err := io.ReadAll(inputReader)
	if err != nil {
		panic(err)
	}

	total := CalculateAllMulInstructions(input)
	fmt.Printf("Total - %d\n", total)

	doBlockTotal := CalculateDoBlockInstructions(input)
	fmt.Printf("Do Block Total - %d\n", doBlockTotal)
}
