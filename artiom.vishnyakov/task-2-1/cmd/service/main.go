package main

import "fmt"

func main() {
	var amount_office int
	_, err := fmt.Scan(&amount_office)
	if err != nil {
		return
	}

	for i := 0; i < amount_office; i++ {
		var amount_emplyee int
		_, err = fmt.Scan(&amount_emplyee)
		if err != nil {
			return
		}
		for j := 0; j < amount_emplyee; j++ {
			var op string
			_, err = fmt.Scan(&op)
			if err != nil {
				return
			}
			if op != "<=" && op != ">=" {
				return
			}
			var num int
			_, err = fmt.Scan(&num)
			if err != nil {
				return
			}
		}
	}

}
