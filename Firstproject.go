package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add(arg1, arg2 int) int {
	return arg1 + arg2
}

func sub(arg1, arg2 int) int {
	return arg1 - arg2
}

func div(arg1, arg2 int) int {
	return arg1 / arg2
}

func mul(arg1, arg2 int) int {
	return arg1 * arg2
}

func main() {
	readerAsVariable := bufio.NewReader(os.Stdin)
	lineFromTerminal, err := readerAsVariable.ReadString('\n')
	if err != nil {
		panic(err)
	}

	lineFromTerminal = strings.TrimSpace(lineFromTerminal)
	splittedLineFromTerminal := strings.Fields(lineFromTerminal)
	if len(splittedLineFromTerminal) != 3 {
		panic("кол-во символов не равно 3")
	}

	firstNumberAsString := splittedLineFromTerminal[0]
	operator := splittedLineFromTerminal[1]
	secondNumberAsString := splittedLineFromTerminal[2]


	var isRoman bool
	if romanToInt(secondNumberAsString) != 0 {
		isRoman = true
	}

	var firstNumberAsInt, secondNumberAsInt int
	if isRoman {
		firstNumberAsInt = romanToInt(firstNumberAsString)
		secondNumberAsInt = romanToInt(secondNumberAsString)
	} else {
		firstNumberAsInt, _ = strconv.Atoi(firstNumberAsString)
		secondNumberAsInt, _ = strconv.Atoi(secondNumberAsString)
	}

	if !isRoman && (firstNumberAsInt > 10 || secondNumberAsInt > 10) {
		panic("арабское число больше 10")
	}
	
	if firstNumberAsInt == 0 || secondNumberAsInt == 0 {
		panic("арабское число не может быть 0")
	}

	var resultAsInt int
	switch operator {
	case "+":
		resultAsInt = add(firstNumberAsInt, secondNumberAsInt)
	case "-":
		resultAsInt = sub(firstNumberAsInt, secondNumberAsInt)
	case "*":
		resultAsInt = mul(firstNumberAsInt, secondNumberAsInt)
	case "/":
		resultAsInt = div(firstNumberAsInt, secondNumberAsInt)
	}

	if resultAsInt < 1 && isRoman {
		panic("результат отрицательный")
	}

	var resultAsString string
	if isRoman {
		resultAsString = intToRoman(resultAsInt)
	} else {
		resultAsString = strconv.Itoa(resultAsInt)
	}

	fmt.Println(resultAsString)
}

func romanToInt(s string) int {
	know := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	lengthOfString := len(s)
	lastElement := s[len(s)-1 : lengthOfString]
	var result int
	result = know[lastElement]
	for i := len(s) - 1; i > 0; i-- {
		if know[s[i:i+1]] <= know[s[i-1:i]] {
			result += know[s[i-1:i]]
		} else {
			result -= know[s[i-1:i]]
		}
	}
	return result
}

func intToRoman(num int) string {
	var roman string = ""
	var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	var index = len(romans) - 1

	for num > 0 {
		for numbers[index] <= num {
			roman += romans[index]
			num -= numbers[index]
		}
		index -= 1
	}

	return roman
}
