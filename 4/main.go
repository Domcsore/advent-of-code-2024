package main

import (
	"bufio"
	"fmt"
	"io"

	"github.com/domcsore/advent-of-code-2024/util"
)

func InputFromFileReader(i *bufio.Reader) [][]rune {
	input := make([][]rune, 1)

	currentRow := 0
	for {
		inputRune, _, err := i.ReadRune()
		if err == io.EOF {
			// Remove potentially empty last row
			if len(input[currentRow]) == 0 {
				input = input[:len(input)-1]
			}
			break
		}

		if inputRune == '\n' {
			currentRow++

			if len(input)-1 != currentRow {
				input = append(input, make([]rune, 0))
			}
		} else {
			input[currentRow] = append(input[currentRow], inputRune)
		}
	}

	return input
}

type WordSearch struct {
	data      [][]rune
	rows      int
	columns   int
	foundMask [][]bool
	crossMask [][]bool
}

func (ws *WordSearch) Print(mask int) {
	if mask == 1 {
		printWordSearch(ws.data, ws.foundMask)
	} else {
		printWordSearch(ws.data, ws.crossMask)
	}
}

func (ws *WordSearch) hasWordsAtPosition(row int, column int, word []rune) (foundCount int) {
	foundCount = 0

	wordLength := len(word)

	wordColumnEndRight := column + wordLength
	wordColumnEndLeft := column - wordLength
	wordRowEndDown := row + wordLength
	wordRowEndUp := row - wordLength

	wordFitsRight := wordColumnEndRight <= ws.columns
	wordFitsLeft := wordColumnEndLeft >= -1
	wordFitsDown := wordRowEndDown <= ws.rows
	wordFitsUp := wordRowEndUp >= -1

	// Look Right
	if wordFitsRight {
		wordFound := true
		for i, letter := range word {
			if ws.data[row][column+i] != letter {
				wordFound = false
				break
			}
		}
		if wordFound {
			for i := range word {
				ws.foundMask[row][column+i] = true
			}
			foundCount++
		}
	}

	// Look Left
	if wordFitsLeft {
		wordFound := true
		for i, letter := range word {
			if ws.data[row][column-i] != letter {
				wordFound = false
				break
			}
		}
		if wordFound {
			for i := range word {
				ws.foundMask[row][column-i] = true
			}
			foundCount++
		}
	}

	// Look Down
	if wordFitsDown {
		wordFound := true
		for i, letter := range word {
			if ws.data[row+i][column] != letter {
				wordFound = false
				break
			}
		}
		if wordFound {
			for i := range word {
				ws.foundMask[row+i][column] = true
			}
			foundCount++
		}
	}

	// Look Up
	if wordFitsUp {
		wordFound := true
		for i, letter := range word {
			if ws.data[row-i][column] != letter {
				wordFound = false
				break
			}
		}
		if wordFound {
			for i := range word {
				ws.foundMask[row-i][column] = true
			}
			foundCount++
		}
	}

	// Look Right Down
	if wordFitsRight && wordFitsDown {
		wordFound := true
		for i, letter := range word {
			if ws.data[row+i][column+i] != letter {
				wordFound = false
				break
			}
		}
		if wordFound {
			for i := range word {
				ws.foundMask[row+i][column+i] = true
			}
			foundCount++
		}
	}

	// Look Left Down
	if wordFitsLeft && wordFitsDown {
		wordFound := true
		for i, letter := range word {
			if ws.data[row+i][column-i] != letter {
				wordFound = false
				break
			}
		}
		if wordFound {
			for i := range word {
				ws.foundMask[row+i][column-i] = true
			}
			foundCount++
		}
	}

	// Look Left Up
	if wordFitsLeft && wordFitsUp {
		wordFound := true
		for i, letter := range word {
			if ws.data[row-i][column-i] != letter {
				wordFound = false
				break
			}
		}
		if wordFound {
			for i := range word {
				ws.foundMask[row-i][column-i] = true
			}
			foundCount++
		}
	}

	// Look Right Up
	if wordFitsRight && wordFitsUp {
		wordFound := true
		for i, letter := range word {
			if ws.data[row-i][column+i] != letter {
				wordFound = false
				break
			}
		}
		if wordFound {
			for i := range word {
				ws.foundMask[row-i][column+i] = true
			}
			foundCount++
		}
	}
	return foundCount
}

func (ws *WordSearch) FindWords(word []rune) (foundCount int) {
	foundCount = 0
	for rowIndex, row := range ws.data {
		for columnIndex := range row {
			foundCount += ws.hasWordsAtPosition(rowIndex, columnIndex, word)
		}
	}

	return foundCount
}

func (ws *WordSearch) hasCrossMasAtPosition(row int, column int) (found bool) {
	if row == 0 || row == (ws.rows-1) || column == 0 || (column == ws.columns-1) {
		return false
	}

	if ws.data[row][column] != 'A' {
		return false
	}

	// TL to BR
	topLeftRune := ws.data[row-1][column-1]

	if topLeftRune != 'M' && topLeftRune != 'S' {
		return false
	}

	bottomRightRune := ws.data[row+1][column+1]

	oppoTopLeftRune, err := OppositeSorM(topLeftRune)
	if err != nil {
		panic(err)
	}

	if bottomRightRune != oppoTopLeftRune {
		return false
	}

	// TR to BL
	topRightRune := ws.data[row-1][column+1]

	if topRightRune != 'M' && topRightRune != 'S' {
		return false
	}

	bottomLeftRune := ws.data[row+1][column-1]

	oppoTopRightRune, err := OppositeSorM(topRightRune)
	if err != nil {
		panic(err)
	}

	if bottomLeftRune != oppoTopRightRune {
		return false
	}

	ws.crossMask[row][column] = true
	ws.crossMask[row+1][column+1] = true
	ws.crossMask[row+1][column-1] = true
	ws.crossMask[row-1][column-1] = true
	ws.crossMask[row-1][column+1] = true

	return true
}

func (ws *WordSearch) FindCrossMasWords() int {
	foundCount := 0
	for ri, r := range ws.data {
		for ci := range r {
			found := ws.hasCrossMasAtPosition(ri, ci)
			if found {
				foundCount++
			}
		}
	}
	return foundCount
}

func OppositeSorM(r rune) (rune, error) {
	if r == 'M' {
		return 'S', nil
	}

	if r == 'S' {
		return 'M', nil
	}

	return 'A', fmt.Errorf("expected rune S or M, recieved %c", r)

}

func printWordSearch(ws [][]rune, mask [][]bool) {
	for ri, r := range ws {
		for ci, c := range r {
			if mask[ri][ci] {
				fmt.Printf("\033[32m%s\033[0m", string(c))
			} else {
				fmt.Printf("%s", string(c))
			}
		}
		fmt.Print("\n")
	}
}

func CreateWordSearchFromFileReader(i *bufio.Reader) (*WordSearch, error) {
	data := InputFromFileReader(i)

	rows := len(data)

	if rows < 1 {
		return &WordSearch{}, fmt.Errorf("word search must have at least one row")
	}

	columns := len(data[0])

	for _, r := range data {
		if len(r) != columns {
			return &WordSearch{}, fmt.Errorf("word search must have the same amount of columns in each row")
		}
	}

	foundMask := make([][]bool, rows)
	for i := 0; i < len(foundMask); i++ {
		foundMask[i] = make([]bool, columns)
	}

	crossMask := make([][]bool, rows)
	for i := 0; i < len(crossMask); i++ {
		crossMask[i] = make([]bool, columns)
	}

	return &WordSearch{
		data,
		rows,
		columns,
		foundMask,
		crossMask,
	}, nil
}

func main() {
	inputReader := util.OpenFileFromArgs()

	ws, err := CreateWordSearchFromFileReader(inputReader)
	if err != nil {
		panic(err)
	}

	wordsFound := ws.FindWords([]rune("XMAS"))
	fmt.Printf("Word found - %d\n", wordsFound)

	crossCount := ws.FindCrossMasWords()
	fmt.Printf("Cross Mas's found - %d\n", crossCount)
}
