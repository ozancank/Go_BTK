package channels

import (
	"fmt"
	"time"
)

//asenkron

func CiftSayilar(ciftSayiCn chan int) {
	toplam := 0
	for i := 0; i <= 10; i += 2 {
		toplam = toplam + i
		fmt.Println("Çift sayı çalışıyor")
		time.Sleep(1 * time.Second)
	}
	ciftSayiCn <- toplam
}

func TekSayilar(tekSayiCn chan int) {
	toplam := 0
	for i := 1; i <= 10; i += 2 {
		toplam = toplam + i
		fmt.Println("Tek sayı çalışıyor")
		time.Sleep(1 * time.Second)
	}
	tekSayiCn <- toplam
}
