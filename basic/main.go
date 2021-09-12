package main

import (
	"fmt"
	"golesson/functions"
)

func main() {
	//variables.Demo1()
	//fmt.Print()
	//conditionals.Demo2()
	//loops.Demo3()
	//arrays.Demo1()
	//slices.Demo2()
	functions.SelamVer()
	var sonuc = functions.Topla(2, 1)
	fmt.Println(sonuc * 10)

	s1, s2, s3, s4 := functions.DortIslem(6, 2)
	fmt.Println(s1, s2, s3, s4)

	v1, v2, v3, _ := functions.DortIslem(6, 2)
	fmt.Println(v1, v2, v3)
}
