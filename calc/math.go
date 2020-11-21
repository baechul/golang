package calc
// package declaration(calc) has nothing to do with the module name(github.com/baechul/golang)

import "errors"

// Add no need to explain.
func Add(numbers ...int) (int, error) {
	sum := 0

	if(len(numbers) < 2) {
		return sum, errors.New("provide more than 2 numbers")
	} else {
		// ignore index hence used _.
		for _, num := range numbers {
			sum += num
		}
		return sum, nil
	}
}


// func Add(i int, j int) int {
// 	return i + j
// }

// Add no need to explain.
// func Add(numbers ...int) int {
// 	sum := 0
// 	for i:=0; i<len(numbers); i++ {
// 		sum += numbers[i]
// 	}
// 	return sum
// }