﻿package main
 
import "io/ioutil"
 
func main() {
    //파일 읽기
    bytes, err := ioutil.ReadFile("F:\\temp\\1.txt")
    if err != nil {
        panic(err)
    }
    //파일 쓰기
    err = ioutil.WriteFile("F:\\temp\\2.txt", bytes, 0)
    if err != nil {
        panic(err)
    }
}