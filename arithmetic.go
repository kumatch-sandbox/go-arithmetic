package arithmetic

func Plus(base int, numbers ...int) int {
	for _, num := range numbers {
		base += num
	}
	return base
}

func Minus(base int, numbers ...int) int {
	for _, num := range numbers {
		base -= num
	}
	return base
}

func Times(base int, numbers ...int) int {
	for _, num := range numbers {
		base *= num
	}
	return base
}

func Divided(base int, numbers ...int) int {
	for _, num := range numbers {
		base = base / num
	}
	return base
}
