package main

import (
	"strconv"
	"strings"
)

func DefaultString(input string, defaultValue string) string {
	if input == "" {
		return defaultValue
	}
	return input
}

func DefaultInt(input string, defaultValue int) int {
	if input == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(input)

	if err != nil {
		return defaultValue
	}

	return i
}

func DefaultSlice(input string, defaultValue []string) []string {
	if input == "" {
		return defaultValue
	}

	split := strings.Split(input, ",")

	for i, item := range split {
		split[i] = strings.ToLower(strings.TrimSpace(item))
	}

	return split
}

func stringInSlice(str string, list []string) bool {
	for _, item := range list {
		if item == str {
			return true
		}
	}

	return false
}

func roundFloat64(input float64, places int) float64 {
	rounding := 1.0

	for i := 0; i < places; i++ {
		rounding *= 10.0
	}

	return float64(int(input*rounding)) / rounding
}
