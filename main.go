package main

import "fmt"

func Demo3() {
	aklimdakiSayi := 80
	tahminEdilenSayi := 0
	tahminSayisi := 0
	var basariDurumu string
	fmt.Println("Bir Sayı Tahmin ediniz")
	fmt.Scanln(&tahminEdilenSayi)
	tahminSayisi = tahminSayisi + 1
	if tahminEdilenSayi > 0 && tahminEdilenSayi <= 100 {
		fmt.Println("Sayı 0-100 arasındadır")
	} else {
		fmt.Println("Sayı 0-100 arasında değildir")
	}
	for aklimdakiSayi != tahminEdilenSayi {
		if aklimdakiSayi > tahminEdilenSayi {
			fmt.Println("Daha büyük bir sayı giriniz")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi = tahminSayisi + 1
		} else if aklimdakiSayi < tahminEdilenSayi {
			fmt.Println("Daha küçük bir sayı giriniz")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi = tahminSayisi + 1
		}
		if tahminSayisi > 0 && tahminSayisi <= 3 {
			basariDurumu = "Süper"
		}
		if tahminSayisi > 3 && tahminSayisi <= 6 {
			basariDurumu = "Başarılı"
		}
		if tahminSayisi > 6 {
			basariDurumu = "Fena değil"
		}
	}
	fmt.Printf("Tebrikler Bildiniz. %v tahminde bildiniz %v ", tahminSayisi, basariDurumu)
}

func main() {
	Demo3()
}
