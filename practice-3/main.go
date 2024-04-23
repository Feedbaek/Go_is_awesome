package main

// 같은 패키지 내에서는 함수, 변수, 상수, 타입 등을 공유할 수 있다.
// 다른 패키지에서 사용하려면 첫 글자를 대문자로 해야 한다.
// 변수, 상수, 타입, 함수 등의 이름이 대문자로 시작하면 public이고, 소문자로 시작하면 private이기 때문이다.

// go run main.go 로 실행하면, sayHello 함수를 찾을 수 없다는 에러가 발생한다.
// go run main.go sayHello.go 로 실행하면, 정상적으로 실행된다. 참조하는 파일을 모두 명시해서 같이 컴파일 해주어야 한다.
// 마치 C언어에서 컴파일 하는 것과 비슷한 것 같다.

const helloMsg string = "Hello"

var score int = 100

func main() {
	sayHello("world")
}
