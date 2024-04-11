package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(x, y int, sign string) int {
	switch sign {
		case "+": return x + y
		case "-": return x - y
		case "*": return x * y
		case "/": return x / y
		default: return 0
	}
}

func romanToInt(s string) int {
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	prev := 0

	for _, char := range s {
		value := romanMap[char]
		if value > prev {
			result += value - 2 * prev
		} else {
			result += value
		}
		prev = value
	}

	return result
}

func intToRoman(num int) string {
    arabic := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    
    result := ""
    for i := 0; i < len(arabic); i++ {
        for num >= arabic[i] {
            result += roman[i]
            num -= arabic[i]
        }
    }
    return result
}

func solve(input string) string{	
	sign := ""

	for i := 0; i < len(input); i++ {
		switch input[i : i+1] {
		case "*", "/", "+", "-":
			isInvalid := sign == "" && i != 0
			if isInvalid {
				sign = input[i : i+1]
			} else {
				panic("Выдача паники, так как формат математической операции не удовлетворяет заданию.")
			}
		}
	}

	if sign != "" {
		splitted := strings.Split(input, sign)
		isNull := splitted[0] == "0" || splitted[1] == "0"

		if isNull {
			panic("Выдача паники, так как строка содержит ноль.")
		}

		left, _ := strconv.Atoi(splitted[0])
		right, _ := strconv.Atoi(splitted[1])

		isArabic := left != 0 && right != 0
		isRoman := left == 0 && right == 0

		if isArabic {
			isInInterval := left > 0 && left < 11 && right > 0 && right < 11
			if isInInterval {
				return strconv.Itoa(calculate(left, right, sign)) 
			} else {
				panic("Выдача паники, так как строка содержит число вне интервала [1...10].")
			}
		} else if isRoman {
			left = romanToInt(splitted[0])
			right = romanToInt(splitted[1])
			isInInterval := left > 0 && left < 11 && right > 0 && right < 11
			if isInInterval {
				res := calculate(left, right, sign)
				if res > 0 {
					return intToRoman(res)
				} else {
					panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
				}
			} else {
				panic("Выдача паники, так как строка содержит число вне интервала [1...10].")
			}

			
		} else {
			panic("Выдача паники, так как используются одновременно разные системы счисления.")
		}

		
	} else {
		panic("Выдача паники, так как строка не является математической операцией.")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Input:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input = strings.Replace(input, " ", "", -1)

		fmt.Println("Output:")
		fmt.Println(solve(input))
	}
}