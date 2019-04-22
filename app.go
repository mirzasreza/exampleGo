package main

import (
	"fmt"
	"errors"
	"math"
)

type person struct {
	name string
	age int
}

func main() {
	fmt.Println("Hello, World!")

	//var x int = 70
	x := 71
	//var y int = 30
	y := 31
	//var sum int = x + y
	addNum := x + y
	fmt.Println(addNum)

	if x > 70 {
		fmt.Println("X greater than 70")
	} else if x > 100 {
		fmt.Println("X greater than 100")
	} else {
		fmt.Println("X less than 70")
	}

	// array
	//var a[5]int
	a := [5]int {1,2,3,4,5}
	a[2] = 7
	fmt.Println(a)


	b:= []int {1,2,3,4,5}
	b = append(b, 6)
	fmt.Println(b)

	// next map
	veritices := make(map[string]int)
	veritices["triangle"] = 3
	veritices["square"] = 2
	veritices["dodecagon"] = 12
	fmt.Println(veritices)
	fmt.Println(veritices["square"])
	delete(veritices, "square")
	fmt.Println(veritices)

	// only loop in go is for loop
	for i:=0; i<5; i++ {
		fmt.Println(i)
	}

	// while loog using for loop
	i := 0
	for i<5 {
		fmt.Println(i)
		i++
	}

	// range for array
	arr := []string{"apple", "ball", "cat"}
	for key, value := range arr {
		fmt.Println("key:", key, "value:", value)
	}
	// range for slice
	m := make(map[string]string)
	m["a"] = "apple"
	m["b"] = "ball"
	m["c"] = "cat"
	for key, value := range m {
		fmt.Println("key:", key, "value:", value)
	}

	// calling external function with single parameter
	sumResult := sum(2, 3)
	fmt.Println(sumResult)

	// calling external function with double parameter
	result, err := sqroot(16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// struct
	p := person {name: "Mirza", age: 40}
	fmt.Println(p)
	fmt.Println(p.age)

}

func sum(x int, y int) int {
	return x + y
}

func sqroot(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined error for negetive number")
	}
	
	return math.Sqrt(x), nil
}
