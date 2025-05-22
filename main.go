package main

import (
	"fmt"
)

type Pengguna struct {
	nama     string
	minat    string
	keahlian string
}

type Karier struct {
	nama     string
	minat    string
	keahlian string
}

// kapasitas maksimum array pengguna
const MAX_PENGGUNA = 100

// kapasitas maksimum array karier
const MAX_KARIER = 10

// array statis pengguna
type DataPengguna [MAX_PENGGUNA]Pengguna

var daftarPengguna DataPengguna

// array statis karier dan beberapa pilihan karier
type DataKarier [MAX_KARIER]Karier

var daftarKarier = DataKarier{
	{"Software Developer", "Teknologi", "Coding"},
	{"Web Developer", "Teknologi", "Coding"},
	{"Game developer", "Gaming", "Coding"},
	{"Jurnalis", "Jurnalistik", "Menulis"},
	{"Web Designer", "Teknologi", "Desain"},
	{"Content Creator", "Hiburan", "Komunikasi"},
	{"Graphic Designer", "Seni", "Desain"},
	{"Digital Marketer", "Bisnis", "Komunikasi"},
	{"Akuntan", "Bisnis", "Angka"},
	{"Guru", "Pendidikan", "Mengajar"},
}

func main() {
	var pilihan int
	// counter jumlah pengguna yang ada
	var jumlahPengguna int = 0

	for {
		tampilMenu()
		fmt.Print("Pilih menu:  ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahData(&daftarPengguna, &jumlahPengguna)
		case 2:
			lihatRekomen(daftarPengguna, jumlahPengguna)
		case 3:
			editData(&daftarPengguna, jumlahPengguna)
		case 4:
			hapusData(&daftarPengguna, &jumlahPengguna)
		case 5:
			cariData(daftarPengguna, daftarKarier, jumlahPengguna)
		case 6:
			urutkanData(&daftarPengguna, jumlahPengguna)
		case 7:
			fmt.Println("Terimakasih telah menggunakan aplikasi ini.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih menu yang tersedia.")
		}
	}
}
