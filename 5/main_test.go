package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func getTestReader() *bufio.Reader {
	testFile, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}
	return bufio.NewReader(testFile)
}

func testUpdate(t *testing.T, index int, expect bool) {
	testReader := getTestReader()

	rules, updates, err := RulesAndUpdatesFromFileReader(testReader)
	if err != nil {
		panic(err)
	}

	updateToTest := updates[index]
	fmt.Printf("%v\n", updateToTest)

	result := isUpdateCorrect(updateToTest, rules)

	if result != expect {
		t.Errorf("update: %v\n", updateToTest)
	}
}

func TestFirstCorrectUpdate(t *testing.T) {
	testUpdate(t, 0, true)
}

func TestSecondCorrectUpdate(t *testing.T) {
	testUpdate(t, 1, true)
}

func TestThirdCorrectUpdate(t *testing.T) {
	testUpdate(t, 2, true)
}

func TestFourthIncorrectUpdate(t *testing.T) {
	testUpdate(t, 3, false)
}

func TestFifthIncorrectUpdate(t *testing.T) {
	testUpdate(t, 4, false)
}

func TestSixthIncorrectUpdate(t *testing.T) {
	testUpdate(t, 5, false)
}
