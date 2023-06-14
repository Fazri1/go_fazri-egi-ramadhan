package calculator

func Add(num ...float64) float64 {
	result := num[0]
	for i := 1; i < len(num); i++ {
		result += num[i]
	}

	return result
}

func Sub(num ...float64) float64 {
	result := num[0]
	for i := 1; i < len(num); i++ {
		result -= num[i]
	}

	return result
}

func Mult(num ...float64) float64 {
	result := num[0]
	for i := 1; i < len(num); i++ {
		result *= num[i]
	}

	return result
}

func Div(num ...float64) float64 {
	result := num[0]
	for i := 1; i < len(num); i++ {
		if num[i] == 0 {
			return 0
		}
		result /= num[i]
	}

	return result
}
