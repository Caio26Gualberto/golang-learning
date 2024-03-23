package main

import "fmt"

func main() {
	day := dayOfWeek(2)
	day2 := dayOfWeek(8)
	fmt.Println(day)
	fmt.Println(day2)

	day3 := dayOfWeek(2)
	day4 := dayOfWeek(8)
	fmt.Println(day3)
	fmt.Println(day4)
}

func dayOfWeek(number int) string {
	switch number {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

func dayOfWeek2(number int) string {

	var dayOfWeek string

	switch {
	case number == 1:
		dayOfWeek = "Domingo"
		fallthrough //Fallthrough causes validation to drop to the next clause
	case number == 2:
		dayOfWeek = "Segunda"
	case number == 3:
		dayOfWeek = "Terça"
	case number == 4:
		dayOfWeek = "Quarta"
	case number == 5:
		dayOfWeek = "Quinta"
	case number == 6:
		dayOfWeek = "Sexta"
	case number == 7:
		dayOfWeek = "Sábado"
	default:
		dayOfWeek = "Número inválido"
	}

	return dayOfWeek
}
