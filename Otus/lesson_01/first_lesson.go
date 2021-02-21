package main

import (
	"fmt"
)

func main() {
	fmt.Println(`Hello, 
	
	
	playground`)

	var a int = 5
	var b = 6
	var bb int8 = 6
	c := 7
	var d = 1
	b, e := 5, 5
	fmt.Printf("a = %d, b = %d, bb = %d, c = %d, d = %d, e = %d\n", a, b, bb, c, d, e)
	fmt.Println(b)

	if _, err := DoSomething(); err != nil {
		println("yes")
	}

	i := 0
	for {
		if i%2 == 1 {
			i++
			continue
		}
		println(i)
		i++
		if i > 10 {
			break
		}
	}
	arr := [...]int{81, 54, 43, 66}
	for i, v := range arr {
		println(i, v)
	}

	swData := 15
	switch swData {
	case 10:
		println("One")
		fallthrough
	case 15:
		println("Two")
		fallthrough
	case 30:
		println("Three")
		fallthrough
	default:
		println("Default!")
	}
}

func DoSomething() (int, error) {
	return 2, fmt.Errorf("division by zero")
}
