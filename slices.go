package learngo

import (
	"slices"
	"sort"
)

func SlicesS() {
	var asd []string
	fmt.Printf("Type: %T, Value: %#v\n", asd, asd)
	as := []string{"yellow", "anna", "zick", "gordon", "light"}
	slices.Sort(as) // mutating current as
	fmt.Println(sort.SearchStrings(as, "v"))
	fmt.Println("length:", len(as))
	fmt.Println("array:", as)
	s1 := make([]int, 4, 5)
	fmt.Printf("Type slice: %T, Value: %v\n", s1, s1)
	sliceNew := new([]int) // not good way
	fmt.Printf("Type: %T, Value: %#v\n", sliceNew, sliceNew)
	fmt.Println(*sliceNew == nil)

	q := [...]string{"yellow", "anna", "zick", "gordon", "light"}
	sortedNumbers := q[:]      // NOT COPY, JUST REFERENCE
	slices.Sort(sortedNumbers) // SORTING (MUTATING ORIGIANL ARRAY(q) TOO !!!!!!)
	fmt.Println("sorted:", sortedNumbers)
	fmt.Println("original array:", q) // SAME RESULT WITH sortedNumbers

	mySlice := []string{"asd", "qwe", "rty"}
	mySliceCopy := mySlice
	newCopy := mySliceCopy
	newCopy[0] = "aaaaaaa"
	fmt.Printf("Original slice Value: %v\n", mySlice)
	fmt.Printf("Mutated slice 1: %v\n", mySliceCopy)
	fmt.Printf("Mutated slice 2: %v\n", newCopy)

	ar := [...]int{4, 5, 6, 3, 7, 8, 0, 2, 7, 23, 457547, 56756} // ORIGINAL ARRAY
	g := ar                                                      // COPY FROM ORIGINAL ARRAY, IF ADD & (g := &ar) OR [:] (g := ar[:])  NOT COPY
	slices.Sort(g[:])                                            // MUTATING ONLY g, because we add [:]. [:] is a slice
	fmt.Println("Original array", ar)                            // NOT MUTATED
	fmt.Println("Mutated array", g)                              // MUTATED

	var twoD [2][3]int // empty array with [[0 0 0] [0 0 0]]
	var twoD2 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(twoD)
	fmt.Println(twoD2)

	mainSlice := []int{1, 2, 3, 4, 5}
	list := []int{1, 2, 3, 4, 5}
	copyMainSlice := mainSlice[:]
	copyMainSlice[0] = 10

	fmt.Println(append(copyMainSlice, 6))

	handle(list)
	double(list)
	fmt.Println(list)
	fmt.Println(mainSlice)
	fmt.Println(copyMainSlice)

	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)
	slice1 := arr1
	slice1[0] = 10
	arr1[1] = 20
	fmt.Println(slice1)
	fmt.Println(arr1)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	b = append(b, "XXX")
	b = append(b, "YYY")
	fmt.Println(a, b)
	fmt.Println(names)

	fmt.Println(example() == example2())

	fmt.Printf("example = %#v, example2 = %#v\n", example(), example2())
	fmt.Println(example())
	fmt.Println(example2())

	// make creates a slice or map with length and capacity ================
	m := make(map[string]int)
	m = map[string]int{
		"Tashkent": 01,
		"Namangan": 02,
		"Andijon":  03,
	}
	fmt.Println(m)
	m["Farg'ona"] = 04
	fmt.Println("Andijon", m["Andijon"])
	fmt.Println("len", len(m))
	delete(m, "Andijon")
	fmt.Println(m["Andijon"]) // result is 0
	_, ok := m["Andij"]
	if ok {
		fmt.Println("Andij is in m")
	}
	fmt.Println("Andij is not in m")
	fmt.Println(m)

	m2 := make([]int, 3, 50)
	fmt.Println(m2)
	fmt.Printf("Length %v\n", len(m2))
	fmt.Printf("Capacity %v\n", cap(m2))

	fmt.Println(allOddSum(1, 100))
}