package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isim := "Ozan"
	fmt.Println(s.HasPrefix(isim, "Oz"))
	fmt.Println(s.HasSuffix(isim, "an"))
	fmt.Println(s.Index(isim, "za"))

	harfler := []string{"o", "z", "a", "n"}
	sonuc := s.Join(harfler, "-")
	fmt.Println(sonuc)

	fmt.Println(s.Replace(sonuc, "-", "+", -1))
	fmt.Println(s.Replace(sonuc, "-", "+", 2))

	fmt.Println(s.Split(sonuc, "-"))
	fmt.Println(s.Repeat(sonuc, 5))
}
