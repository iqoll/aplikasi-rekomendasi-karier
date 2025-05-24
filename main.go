package main

import (
	"fmt"
	"strings" // kami menggunakan strings untuk memudahkan pencocokan data
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

// Procedure Menambahkan Data Pengguna
func tambahData(pengguna *DataPengguna, jumlah *int) {
	// {I.S. Terdefinisi array pengguna sebanyak jumlah data.
	// F.S. Menambahkan satu data pengguna baru ke dalam array, dan memperbarui nilai jumlah}
	if *jumlah >= MAX_PENGGUNA {
		fmt.Println("Data sudah penuh")
		return
	}

	fmt.Print("Masukkan Nama:  ")
	fmt.Scanln(&pengguna[*jumlah].nama)

	fmt.Print("Masukkan minat (contoh: Teknologi, Seni, Bisnis, dll):  ")
	fmt.Scanln(&pengguna[*jumlah].minat)

	fmt.Print("Masukkan keahlian (contoh: Coding, Desain, Analisis Data, dll): ")
	fmt.Scanln(&pengguna[*jumlah].keahlian)

	// menambahkan indeks atau penghitung
	*jumlah++
}

// Procedure Rekomendasi Karier
func lihatRekomen(pengguna DataPengguna, jumlah int) {
	// {I.S. Terdefinisi array pengguna sebanyak jumlah data.
	// F.S. Menampilkan daftar pengguna, meminta pilihan pengguna dari user, dan menampilkan rekomendasi karier berdasarkan minat dan keahlian pengguna tersebut.}
	var angka int
	tampilMenuPengguna(&pengguna, &jumlah)
	fmt.Print("Pilih nomor pengguna untuk melihat rekomendasi: ")
	fmt.Scan(&angka)

	if angka < 1 || angka > jumlah {
		fmt.Println("Nomor tidak valid.")
		return
	}

	userTerpilih := pengguna[angka-1]
	tampilkanRekomendasi(userTerpilih)

	var kembali string
	fmt.Print("\nKetik apa saja lalu tekan ENTER untuk kembali ke menu utama: ")
	fmt.Scanln(&kembali)
	fmt.Println("Kembali ke menu utama...")
}

// Logic menggunakan algoritma sequential search untuk mencocokan data pengguna dan karier
func tampilkanRekomendasi(p Pengguna) {
	// {I.S. Terdefinisi sebuah data pengguna p dengan atribut nama, minat, dan keahlian.
	// F.S. Menampilkan daftar karier yang sesuai dengan minat dan keahlian pengguna, atau pesan jika tidak ada yang cocok. Menunggu input sebelum kembali ke menu utama.}
	fmt.Printf("=== Rekomendasi Karier untuk %s  ===\n", p.nama)

	cocokDitemukan := false
	for i := 0; i < MAX_KARIER; i++ {
		if strings.ToLower(p.minat) == strings.ToLower(daftarKarier[i].minat) && strings.ToLower(p.keahlian) == strings.ToLower(daftarKarier[i].keahlian) {
			fmt.Println("- " + daftarKarier[i].nama)
			cocokDitemukan = true
		}
	}
	if !cocokDitemukan {
		fmt.Println("Maaf, belum ada rekomendasi yang cocok.")
	}

	var kembali string
	fmt.Print("\n Ketik apa saja lalu tekan ENTER untuk kembali ke menu utama: ")
	fmt.Scan(&kembali)
}

// Procedure Mengedit Data Pengguna
func editData(pengguna *DataPengguna) {
	// {I.S. Terdefinisi array pengguna
	// F.S. Mengedit salah satu data pengguna (nama, minat, atau keahlian) apabila ditemukan berdasarkan nama yang diinputkan.}
	var nama string
	fmt.Println("=== Edit Data Pengguna ===")
	fmt.Print("Masukkan nama pengguna yang ingin diedit: ")
	fmt.Scan(&nama)
	var angka int
	cocokDitemukan := false

	// mencari nama menggunakan algoritma sequential search
	for i := 0; i < MAX_PENGGUNA; i++ {
		if strings.ToLower(nama) == strings.ToLower(pengguna[i].nama) {
			fmt.Print("Data ditemukan: \n\n")
			fmt.Printf("Nama :%s \n", pengguna[i].nama)
			fmt.Printf("Minat :%s \n", pengguna[i].minat)
			fmt.Printf("Keahlian :%s \n", pengguna[i].keahlian)
			cocokDitemukan = true

			// Pemilihan kategori data yang akan diedit
			tampilMenuData()
			fmt.Print("Pilih nomor data yang mau di edit: ")
			fmt.Scan(&angka)
			if angka < 1 || angka > 3 {
				fmt.Println("Nomor tidak valid.")
			} else if angka == 1 {
				fmt.Printf("Masukkan data nama yang baru: ")
				fmt.Scan(&pengguna[i].nama)
			} else if angka == 2 {
				fmt.Printf("Masukkan data minat yang baru: ")
				fmt.Scan(&pengguna[i].minat)
			} else {
				fmt.Printf("Masukkan data keahlian yang baru: ")
				fmt.Scan(&pengguna[i].keahlian)
			}
		}
	}
	if !cocokDitemukan {
		fmt.Printf("Maaf, data dengan nama '%s' tidak ditemukan.", nama)
	}
	var kembali string
	fmt.Print("\n Ketik apa saja lalu tekan ENTER untuk kembali ke menu utama: ")
	fmt.Scanln(&kembali)
	fmt.Println("Kembali ke menu utama...")
}

// Procedure Menghapus Data Pengguna
func hapusData(pengguna *DataPengguna, jumlah *int) {
	// {I.S. Terdefinisi array pengguna sebanyak jumlah data.
	//F.S. Menghapus data pengguna berdasarkan nama, lalu menggeser elemen array agar tidak ada celah, dan mengurangi nilai jumlah.}
	var nama, yakin string
	var j int
	fmt.Println("=== Hapus Data Pengguna ===")
	fmt.Print("Masukkan nama pengguna yang ingin dihapus: ")
	fmt.Scan(&nama)
	cocokDitemukan := false

	for i := 0; i < MAX_PENGGUNA; i++ {
		if strings.ToLower(nama) == strings.ToLower(pengguna[i].nama) {
			fmt.Print("Data ditemukan: \n\n")
			fmt.Printf("Nama :%s \n", pengguna[i].nama)
			fmt.Printf("Minat :%s \n", pengguna[i].minat)
			fmt.Printf("Keahlian :%s \n", pengguna[i].keahlian)
			cocokDitemukan = true

			fmt.Print("Apakah Anda yakin mau menghapus data ini? (Yes/ No)")
			fmt.Scan(&yakin)

			if yakin == "Yes" {
				j = i
				for j <= *jumlah-2 {
					pengguna[j] = pengguna[j+1]
					j = j + 1
				}
				*jumlah = *jumlah - 1
				fmt.Printf("Data telah berhasil dihapus")
			} else {
				fmt.Printf("Data tidak jadi dihapus")
			}
		}
	}
	if !cocokDitemukan {
		fmt.Printf("Maaf, data dengan nama '%s' tidak ditemukan.", nama)
	}
	var kembali string
	fmt.Print("\n Ketik apa saja lalu tekan ENTER untuk kembali ke menu utama: ")
	fmt.Scanln(&kembali)
	fmt.Println("Kembali ke menu utama...")
}

// Procedure Mencari Data Pengguna berdasarkan Kategori Data
func cariData(pengguna DataPengguna, karier DataKarier, n int) {
	// {I.S. Terdefinisi array pengguna sebanyak n data, dan array karier berisi daftar karier yang tersedia.
	// F.S. Menampilkan daftar pengguna yang sesuai dengan suatu minat dan keahlian yang diminta user.}
	var i, angka int
	var cari string

	tampilMenuData()
	fmt.Println("Cari data berdasarkan: ")
	fmt.Scanln(&angka)
	if angka == 1 {
		fmt.Println("Masukkan nama yang dicari: ")
		fmt.Scanln(&cari)
		for i = 0; i < n; i++ {
			if pengguna[i].nama == cari {
				fmt.Printf("Nama: %s \n", pengguna[i].nama)
				fmt.Printf("Minat: %s \n", pengguna[i].minat)
				fmt.Printf("Keahlian: %s \n", pengguna[i].keahlian)
			}
		}
	} else if angka == 2 {
		fmt.Println("Masukkan minat yang dicari: ")
		fmt.Scanln(&cari)
		for i = 0; i < MAX_KARIER; i++ {
			if karier[i].minat == cari {
				fmt.Println("- Rekomendasi Karir: ", karier[i].nama)
				fmt.Println("- keahlian yang dibutuhkan: ", karier[i].keahlian)
				fmt.Println()
			}
		}
	} else {
		fmt.Println("Masukkan keahlian yang dicari: ")
		fmt.Scanln(&cari)
		for i = 0; i < MAX_KARIER; i++ {
			if karier[i].keahlian == cari {
				fmt.Println("-Rekomendasi Karir: ", karier[i].nama)
				fmt.Println("-jika Anda berminat pada bidang", karier[i].minat)
				fmt.Println()
			}
		}
	}
}

// Procedure Mengurutkan Data
func urutkanData(pengguna *DataPengguna, n int) {
	// {I.S. Terdefinisi array data pengguna sebanyak n
	// F.S. Array pengguna terurut berdasarkan pilihan}

	var angka, pass, i, j, min, max int
	var temp Pengguna

	fmt.Println("Urutkan data berdasarkan:")
	fmt.Println("1. Urutan A-Z")
	fmt.Println("2. Urutan Z-A")
	fmt.Print("Pilihan: ")
	fmt.Scanln(&angka)

	if angka == 1 {
		// Selection Sort A-Z
		for pass = 0; pass < n; pass++ {
			min = pass
			for i = pass + 1; i < n; i++ {
				if pengguna[i].nama < pengguna[min].nama {
					min = i
				}
			}
			temp = pengguna[pass]
			pengguna[pass] = pengguna[min]
			pengguna[min] = temp
		}
	} else if angka == 2 {
		// Selection Sort Z-A
		for pass = 0; pass < n; pass++ {
			max = pass
			for i = pass + 1; i < n; i++ {
				if pengguna[i].nama > pengguna[max].nama {
					max = i
				}
			}
			temp = pengguna[pass]
			pengguna[pass] = pengguna[max]
			pengguna[max] = temp
		}
	} else {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	// Tampilkan hasil setelah diurut
	fmt.Println("=== Daftar Pengguna ===")
	for j = 0; j < n; j++ {
		fmt.Printf("%d. %s\n", j+1, pengguna[j].nama)
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
			editData(&daftarPengguna)
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
