package main

import (
	"fmt"
	"strconv"
	"strings"
)

func arabicToRoman(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symb := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	romanNumeral := ""
	i := 0
	for num > 0 {
		for num >= val[i] {
			romanNumeral += symb[i]
			num -= val[i]
		}
		i++
	}
	return romanNumeral
}

func romanToArabic(roman string) (int, error) {
	romanDict := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	arabic := 0
	prevValue := 0
	for _, numeral := range roman {
		value, exists := romanDict[numeral]
		if !exists {
			return 0, fmt.Errorf("Неверная римская цифра: %c", numeral)
		}
		if value > prevValue {
			arabic += value - 2*prevValue // Вычитаем дважды предыдущее значение
		} else {
			arabic += value
		}
		prevValue = value
	}
	return arabic, nil
}

func calculate(expression string) (int, error) {
	operators := []string{"+", "-", "*", "/"}
	var operator string
	for _, op := range operators {
		if strings.Contains(expression, op) {
			operator = op
			break
		}
	}
	if operator == "" {
		return 0, fmt.Errorf("Неверная арифметическая операция")
	}

	nums := strings.Split(expression, operator)
	if len(nums) != 2 {
		return 0, fmt.Errorf("Неверный формат выражения")
	}

	num1Str := strings.TrimSpace(nums[0])
	num2Str := strings.TrimSpace(nums[1])

	var num1, num2 int
	var err error

	if strings.ContainsAny(num1Str, "IVXLCDM") && strings.ContainsAny(num2Str, "IVXLCDM") {
		return 0, fmt.Errorf("Калькулятор умеет работать только с арабскими или римскими цифрами одновременно")
	}

	if strings.ContainsAny(num1Str, "IVXLCDM") {
		num1, err = romanToArabic(num1Str)
		if err != nil {
			return 0, err
		}
		num2, err = strconv.Atoi(num2Str)
		if err != nil {
			return 0, err
		}
	} else {
		num1, err = strconv.Atoi(num1Str)
		if err != nil {
			return 0, err
		}
		num2, err = strconv.Atoi(num2Str)
		if err != nil {
			return 0, err
		}
	}

	var result int
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("Деление на ноль")
		}
		result = num1 / num2
	}

	if result < 0 {
		return 0, fmt.Errorf("Результатом работы калькулятора с арабскими числами могут быть отрицательные числа и ноль")
	}
	if result > 10 {
		return 0, fmt.Errorf("Калькулятор должен принимать на вход числа от 1 до 10 включительно")
	}

	return result, nil
}

func main() {
	var expression string
	fmt.Print("Введите выражение: ")
	fmt.Scanln(&expression)

	result, err := calculate(expression)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}
}
