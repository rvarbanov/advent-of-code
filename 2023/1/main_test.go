package main

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type MainTestSuite struct {
	suite.Suite
}

func (suite *MainTestSuite) SetupTest() {
}
func (suite *MainTestSuite) TearDownTest() {
}

func (suite *MainTestSuite) Test_calcOne() {
	input := []byte("1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet\nasdf")
	expectedSum := 12 + 38 + 15 + 77
	sum := calculate(input)
	suite.Equal(expectedSum, sum)
}

func (suite *MainTestSuite) Test_calcTwo() {
	input := []byte("two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen")
	expectedSum := 29 + 83 + 13 + 24 + 42 + 14 + 76
	sum := calculate(input)
	suite.Equal(expectedSum, sum)
}

func (suite *MainTestSuite) Test_getIntNumbersFromLine() {
	testCases := map[string]struct {
		line     string
		expected []Number
	}{
		"1abc2": {
			line:     "1abc2",
			expected: []Number{{0, 1}, {4, 2}},
		},
		"a1abc2": {
			line:     "a1abc2",
			expected: []Number{{1, 1}, {5, 2}},
		},
		"aabc2": {
			line:     "aabc2",
			expected: []Number{{4, 2}},
		},
		"aabc": {
			line:     "aabc",
			expected: nil,
		},
	}

	for name, tc := range testCases {
		suite.Run(name, func() {
			numbers := getIntNumbersFromLine(tc.line)
			suite.Equal(tc.expected, numbers)
		})
	}
}

func (suite *MainTestSuite) Test_getCharNumbersFromLine() {
	testCases := map[string]struct {
		line     string
		expected []Number
	}{
		"two1nine": {
			line:     "two1nine",
			expected: []Number{{0, 2}, {4, 9}},
		},
		"eightwothree": {
			line:     "eightwothree",
			expected: []Number{{0, 8}, {4, 2}, {7, 3}},
		},
	}

	for name, tc := range testCases {
		suite.Run(name, func() {
			numbers := getCharNumbersFromLine(tc.line)
			suite.Equal(len(tc.expected), len(numbers))
			suite.Equal(tc.expected, numbers)
		})
	}
}

func (suite *MainTestSuite) Test_getLineNumbers() {
	//_ = []byte("\n\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen")
	//[]byte("")
	testCases := map[string]struct {
		iNumbers []Number
		cNumbers []Number
		expected []int
	}{
		"two1nine": {
			iNumbers: getIntNumbersFromLine("two1nine"),
			cNumbers: getCharNumbersFromLine("two1nine"),
			expected: []int{2, 9},
		},
		"eightwothree": {
			iNumbers: getIntNumbersFromLine("eightwothree"),
			cNumbers: getCharNumbersFromLine("eightwothree"),
			expected: []int{8, 3},
		},
		"abcone2threexyz": {
			iNumbers: getIntNumbersFromLine("abcone2threexyz"),
			cNumbers: getCharNumbersFromLine("abcone2threexyz"),
			expected: []int{1, 3},
		},
		"xtwone3four": {
			iNumbers: getIntNumbersFromLine("xtwone3four"),
			cNumbers: getCharNumbersFromLine("xtwone3four"),
			expected: []int{2, 4},
		},
		"4nineeightseven2": {
			iNumbers: getIntNumbersFromLine("4nineeightseven2"),
			cNumbers: getCharNumbersFromLine("4nineeightseven2"),
			expected: []int{4, 2},
		},
		"zoneight234": {
			iNumbers: getIntNumbersFromLine("zoneight234"),
			cNumbers: getCharNumbersFromLine("zoneight234"),
			expected: []int{1, 4},
		},
		"7pqrstsixteen": {
			iNumbers: getIntNumbersFromLine("7pqrstsixteen"),
			cNumbers: getCharNumbersFromLine("7pqrstsixteen"),
			expected: []int{7, 6},
		},
		"7": {
			iNumbers: getIntNumbersFromLine("7"),
			cNumbers: getCharNumbersFromLine("7"),
			expected: []int{7, 7},
		},
		"six": {
			iNumbers: getIntNumbersFromLine("six"),
			cNumbers: getCharNumbersFromLine("six"),
			expected: []int{6, 6},
		},
		"n7": {
			iNumbers: getIntNumbersFromLine("n7"),
			cNumbers: getCharNumbersFromLine("n7"),
			expected: []int{7, 7},
		},
		"2dcvcqcbpshsixone3": {
			iNumbers: getIntNumbersFromLine("2dcvcqcbpshsixone3"),
			cNumbers: getCharNumbersFromLine("2dcvcqcbpshsixone3"),
			expected: []int{2, 3},
		},
	}

	for testName, testCase := range testCases {
		suite.Run(testName, func() {
			numbers := getLineNumbers(testCase.iNumbers, testCase.cNumbers)
			suite.Equal(testCase.expected, numbers)
		})
	}
}

func (suite *MainTestSuite) Test_sortNumbers() {
	testCases := map[string]struct {
		input    []Number
		expected []Number
	}{
		"ordered": {
			input:    []Number{{0, 2}, {1, 9}},
			expected: []Number{{0, 2}, {1, 9}},
		},
		"not ordered": {
			input:    []Number{{5, 2}, {1, 9}},
			expected: []Number{{1, 9}, {5, 2}},
		},
		"mixed": {
			input:    []Number{{0, 3}, {5, 2}, {1, 9}},
			expected: []Number{{0, 3}, {1, 9}, {5, 2}},
		},
	}

	for testName, testCase := range testCases {
		suite.Run(testName, func() {
			numbers := sortNumbers(testCase.input)
			suite.Equal(testCase.expected, numbers)
		})
	}
}

func TestMainTestSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}
