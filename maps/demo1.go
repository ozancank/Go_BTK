package maps

import "fmt"

func Demo1() {
	sozluk := make(map[string]string)
	sozluk["table"] = "Masa"

	fmt.Println(sozluk["table"])
	fmt.Println(len(sozluk))
	fmt.Println(sozluk)

	deger, varMi := sozluk["table"]
	fmt.Println(deger)
	fmt.Println(varMi)
}
