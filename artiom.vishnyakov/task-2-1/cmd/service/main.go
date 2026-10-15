package main

import "fmt"

func calcTemp(operation string, temp int, maxTemp *int, minTemp *int) int {
	if operation == "<=" {
		if temp < *maxTemp {
			*maxTemp = temp
		}
	}

	if operation == ">=" {
		if temp > *minTemp {
			*minTemp = temp
		}
	}

	if *minTemp > *maxTemp {
		return -1
	}

	return *minTemp
}

func main() {
	var amountOffice int

	_, err := fmt.Scan(&amountOffice)
	if err != nil {
		return
	}

	for range amountOffice {
		var amountEmployee int

		_, err = fmt.Scan(&amountEmployee)
		if err != nil {
			return
		}

		minTemp := 15
		maxTemp := 30

		for range amountEmployee {
			var operation string

			_, err = fmt.Scan(&operation)
			if err != nil {
				return
			}

			if operation != "<=" && operation != ">=" {
				return
			}

			var num int

			_, err = fmt.Scan(&num)
			if err != nil {
				return
			}

			fmt.Println(calcTemp(operation, num, &maxTemp, &minTemp))
		}
	}
}
