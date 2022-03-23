package main

import (
	"fmt"
	"strings"
)

type RomanNumeral struct {
	Arabic int
	Roman  string
}

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbols ...byte) int {
	Roman := string(symbols)
	for _, s := range r {
		if s.Roman == Roman {
			return s.Arabic
		}
	}
	return 0
}

var allRomanNumeral = RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRomanNumeral(arabic int) string {
	var result strings.Builder

	for _, romanNumeral := range allRomanNumeral {
		for arabic >= romanNumeral.Arabic {
			result.WriteString(romanNumeral.Roman)
			arabic -= romanNumeral.Arabic
		}
	}

	return result.String()
}

func couldBeSubtractive(index int, currentSymbol byte, roman string) bool {
	isSubtractiveSymbol := string(currentSymbol) == "I" || string(currentSymbol) == "X" || string(currentSymbol) == "C"
	return index+1 < len(roman) && isSubtractiveSymbol
}

func ConvertToArabic(roman string) (total int) {
	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		if couldBeSubtractive(i, symbol, roman) {
			nextSymbol := roman[i+1]

			if value := allRomanNumeral.ValueOf(symbol, nextSymbol); value != 0 {
				total += value
				i++ // move past this character too for the next loop
			} else {
				total += allRomanNumeral.ValueOf(symbol)
			}
		} else {
			total += allRomanNumeral.ValueOf(symbol)
			fmt.Println("Proud of this work")
		}
	}

	return
}
