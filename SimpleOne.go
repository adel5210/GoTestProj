package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func swap(a, b string) (string, string) {
	return b, a
}

func knowToReturn(x, y int) {
	x = 1
	y = 2
	//	short function only
	return
}

type Vertex struct {
	X int
	Y int
}

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
	add(1, 2)
	x, y := swap("Swapped", "me")
	fmt.Println(x, y)

	i := 42
	p := &i
	fmt.Println(*p)
	v := Vertex{1, 2}
	v.Y
	pv := &v

	arrrr := make([]int, 0, 6)
	for i, val := range arrrr {

	}

}
