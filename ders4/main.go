package main

import "fmt"

func main() {

	// 5 elemanli integer tipinde array

	var sayilar [5]int
	sayilar = [5]int{0, 1, 2, 3, 4}

	yenisayilar := sayilar[:]

	yenisayilar = append(yenisayilar, 5)
	fmt.Print(yenisayilar, len(yenisayilar), len(sayilar))

}
