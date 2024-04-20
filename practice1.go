package main

import "fmt"

// go에서 어레이는 정적 배열, 슬라이스는 동적 배열이다.
// 슬라이스는 배열의 포인터를 가지고 있으며, 길이와 용량을 가지고 있다.
// 길이는 슬라이스에 저장된 요소의 개수, 용량은 슬라이스의 길이가 아닌 할당된 배열의 크기를 의미한다.

// 아래는 슬라이스의 append()를 사용할 때, 언제 재할당이 발생하는지 확인하는 예제이다.

func main() {
	// 슬라이스에 append()를 사용할 때, append의 반환 값으로 기존 슬라이스를 덮어써야 한다.
	// 그렇지 않으면 append() 사용 시, cap이 부족한 경우 새로운 슬라이스를 생성하게 되고, 기존 슬라이스와는 다른 주소를 가리키게 된다.

	var slice []int = []int{1, 2, 3, 4, 5}
	// cap이 부족하기에 새로운 슬라이스를 생성하고, 새로운 슬라이스에 6을 추가한다.
	newSlice := append(slice, 6)
	fmt.Println("first append: ", slice, newSlice)
	// slice의 첫 번째 요소를 변경하면, newSlice에는 영향을 미치지 않는다.
	slice[0] = 100
	fmt.Println("slice changed: ", slice, newSlice)

	// newSlice에 7을 추가한다. 이때, cap이 부족하지 않기에 새로운 슬라이스를 생성하지 않는다.
	new2Slice := append(newSlice, 7)
	fmt.Println("second append: ", slice, newSlice, new2Slice)
	// 같은 주소를 가르키기 때문에 newSlice의 첫 번째 요소를 변경하면, new2Slice에도 영향을 미친다.
	newSlice[0] = 200
	fmt.Println("newSlice changed: ", slice, newSlice, new2Slice)

	// 주소를 출력해보면, newSlice와 new2Slice가 포인터로 같은 배열 주소를 가리키는 것을 확인할 수 있다.
	fmt.Printf("[array address] slice : %p, newSlice: %p, new2Slice: %p\n", slice, newSlice, new2Slice)
	// 하지만 실제 슬라이스의 주소는 다르다. 다른 객체이기 때문이다.
	fmt.Printf("[slice address] slice: %p, newSlice: %p, new2Slice: %p\n", &slice, &newSlice, &new2Slice)
	// 길이와 용량을 출력해보면, newSlice와 new2Slice의 길이가 다르고 용량은 같다는 것을 확인할 수 있다.
	fmt.Println(len(slice), cap(slice), len(newSlice), cap(newSlice), len(new2Slice), cap(new2Slice))
}
