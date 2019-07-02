package main

import "fmt"

func testLiteral() {
	// Raw String Literal. 복수라인.
	rawLiteral := `아리랑\n
아리랑\n
  아라리요`

	// Interpreted String Literal
	interLiteral := "아리랑아리랑\n아리리요"
	// 아래와 같이 +를 사용하여 두 라인에 걸쳐 사용할 수도 있다.
	// interLiteral := "아리랑아리랑\n" +
	//                 "아리리요"

	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)
}

func testTypeConversion() {
	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)
	println(f, u)

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes, str2)
}

func testGoCondition() {
	var k int = 1
	if k == 1 { //같은 라인
		println("One")
	}
	if k == 1 {
		println("One")
	} else if k == 2 { //같은 라인
		println("Two")
	} else { //같은 라인
		println("Other")
	}

	var i int = 3
	var max int = 10
	var val int
	println("testGoCondition")
	if val := i * 2; val < max {
		println(val)
	}

	// 아래 처럼 사용하면 Scope 벗어나 에러
	val++
	println(val)

	var name string
	var category = 1

	switch category {
	case 1:
		name = "Paper Book"
	case 2:
		name = "eBook"
	case 3, 4:
		name = "Blog"
	default:
		name = "Other"
	}
	println(name)

	// Expression을 사용한 경우
	switch x := category << 2; x - 1 {
	//...
	}
}

func grade(score int) {
	switch {
	case score >= 90:
		println("A")
	case score >= 80:
		println("B")
	case score >= 70:
		println("C")
	case score >= 60:
		println("D")
	default:
		println("No Hope")
	}
}

// func typechecker(v int) {
// 	switch v.(type) {
// 	case int:
// 		println("int")
// 	case bool:
// 		println("bool")
// 	case string:
// 		println("string")
// 	default:
// 		println("unknown")
// 	}
// }

func check(val int) {
	switch val {
	case 1:
		fmt.Println("1 이하")
		fallthrough
	case 2:
		fmt.Println("2 이하")
		fallthrough
	case 3:
		fmt.Println("3 이하")
		fallthrough
	default:
		fmt.Println("default 도달")
	}
}

func testFor1() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum)
}

func testFor2() {
	n := 1
	for n < 100 {
		n *= 2
		//if n > 90 {
		//   break
		//}
	}
	println(n)
}

func testFor3() {
	for {
		println("Infinite loop")
	}
}

func testFor4() {
	names := []string{"홍길동", "이순신", "강감찬"}

	for index, name := range names {
		println(index, name)
	}
}

func testFor5() {
	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue // for루프 시작으로
		}
		a++
		if a > 10 {
			break //루프 빠져나옴
		}
	}
	if a == 11 {
		goto END //goto 사용예
	}
	println(a)

END:
	println("End")
}

func testFor6() {
	i := 0

L1:
	for {

		if i == 0 {
			break L1
		}
	}

	println("OK")
}

func main() {
	testLiteral()
	testTypeConversion()
	testGoCondition()
	//grade(90)
	//grade(80)
	grade(70)

	//typechecker(980)
	check(2)

	testFor1()
	testFor2()
	//testFor3()
	testFor4()
	testFor5()
	testFor6()
}
