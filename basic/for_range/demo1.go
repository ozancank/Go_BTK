package for_range

import "fmt"

func Demo1() {
	sehirler := []string{"Ankara", "İstanbul"}
	for i := 0; i < len(sehirler); i++ {
		println(sehirler[i])
	}

	sozluk := map[string]string{"book": "Kitap", "table": "masa"}
	for k, v := range sozluk {
		fmt.Print(k)
		fmt.Print(" : ")
		fmt.Print(v)
		fmt.Println()
	}
}
