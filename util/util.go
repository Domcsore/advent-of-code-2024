package util

import (
	"bufio"
	"os"
	"strconv"
)

func OpenFileFromArgs() *bufio.Reader {
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

func IntFromBytes(b []byte) (int, error) {
	return strconv.Atoi(string(b))
}
