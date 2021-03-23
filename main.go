package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string

	fmt.Print("Input value: ")
	fmt.Scan(&input)

	grade, err := gradeDeterminer(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(grade)
}

func gradeDeterminer(input string) (string, error) {
	i, err := strconv.Atoi(input)
	if err != nil {
		return "", fmt.Errorf("Input isn't number")
	}

	if i >= 90 {
		return "A", nil
	} else if i >= 80 && i <= 89 {
		return "B", nil
	} else if i >= 70 && i <= 79 {
		return "c", nil
	} else if i >= 60 && i <= 69 {
		return "D", nil
	}
	return "E", nil
}
