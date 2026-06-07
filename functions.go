package learngo

import (
	"fmt"
	"time"
)

func functions() {
	var personPrint = func(userName string) {
		fmt.Println(userName)
	}
	var personGreet2 = func(firstName, lastName string) {
		var userName = fmt.Sprintf("%v %v", firstName, lastName)
		fmt.Println(userName)
	}
	personPrint(userName)
	personGreet2(userName, "Doe")
	// closure functions in go
	var dollar = 30

	getDollar := func() int {
		return dollar
	}
	fmt.Println(getDollar())
	dollar = 40
	fmt.Println(getDollar())

	var DivideBy2 = createDivider(2)
	fmt.Println(DivideBy2(10))
	fmt.Println(DivideBy2(20))

	var Sum = calculate(1, 2, add)
	fmt.Println("Sum = ", Sum)

	returnUsername = func(name string) string { return name }

	fmt.Println(returnUsername("John"))

	returnAge = func(birthYear uint) (age int, currentYear int) {
		currentYear = time.Now().Year()
		age = currentYear - int(birthYear)
		return
		// return age, cyear
	}

	age, currentYear := returnAge(2000)
	fmt.Printf("Age is %v and current year is %v\n", age, currentYear)
}