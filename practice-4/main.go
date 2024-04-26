package main

import "fmt"

// int, string, 배열 등의 값은 함수 내에서 변경되어도 원본 값에는 영향을 미치지 않는다.
func changeValues(age int, name string, hobby [3]string) {
	age = 20
	name = "minskim3"
	hobby[0] = "football"
}

// 슬라이스와 맵은 포인터를 가지고 있기 때문에, 함수 내에서 변경되면 원본 값에도 영향을 미친다.
func changePointerValues(fruits []string, prices map[string]int) {
	fruits[0] = "orange"
	prices["carrot"] = 200
}

func main() {
	// Non-Pointer Values
	age := 10
	name := "minskim2"
	hobby := [3]string{"soccer", "baseball", "basketball"}

	// Pointer Wrapper Values
	fruits := []string{"apple", "banana", "cherry"}
	prices := map[string]int{"carrot": 100, "cucumber": 150, "lettuce": 200}

	// Non-Pointer Values는 함수 내에서 변경되어도 원본 값에는 영향을 미치지 않는다.
	changeValues(age, name, hobby)

	fmt.Println("age: ", age)
	fmt.Println("name: ", name)
	fmt.Println("hobby: ", hobby)

	// Pointer Wrapper Values는 함수 내에서 변경되면 원본 값에도 영향을 미친다.
	changePointerValues(fruits, prices)

	fmt.Println("fruits: ", fruits)
	fmt.Println("prices: ", prices)
}
