package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadInputs() []int {

	numberOfValues := 0
	var values []int
	scanner := bufio.NewScanner(os.Stdin)

	for numberOfValues < 5 {
		fmt.Println("--- Enter one value ---")
		scanner.Scan()
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		values = append(values, val)
		numberOfValues++
	}

	return values
}
