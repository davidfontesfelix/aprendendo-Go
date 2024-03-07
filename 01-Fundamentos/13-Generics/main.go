package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func sum[T Number](m map[string]T) T {
	var sum T
	for _, value := range m {
		sum += value
	}
	return sum
}

func compare[T Number](a T, b T) bool {
	if a == b {
		return true
	} else {
		return false
	}
}

// É ridoculo usar criar outra função igual por causa de tipos, por isso existe generics

// func sumFloat(m map[string]float64) float64 {
// 	var sum float64
// 	for _, value := range m {
// 		sum += value
// 	}
// 	return sum
// }

func main() {
	m := map[string]int{"Wesley": 1000, "Maria": 1600, "Lucas": 2000}
	m2 := map[string]float64{"Wesley": 100.20, "Maria": 1600.70, "Lucas": 2000.00}

	m3 := map[string]MyNumber{"Wesley": 1000, "Maria": 1600, "Lucas": 2000}
	println(sum(m))
	println(sum(m2))
	println(sum(m3))

	println(compare(10, 10.0))
}
