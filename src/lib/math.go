package lib

var Public = 100
var private = 1000

func Average(s []int) int {
	total := 0

	for _, i := range s {
		total += i
	}
	return int(total / len(s))
}
