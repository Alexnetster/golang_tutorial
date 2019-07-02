package main

import (
	"fmt"
	"runtime"
)

func GetFunctionName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func testMap() {
	fmt.Println("FnName:", GetFunctionName())

	var m map[int]string

	m = make(map[int]string)
	//추가 혹은 갱신
	m[901] = "Apple"
	m[134] = "Grape"
	m[777] = "Tomato"

	// 키에 대한 값 읽기
	str := m[134]
	println(str)

	noData := m[999] // 값이 없으면 nil 혹은 zero 리턴
	println(noData)

	// 삭제
	delete(m, 777)
}

func testMap2() {
	fmt.Println("FnName:", GetFunctionName())

	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "FaceBook",
		"AMZN": "Amazon",
	}

	// map 키 체크
	val, exists := tickers["MSFT1"]
	if !exists {
		println("No MSFT ticker")
	} else {
		println(val)
	}
}

func testMap3() {
	fmt.Println("FnName:", GetFunctionName())

	myMap := map[string]string{
		"A": "Apple",
		"B": "Banana",
		"C": "Charlie",
	}

	// for range 문을 사용하여 모든 맵 요소 출력
	// Map은 unordered 이므로 순서는 무작위
	for key, val := range myMap {
		fmt.Println(key, val)
	}
}

func main() {
	testMap()
	testMap2()
	testMap3()
}
