package main

import (
	"fmt"
	"math"
)

// Sqrt 함수는 뉴턴 방법을 사용하여 x의 제곱근을 계산한다.
// 뉴턴 방법은 z = z - (z^2 - x) / 2z 를 반복하여 z^2이 x와 같아질 때까지 z를 찾는 방법이다.
// math.Sqrt 함수와 비교하면서 오차 값을 어느 정도로 설정해야 똑같아지는지 확인해보자.
func Sqrt(x float64) float64 {
	z := 1.0
	for dif := math.Abs(z*z - x); dif > 0.000000000001; dif = math.Abs(z*z - x) {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
