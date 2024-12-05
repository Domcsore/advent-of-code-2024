package main

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"

	"github.com/domcsore/advent-of-code-2024/util"
)

func ruleFromBytes(b []byte) ([]int, error) {
	x, err := strconv.Atoi(string(b[:2]))
	if err != nil {
		return nil, err
	}

	y, err := strconv.Atoi(string(b[3:]))
	if err != nil {
		return nil, err
	}

	return []int{x, y}, nil
}

func updateFromBytes(b []byte) ([]int, error) {
	updates := make([]int, 0)
	for i := 0; i < len(b); i += 3 {
		update, err := strconv.Atoi(string(b[i : i+2]))
		if err != nil {
			return nil, err
		}

		updates = append(updates, update)
	}

	return updates, nil
}

func isUpdateCorrect(update []int, rules map[int][]int) bool {
	for i := 1; i < len(update); i++ {
		if mustBeBefore, ok := rules[update[i]]; ok {
			toCheck := update[:i]
			for _, v := range mustBeBefore {
				if slices.Contains(toCheck, v) {
					return false
				}
			}
		}
	}

	return true
}

func fixUpdate(update []int, rules map[int][]int) []int {

}

func middleIndex(length int) int {
	return int((length - 1) / 2)
}

func RulesAndUpdatesFromFileReader(fr *bufio.Reader) (rules map[int][]int, updates [][]int, err error) {
	rules = make(map[int][]int)
	updates = make([][]int, 0)

	for {
		lineWithNewLine, err := fr.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, nil, err
		}

		line := lineWithNewLine[:len(lineWithNewLine)-1]

		if len(line) == 0 {
			continue
		}

		isRule := line[2] == '|'
		if isRule {
			rule, err := ruleFromBytes(line)
			if err != nil {
				panic(err)
			}
			if _, ok := rules[rule[0]]; ok {
				rules[rule[0]] = append(rules[rule[0]], rule[1])
			} else {
				rules[rule[0]] = []int{rule[1]}
			}
		}

		isUpdate := line[2] == ','
		if isUpdate {
			update, err := updateFromBytes(line)
			if err != nil {
				panic(err)
			}
			updates = append(updates, update)
		}
	}

	return rules, updates, nil
}

func TotalOfCorrectUpdatesMidPoint(updates [][]int, rules map[int][]int) int {
	total := 0
	for _, u := range updates {
		if isUpdateCorrect(u, rules) {
			midIndex := middleIndex(len(u))
			total += u[midIndex]
		}
	}

	return total
}

func main() {
	inputReader := util.OpenFileFromArgs()

	rules, updates, err := RulesAndUpdatesFromFileReader(inputReader)
	if err != nil {
		panic(err)
	}

	for _, update := range updates {
		fmt.Printf("%v\n", update)
	}

	fmt.Println(TotalOfCorrectUpdatesMidPoint(updates, rules))

}
