// 1.0 53859
// 1.1 53866
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	strNums = map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}
)

type Number struct {
	Idx    int
	Number int
}

func sortNumbers(numbers []Number) []Number {
	for i := 0; i < len(numbers); i++ {
		for j := i; j < len(numbers); j++ {
			if numbers[i].Idx > numbers[j].Idx {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}

	return numbers
}

func main() {
	input, err := os.ReadFile("inputTwo")
	if err != nil {
		panic(err)
	}

	total := calculate(input)
	fmt.Println(total)
}

func calculate(input []byte) int {
	var allNumbers []int

	for lineIdx, line := range strings.Split(string(input), "\n") {
		fmt.Println(lineIdx, line)

		iNumbers := getIntNumbersFromLine(line)
		fmt.Println(iNumbers)
		cNumbers := getCharNumbersFromLine(line)
		fmt.Println(cNumbers)

		lNumbers := getLineNumbers(iNumbers, cNumbers)
		fmt.Println(lNumbers)

		if len(lNumbers) < 1 {
			continue
		}
		lineNumSum := lNumbers[0]*10 + lNumbers[1]
		fmt.Println(lineNumSum)
		allNumbers = append(allNumbers, lineNumSum)
	}

	totalSum := 0
	for _, sum := range allNumbers {
		totalSum += sum
	}

	return totalSum
}

func getIntNumbersFromLine(line string) []Number {
	var numbers []Number
	for chIdx, ch := range line {
		char := string(ch)
		n, ok := strconv.Atoi(char)
		if ok != nil {
			continue
		}
		number := Number{
			Idx:    chIdx,
			Number: n,
		}
		numbers = append(numbers, number)
	}

	return numbers
}

func getCharNumbersFromLine(line string) []Number {
	var cNumbers []Number

	for intNum, strNum := range strNums {
		numIdx := strings.Index(line, strNum)
		if numIdx == -1 {
			continue
		}

		rNumber := Number{numIdx, intNum}
		cNumbers = append(cNumbers, rNumber)
	}

	cNumbers = sortNumbers(cNumbers)
	return cNumbers
}

func getLineNumbers(iNumbers []Number, cNumbers []Number) []int {
	firstNum := 0
	lastNum := 0
	var lNumbers []int

	if len(iNumbers) < 1 && len(cNumbers) < 1 {
		return lNumbers
	}

	if len(iNumbers) == 0 {
		firstNum = cNumbers[0].Number
		lastNum = cNumbers[len(cNumbers)-1].Number
		lNumbers = []int{firstNum, lastNum}
		return lNumbers
	}
	if len(cNumbers) == 0 {
		firstNum = iNumbers[0].Number
		lastNum = iNumbers[len(iNumbers)-1].Number
		lNumbers = []int{firstNum, lastNum}
		return lNumbers
	}

	// find the first number
	if iNumbers[0].Idx < cNumbers[0].Idx {
		firstNum = iNumbers[0].Number
	}

	if cNumbers[0].Idx < iNumbers[0].Idx {
		firstNum = cNumbers[0].Number
	}

	// find the last number
	if iNumbers[len(iNumbers)-1].Idx > cNumbers[len(cNumbers)-1].Idx {
		lastNum = iNumbers[len(iNumbers)-1].Number
	}

	if cNumbers[len(cNumbers)-1].Idx > iNumbers[len(iNumbers)-1].Idx {
		lastNum = cNumbers[len(cNumbers)-1].Number
	}

	lNumbers = []int{firstNum, lastNum}

	return lNumbers
}
