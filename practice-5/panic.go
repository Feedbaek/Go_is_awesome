package main

import "fmt"

// f() 함수 안에서 defer를 사용해서 d() 함수를 스택에 등록하면, f() 함수가 종료되기 직전에 d() 함수가 실행된다.
// 파일을 열었을 때 파일을 닫는 등의 작업을 할 때 유용하다.

// panic 함수는 현재 함수를 즉시 멈추고 현재 함수 스택의 defer 함수들을 모두 실행한 후 즉시 리턴한다.
// 이 작업은 모든 호출된 함수에서 반복되며, 마지막에는 프로그램이 에러 메시지와 함께 종료된다.

// recover 함수는 defer 함수 안에서만 사용 의미가 있으며, panic 함수에 의한 패닉 상태를 다시 정상 상태로 되돌리는 함수이다.

// defer, panic, recover를 함께 사용하면 예외 처리를 할 때 유용하다.
func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
