package main

func say1(msg string) {
	println(msg)
}

func say2(msg *string) {
	println(*msg)
	*msg = "Changed" //메시지 변경
}

func say3(msg ...string) {
	for _, s := range msg {
		println(s)
	}
}

func sum(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func sum2(nums ...int) (int, int) {
	s := 0     // 합계
	count := 0 // 요소 갯수
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}

func sum3(nums ...int) (count int, total int) {
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return
}

func testAnonymousFunction() {
	sum := func(n ...int) int { //익명함수 정의
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}

	result := sum(1, 2, 3, 4, 5) //익명함수 호출
	println(result)
}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

// 원형 정의
type calculator func(int, int) int

// calculator 원형 사용
func calc2(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}

func testLevelFunc() {
	//변수 add 에 익명함수 할당
	add := func(i int, j int) int {
		return i + j
	}

	// add 함수 전달
	r1 := calc(add, 10, 20)
	println(r1)

	// 직접 첫번째 파라미터에 익명함수를 정의함
	r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
	println(r2)

	// add 함수 전달
	ra1 := calc2(add, 10, 20)
	println(ra1)

	// 직접 첫번째 파라미터에 익명함수를 정의함
	ra2 := calc2(func(x int, y int) int { return x - y }, 10, 20)
	println(ra2)
}

func main() {
	msg := "Hello"
	say1(msg)
	say2(&msg)
	println(msg) //변경된 메시지 출력

	say3("This", "is", "a", "book")
	say3("Hi")

	total := sum(1, 7, 3, 5, 9)
	println(total)

	count, total := sum2(1, 7, 3, 5, 9)
	println(count, total)

	count3, total3 := sum3(1, 7, 3, 5, 9)
	println(count3, total3)

	testAnonymousFunction()
	testLevelFunc()
}
