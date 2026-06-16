package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync/atomic"

	// "math/rand"
	"runtime"
	"strings"
	"sync"
	"time"

	learngo "github.com/keyboy21/learngo/lessons"
)

var (
	UserName = "John"
	userAge  = 23
	userJob  = "Developer"
)

const (
	size   float64 = 123.123
	width  float32 = 1.23123
	lenght float32 = 123123.123123123
	space          = 123123.123123123
)

var (
	i8  int8  = -128 | 127
	i16 int16 = -32768 | 32768
	i32 int32 = -2147483648 | 2147483648
	i64 int64 = -9223372036854775808 | -9223372036854775808
)

var (
	u8  uint8  = 0 | 255
	u16 uint16 = 0 | 65535
	u32 uint32 = 0 | 4294967295
	u64 uint64 = 0 | 18446744073709551615
)

const (
	Monday = 1 + iota
	TuesDay
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

// multiple results
func swap(x, y string) (string, string, string) {
	c := x + y
	return x, y, c
}

func swap2(x, y string) string {
	return x + y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return x, y
}

var c, python, java bool

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
		fmt.Printf("Iteration %d: z = %v\n", i+1, z)
	}

	return z
}

func oSystem() string {
	os := runtime.GOOS
	return fmt.Sprintf("Operation system is %v", os)
}

func helloTime() {
	t := time.Now().Hour()

	switch {
	case t < 12:
		fmt.Println("Good morning!")
	case t < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

type Developer struct {
	devType string
}

type Product struct {
	id    uint
	name  string
	price float32
}

type User struct {
	id          uint
	name, email string
	phones      []string
	cart        []Product
	Developer
}

func (uAuth *User) getUserToken() (token string, ok bool) {

	userEmail, userName := uAuth.email, uAuth.name

	if userEmail == "" || userName == "" {
		return "", false
	}

	token = fmt.Sprintf("Bearer: %v:%v", uAuth.email, uAuth.name)
	ok = true

	return token, ok
}

func bubbleSort(arr []int) {
	l := len(arr)
	fmt.Println("Array length", l)

	for i := 0; i < l-1; i++ {
		for j := 0; j < l-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func binarySearch(arr []int, target int) int {
	f := 0
	l := len(arr) - 1

	for i := 0; i < len(arr); i++ {
		mid := (f + l) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			f = arr[mid] + 1
		} else if arr[mid] > target {
			l = arr[mid] - 1
		}
	}

	return -1
}

func square(x int) int {
	return x * x
}

func squarePoint(x *int) {
	*x = *x * *x
}

var returnUsername func(string) string
var returnAge func(birthYear uint) (age int, currentYear int)

const (
	min = 1
	max = 8
)

type Exapmle struct {
	Value string
}

func example() any {
	var e *Exapmle

	return e
}

func example2() any {
	return nil
}

func double(nums []int) {
	// res := make([]int,0, len(nums))

	for _, num := range nums {
		num *= 2
	}

	// return res
}

func handle(list []int) {
	list[1] = 10
}

type Square struct {
	Side int
}

// In go you can call methods with pointer receiver or value receiver, go automatically
// handle them based on the type of the receiver

// method with value receiver
func (s Square) Perimeter() {
	fmt.Printf("%T, %#v", s, s)
	fmt.Printf("Perimeter: %d\n", s.Side*4)
}

// method with pointer receiver
func (s *Square) Scale(mutiplier int) {
	fmt.Printf("%T, %#v\n", s, s)
	s.Side = s.Side * mutiplier
}

type Runner interface {
	Run() (canRun string, ok bool)
}
type Swimmer interface {
	Swim() string
}
type Flyer interface {
	Fly() string
}
type Ducker interface {
	Runner
	Swimmer
	Flyer
}

type Duck struct {
	name string
}

func (u User) Run() (canRun string, ok bool) {
	if u.name == "" {
		return "", false
	}

	canRun = fmt.Sprintf("User: %v is running", u.name)
	ok = true

	return
}

func (u User) WriteCode() string {

	return fmt.Sprintf("User: %v is writing code", u.name)
}

func (d Duck) Run() (canRun string, ok bool) {
	if d.name == "" {
		return "", false
	}

	canRun = fmt.Sprintf("Duck: %v is running", d.name)
	ok = true

	return canRun, ok
}

func (d Duck) Fly() string {

	return fmt.Sprintf("Duck: %v is flying", d.name)
}

func (u User) Fly() string {
	return fmt.Sprintf("User: %v can not fly", u.name)
}

// Important: If we accept interface his type and value can be nil.
// If it type or value nil it will be panic.
// For this reason we should use type assertion to check.
func typeAssertion(runner Runner) {
	fmt.Printf("%T, %#v\n", runner, runner)

	if _, ok := runner.(*User); ok {
		fmt.Println(runner.Run())
	} else {
		fmt.Println("User is nil")
	}

	switch v := runner.(type) {
	case *User:
		fmt.Println(v.WriteCode())
	case *Duck:
		canRun, ok := v.Run()
		if ok {
			fmt.Println(canRun)
		} else {
			fmt.Println("can't run")
		}
	default:
		fmt.Println("unknown type")
	}
}

func polymorphism(runner Runner) {
	// will be panic if type or value is nil
	fmt.Println(runner.Run())
}

type Person struct {
	name string
	age  int
}

type WorkExperience struct {
	year int
}

type WoodBuilder struct {
	Person
	WorkExperience
}

type BuildingDestroyer struct {
	Person
	WorkExperience
	Destroyer
}

type BrickBuilder struct {
	Person
}

type Builder interface {
	Build()
}

type Destroyer interface {
	Destroy()
}

type Building struct {
	Builder
	Destroyer
	name string
}

func (p Person) printName() {
	fmt.Println(p.name)
}

func (w WoodBuilder) printName() {
	fmt.Println(w.name)
}

// implements Builder interface for WoodBuilder
func (w WoodBuilder) Build() {
	fmt.Println("Building House from Wood")
}

// implements Builder interface for BrickBuilder
func (b BrickBuilder) Build() {
	fmt.Println("Building House from Brick")
}

// implement Destroyer interface for BuildDestroyer
func (b Building) Destroy() {
	if b.Destroyer == nil {
		fmt.Println("No destroyer provided")
		return
	}

	fmt.Println("Destroying House")
}

func showAllElements(arr ...int) {
	fmt.Printf("Length: %d\n Capacity: %d\n", len(arr), cap(arr))

	for _, value := range arr {
		fmt.Printf("Value: %d\n", value)
	}
}

func showNumbers(num int) {
	for i := 0; i <= num; i++ {
		fmt.Println(i)
	}
}

func sum(x, y int) (sum int) {

	defer func() { // deferred function calls before returning
		fmt.Println("deferred function")
		sum *= 2
	}()

	sum = x + y
	return sum
}

func defferedValues() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("deferred value", i)
	}

	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println("deferred func", i)
		}()
	}
}

func makePanic() {

	defer func() {
		panicValue := recover()

		fmt.Println(panicValue)
	}()

	panic("make panic")
	fmt.Println("after panic")
}

type Fruit struct {
	name string
}

func collectValues(fruits []Fruit) map[string]int {
	result := make(map[string]int, len(fruits))

	for _, fruit := range fruits {
		if result[fruit.name] == 0 {
			result[fruit.name] = 1
		} else {
			result[fruit.name] += 1
		}
	}

	return result

}

func WordCount(s string) map[string]int {

	words := strings.Fields(s)
	counts := make(map[string]int, len(words))

	for _, word := range words {
		counts[word]++
	}
	return counts
}

func variables() {
	var userName string // default ""
	var ourBool bool    // default false
	myNumber := 1
	fmt.Println(userName)
	fmt.Println(ourBool)
	fmt.Println(myNumber)
	userName = "John"
	ourBool = true
	fmt.Println(userName)
	fmt.Printf("Type: %T, Value: %v\n", userName, userName)
	fmt.Printf("Type: %T, Value: %v\n", ourBool, ourBool)

	var num1 int = 1
	var num2 uint = 2
	var sum = num1 + int(num2)
	fmt.Println(sum)
}

func conditionalExpression() {
	if age := 20; age > 18 {
		fmt.Println("User is adult")
	} else {
		fmt.Println("User is too young")
	}

	age := 17
	if age > 18 {
		fmt.Println("User is adult")
	} else {
		fmt.Println("User is too young")
	}
}

func switchStatement() {
	// rand.Seed((time.Now()).UnixNano()) // old way
	// r := rand.New(rand.NewSource(42)) // custom seed

	randValue := rand.Intn(10)

	switch {
	case randValue >= 1 && randValue <= 2:
		fmt.Printf("randValue = %v, adound 1 and 2", randValue)
	case randValue >= 3 && randValue <= 4:
		fmt.Println("randValue3,4 = ", randValue)
	case randValue == 5:
		fmt.Printf("randValue5 = %v, and fallthrough\n", randValue)
		fallthrough
	case randValue == 6:
		fmt.Println("randValue6 = ", randValue)
	case randValue >= 7 && randValue <= 8:
		fmt.Println("randValue7,8 = ", randValue)
	default:
		fmt.Println("default = ", randValue)
	}

	switch num := rand.Intn(10); num {
	case 1:
		fmt.Printf("Number is: %v", num)
	case 2:
		fmt.Printf("Number is: %v", num)
	case 3:
		fmt.Printf("Number is: %v", num)
	default:
		fmt.Printf("Default number is: %v", num)
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Printf("Operation system is %v", os)
	case "freebsd":
		fmt.Printf("Operation system is %v", os)
	case "linux":
		fmt.Printf("Operation system is %v", os)
	default:
		fmt.Printf("Operation system is %v\n", os)
	}

	whatAmI := func(i any) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	today := time.Now().Weekday()

	switch time.Sunday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func pointers() {
	var intPointer *int   // nil pointer
	var uIntPointer *uint // nil pointer
	// default value of pointer is nil
	fmt.Printf(" %T, %v\n", intPointer, intPointer)                      // *int, <nil>
	fmt.Printf(" %T, %v\n", uIntPointer, uIntPointer)                    // *uint, <nil>
	fmt.Printf("check pointer has value or not %v\n", intPointer == nil) // true
	// fmt.Printf("%T, %v\n", intPointer, *intPointer)                      // don't do this, because intPointer is nil and it will be panic
	intPointer = &userAge    // set address of userAge to intPointer
	fmt.Println(intPointer)  // get address of userAge 0xc0000b6010
	fmt.Println(*intPointer) // get value of userAge 23 use * before pointer

	var b = 255
	bPointer := &b                           // bPointer is a pointer to b
	fmt.Println("Value of b is", *bPointer)  // value 255
	fmt.Println("Address of b is", bPointer) // address 0xc0000b6010

	var newPointer = new(int) // new() function creates a pointer to the type and initializes it to default value
	fmt.Printf("newPointer type: %T, newPointer value: %v\n newPointer address: %v\n", newPointer, *newPointer, newPointer)
	*newPointer = 10 // to set value to pointer use * before pointer
	fmt.Println("newPointer value after set", *newPointer)

	var x int = 10
	square(x) // square doesn't have side effects because it doesn't change the value of x (pass by value)
	fmt.Println("x after square", x)
	squarePoint(&x) // squarePoint has side effects because it changes the value of x (pass by reference)
	fmt.Println("x after squarePoint", x)
}

func arrays() {
	var myint [3]int
	fmt.Printf("Type: %T, Value: %v\n", myint, myint)
	myint[2] = 4
	myint[0] = 1
	fmt.Printf("Type: %T, Value: %v\n", myint, myint)

	phones := [...]string{"John", "Tony", "Stark"}
	fmt.Printf("Capacity: %v, Length: %v\n", cap(phones), len(phones))

	for i := 0; i < len(phones); i++ {
		fmt.Println(phones[i])
	}
	for key, value := range phones {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}
	// people := [...]User{
	// 	{
	// 		id:    1,
	// 		name:  "John",
	// 		email: "asd@gmail.com",
	// 		// phones: phones,
	// 	},
	// }
}

func copySlice() {
	// copy ========
	destination := make([]string, 3, 3)
	fmt.Printf("Destination Value: %#v\n", destination)
	superHeroesSlice := []string{"Hulk", "Iron Man", "Ant man"}
	fmt.Println("Copied:", copy(destination, superHeroesSlice))
	fmt.Println("Cap", cap(destination), "Len", len(destination))
	fmt.Printf("destination: %#v", destination)
	for _, value := range destination {
		fmt.Println(value)
	}
}

func maps() {
	var defaultMap map[int]string // default nil
	fmt.Printf("Type: %T, Value: %#v\n", defaultMap, defaultMap)
	// fmt.Printf("Len: %d\n", len(defaultMap))

	// by make
	mapByMake := make(map[string]int, 4)
	fmt.Printf("Type: %T, Value: %#v\n", mapByMake, mapByMake)
	fmt.Printf("Len: %d\n", len(mapByMake))
	mapByMake["Antony"] = 30
	fmt.Printf("Value: %#v\n", mapByMake)
	mapByMake["John"] = 20
	fmt.Printf("Value: %#v\n", mapByMake)

	// by literal
	literalMap := map[string]int{"John": 12, "Tony": 20}
	fmt.Printf("Type: %T, Value: %#v\n", literalMap, literalMap)
	fmt.Printf("Len: %d\n", len(literalMap))

	// by new
	p := new(map[string]int)
	fmt.Printf("Type: %T, Value: %#v\n", p, p)
	mapByNew := *p
	// need linter
	// mapByNew["Toby"] = 30 // panic: assignment to entry in nil map
	fmt.Println(mapByNew)

	// get
	value, ok := mapByMake["Henry"]
	fmt.Printf("Value: %v, isExist: %v\n", value, ok)

	// delete
	delete(mapByMake, "John")
	fmt.Printf("Type: %T, Value: %#v\n", mapByMake, mapByMake)

	// map iteration
	mapByMake["Bryuce"] = 45
	mapByMake["Anna"] = 30
	for value, key := range mapByMake {
		// print randomly
		fmt.Printf("Key: %v, value: %v\n", value, key)
	}

	users := []struct {
		id   int32
		name string
	}{
		{
			id:   3,
			name: "John",
		},
		{
			id:   5,
			name: "Dode",
		},
		{
			id:   5,
			name: "Tony",
		},
		{
			id:   3,
			name: "Stark",
		},
	}

	type User struct {
		id   int32
		name string
	}

	uniqueMap := make(map[int64]struct {
		id   int32
		name string
	}, len(users))

	for _, user := range users {
		if _, ok := uniqueMap[int64(user.id)]; !ok {
			uniqueMap[int64(user.id)] = struct {
				id   int32
				name string
			}{
				id:   user.id,
				name: user.name,
			}
		}
	}

	fmt.Printf("Unique map: %#v\n, Len: %v", uniqueMap, len(uniqueMap))

	findUserBySlice := func(id int, users []struct {
		id   int32
		name string
	}) *struct {
		id   int
		name string
	} {
		for _, user := range users {
			if user.id == int32(id) {
				return &struct {
					id   int
					name string
				}{
					id:   int(user.id),
					name: user.name,
				}
			}
		}
		return nil
	}
	findUserBySlice(3, users)

	// =========================================================
	usersByMap := make(map[int32]User, len(users))

	for _, user := range users {
		usersByMap[user.id] = user
	}

	findUserByMap := func(userId int, users map[int32]User) *User {
		value, ok := users[int32(userId)]
		if ok {
			return &value
		}
		return nil
	}

	findUserByMap(5, usersByMap)

}

func goRoutines() {
	// MARK: go routines
	fmt.Println("Number of logical cores:", runtime.NumCPU()) //get number of logical cores
	// runtime.GOMAXPROCS(1) // set number of logical cores to use
	go showNumbers(100) // may not show all numbers
	// runtime.Gosched() // yield the processor, let other goroutines run
	fmt.Println("End of main") // and back to the main goroutine

	time.Sleep(time.Second * 10)

	fmt.Println("Exit")

	// fmt.Println(sum(1, 2)) // returns deferred function result
	// defferedValues()
	makePanic()

	// waitGroup(100)
	// withConcurrent()
	// readWithMutex()
}

func withOutWait() {
	// main go routine without waiting these go routines closes
	for i := 0; i < 10; i++ {
		go fmt.Println(i + 1)
	}

	// runtime.Gosched()
}

// print numbers in a goroutine, but not sequentially
func waitGroup(num int) {
	var wg sync.WaitGroup
	wg.Add(num) // set number of go routines to group

	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done() // finish goroutine when done, if not deadlock
			fmt.Println(i)
		}()
	}

	wg.Wait() // wait for all goroutines to finish
	fmt.Println("End of loop")
}

func withConcurrent() {
	start := time.Now()
	var wg sync.WaitGroup // for adding goroutines to schedule
	var mu sync.Mutex     // for synchronizing access to counter

	wg.Add(5000)
	var counter int
	for i := 0; i < 5000; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond)
			mu.Lock() // lock to run one goroutine at a time
			counter++
			mu.Unlock() // unlock after use of counter, let other goroutines run
		}()
	}
	wg.Wait()

	fmt.Println("Counter:", counter)
	fmt.Println("Time taken:", time.Now().Sub(start).Seconds())
}

func withRWMutex() {
	start := time.Now()
	var (
		counter int
		wg      sync.WaitGroup
		mu      sync.RWMutex // read-write mutex for synchronizing reads and writes
	)

	wg.Add(100)

	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			mu.RLock() // not block reading for other go routines
			time.Sleep(time.Millisecond)
			_ = counter
			mu.RUnlock()
		}()
	}

	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			mu.Lock() // lock other go routines to write, even read mutex waiting to access
			time.Sleep(time.Millisecond)
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println("Counter:", counter)
	fmt.Println("Time taken:", time.Since(start))

}

func makeRequest(num int) <-chan string {
	reqChan := make(chan string)

	go func() {
		time.Sleep(time.Second)
		reqChan <- fmt.Sprintf("response: %v", num)
	}()

	return reqChan
}

func WithAtomicAdd() {
	start := time.Now()
	var (
		wg      sync.WaitGroup
		counter int64
	)

	wg.Add(5000)

	for i := 0; i < 5000; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func CompareAndSwapAtomic() {
	var (
		counter int64
		wg      sync.WaitGroup
	)

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			if !atomic.CompareAndSwapInt64(&counter, 0, 1) {
				return
			}

			fmt.Println("Swapped goroutine number is", i)
		}(i)
	}

	wg.Wait()

	fmt.Println(counter)
}

func AtomicVal() {
	var val atomic.Value

	val.Store(1)
	fmt.Println(val.Load())
	fmt.Println(val.Swap(2))
	fmt.Println(val.Load())

	fmt.Println(val.CompareAndSwap(2, 3))
}

func main() {

  learngo.Input()
	// learngo.OsFile()
	
	// learngo.SimpleWriter()
	// learngo.SimpleReader()
	// learngo.RowsReader()

	// ints := []int64{1, 2, 35}
	// floats := []float64{2, 235, 346, 457, 457}
	// intsSum := learngo.SumGeneric(ints)
	// fmt.Println(intsSum)
	// fmt.Println(learngo.SumGeneric[float64](floats))
	// fmt.Println(contains(ints, 2))

	// users := []Person{
	// 	{
	// 		name: "John",
	// 		age:  25,
	// 	},
	// 	{
	// 		name: "Doe",
	// 		age:  23,
	// 	},
	// }

	// fmt.Println(contains(users, Person{
	// 	name: "Doe",
	// 	age:  23,
	// }))

	// AtomicVal()

	// firstRes := makeReques(1)
	// secondReq := makeReques(2)

	// fmt.Println("FirstResponse:", <-firstRes)
	// fmt.Println("FirstResponse:", <-secondReq)

	// chanAsMutex()
	// learngo.WithoutErrorGroup()
	// withErrorGroup()

	// baseContext()
	// workerPool()

	// ===================================
	// channels()
	// bufferedChannel()
	// chanWithRange()
	// baseSelect()
	// graceFullShoutDown()

	// ===================================
	// withOutWait()
	// waitGroup(10)
	// withConcurrent()
	// withAtomicAdd()
	// readWithMutex()

	// Custom types ===========================================
	// type CustomString string
	// type CustomInt int

	// ageInt := CustomInt(4)
	// var myint = int(int(4))
	// fmt.Println(myint)

	// var newDeveloper Developer
	// John := Developer{}
	// fmt.Printf("Developer %#v\n", John)
	// fmt.Printf("new Dev: %#v", newDeveloper)
	// Structure, Interface ================================================================
	// var runner Runner // nil, nil
	// fmt.Printf("Type: %T, Value: %v", runner, runner)
	// var user *User
	// var developer *Developer
	// runner = user
	// runner.Run()
	// runner = developer
	// polymorphism(runner)

	// building := Building{
	// 	Builder: WoodBuilder{
	// 		Person: Person{
	// 			name: "John",
	// 			age:  20,
	// 		},
	// 		WorkExperience: WorkExperience{
	// 			year: 5,
	// 		},
	// 	},
	// 	name: "Wooden Building",
	// 	Destroyer: BuildingDestroyer{
	// 		Person: Person{
	// 			name: "Rick",
	// 			age: 34,
	// 		},

	// 		WorkExperience: WorkExperience{
	// 			year: 10,
	// 		},

	// 	},
	// }

	// building.Build()
	// building.Destroy()

	// WoodBuilder and BrickBuilder are implementing Builder interface
	// we can pass WoodBuilder or BrickBuilder to the Build() method
	// woodenBuilding := Building{
	// 	Builder: WoodBuilder{
	// 		Person:         Person{name: "John", age: 20},
	// 		name:           "Wood",
	// 		WorkExperience: WorkExperience{name: "John", age: 20},
	// 	}}
	// woodenBuilding.Build()

	// brickBuilding := Building{
	// 	Builder: BrickBuilder{Person{name: "John", age: 20}},
	// 	name:    "Brick",
	// }
	// brickBuilding.Build()

	// Embedding =============================================================
	// builder := WoodBuilder{
	// 	Person{name: "John", age: 20},
	// 	"Wood",
	// 	WorkExperience{name: "John", age: 20}}
	// fmt.Printf("Type:%T, Value:%#v\n", builder, builder)

	// shadowing ===
	// fmt.Println(builder.Person.name)
	// fmt.Println(builder.name)
	// builder.Person.printName()
	// builder.printName()

	// Interface ===========================================================
	// var runner Runner // by default interface is nil
	// fmt.Printf("Type:%T, Value:%#v\n", runner, runner)
	// var john = User{id: 1, name: "John"}
	// runner = john
	// typeAssertion(runner)
	// runner.Run()
	// fmt.Printf("Type:%T, Value:%#v\n", runner, runner)
	// fmt.Printf("Type:%T, Value:%#v\n", john, john)

	// canRun, ok := john.Run()

	// if ok {
	// 	fmt.Println(canRun)
	// } else {
	// 	fmt.Println("can't run")
	// }

	// blackDuck := Duck{name: "BlackDuck"}
	// runner = blackDuck
	// typeAssertion(runner)
	// runner.Run()
	// canRun, ok = blackDuck.Run()
	// if ok {
	// 	fmt.Println(canRun)
	// } else {
	// 	fmt.Println("can't run")
	// }

	// var emptyInterface any = john
	// fmt.Printf("Type:%T, Value:%#v\n", emptyInterface, emptyInterface)

	// Yusufboy := User{
	// 	id:    1,
	// 	name:  "Yusufboy",
	// 	email: "asas@gmail.com",
	// 	phones: []uint{
	// 		99889359184,
	// 		141414124124123,
	// 	},
	// 	cart: []Product{
	// 		{
	// 			id:    1,
	// 			name:  "Samsung",
	// 			price: 222123,
	// 		},
	// 		{
	// 			id:    2,
	// 			name:  "Xiaomi",
	// 			price: 123123,
	// 		},
	// 	},
	// 	Developer: Developer{
	// 		devType: "Frontend",
	// 	},
	// }

	// userToken, ok := Yusufboy.getUserToken()
	// if ok {
	// 	fmt.Println(userToken)
	// }

	// fmt.Println(Yusufboy)
	// m := Yusufboy.email
	// if m == "" {
	// 	fmt.Println("Don't have email")
	// }
	// fmt.Println(Yusufboy.cart)

	// Anonim struct
	// car := struct {
	// 	Name, Model, Time, Certificate string
	// }{
	// 	Name:        "BMW",
	// 	Model:       "X5",
	// 	Time:        fmt.Sprint(time.Now().ISOWeek()),
	// 	Certificate: time.Now().String(),
	// }

	// fmt.Println(car)
	// fmt.Printf("Time: %v \n", car.Time)
	// fmt.Printf("Certificate %v \n", car.Certificate)
	// car.Name = "Ferrari"
	// fmt.Printf("New name: %v", car.Name)

	//=====================================================
	// var ar = []Fruit{{"apple"}, {"banana"}, {"orange"}, {"apple"}, {"banana"}}
	// fmt.Println(collectValues(ar))

	// str := "hello world. this is a test"
	// fmt.Println(WordCount(str))

	// =====================================
	// var a = []int{1, 2, 3, 4, 5}
	// handle(a) // by default slice is passed by reference, so original slice is modified
	// fmt.Println(a)

	//=========================================================
	// arr := []int{64, 34, 25, 12, 22, 11, 90}
	// fmt.Println("Unsorted slice:", arr)
	// bubbleSort(arr)
	// fmt.Println("Sorted slice:", arr)

	// MARK: make

	// MARK: Bitwise
	// const (
	// 	isAdmin = 1 << iota
	// 	isHeadquartes
	// 	isCTO

	// 	canSeeEurope
	// 	canSeeAsia
	// 	canSeeAfrica
	// )

	// userRoles := canSeeAfrica | canSeeAsia | canSeeEurope
	// user1 := canSeeAfrica | canSeeAsia
	// user2 := canSeeEurope

	// fmt.Println("user1 can see Europe", userRoles&user1 == canSeeEurope)
	// fmt.Println("user2 see Europe", userRoles&user2 == canSeeEurope)

	// user 1
	// 32 | 16 (canSeeAfrica | canSeeAsia) & 8(canSeeEurope)
	// 100000 = 32  // 10000 = 16
	// 1000 = 8		// 1000 = 8
	// -------		// ------
	// 100000 = 32	// 10000 = 16

	// 32 | 16 == 8(canSeeEurope) can not see canSeeEurope

	// strings operations ==========================================
	// a := "Hello! This is string in go"
	// b := "Hello! This is string in go"

	// fmt.Println(strings.Compare(a, b)) // use this: == instead this strings.Compare
	// fmt.Println(strings.Contains(a, "in go"))
	// fmt.Println(strings.Count(a, "s")) // retruned how many are there
	// fmt.Println(strings.Index(a, "H")) // returns 0, because Go uses a zero-based indexing system
	// fmt.Println(strings.Split(a," ")) // returns an array
	// fmt.Println(strings.ReplaceAll(a,"go","Go")) // replace all go with Go
	// fmt.Println(strings.ToUpper(b)) // returns Hello! THIS IS STRING IN GO

	// array,defer, len, cap ======================================================
	// names := [3]string{"Yusuf", "Arnold", "John"}
	// famylies := []string{"Yusuf", "Arnold", "John"}
	// names[0] = "Alex"
	// famylies = append(famylies, "Clark")

	// r:= famylies[:2] // get Yusuf and Aronld
	// fmt.Println(r)
	// fmt.Println(famylies)
	// fmt.Println(cap(nNames))
	// fmt.Println(len(nNames))
	// fmt.Println(nNames)
	// fmt.Println(names)
	// defer fmt.Println(famylies[1])
	// defer helloTime()

	// ----------------------------------------------------
	// fmt.Println("When's Saturday?")
	// fmt.Println(time.Monday)

	// m:= time.Monday
	// fmt.Println(m)

	//====================================================================
	// var String string = "admin"
	// fmt.Println(String)
	// String = "asdasdasdasd"
	// fmt.Println(String)
	// var Number int = 12312312312312
	// var Boolean bool = true

	// fmt.Println(String,Number,Boolean)
	// var n bool

	// a, b := swap("aaa", "bbbb")
	// var a string = aaa
	// var b string = bbbb
	// fmt.Println(a, b)
	// fmt.Println(10 / 3)

	// Strig to bytes =============================
	// s:="this is string"
	// b:= []byte(s)
	// s2 := string(b)
	// fmt.Println(b)
	// fmt.Println(s2)

	// MARK: Runes
	// runes is a sequence of UTF-8 encoded code points.
	// r:= 'A'
	// fmt.Printf("%T, %v", r,r)

	// =============================
	// var a, b, c = swap("a", "b")
	// var g = swap2("a", "b")
	// fmt.Println(a, b, c)
	// fmt.Println(g)

	// MARK:funcitons
	// var x, y int = 3, 5
	// var d = x*x + y*y -> 34
	// var h = float64(d) -> float64 type 34
	// var g = math.Sqrt(h)-> math.Sqrt(x float64) float64 = 5.830951894845301
	// var e = uint(g) -> 5
	// var f = uint(math.Sqrt(float64(x*x + y*y)))
	// fmt.Println(x, y, f)

	//========================================
	// sum := 234
	// fmt.Println(sqrt(float64(sum)), sqrt(-4))

	//===========================================
	// fmt.Println(
	// 	pow(3, 2, 10),
	// 	pow(3, 3, 20),
	// )

	//MARK: bitwise operators
	// a := 10 // 1010
	// b := 3  // 0011
	// fmt.Println(10 & 3) // 0010 -> 2
	// fmt.Println(10 | 3) // 1011 -> 11
	// fmt.Println(10 ^ 3) // 1001 -> 9
	// fmt.Println(10 &^ 3) // 1000 -> 8

}

// func getUserToken() {
// 	panic("unimplemented")
// }

//========================================
// 10 & 3 --> "and" operator
// if 1 & 0 = 0
// if 0 & 0 = 0
// if 1 & 1 = 1
// if 0 & 1 = 0
//-----------------
// 1010
// 0011
// 0010 = 2 -> binary 0010 to decimal 2

//============================================
// 10 | 3 --> "or" operator
// if  1 | 0 = 1
// if  0 | 0 = 0
// if  1 | 1 = 1
// if  0 | 1 = 1
//-----------------
// 1010
// 0011
// 1011 = 11 -> binary 1011 to decimal 11

//=====================================
// 10 ^ 3 --> "exclusive or" operator. If two 1 result 0
// if 1 ^ 0 =  1
// if 0 ^ 0 =  0
// if 1 ^ 1 =  0
// if 0 ^ 1 =  1
//----------------------
// 1010
// 0011
// 1001 = 9 -> binary 1001 to decimal 9

//============================================
// 10 &^ 3 --> "and non" operator. If two 0 result 1
// if  1 &^ 0 = 0
// if  0 &^ 0 = 1
// if  1 &^ 1 = 0
// if  0 &^ 1 = 0
//-------------------------------
// 1010
// 0011
// 0100 = 4 -> binary 0100 to decimal 4
