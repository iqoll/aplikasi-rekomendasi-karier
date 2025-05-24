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

// Menu tampilan utama
func tampilMenu() {
	// I.S F.S menampilkan menu interface untuk user memilih jalannya aplikasi
	fmt.Println("=== Aplikasi Rekomendasi Karier ===")
	fmt.Println("1. Tambah Data Pengguna")
	fmt.Println("2. Lihat Rekomendasi Karier")
	fmt.Println("3. Edit Data")
	fmt.Println("4. Hapus Data")
	fmt.Println("5. Cari Data")
	fmt.Println("6. Urutkan Data")
	fmt.Println("7. Keluar")
}

// Tampil Menu Kategori Data
func tampilMenuData() {
	// I.S F.S menampilkan menu kategori data yang ada
	fmt.Println("=== Kategori Data ===")
	fmt.Println("1.Nama")
	fmt.Println("2.Minat")
	fmt.Println("3.Keahlian")
}

// Tampil Menu Daftar Pengguna
func tampilMenuPengguna(pengguna *DataPengguna, jumlah *int) {
	// I.S F.S menampilkan menu pengguna yang ada
	fmt.Println("=== Daftar Pengguna ===")
	for i := 0; i < *jumlah; i++ {
		fmt.Printf("%d. %s \n", i+1, pengguna[i].nama)
	}
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
