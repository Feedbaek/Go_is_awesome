package main

import "fmt"

// zero value는 변수가 명시적인 초기 값 없이 선언된 경우 갖는 값이다.
// int의 zero value는 0, string의 zero value는 "", 포인터, 인터페이스, 슬라이스, 맵, 채널, 함수의 zero value는 nil이다.
// 아래에서 자료별 zero value 값을 확인해보자.
func main() {
	// int의 zero value는 0
	var i int
	fmt.Println(i)

	// string의 zero value는 ""
	var str string
	fmt.Println(str)

	// 포인터의 zero value는 nil
	var p *int
	fmt.Printf("%v, %v\n", p, p == nil)

	// 인터페이스의 zero value는 nil
	var itf interface{}
	fmt.Printf("%v, %v\n", itf, itf == nil)

	// 슬라이스의 zero value는 nil
	var a []int
	fmt.Printf("%v, %v\n", a, a == nil)
	b := make([]int, 0)
	c := []int{}
	var d []int
	d = nil
	e := make([]int, 0, 0)
	f := make([]int, 0, 1)

	fmt.Println(a == nil, // true
		b == nil, // false
		c == nil, // false
		d == nil, // true
		e == nil, // false
		f == nil) // false
}
