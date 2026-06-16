package main

import "fmt"

const NMAX int = 100

type pemain struct {
	id            int
	nama          string
	posisi        string
	klubAsal      string
	klubTujuan    string
	nilaiTransfer float64
	rating        int
	usia          int
	status        string
}

type arrPemain [NMAX]pemain

func menu() {

	fmt.Println()
	fmt.Println("========= SISTEM TRANSFER PEMAIN =========")
	fmt.Println("1. Tambah Pemain")
	fmt.Println("2. Edit Pemain")
	fmt.Println("3. Hapus Pemain")
	fmt.Println("4. Cari Pemain (Sequential - Nama)")
	fmt.Println("5. Cari Pemain (Binary - Rating)")
	fmt.Println("6. Tampilkan Pemain (Menu Sorting)")
	fmt.Println("7. Statistik Transfer")
	fmt.Println("8. Exit")
	fmt.Println("==========================================")
}

func main() {

	var dataPemain arrPemain
	var nPemain int

	var pilih int
	var selesai bool

	// Inisialisasi data awal biar nggak kosong
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

			menuTampilPemain(&dataPemain, nPemain)

		} else if pilih == 7 {

			statistik(dataPemain, nPemain)

		} else if pilih == 8 {

			selesai = true
			fmt.Println("Program selesai")

		} else {

			fmt.Println("Menu tidak tersedia")
		}
	}
}

func inisialisasiData(A *arrPemain, n *int) {

	A[0] = pemain{1, "Erling_Haaland", "ST", "Man_City", "Real_Madrid", 180.0, 92, 24, "Proses_Transfer"}
	A[1] = pemain{2, "Kylian_Mbappe", "LW", "Real_Madrid", "Bayern_Munich", 200.0, 94, 25, "Tersedia"}
	A[2] = pemain{3, "Pedri", "CM", "Barcelona", "-", 120.0, 88, 22, "Tersedia"}
	A[3] = pemain{4, "Jude_Bellingham", "CAM", "Real_Madrid", "Man_City", 150.0, 91, 21, "Proses_Transfer"}

	*n = 4
}

func tambahPemain(A *arrPemain, n *int) {

	var pilStatus int

	fmt.Println()
	fmt.Println("=== Tambah Pemain ===")

	fmt.Print("ID Pemain    : ")
	fmt.Scan(&A[*n].id)

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

	fmt.Print("Rating       : ")
	fmt.Scan(&A[*n].rating)

	fmt.Print("Usia         : ")
	fmt.Scan(&A[*n].usia)

	fmt.Println("1=Tersedia  2=Proses_Transfer  3=Selesai")
	fmt.Print("Pilih Status : ")
	fmt.Scan(&pilStatus)

	if pilStatus == 2 {
		A[*n].status = "Proses_Transfer"
	} else if pilStatus == 3 {
		A[*n].status = "Selesai"
	} else {
		A[*n].status = "Tersedia"
	}

	*n = *n + 1

	fmt.Println("Data pemain berhasil ditambahkan")
}

func tampilPemain(A arrPemain, n int) {

	var i int

	fmt.Println()
	fmt.Println("===== DATA PEMAIN =====")

	if n == 0 {

		fmt.Println("Data kosong")

	} else {

		for i = 0; i < n; i++ {

			fmt.Println(A[i].id,
				A[i].nama,
				A[i].posisi,
				A[i].klubAsal,
				A[i].nilaiTransfer,
				A[i].rating,
				A[i].usia,
				A[i].status)
		}
	}
}

func menuTampilPemain(A *arrPemain, n int) {

	var pilih int

	fmt.Println()
	fmt.Println("=== TAMPILKAN PEMAIN ===")
	fmt.Println("1. Tampilkan Semua (Tanpa Urut)")
	fmt.Println("2. Nilai Transfer Ascending (Selection)")
	fmt.Println("3. Nilai Transfer Descending (Selection)")
	fmt.Println("4. Usia Ascending (Insertion)")
	fmt.Println("5. Usia Descending (Insertion)")

	fmt.Print("Pilih : ")
	fmt.Scan(&pilih)

	if pilih == 1 {

		tampilPemain(*A, n)

	} else if pilih == 2 {

		selectionNilaiAsc(A, n)
		tampilPemain(*A, n)

	} else if pilih == 3 {

		selectionNilaiDesc(A, n)
		tampilPemain(*A, n)

	} else if pilih == 4 {

		insertionUsiaAsc(A, n)
		tampilPemain(*A, n)

	} else if pilih == 5 {

		insertionUsiaDesc(A, n)
		tampilPemain(*A, n)

	} else {

		fmt.Println("Pilihan tidak tersedia")
	}
}

func sequentialSearchID(A arrPemain, n int, x int) int {

	var i int
	var idx int

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

func sequentialSearchNama(A arrPemain, n int, x string) int {

	var i int
	var idx int

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

func binarySearchRating(A arrPemain, n int, x int) int {

	var left, right, mid int
	var idx int

	left = 0
	right = n - 1
	idx = -1

	for left <= right {

		mid = (left + right) / 2

		if A[mid].rating == x {

			idx = mid
			left = right + 1

		} else if x < A[mid].rating {

			right = mid - 1

		} else {

			left = mid + 1
		}
	}

	return idx
}

func cariPemainNama(A arrPemain, n int) {

	var x string
	var idx int

	fmt.Print("Masukkan Nama Pemain : ")
	fmt.Scan(&x)

	idx = sequentialSearchNama(A, n, x)

	if idx == -1 {

		fmt.Println("Data tidak ditemukan")

	} else {

		fmt.Println("Data ditemukan")
		fmt.Println(A[idx].id,
			A[idx].nama,
			A[idx].posisi,
			A[idx].klubAsal,
			A[idx].nilaiTransfer,
			A[idx].status)
	}
}

func cariPemainRating(A arrPemain, n int) {

	var x int
	var idx int

	// Binary search mewajibkan data di-sort dulu (Ascending)
	selectionRatingAsc(&A, n)

	fmt.Print("Masukkan Rating Pemain : ")
	fmt.Scan(&x)

	idx = binarySearchRating(A, n, x)

	if idx == -1 {

		fmt.Println("Data tidak ditemukan")

	} else {

		fmt.Println("Data ditemukan")
		fmt.Println(A[idx].id,
			A[idx].nama,
			A[idx].posisi,
			A[idx].klubAsal,
			A[idx].nilaiTransfer,
			A[idx].rating)
	}
}

func editPemain(A *arrPemain, n int) {

	var x int
	var idx int
	var pilStatus int

	fmt.Print("Masukkan ID Pemain yang diedit : ")
	fmt.Scan(&x)

	idx = sequentialSearchID(*A, n, x)

	if idx == -1 {

		fmt.Println("Data tidak ditemukan")

	} else {

		fmt.Print("Klub Asal baru   : ")
		fmt.Scan(&A[idx].klubAsal)

		fmt.Print("Klub Tujuan baru : ")
		fmt.Scan(&A[idx].klubTujuan)

		fmt.Print("Nilai baru       : ")
		fmt.Scan(&A[idx].nilaiTransfer)

		fmt.Println("1=Tersedia  2=Proses_Transfer  3=Selesai")
		fmt.Print("Status baru      : ")
		fmt.Scan(&pilStatus)

		if pilStatus == 2 {
			A[idx].status = "Proses_Transfer"
		} else if pilStatus == 3 {
			A[idx].status = "Selesai"
		} else {
			A[idx].status = "Tersedia"
		}

		fmt.Println("Data berhasil diubah")
	}
}

func hapusPemain(A *arrPemain, n *int) {

	var x int
	var idx int
	var i int

	fmt.Print("Masukkan ID Pemain yang dihapus : ")
	fmt.Scan(&x)

	idx = sequentialSearchID(*A, *n, x)

	if idx == -1 {

		fmt.Println("Data tidak ditemukan")

	} else {

		for i = idx; i < *n-1; i++ {

			A[i] = A[i+1]
		}

		*n = *n - 1

		fmt.Println("Data berhasil dihapus")
	}
}

// === KUMPULAN FUNGSI SORTING ===

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

func statistik(A arrPemain, n int) {

	var i int
	var maxN float64
	var iTermahal int
	var jT, jP, jS int

	if n == 0 {

		fmt.Println("Data kosong")

	} else {

		maxN = A[0].nilaiTransfer
		iTermahal = 0

		for i = 0; i < n; i++ {

			if A[i].nilaiTransfer > maxN {

				maxN = A[i].nilaiTransfer
				iTermahal = i
			}

			if A[i].status == "Tersedia" {
				jT = jT + 1
			} else if A[i].status == "Proses_Transfer" {
				jP = jP + 1
			} else {
				jS = jS + 1
			}
		}

		fmt.Println()
		fmt.Println("===== STATISTIK TRANSFER =====")
		fmt.Println("Total Pemain :", n)
		fmt.Println("Termahal     :", A[iTermahal].nama, "(", maxN, ")")
		fmt.Println("Tersedia     :", jT)
		fmt.Println("Proses       :", jP)
		fmt.Println("Selesai      :", jS)
	}
}