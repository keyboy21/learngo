package learngo

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func loops() {
	for x := 1.0; x <= 5.0; x++ {
		result := sqrt(x)
		fmt.Printf("Square root of %v is %v\n", x, result)
		fmt.Printf("Difference from math.Sqrt: %v\n, %v", math.Sqrt(x), result)
		fmt.Println()
	}

	sum := 0
	for i := 0; i < 10; i++ {
		sum += 10
	}
	fmt.Println(sum)

	sum2 := 1
	for sum2 < 200 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// continue , break, labels ==============================================
	// continue will skip the current iteration and continue to the next one
	for i := 0; i < 20; i++ {
		if i%2 == 1 {
			continue
		}

		fmt.Println("i = ", i)
	}

	// 'break' will break the loop and end the loop
	for i := 1; i < 20; i++ {
		if i > 10 {
			break
		}
		fmt.Println("i = ", i)
	}

	// 'labels' can be used to break or continue outer loops
Outer:
	for i := 1; i <= 20; i++ {
	Inner:
		for j := 1; j <= 10; j++ {
			fmt.Printf("i = %v, j = %v\n", i, j)
			if i == 10 {
				// you can use break and continue(skip) with labels
				continue Outer
			}
			if j == 15 {
				continue Inner
			}
		}
	}
}
