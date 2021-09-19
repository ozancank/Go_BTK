package arrays

import "fmt"

func Demo1() {
	var sayilar [5]int
	sayilar[2] = 50
	fmt.Println(sayilar)
	fmt.Println(sayilar[2])

	sayilar2 := [5]int{20, 30, 40, 50, 2}
	fmt.Println(sayilar2)

	var sayilar3 [2][3]int
	sayilar3[0][0] = 1
	sayilar3[0][1] = 3
	sayilar3[0][2] = 5
	sayilar3[1][0] = 2
	sayilar3[1][1] = 4
	sayilar3[1][2] = 6

	fmt.Println(sayilar3)

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(sayilar3[i][j])
		}
	}
}
