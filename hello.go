package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strings"
)

var globalVariable = "i am global"

func main() {
	// var a, b, c int
	// var d string
	// fmt.Println(a, b, c, d)
	// fmt.Println("len of string = ", len(d))
	// a = 45
	// b = 45
	// fmt.Print("equality = ", a == b, " this is it")

	// var a, b, c = 24, 7.24, "okay"
	// fmt.Println(a, b, c)

	// d, e, f, g:= 934, 99.97, "iaudbfvi", false
	// fmt.Println(d, e, f, g)

	const PI = math.Pi
	// const newPi := math.Pi // error
	// fmt.Println(PI)
	// fmt.Printf("%T", PI)

	const NAME, IS_POSSIBLE = "john wick", true
	// fmt.Println(NAME, IS_POSSIBLE)

	// the below const should be declared above main()
	const (
		ENVIRONMENT = "PROD"
		CORRECT     = false
		DISCOUNT    = 34.35
	)
	// fmt.Println(ENVIRONMENT, CORRECT, DISCOUNT)

	// firstNum, secondNum:= 34, 20
	// fmt.Println(firstNum &^ secondNum) // bitwise not of second opeartor and then bitwise not

	numvalue := 4 // pointers in go, just like C++
	address := &numvalue

	// fmt.Println(*address)
	*address = 7
	// fmt.Println(numvalue)

	// display()

	// var globalVariable int = 8374
	// fmt.Println(globalVariable)

	// var numerator, denominator int = 34, 8
	// var result = float64(numerator) / float64(denominator)
	// fmt.Println(result)

	// var firstnum int = 98
	// var secondnum float64 = 398.24
	// var result = firstnum + secondnum // error
	// result := float64(firstnum) + secondnum
	// fmt.Println(result)

	// secondnum = float64(firstnum) // assignment also done by type-casting
	// fmt.Println(secondnum)

	// var thirdnum int64 = 924
	// firstnum = int(thirdnum) // int64 and int32

	// var (
	// 	firstNum = 24
	// 	secondNum = 242
	// 	thirdNum = 98
	// 	fourthNum = 92
	// )
	// fmt.Println(firstNum + secondNum + thirdNum + fourthNum)

	// var x int
	// fmt.Println(x)

	p, q := 98, 24
	fmt.Println(p, q)
	// p, q := 87384, 133 // error, atleast 1 new variable required for assignment using :=
	p, q, r := 87384, 133, 98249 // works fine
	fmt.Println(p, q, r)
	r = 983
	fmt.Println(r)

	if p > 100 {
		if q < 100 {
			fmt.Println("fist if")
		} else if q > 100 {
			fmt.Println("secod if")
		}
	} else {
		fmt.Println("last if")
	}

	// For loop - TYPE 1:

	for i := 0; i < 100; i++ {
		fmt.Println("hello world")
	}

	// For loop - TYPE 2: (infinite loop)

	// for {
	// 	fmt.Println("i am infinite")
	// }

	// For loop - TYPE 3: (while loop)

	i := 1
	for i < 10 {
		fmt.Println(i)
		i += 1
	}

	// For loop - TYPE 4: (range rvariable in for loop)

	rvariable := []string{"first", "second", "third"}

	for i, j := range rvariable {
		fmt.Println(i, j)
	}
	// i -> index, j -> actual value of rvariable at index i

	// For loop - TYPE 5: (over maps)

	rmap := map[int]string{
		22: "first",
		23: "second",
		98: "last",
	}
	for key, value := range rmap {
		fmt.Println(key, value)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 7 {
			i = i + 2
			continue
		}
		if i == 8 {
			break
		}
	}

	// expression-switch statements;

	switch day := 2; day { // both day:=2 and day are optional in this statement.
	case 1:
		fmt.Println("frist")
	case 2:
		fmt.Println("second")
	case 3:
		fmt.Println("third")
	default:
		fmt.Println("invalid")
	}

	var x int = 2

	switch { // no expression, no statement
	case x == 0:
		fmt.Println("x=0")
	case x == 1:
		fmt.Println("x=1")
	case x == 2:
		fmt.Println("x=2")
	default:
		fmt.Println("invalid")
	}

	var value string = "five"

	switch value { // expression only
	case "one":
		fmt.Println("python")
	case "two", "three":
		fmt.Println("java")
	case "four", "five", "six":
		fmt.Println("golang")
	}

	// type switch statement : (used to compare types only)

	var val interface{}

	switch q := val.(type) {
	case bool:
		fmt.Println("boolean")
	case float64:
		fmt.Println("float64")
	case int:
		fmt.Println("int type")
	default:
		fmt.Printf("%T", q)
	}

	// functions in go language :

	x, y := 2, 5
	fmt.Println(x, y)

	swap_by_val(x, y)
	fmt.Println(x, y)

	swap_by_ref(&x, &y)
	fmt.Println(x, y)

	// anonymous functions

	func(ele string) {
		fmt.Println(ele)
	}("geeks")

	sum := func(a, b int) int { // type of sum is function
		return a + b
	}

	res := sum(5, 7)
	fmt.Println(res)

	modify := func(Z *int) {
		*Z = 10
	}

	Z := 40
	modify(&Z)
	fmt.Println("Z = ", Z)

	// passing & returning anonymous functions as arguments (callbacks) to/from function for later

	// returning mulitiple values from functions :
	// method 1:

	perform_1 := func(a, b int) (int, int, int, int) {
		return a + b, a - b, a * b, a % b
	}

	var sumof, diffof, mulof, remof = perform_1(13, 2)
	fmt.Println(sumof, diffof, mulof, remof)

	var ww, xx, yy, zz = perform_2(29, 3)
	fmt.Println(ww, xx, yy, zz)

	var a1, a2, a3, a4 = perform_3(98, 23)
	fmt.Println(a1, a2, a3, a4)

	var p1, q1, _, _ = perform_3(93, 24) // _ blank identifier for no warning.
	fmt.Println(p1, q1)

	// defer keyword

	fmt.Println("starting")      // (1)
	fmt.Println(multiply(12, 8)) // (2)
	// below 2 statements follow LIFO
	// defer fmt.Println("ending") // this defer gets executed last (4)
	// defer fmt.Println(multiply(9, 7)) // this defer gets executed first (3)

	// struct in Go ( always defined using type keyword ):

	var adrs Address
	fmt.Println("a => ", adrs) // all values set to 0 equivalent(0,"",nil etc) depending on their type

	// initializing of all fields in order DEFINED in struct (all fields mandatory)
	adrs1 := Address{"Akshay", "Dehradun", 9834}
	fmt.Println("Address1 = > ", adrs1)

	// initializing of some/all fields using key:value

	adrs2 := Address{Name: "Modi", Pincode: 2872}
	fmt.Println("Address2 = > ", adrs2, " and city = ", adrs2.city)
	adrs2.city = "Delhi"

	// pointer to a struct (fields can be accessed with/without *)
	employee1 := &Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("employee pointer value employee1.=> ", *employee1) // {Sam Anderson 55 6000}
	fmt.Println("accessing age without *", employee1.age)           // 55
	fmt.Println("accessing age with *", (*employee1).age)           // 55

	employee1.firstName = "Mark"
	fmt.Println("new first name of employee", employee1.firstName) // Mark

	// compare 2 struct of same type:
	// method 1: using ==

	adrs3 := Address{"hello", "pune", 9824}
	adrs4 := Address{"hello", "pune", 9824}

	fmt.Println(adrs3 == adrs4) // true
	fmt.Println(adrs3 == adrs2) // false

	// method 2: usign DeepEqual() in "reflect" package
	fmt.Println(reflect.DeepEqual(adrs3, adrs4)) // true
	fmt.Println(reflect.DeepEqual(adrs3, adrs2)) // false

	// nested struct : 2 ways
	// type 1: name of nested struct field as a variable
	thishr := HR{
		salary: 40,
		details: Author{
			name:   "somenone",
			branch: "CSE",
			year:   2021,
		},
	}
	fmt.Println("thishr = ", thishr)
	fmt.Println(thishr.salary, thishr.details.branch, thishr.details.name)

	// type 2: name of nested struct field as another struct itself
	myperson := Person{
		FirstName: "john",
		LastName:  "wick",
		Department: Department{
			Name: "CSE",
			City: "London",
		},
	}
	// accessing the nested struct fields without using "Promoted fields"
	fmt.Println(myperson.FirstName, myperson.Department.City)
	// accessing the nested struct fields using "Promoted fields"
	fmt.Println(myperson.FirstName, myperson.City) // City directly available in parent struct due to no-name conflicts.

	// to add a method to a struct, define outside the struct and pass the instance of the struct as the argument to the fucntion. this funtion then cannot be directly called from the main fucntion
	fmt.Println(myperson.Department.getCityName()) // calling the method binded to struct Person

	// PROMOTED METHODS IN GO (calling the child struct binded method directly from the parent struct)
	fmt.Println(myperson.getCityName())

	//----------------------------------------------------------------
	// method of having field as of function type :
	myman := Man{
		name:      "sumurth",
		language:  "en",
		increment: 12,
		base:      500,
		salary: func(base, increment int) int {
			return base * increment
		},
	}
	fmt.Println(myman.salary(myman.base, myman.increment))
	fmt.Println(myman.name)

	car := Car{
		name:   "alto",
		speed:  100,
		factor: 20,
		fast: func(factor int, speed int) int {
			return factor + speed
		},
	}
	fmt.Println(car.fast(car.factor, car.speed))
	fmt.Println(car.name)

	calculator := Calculator{
		a: 20,
		b: 30,
		add: func(i1, i2 int) int {
			return i1 + i2
		},
		sub: func(i1, i2 int) int {
			return i1 - i2
		},
		mul: func(i1, i2 int) int {
			return i1 * i2
		},
		div: func(i1, i2 int) int {
			return i1 / i2
		},
	}
	fmt.Println(calculator.add(calculator.a, calculator.b))
	fmt.Println(calculator.sub(calculator.a, calculator.b))
	fmt.Println(calculator.mul(calculator.a, calculator.b))
	fmt.Println(calculator.div(calculator.a, calculator.b))

	// fmt.Println(genericFunc(87, 2)) # does not provide any functionality

	// ----------------------------------------------------------------
	// ARRAYS
	// static, fixed lenght, len can NOT be changed, of same data type only

	// initializaiton:
	// type 1 : using var:
	var arr1 = [3]int{1, 2, 3}
	for i := 0; i < 3; i++ {
		fmt.Println(arr1[i])
	}
	// type 2: shortahand
	names := [4]string{"a", "b", "c", "d"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	// multi-dimaensional :

	arr2 := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr2[i]); j++ {
			fmt.Printf("%d ", arr2[i][j])
		}
		fmt.Println()
	}

	arr2[1][2] = 39
	fmt.Println(arr2)
	fmt.Println(arr2[1])

	// when don't know the len of array beforehand, use ...
	cars := [...]string{"a", "b", "c", "d", "e", "f"}
	for i := 0; i < len(cars); i++ {
		fmt.Println(cars[i])
	}

	arr100 := [2]string{"a", "b"}
	arr200 := [2]string{"a", "b"}
	arr300 := [2]string{"c", "d"}

	oldarr := [3]string{"a", "b", "c"}
	newarr := oldarr // creating a copy of old array ( doesn't reference to old, completely new in memory)

	fmt.Println(arr100 == arr200) // true
	fmt.Println(arr200 == arr300) // false
	// fmt.Println(arr100 == oldarr) // error mis-matched types due to length

	newarr[2] = "d" // no change in oldarr due to this
	fmt.Println(oldarr)
	fmt.Println(newarr)

	// creating new array using pointer to old array, old array changes on chaing new array
	referenceArr := &oldarr
	(*referenceArr)[0] = "z" // OR below
	referenceArr[1] = "y"
	fmt.Println(oldarr)
	fmt.Println(*referenceArr)

	// copying array 3 methods:
	// method 1: using loop
	originalarr := [5]int{1, 2, 3, 4, 5}
	copyarr := make([]int, len(originalarr))

	for i := 0; i < len(originalarr); i++ {
		copyarr[i] = originalarr[i]
	}
	fmt.Println(copyarr)
	copyarr[2] = 10 // no chagne in origiaal
	fmt.Println(copyarr)
	fmt.Println(originalarr)

	// method 2 : using copy fucntion
	copyarr2 := make([]int, len(originalarr))

	copy(copyarr2, originalarr[:]) // copy
	fmt.Println(copyarr2)

	// passing array to a funciton
	var sumArr int = getTotalSumDefinedIndex(originalarr)
	fmt.Println(sumArr)

	myArrOfInt := []int{1, 2, 3, 4, 5, 6, 7}
	newSum := getTotalSumNoDefinedIndex(myArrOfInt)
	println(newSum)

	//----------------------------------------------------------------
	// SLICES
	// dynamic arrays, just don't write the len in the array literal.

	myarr := [4]int{1, 2, 3, 4} // array
	myslice := []int{1, 2, 3, 4}

	// creaing a slice from another array and slice, refer from the origianl in memory. changes reflect

	slice1 := myarr[1:3]
	fmt.Println(slice1)
	slice2 := myslice[2:3]
	fmt.Println(slice2)

	// appending elements to the end of the slice
	slice2 = append(slice2, 7, 8, 9)
	fmt.Println(slice2)
	fmt.Println(myslice) // no change on the myslice

	// len(slice) : len of the slie
	// cap(slice) : max capacity of slice (from the start of the slice to the end of the original arr/slice)

	// creating slice using make() fucntion and creaiton of dynamic arryas (slices) :

	dynamicslice := make([]int, 10)
	dynamicslice[4] = 987
	fmt.Println(dynamicslice) // [0 0 0 0 987 0 0 0 0 0]

	/// iterating over a slice using the len or range :
	// using len :
	for i := 0; i < len(dynamicslice); i++ {
		fmt.Println(dynamicslice[i])
	}
	// using range :
	for _, value := range dynamicslice {
		fmt.Println(value)
	}
	// passing a slice to a fucntion :
	odd := countOdd(dynamicslice)
	fmt.Println(odd)

	// to compare 2 slice, use DeepEqual only, == works only with nil in case of slice

	// multi-dimensional slice :
	multislice := [][]int{
		{3, 2, 1},
		{6, 5, 4},
	}
	fmt.Println(multislice)

	multislice2 := [][]string{
		{"a", "b", "c"},
		{"d", "e", "f"},
	}
	fmt.Println(multislice2)

	multislice[0] = append(multislice[0], 10)
	fmt.Println(multislice) // [[3 2 1 10] [6 5 4]]

	// Sorting a slice :

	sort.Ints(multislice[0])
	fmt.Println(multislice) // [[1 2 3 10] [6 5 4]]
	sort.Strings(multislice2[1])

	// copy a slice into another using make() and copy()

	sourceslice := []int{1, 23, 4, 66}
	destinationslice := make([]int, len(sourceslice))

	copy(destinationslice, sourceslice)
	fmt.Println(destinationslice) // [1 23 4 66]

	//----------------------------------------------------------------
	// STRINGS :

	mystr := "ianvipanvienvinadipfvnadipfuviud"
	for i := 0; i < len(mystr); i++ {
		fmt.Printf("%c ", mystr[i])
	}
	for _, value := range mystr {
		fmt.Println(value) // prints bytecode
	}
	r1, r2, r3 := "abc", "abc", "def"
	fmt.Println(r1 == r2)
	fmt.Println(r2 == r3)

	// string concatenation:
	// method 1:
	fmt.Println(r1 + " " + r2)
	// method 2:
	r1 += r2
	fmt.Println(r1)

	// joing to create a stirng using a sepeator from a list -> strings.Join()
	stringArr := []string{"abc", "def", "ghi"}

	var result string = strings.Join(stringArr, " : ")
	fmt.Println(result) // abc : def : ghi

	// trimming in strings:
	// trim white-space :

	originalstr := "     anafiadvuia  piandfiv    *   "
	spacestr := strings.TrimSpace(originalstr)
	fmt.Println(spacestr, len(spacestr)) // anafiadvuia  piandfiv    * 26
	leftstr := strings.TrimLeft(originalstr, " ")
	fmt.Println(leftstr, len(leftstr)) // anafiadvuia  piandfiv    *    29
	tobedel := "aipniupan aipudnpaidnf aiducviadunfv"
	suffixstr := strings.TrimSuffix(tobedel, "viadunfv")
	fmt.Println(suffixstr) // aipniupan aipudnpaidnf aiduc

	// split a string into array of string :
	// method 1: using split()

	originalstr = "aioun,a,fg,sfga,ert,er,g,agf,erg,a,rtaer,ga,rtge"
	listofstrings := strings.Split(originalstr, ",")
	fmt.Println(listofstrings) // [aioun a fg sfga ert er g agf erg a rtaer ga rtge]
	listofstrings = strings.SplitAfter(originalstr, ",")
	fmt.Printf("listofstrings: %v\n", listofstrings) // listofstrings: [aioun, a, fg, sfga, ert, er, g, agf, erg, a, rtaer, ga, rtge]

	// check the existence of a string in another string using strings.Contains()
	originalstr = "heheh lolol hehehe"
	isPresent := strings.Contains(originalstr, "hehe")
	fmt.Println(isPresent) // true

	// contains any: any character from checkstring present in originalstring
	isGorFpresent := strings.ContainsAny(originalstr, "gf") // false
	fmt.Println(isGorFpresent)
	isGorFpresent = strings.ContainsAny(originalstr, "lz") // true , l present
	fmt.Println(isGorFpresent)

	// repeat a string for N times using Repeat()
	ag := "ok "
	ag4 := strings.Repeat(ag, 4)
	fmt.Println(ag4) // ok ok ok ok
	// or simply using a loop:
	var ag4loop string
	for i := 0; i < 4; i++ {
		ag4loop += ag
	}
	fmt.Println(ag4loop) // ok ok ok ok
	// find index of substirng within a string : -1
	var index int = strings.Index(ag4, " ok")
	indexnotpresent := strings.Index(ag4, " not present")
	fmt.Println(index)           // 2
	fmt.Println(indexnotpresent) // -1

	countOk := strings.Count(ag4, "ok")
	fmt.Println(countOk) // 4

	arr := []int{9, 8, 7, 6, 5, 4, 2, 3, 1, 0}
	newsortedarr := bubbleSortReturn(arr)
	fmt.Println(newsortedarr)
	bubbleSortNoReturn(arr)
	fmt.Println(arr)

	//----------------------------------------------------------------
	// POINTERS

	num := 10
	// declaration of a pointer:
	var ptr *int // nil type
	// initialization of a pointer:
	ptr = &num
	// accessing the pointer
	fmt.Println(*ptr) // 10

	// using shorthand :
	strptr := &ag4
	fmt.Println(*strptr) // ok ok ok ok

	// pointer to a pointer ( double pointer )
	var ptr2 **int = &ptr
	fmt.Println("num = ", num) // 10
	*ptr = 20
	fmt.Println("num = ", num) // 20
	**ptr2 = 30
	fmt.Println("num = ", num) // 30

	// returning pointer from a function
	valPtrFunc := pointerReturnFunc()
	fmt.Println(*valPtrFunc) // 10

	//----------------------------------------------------------------
	// INTERFACES

	// a collection of methods and is a custom type, purely abstract, cannot be used to instantiate
	myvalinstance := myvalue{24.9, 92.5}
	fmt.Println(myvalinstance.Tarea())  // 7232.204999999999
	fmt.Println(myvalinstance.Volume()) // 360163.80899999995

	// IMPORTANT: dynamic type variables using interface

	var dynamicVar interface{}

	dynamicVar = 20
	fmt.Println(dynamicVar) // 20
	dynamicVar = 29.9834
	fmt.Println(dynamicVar) // 29.9834
	dynamicVar = "aiunf"
	fmt.Println(dynamicVar) // aiunf

	// passing interface as argument to function :
	var newdynamicvar interface{} = 98
	interfaceCheckFunc(newdynamicvar)

	// doing type-based opeartion using type-switch

	interfaceRes := typeFunc(984)
	fmt.Println(interfaceRes)

	interfaceRes = typeFunc(93.934)
	fmt.Println(interfaceRes)

	interfaceRes = typeFunc("world")
	fmt.Println(interfaceRes)

	interfaceRes = typeFunc(false)
	fmt.Println(interfaceRes)

	/*
		a is int
		33456
		a is float
		98.914
		a is stirng
		hello world
		unknown type
		<nil>
	*/

	// mutliple interfaces

	values := author{
		name:      "mickey",
		branch:    "cse",
		college:   "kiet",
		salary:    1000,
		year:      2024,
		particles: 1000,
		tarticles: 2500,
	}
	var authordetails AuthorDetails = values
	authordetails.details()
	var authorarticles AuthorArticles = values
	authorarticles.articles()
	fmt.Println()

	// mutiple interfaces and nested interfaces using a struct pointer

	// var doc *Document = &Document{content: "initial content"} OR
	doc := &Document{content: "initial content"}

	fmt.Println(doc.content) // initial content
	// note the initialization of the Reader instance with that of the implementing struct's object/pointer
	var reader Reader = doc
	fmt.Println(reader.Read()) // initial content

	var writer Writer = doc
	writer.Write("new content")

	var readerwriter ReaderWriter = doc
	fmt.Println(readerwriter.Read()) // new content
	readerwriter.Write("latest content")
	fmt.Println(readerwriter.Read()) // latest content
}

func display() {
	// localVariable := "i am local to display func"
	fmt.Println(globalVariable)
	// fmt.Print(localVariable) // local to this function only
}

func swap_by_val(a, b int) {
	var q = a
	a = b
	b = q
}

func swap_by_ref(a, b *int) {
	var q = *a
	*a = *b
	*b = q
}

func perform_2(a, b int) (s, d, m, r int) {
	s = a + b
	d = a - b
	m = a * b
	r = a % b
	return
	// OR simply return a + b, a - b, a * b, a % b
}

func perform_3(a, b int) (s int, d int, m float64, r bool) {
	// s := a + b error, s already declared
	s = a + b
	d = a - b
	m = float64(a) / float64(b)
	r = a > b
	return
}

func multiply(a, b int) int {
	return a * b
}

// structures in go ( type keyword )

type Address struct {
	Name    string
	city    string
	Pincode int
}

type Employee struct {
	firstName, lastName string
	age, salary         int
}

// below 2 structs are for nested struct having another stuct as another varible name
type Author struct {
	name, branch string
	year         int
}

type HR struct {
	salary  int
	details Author
}

// below 2 structs are for nested struct having another stuct as having struct name as varible name
type Department struct {
	Name string
	City string
}

//	type Person struct {
//		FirstName string
//		LastName string
//		Department Department
//	}
//
// OR using the anonymous field for the structue ( PROMOTED FIELDS )
type Person struct {
	FirstName string
	LastName  string
	Department
}

func (d Department) getCityName() string {
	return "city name = " + d.City
}

// struct having field as function -------------------
type Man struct {
	name      string
	language  string
	increment int
	base      int
	salary    getTotalSalary
}
type getTotalSalary func(int, int) int

// a more generic implementation of the above approach : ------
type Car struct {
	name   string
	speed  int
	factor int
	fast   func(int, int) int
}

// a even more generic implementation of the above approach: ------
type genericFunc func(int, int) int
type Calculator struct {
	a   int
	b   int
	add genericFunc
	sub genericFunc
	mul genericFunc
	div genericFunc
}

// ----------------------------------------------------------------
func getTotalSumDefinedIndex(arr [5]int) int { // operatinng for a array
	sum := 0
	for _, value := range arr { // _ for index
		sum += value
	}
	return sum
}
func getTotalSumNoDefinedIndex(arr []int) int { // operation for a slice
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}
func countOdd(arr []int) int {
	cnt := 0
	for _, value := range arr {
		if (value & 1) == 1 {
			cnt += 1
		}
	}
	return cnt
}

func bubbleSortReturn(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap_by_ref(&arr[j], &arr[j+1])
			}
		}
	}
	return arr
}
func bubbleSortNoReturn(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap_by_ref(&arr[j], &arr[j+1])
			}
		}
	}
}
func pointerReturnFunc() *int {
	newval := 10
	return &newval
}

// interface functions : --------------------------------
type tank interface {
	Tarea() float64
	Volume() float64
}
type myvalue struct {
	radius float64
	hieght float64
}

// implementing the methods of the interface :
func (m myvalue) Tarea() float64 {
	return m.radius * 3.14 * m.hieght
}
func (m myvalue) Volume() float64 {
	return 2 * 3.14 * m.radius * m.radius * m.hieght
}

// ------------ passing interface to a function :
func interfaceCheckFunc(a interface{}) {
	value, ok := a.(int)
	fmt.Println(value, ok) // 98 true
}
func typeFunc(a interface{}) interface{} {
	switch a.(type) {
	case int:
		fmt.Println("a is int")
		return 34 * a.(int) // important
	case float64:
		fmt.Println("a is float")
		return 4.98 + a.(float64)
	case string:
		fmt.Println("a is stirng")
		return "hello " + a.(string)
	default:
		fmt.Println("unknown type")
	}
	return nil
}

// -- multiple interfaces :
type author struct {
	name      string
	branch    string
	college   string
	year      int
	salary    int
	particles int
	tarticles int
}
type AuthorDetails interface {
	details()
}
type AuthorArticles interface {
	articles()
}

// implementing the interface functions:
func (a author) details() {
	fmt.Printf("Author Name: %s", a.name)
	fmt.Printf("\nBranch: %s and passing year: %d", a.branch, a.year)
	fmt.Printf("\nCollege Name: %s", a.college)
	fmt.Printf("\nSalary: %d", a.salary)
	fmt.Printf("\nPublished articles: %d", a.particles)
}
func (a author) articles() {
	pendingarticles := a.tarticles - a.particles
	fmt.Printf("\nPending articles: %d", pendingarticles)
}

// mutiple interfaces and nested interface
type Reader interface {
	Read() string
}
type Writer interface {
	Write(string)
}
type ReaderWriter interface { // nested interfaces
	Reader
	Writer
}
type Document struct {
	content string
}

func (d Document) Read() string {
	return d.content
}
func (d *Document) Write(content string) {
	d.content = content
}
