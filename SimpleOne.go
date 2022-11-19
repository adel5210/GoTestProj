package main

import "fmt"

func main() {
	fmt.Println("Start", "Here")
	a := true
	fmt.Println(a)

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}

		if j := i; j == 10 {
			fmt.Println("Its 10")
		}
	}

	testSwitch := 'a'
	switch testSwitch {
	case 'a':
		fmt.Println('a')
	case 'b':
		fmt.Println('b')

	}

	testType := func(i interface{}) {
		switch t := i.(type) {
		default:
			fmt.Println(t)
		}
	}
	testType("a")
	arr := make([]int, 1)
	mp := make(map[int]string)
	for i := 0; i < 3; i++ {
		arr = append(arr, i)
		mp[i] = "yes"
	}

}
