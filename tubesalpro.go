package main

import "fmt"

const NMAX int = 100
type pemain struct {
	id, rating, usia int
	nama, posisi, klubAsal, klubTujuan, status string
	nilaiTransfer float64
}
type arrPemain [NMAX]pemain

// Fungsi untuk menampilkan menu utama yang lebih rapi
func menu() {
	fmt.Println("\n==========================================")
	fmt.Println("       SISTEM TRANSFER PEMAIN 2026        ")
	fmt.Println("==========================================")
	fmt.Println(" 1. Tambah Pemain")
	fmt.Println(" 2. Edit Pemain")
	fmt.Println(" 3. Hapus Pemain")
	fmt.Println(" 4. Cari Pemain (Sequential - Nama)")
	fmt.Println(" 5. Cari Pemain (Binary - Rating)")
	fmt.Println(" 6. Tampilkan Pemain (Menu Sorting)")
	fmt.Println(" 7. Exit Program")
	fmt.Println("==========================================")
}

func main() {
	var dataPemain arrPemain
	var nPemain int
	var pilih int
	var selesai bool

	// Pemanggilan inisialisasi data awal agar array tidak kosong
	inisialisasiData(&dataPemain, &nPemain)
	selesai = false

	for selesai == false {
		menu()
		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			tambahPemain(&dataPemain, &nPemain)
		} else if pilih == 2 {
			editPemain(&dataPemain, nPemain)
		} else if pilih == 3 {
			hapusPemain(&dataPemain, &nPemain)
		} else if pilih == 4 {
			cariPemainNama(dataPemain, nPemain)
		} else if pilih == 5 {
			cariPemainRating(dataPemain, nPemain)
		} else if pilih == 6 {
			// [PERBAIKAN] Tidak menggunakan '&' agar data asli tidak ikut ter-sort
			menuTampilPemain(dataPemain, nPemain)
		} else if pilih == 7 {
			selesai = true
			fmt.Println("\nTerima kasih! Program selesai.")
		} else {
			fmt.Println("\n[!] Pilihan menu tidak tersedia.")
		}
	}
}

// Fungsi untuk mengisi array dengan data awal
func inisialisasiData(A *arrPemain, n *int) {
	A[0] = pemain{1, 92, 24, "Erling_Haaland", "ST", "Man_City", "Real_Madrid", "Proses_Transfer", 180.0}
	A[1] = pemain{2, 94, 25, "Kylian_Mbappe", "LW", "Real_Madrid", "Bayern_Munich", "Tersedia", 200.0}
	A[2] = pemain{3, 88, 22, "Pedri", "CM", "Barcelona", "-", "Tersedia", 120.0}
	A[3] = pemain{4, 91, 21, "Jude_Bellingham", "CAM", "Real_Madrid", "Man_City", "Proses_Transfer", 150.0}
	*n = 4
}

// Fungsi untuk menambahkan data pemain baru ke dalam array
func tambahPemain(A *arrPemain, n *int) {
	var pilih int

	fmt.Println("\n=== TAMBAH DATA PEMAIN ===")
	fmt.Print("ID Pemain    : ")
	fmt.Scan(&A[*n].id)
	fmt.Print("Rating       : ")
	fmt.Scan(&A[*n].rating)
	fmt.Print("Usia         : ")
	fmt.Scan(&A[*n].usia)
	fmt.Print("Nama         : ")
	fmt.Scan(&A[*n].nama)
	fmt.Print("Posisi       : ")
	fmt.Scan(&A[*n].posisi)
	fmt.Print("Klub Asal    : ")
	fmt.Scan(&A[*n].klubAsal)
	fmt.Print("Klub Tujuan  : ")
	fmt.Scan(&A[*n].klubTujuan)
	fmt.Print("Nilai (Juta) : ")
	fmt.Scan(&A[*n].nilaiTransfer)

	fmt.Println("-------------------------")
	fmt.Println("1=Tersedia | 2=Proses_Transfer | 3=Selesai")
	fmt.Print("Pilih Status : ")
	fmt.Scan(&pilih)

	if pilih == 2 {
		A[*n].status = "Proses_Transfer"
	} else if pilih == 3 {
		A[*n].status = "Selesai"
	} else {
		A[*n].status = "Tersedia"
	}

	*n = *n + 1
	fmt.Println("\n[+] Data pemain berhasil ditambahkan!")
}

// Fungsi untuk mencetak seluruh isi array pemain dengan format yang rapi
func tampilPemain(A arrPemain, n int) {
	var i int

	fmt.Println("\n========================================================================================================")
	fmt.Printf("%-4s | %-18s | %-6s | %-15s | %-12s | %-6s | %-4s | %-16s\n", "ID", "Nama Pemain", "Posisi", "Klub Asal", "Nilai (Jt E)", "Rating", "Usia", "Status Transfer")
	fmt.Println("--------------------------------------------------------------------------------------------------------")

	if n == 0 {
		fmt.Println("                                      Data pemain masih kosong                                          ")
	} else {
		// Looping untuk menampilkan setiap elemen array dalam format baris tabel
		for i = 0; i < n; i++ {
			fmt.Printf("%-4d | %-18s | %-6s | %-15s | %-12.2f | %-6d | %-4d | %-16s\n", A[i].id, A[i].nama, A[i].posisi, A[i].klubAsal, A[i].nilaiTransfer, A[i].rating, A[i].usia, A[i].status)
		}
	}
	fmt.Println("========================================================================================================")
}

// Menu turunan untuk memilih metode penampilan (sorting) data
// [PERBAIKAN] Parameter A dirubah menjadi 'pass by value' (tanpa pointer)
func menuTampilPemain(A arrPemain, n int) {
	var pilih int

	fmt.Println("\n=== MENU TAMPILKAN PEMAIN ===")
	fmt.Println("1. Tampilkan Semua (Sesuai Input Awal)")
	fmt.Println("2. Urut Nilai Transfer - Ascending (Selection)")
	fmt.Println("3. Urut Nilai Transfer - Descending (Selection)")
	fmt.Println("4. Urut Usia - Ascending (Insertion)")
	fmt.Println("5. Urut Usia - Descending (Insertion)")
	fmt.Print("Pilih pengurutan : ")
	fmt.Scan(&pilih)

	// Karena A bukan pointer, sorting di sini hanya merubah data fotokopian sementara
	if pilih == 1 {
		tampilPemain(A, n)
	} else if pilih == 2 {
		selectionNilaiAsc(&A, n)
		tampilPemain(A, n)
	} else if pilih == 3 {
		selectionNilaiDesc(&A, n)
		tampilPemain(A, n)
	} else if pilih == 4 {
		insertionUsiaAsc(&A, n)
		tampilPemain(A, n)
	} else if pilih == 5 {
		insertionUsiaDesc(&A, n)
		tampilPemain(A, n)
	} else {
		fmt.Println("\n[!] Pilihan tidak tersedia.")
	}
}

// Algoritma Sequential Search untuk mencari indeks berdasarkan ID
func sequentialSearchID(A arrPemain, n int, x int) int {
	var i, idx int
	idx = -1
	i = 0

	for i < n {
		if A[i].id == x {
			idx = i
		}
		i++
	}
	return idx
}

// Algoritma Sequential Search untuk mencari indeks berdasarkan Nama
func sequentialSearchNama(A arrPemain, n int, x string) int {
	var i, idx int
	idx = -1
	i = 0

	for i < n {
		if A[i].nama == x {
			idx = i
		}
		i++
	}
	return idx
}

// Algoritma Binary Search untuk mencari indeks berdasarkan Rating
func binarySearchRating(A arrPemain, n int, x int) int {
	var kiri, kanan, tengah, idx int
	kiri = 0
	kanan = n - 1
	idx = -1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		if A[tengah].rating == x {
			idx = tengah
			kiri = kanan + 1
		} else if x < A[tengah].rating {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	return idx
}

// Subprogram untuk mencari pemain berdasarkan Nama dan menampilkan detailnya
func cariPemainNama(A arrPemain, n int) {
	var x string
	var idx int

	fmt.Println("\n=== CARI PEMAIN (NAMA) ===")
	fmt.Print("Masukkan Nama Pemain : ")
	fmt.Scan(&x)

	idx = sequentialSearchNama(A, n, x)

	if idx == -1 {
		fmt.Println("\n[-] Data pemain tidak ditemukan.")
	} else {
		fmt.Println("\n[+] Data Pemain Ditemukan:")
		fmt.Println("--------------------------------")
		fmt.Println("ID             :", A[idx].id)
		fmt.Println("Nama Lengkap   :", A[idx].nama)
		fmt.Println("Posisi         :", A[idx].posisi)
		fmt.Println("Klub Asal      :", A[idx].klubAsal)
		fmt.Println("Nilai Transfer :", A[idx].nilaiTransfer, "Juta Euro")
		fmt.Println("Status         :", A[idx].status)
		fmt.Println("--------------------------------")
	}
}

// Subprogram untuk mencari pemain berdasarkan Rating dan menampilkan detailnya
func cariPemainRating(A arrPemain, n int) {
	var x, idx int

	// Binary search WAJIB disorting terlebih dahulu
	selectionRatingAsc(&A, n)

	fmt.Println("\n=== CARI PEMAIN (RATING) ===")
	fmt.Print("Masukkan Rating Pemain : ")
	fmt.Scan(&x)

	idx = binarySearchRating(A, n, x)

	if idx == -1 {
		fmt.Println("\n[-] Data pemain tidak ditemukan.")
	} else {
		fmt.Println("\n[+] Data Pemain Ditemukan:")
		fmt.Println("--------------------------------")
		fmt.Println("ID             :", A[idx].id)
		fmt.Println("Nama Lengkap   :", A[idx].nama)
		fmt.Println("Posisi         :", A[idx].posisi)
		fmt.Println("Klub Asal      :", A[idx].klubAsal)
		fmt.Println("Nilai Transfer :", A[idx].nilaiTransfer, "Juta Euro")
		fmt.Println("Rating Pemain  :", A[idx].rating)
		fmt.Println("--------------------------------")
	}
}

// Fungsi untuk mengubah beberapa atribut pemain berdasarkan ID yang dicari
func editPemain(A *arrPemain, n int) {
	var x, idx, pilih int

	fmt.Println("\n=== EDIT DATA PEMAIN ===")
	fmt.Print("Masukkan ID Pemain yang diedit : ")
	fmt.Scan(&x)

	idx = sequentialSearchID(*A, n, x)

	if idx == -1 {
		fmt.Println("\n[-] Data pemain tidak ditemukan.")
	} else {
		fmt.Println("Ditemukan pemain bernama:", A[idx].nama)
		fmt.Println("-------------------------")
		
		fmt.Print("Klub Asal baru   : ")
		fmt.Scan(&A[idx].klubAsal)
		fmt.Print("Klub Tujuan baru : ")
		fmt.Scan(&A[idx].klubTujuan)
		fmt.Print("Nilai baru       : ")
		fmt.Scan(&A[idx].nilaiTransfer)

		fmt.Println("-------------------------")
		fmt.Println("1=Tersedia | 2=Proses_Transfer | 3=Selesai")
		fmt.Print("Status baru      : ")
		fmt.Scan(&pilih)

		if pilih == 2 {
			A[idx].status = "Proses_Transfer"
		} else if pilih == 3 {
			A[idx].status = "Selesai"
		} else {
			A[idx].status = "Tersedia"
		}
		fmt.Println("\n[+] Data berhasil diperbarui!")
	}
}

// Fungsi untuk menghapus satu pemain dari array berdasarkan ID
func hapusPemain(A *arrPemain, n *int) {
	var x, idx, i int

	fmt.Println("\n=== HAPUS DATA PEMAIN ===")
	fmt.Print("Masukkan ID Pemain yang dihapus : ")
	fmt.Scan(&x)

	idx = sequentialSearchID(*A, *n, x)

	if idx == -1 {
		fmt.Println("\n[-] Data pemain tidak ditemukan.")
	} else {
		for i = idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n = *n - 1
		fmt.Println("\n[+] Data berhasil dihapus secara permanen!")
	}
}

// FUNGSI SORTING

// Algoritma Selection Sort untuk mengurutkan Nilai Transfer (Terkecil ke Terbesar)
func selectionNilaiAsc(A *arrPemain, n int) {
	var pass, idx, i int
	var temp pemain

	for pass = 0; pass < n-1; pass++ {
		idx = pass
		for i = pass + 1; i < n; i++ {
			if A[i].nilaiTransfer < A[idx].nilaiTransfer {
				idx = i
			}
		}
		temp = A[pass]
		A[pass] = A[idx]
		A[idx] = temp
	}
}

// Algoritma Selection Sort untuk mengurutkan Nilai Transfer (Terbesar ke Terkecil)
func selectionNilaiDesc(A *arrPemain, n int) {
	var pass, idx, i int
	var temp pemain

	for pass = 0; pass < n-1; pass++ {
		idx = pass
		for i = pass + 1; i < n; i++ {
			if A[i].nilaiTransfer > A[idx].nilaiTransfer {
				idx = i
			}
		}
		temp = A[pass]
		A[pass] = A[idx]
		A[idx] = temp
	}
}

// Algoritma Selection Sort untuk mengurutkan Rating (Terkecil ke Terbesar)
func selectionRatingAsc(A *arrPemain, n int) {
	var pass, idx, i int
	var temp pemain

	for pass = 0; pass < n-1; pass++ {
		idx = pass
		for i = pass + 1; i < n; i++ {
			if A[i].rating < A[idx].rating {
				idx = i
			}
		}
		temp = A[pass]
		A[pass] = A[idx]
		A[idx] = temp
	}
}

// Algoritma Insertion Sort untuk mengurutkan Usia (Terkecil ke Terbesar)
func insertionUsiaAsc(A *arrPemain, n int) {
	var pass, i int
	var temp pemain

	for pass = 1; pass < n; pass++ {
		temp = A[pass]
		i = pass
		for i > 0 && temp.usia < A[i-1].usia {
			A[i] = A[i-1]
			i = i - 1
		}
		A[i] = temp
	}
}

// Algoritma Insertion Sort untuk mengurutkan Usia (Terbesar ke Terkecil)
func insertionUsiaDesc(A *arrPemain, n int) {
	var pass, i int
	var temp pemain

	for pass = 1; pass < n; pass++ {
		temp = A[pass]
		i = pass
		for i > 0 && temp.usia > A[i-1].usia {
			A[i] = A[i-1]
			i = i - 1
		}
		A[i] = temp
	}
}
