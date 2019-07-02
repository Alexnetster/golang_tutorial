package main

import (
	"fmt"
	"runtime"
)

// func GetFunctionName(i interface{}) string {
// 	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
// }

func GetFunctionName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func testArray() {
	fmt.Println("name:", GetFunctionName())
	var a [3]int //정수형 3개 요소를 갖는 배열 a 선언
	a[0] = 1
	a[1] = 2
	a[2] = 3
	println(a[1]) // 2 출력

	var a1 = [3]int{1, 2, 3}
	var a3 = [...]int{1, 2, 3} //배열크기 자동으로
	println(a1[1])             // 2 출력
	println(a3[1])             // 2 출력

	var multiArray [3][4][5]int // 정의
	multiArray[0][1][2] = 10    // 사용

	println(multiArray[0][1][2]) // 10 출력
	//println(multiArray[0][0][0]) // 10 출력
	fmt.Println(multiArray)

	var b = [2][3]int{
		{1, 2, 3},
		{4, 5, 6}, //끝에 콤마 추가
	}
	println(b[1][2])
}

func testArraySlice() {
	fmt.Println("name:", GetFunctionName())
	var a []int        //슬라이스 변수 선언
	a = []int{1, 2, 3} //슬라이스에 리터럴값 지정
	a[1] = 10
	fmt.Println(a) // [1, 10, 3]출력
}

func testMakeSlice() {
	fmt.Println("name:", GetFunctionName())
	s := make([]int, 5, 10)
	fmt.Println(s)          // [1, 10, 3]출력
	println(len(s), cap(s)) // len 5, cap 10
}

func testNilSlice() {
	fmt.Println("name:", GetFunctionName())
	var s []int

	if s == nil {
		println("Nil Slice")
	}
	println(len(s), cap(s)) // 모두 0
}

func testSubSlice() {
	fmt.Println("name:", GetFunctionName())
	s := []int{0, 1, 2, 3, 4, 5}
	s = s[2:5]
	fmt.Println(s) //2,3,4 출력
	s = s[1:]      // 3, 4
	fmt.Println(s) // 3, 4 출력
}

func testSliceAppend() {
	fmt.Println("name:", GetFunctionName())
	s := []int{0, 1}

	// 하나 확장
	s = append(s, 2) // 0, 1, 2
	// 복수 요소들 확장
	s = append(s, 3, 4, 5) // 0,1,2,3,4,5

	fmt.Println(s)
}

func testSliceMake1() {
	fmt.Println("name:", GetFunctionName())
	// len=0, cap=3 인 슬라이스
	sliceA := make([]int, 0, 3)

	// 계속 한 요소씩 추가
	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		// 슬라이스 길이와 용량 확인
		fmt.Println(len(sliceA), cap(sliceA))
	}

	fmt.Println(sliceA) // 1 부터 15 까지 숫자 출력
}

func testSliceAppend2() {
	fmt.Println("name:", GetFunctionName())
	sliceA := []int{1, 2, 3}
	sliceB := []int{4, 5, 6}

	sliceA = append(sliceA, sliceB...)
	//sliceA = append(sliceA, 4, 5, 6)

	fmt.Println(sliceA) // [1 2 3 4 5 6] 출력
}

func testSliceCopy() {
	fmt.Println("name:", GetFunctionName())
	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)
	copy(target, source)
	fmt.Println(target)               // [0 1 2 ] 출력
	println(len(target), cap(target)) // 3, 6 출력
}

func main() {
	testArray()
	testArraySlice()
	testMakeSlice()
	testNilSlice()
	testSubSlice()
	testSliceAppend()
	testSliceMake1()
	testSliceAppend2()
	testSliceCopy()
}
