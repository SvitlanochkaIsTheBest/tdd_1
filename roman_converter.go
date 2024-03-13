package main

import (
	"errors"
	"sort"
)

type numeral struct {
	value  int
	symbol string
}

var romanNumerals = []numeral{
	{1, "I"},
	{4, "IV"},
	{5, "V"},
	{9, "IX"},
	{10, "X"},
	{40, "XL"},
	{50, "L"},
	{90, "XC"},
	{100, "C"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
}

func init() {
	sort.Slice(romanNumerals, func(i, j int) bool {
		return romanNumerals[i].value > romanNumerals[j].value
	})
}

func toRoman(numeral int) (string, error) {
	if numeral < 1 || numeral > 3999 {
		return "", errors.New("number out of range (1-3999)")
	}

	result := ""
	for _, n := range romanNumerals {
		for numeral >= n.value {
			result += n.symbol
			numeral -= n.value
		}
	}
	return result, nil
}
