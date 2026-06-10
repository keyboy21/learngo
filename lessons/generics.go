package learngo

import "fmt"

func SumGeneric[V ~int64 | float64](numbers []V) (sum V) {
	for _, num := range numbers {
		sum = sum + num
	}
	return
}

func Contains[T comparable](elements []T, element T) bool {
	for _, val := range elements {
		if val == element {
			return true
		}
	}
	return false
}

func ShowAny[T any](val ...T) {
	fmt.Println(val)
}

type MyNumber interface {
	~int32 | int64 | float64
}

type Numbers[T MyNumber] []T

func unionInterfaceAndType() {
	var ints Numbers[int32]
	ints = append(ints, []int32{124, 235, 23, 523}...)

	floats := Numbers[float64]{1.24, 345, 456.456}
	fmt.Println(floats)

}

type CustomInt int64

func typeApproximation() {
	customInts := []CustomInt{2, 3, 5}

	intsSum := SumGeneric(customInts)
	fmt.Println(intsSum)
}
