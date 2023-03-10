package app

import (
	"log"
	"strconv"
)

func stringToInt(s string) int {
	stringConverted, e := strconv.Atoi(s)

	if e != nil {
		log.Fatalf("Error converting %s to int", s)
	}

	return stringConverted
}

func intToString(i int) string {
	intConverted := strconv.Itoa(i)

	return intConverted
}

func elementExistInSlice(s []int, str int) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
