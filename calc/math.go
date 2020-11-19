package calc
// package declaration(calc) has nothing to do with the module name(github.com/baechul/golang)

// func Add(i int, j int) int {
// 	return i + j
// }

// Add no need to explain.
func Add(numbers ...int) int {
	sum := 0
	for i:=0; i<len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}