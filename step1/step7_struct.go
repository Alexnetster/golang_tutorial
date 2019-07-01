package main

import "fmt"

// struct 정의
type person struct {
	name string
	age  int
}

func testStruct1() {
	// person 객체 생성
	p := person{}

	// 필드값 설정
	p.name = "Lee"
	p.age = 10

	fmt.Println(p)

	var p1 person
	p1 = person{"Bob", 20}
	fmt.Println(p1)

	p2 := person{name: "Sean", age: 50}
	fmt.Println(p2)

	px := new(person)
	px.name = "Lee" // p가 포인터라도 . 을 사용한다
	fmt.Println(px)
}

type dict struct {
	data map[int]string
}

//생성자 함수 정의
func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d //포인터 전달
}

func testStruct2() {
	dic := newDict() // 생성자 호출
	dic.data[1] = "A"
	fmt.Println(dic)
}

//Rect - struct 정의
type Rect struct {
	width, height int
}

//Rect의 area() 메소드
func (r Rect) area() int {
	r.width++
	return r.width * r.height
}

func testStructRect() {
	rect := Rect{10, 20}
	area := rect.area() //메서드 호출
	println(rect.width, area)
}

// 포인터 Receiver
func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

func testStructRect2() {
	rect := Rect{10, 20}
	area := rect.area2()      //메서드 호출
	println(rect.width, area) // 11 220 출력
}

func main() {
	testStruct1()
	testStruct2()
	testStructRect()
	testStructRect2()
}
