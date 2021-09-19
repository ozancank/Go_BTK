package loops

import "fmt"

func Demo3() {
	aklimdakiSayi := 80
	tahminEdilenSayi := 0

	fmt.Println("Bir sayı tahmin ediniz")
	fmt.Scanln(&tahminEdilenSayi)

	for aklimdakiSayi != tahminEdilenSayi {
		if tahminEdilenSayi < aklimdakiSayi {
			fmt.Println("Daha büyük bir sayı giriniz")
			fmt.Scanln(&tahminEdilenSayi)
		}

		if tahminEdilenSayi > aklimdakiSayi {
			fmt.Println("Daha küçük bir sayı giriniz")
			fmt.Scanln(&tahminEdilenSayi)
		}
	}

	if tahminEdilenSayi == aklimdakiSayi {
		fmt.Println("Bravo")
	}
}
