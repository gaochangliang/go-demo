package math

func adder(sc ...int) int {
	var total int
	for _, v := range sc {
		total += v
	}
	return total
}
