package string_functions

import (
	"fmt"
	s "strings"
)

func Demo1() {
	isim := "Ozan Can"
	fmt.Println(s.Count(isim, "z"))
	fmt.Println(s.Count(isim, "a"))
	fmt.Println(s.Index(isim, "a"))
	sonuc := s.Index(isim, "e")

	if sonuc != -1 {
		fmt.Println("e harfi var")
	} else {
		fmt.Println("e harfi yok")
	}

	fmt.Println(s.ToLower(isim))
	fmt.Println(s.ToUpper(isim))
}
