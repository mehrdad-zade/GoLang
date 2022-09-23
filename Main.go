package main

/*

 */

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect" // for struct tags
	"strconv"
)

var x int = 2
var DescriptiveName string = "loyal"

func main() {
	i := 1
	var c complex64 = 10 + 2i // another way to initialise is : var c complex64 = complex(10, 2). complex42 is based on 2 float32s
	var r rune = 'a'          // runes are of type int32
	//text in Go are either string (UTF-8) or rune (UFT-32, alias for int32)
	fmt.Println("let's Go")

	fmt.Printf("--values: %v, %v, %v, %T, %T\n", x, DescriptiveName, strconv.Itoa(int(real(c))+x), c, r)

	fmt.Println("--Logic:")
	fmt.Println(x & i) // 10 & 01 = 0
	fmt.Println(x | i) // 11 -> 3
	fmt.Println(x ^ i)
	fmt.Println(x &^ i)
	fmt.Println(x << 3) // x * 2^3
	fmt.Println(x >> 3) // x / 2^3

	//----------------------------------------------------
	//constants: naming conventiosn are like vars; camel-casing. in case they should be exported, first char will be capitalized like vars
	//cannot initialize const with something that should be determined in run-time; such as a method. i.e this is not accepted: const c := math.sin(2)
	const myConst int = 100
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	) // iota starts from zero and keeps incrementing; scoped to a block
	fmt.Printf("--Constants: %v, %v, %v\n", c1, c2, c3)

	const (
		_  = iota // if you don't care about the zero value, iota will generate 0 but won't assing it to a const that can be retrieved
		KB = 1 << (10 * iota)
		MB // compiler infers the pattern being iota
		GB
	)
	fileSize := 400000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)

	const (
		isAdmin = 1 << iota //assing each rol to one byte
		hasReadAccess
		hasWriteAccess
		isManager
	)
	var roles byte = isAdmin | hasReadAccess | hasWriteAccess
	fmt.Printf("%b = %v\n", roles, roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin) // true, because was added to the role
	fmt.Printf("Is Manager? %v\n", isManager&roles == isManager)

	//----------------------------------------------------Arrays and Slice
	//Arrays: call by value
	grades := [3]int{2, 4, 6}
	fmt.Println("--Arrays:", grades)

	newGrades := [...]int{1, 2, 3, 4, 5, 6} //three-dots makes the array large enough to house the size of initialized values
	newGrades[0] = len(newGrades)
	fmt.Println("--Arrays a bit flexible:", newGrades)

	image := [2][2]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println("--Arrays 2d:", image[0][0])

	arrPassVal := grades
	arrPassRef := &grades
	arrPassRef[1] = 200
	fmt.Printf("--Array is pass by value: %v, %v, %v\n", grades, arrPassVal, arrPassRef)

	//slice: call by ref. be mindful of len and cap. cap is the max memory size that is assigned to this slice, which grows in a ^2 fasion. len is just the num of elements in the slice
	gradeSlice := []int{2, 2, 2, 2, 2, 2} //slice size is not fixed like arrays
	gradePass := gradeSlice
	gradePass[0] = 100
	fmt.Println("--Slice is pass by ref:", gradeSlice[0:3])

	//create slice with make()
	makeGrades := make([]int, 3)                       //third arg is optional. to define the size, and make it an array
	makeGrades = append(makeGrades, 3, 3, 3, 3, 3)     // append is an expensive opt
	makeGrades = append(makeGrades, []int{5, 5, 5}...) // to add a slice to a slice you need three-dots because your slice was supposed to get int not slice
	makeGradesPopFromLeft := makeGrades[1:]
	makeGradesPopFromRight := makeGrades[:len(makeGrades)-1]
	makeGradesPopFromMiddle := append(makeGrades[:2], makeGrades[3:]...)
	fmt.Println("--Slice with make: ", makeGrades, makeGradesPopFromLeft, makeGradesPopFromRight, makeGradesPopFromMiddle) // if there are refrences the outcome could be unexpected

	//----------------------------------------------------Maps and Struct
	//maps: order is not guaranteed. call by reference.

	//initialize a map without values: make
	myMap2 := make(map[string]int)
	fmt.Println("Empty map: ", myMap2)

	myMap := map[string]int{
		"key1": 10, //break is implied. if you need fallthrough, it has to be explicitly added, but it will avoid the condition. 
		"key2": 20,
		"key3": 30,
	}
	myMap["key4"] = 40
	delete(myMap, "key1") //this will remove key1 from map, but if you pull the value for key1 it will return "1"
	fmt.Println("Maps: ", myMap, myMap["key1"], len(myMap))

	//check to see if a key-val is present
	val, ok := myMap["key1"]
	fmt.Printf("Value = %v, was is present? %v\n", val, ok)

	//struct: call by value.
	type Doctor struct {
		number  int
		name    string
		friends []string
	}
	aDoctor := Doctor{
		number:  3,
		name:    "John Doe",
		friends: []string{"me", "myself"},
	}
	fmt.Println("Struct: ", aDoctor, aDoctor.friends[0])

	//ananymous struct: has no name, so only used locally to the current scope
	myStruct := struct{ name string }{name: "John Doe"}
	fmt.Println("Ananym Struct: ", myStruct)

	//GO doesn't support inheritance, instead you can do Composition, which means, you can Embed a "parent" type within a struct
	type Animal struct {
		name   string
		origin string
	}
	type Bird struct {
		Animal        // composition and embeding
		name   string `required max:"100"` //tags can provide some validation mechanism on the structs. tags by themselves are meaningless to GO; but you can get the tags for the fileds using "reflect" lib and build your logic
		speed  float32
		canFly bool
	}
	bird := Bird{
		name:   "Gonjishk",
		speed:  40.,
		canFly: true,
	}
	bird.Animal.name = "Parande"
	bird.Animal.origin = "Iran"
	t := reflect.TypeOf(Bird{})
	field, _ := t.FieldByName("name")
	fmt.Println("Bird/Tag: ", bird, field.Tag)

	//----------------------------------------------------IF and SWITCH

	if true {
		fmt.Println()
	}
	if popVal, ok := myMap["key2"]; ok { // notice ok has the bool val, and the statement generating the params are ended with ";". popVal is only valid within the if block's scope.
		fmt.Println("If with ok statemets: ", popVal)
	}

	switch "xyz" {
	case "xyz", "abc":
		fmt.Println("ok") // if first or second item are equal => true
	case "XYZ":
		fmt.Println("n/a")
	default:
		fmt.Println("nope")
	}

	switch i := 1 + 2 + 3; i { // notice how "i" is the switch param after ";"
	case 1, 5, 6:
		fmt.Println("ok")
	case 10, 12, 15:
		fmt.Println("n/a")
	default:
		fmt.Println("nope")
	}

	i = 10
	switch { // notice how "i" is the switch param after ";"
	case i == 1:
		fmt.Println("ok")
	case i <= 10:
		fmt.Println("n/a")
	default:
		fmt.Println("nope")
	}

	var typ interface{} = 2
	switch typ.(type) {
	case int:
		fmt.Println("typ is int")
	case float32:
		fmt.Println("typ is float32")
	case bool:
		fmt.Println("typ is bool")
	}	
	
	//----------------------------------------------------Loop
	for k:=0; k<10; k++ {
		fmt.Println("loop: ", k)
	}

	//scope of vars in a loop
	p := 0
	for ; p<10; p++{
		if p == 9 {
			break
		}
	}
	fmt.Println("P is defined outside of for loop: ", p)

	//loop through maps and slices/arrays
	for k, v := range gradeSlice {
		fmt.Println(k,v)
	}
	for k, v := range myMap {
		fmt.Println(k,v)
	}



	//----------------------------------------------------Defer, Panic and Recover
	//Defer: parameters passed to a defered line, are the ones that are passed to it sequentially, not the final value of such params/args
	fmt.Println("one")
	defer fmt.Println("defer two") // this will run AFTER main func, but BEFORE it returns. defered statements run in LIFO fasion due to dependency
	fmt.Println("three")

	//error
	res, err := http.Get("http://google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() //good option to be defered when there a re a few statements for res. without "defer" keyword in this line, this piece will fail because we closed res before reading it

	robots, err := ioutil.ReadAll(res.Body)	
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Robots http: %s", robots)

	//panic: similar to the use of exception. panic runs after defer. to prevent the app from stopping after panicing, you can have an internal ananymos func with recover() to handle the panic. if the error is not handleable you can panic again
	nom, denom := 2, 1 // change 1 to 0 to see how panic works
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Application paniced but I handled it")
		}
	}() //the "()" is the way to execute the ananymous function
	if denom == 0 {
		fmt.Println("--------------------------")
		panic("Division by zero! -- " + err.Error())
	}else {
		fmt.Println(nom/denom)
	}

	//----------------------------------------------------simple web app
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Go Says Hi.."))
	})
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err.Error())
	}
}
