// WOI INI YANG DIPAKE YA YANG TERBARU

/*
Nama Anggota :
 1. Thania Evelina Munte (103012500239)
 2. Grasfi Imanuela Cahyaningtyas (103012500301)

Judul Topik : Aplikasi Manajemen Kendaraan

Deskripsi Program :
Program ini merupakan aplikasi sederhana manajemen kendaraan berdasarkan untuk mengelola data kendaraan dan pemeliharaan
teknis secara berkala. Data utama yang digunakan adalah data kendaraan, data pemilik, dan data riwayat servis. Dimana
pengguna aplikasi adalah admin bengkel atau manajer operasional armada.

Masalah yang dihadapi: Setiap pertemuan ada perubahan program

Fitur :

 1. Manajemen Kendaraan:
    -Menampilkan daftar data kendaraan
    -Menambah Data Kendaraan & Data Pemilik
    -Mencari Data Kendaraan berdaasarkan Plat, Merk, Model, Tahun
    -Mengupdate Data Kendaraan
    -Menghapus Data Kendaraan
    -Menyortir Data Kendaraan berdasarkan Plat, Merk, Model,Tahun, Warna secara ascending maupun descending

 2. Tambah Riwayat Servis Kendaraan
    -Tambah riwayat servis kendaraan yang sudah dilakukan

 3. Riwayat Servis
    Menampilkan riwayat servis kendaraan

 4. Statistik Servis Kendaraan
    Menampilkan statistik jumlah kendaraan yang diservis perbulan dan kategoori kerusakan yang paling sering
    muncul.
*/
package main

import "fmt"

const NMAX int = 999 // Konstanta untuk jumlah maksimal array

type kendaraan struct { // Struct menyimpan data kendaraan
	idPemilik int
	plat      string
	merk      string
	model     string
	warna     string
	tahun     int
}
type tabKendaraan [NMAX]kendaraan // Tipe data array untuk menyimpan data kendaraan

type servis struct { // Struct untuk menyimpan data servis kendaraan
	idServis         int
	plat             string
	merk, warna      string
	model            string
	tahun            int
	day, month, year int
	jenisServis      string
	namaPemilik      string
	no_telp          string
	keterangan       string
	bulan            int
}
type tabServis [NMAX]servis // Tipe data array untuk menyimpan data booking dan riwayat servis kendaraan

type pemilik struct { //Struct untuk menyimpan data pemilik kendaraan
	namaPemilik string
	no_telp     string
	idPemilik   int
}
type tabPemilik [NMAX]pemilik // Tipe data array untuk menyimpan data pemilik kendaraan

func main() {
	var kendaraan tabKendaraan //Data kendaraan
	var servis tabServis       // Data servis
	var pemilik tabPemilik     // Data pemilik kendaraan
	var nk, ns, np int         //nk : jumlah data kendaraan, ns : jumlah data servis, np : jumlah data pemilik
	var pilih int              // Input integer

	for {
		menu_utama()
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			manajemenKendaraan(&kendaraan, &nk, &pemilik, &np)
		case 2:
			tambahRiwayatServis(kendaraan, pemilik, &servis, nk, np, &ns)
		case 3:
			riwayatServis(servis, ns)
		case 4:
			statistikServis(servis, ns)
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini")
			return
		}
	}

}
func menu_utama() {
	/*
	   Menampilkan pilihan menu utama aplikasi manajemen kendaraan
	*/
	fmt.Println("==============================================================")
	fmt.Printf("|%35s%25s|\n", "AUTOCARE 0.5", "")
	fmt.Printf("|%44s%16s|\n", "APLIKASI MANAJEMEN KENDARAAN", "")
	fmt.Println("==============================================================")
	fmt.Printf("| %-58s |\n", "[1] Manajemen Kendaraan")
	fmt.Printf("| %-58s |\n", "[2] Tambah Riwayat Servis")
	fmt.Printf("| %-58s |\n", "[3] Riwayat Servis")
	fmt.Printf("| %-58s |\n", "[4] Statistik Servis")
	fmt.Printf("| %-58s |\n", "[0] Exit")
	fmt.Println("==============================================================")
	fmt.Print("Pilih [1/2/3/4/0]? ")
}

func manajemenKendaraan(kendaraan *tabKendaraan, n *int, pemilik *tabPemilik, np *int) {
	/*
	   IS 	  : Terdefinisi array kendaraan dengan jumlah data n dan array pemilik dengan jumlah data np.
	   Proses : Menampilkan menu untuk melihat, menambah, mencari,mengupdate, menghapus, dan mengurutkan data kendaraan.
	   FS 	  : Data kendaraan diproses sesuai pilihan pengguna atau keluar dari menu
	*/
	var pilih, idx int
	var x string
	var sorting int
	var ascending bool
	var tahun int

	for {
		fmt.Println()
		fmt.Println("================================")
		fmt.Printf("| %-28s |\n", "MANAJEMEN KENDARAAN")
		fmt.Println("================================")
		fmt.Printf("| %-28s |\n", "[1] Daftar Kendaraan")
		fmt.Printf("| %-28s |\n", "[2] Tambah Kendaraan")
		fmt.Printf("| %-28s |\n", "[3] Cari Kendaraan")
		fmt.Printf("| %-28s |\n", "[4] Update Kendaraan")
		fmt.Printf("| %-28s |\n", "[5] Hapus Kendaraan")
		fmt.Printf("| %-28s |\n", "[6] Sorting Kendaraan")
		fmt.Printf("| %-28s |\n", "[0] Exit")
		fmt.Println("================================")
		fmt.Print("Pilih [1/2/3/4/5/6/0]? ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			daftarKendaraan(*kendaraan, *n, *pemilik, *np) // tambah pemilik gaa
		case 2:
			tambahKendaraan(kendaraan, n, pemilik, np)

		case 3:
			fmt.Println("[1] Search berdasarkan plat kendaraan")
			fmt.Println("[2] Search berdasarkan merk kendaraan")
			fmt.Println("[3] Search berdasarkan model kendaraan")
			fmt.Println("[4] Search berdasarkan tahun kendaraan")
			fmt.Print("Pilih [1/2/3/4]? ")
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				fmt.Print("Plat kendaraan yang dicari  : ")
				fmt.Scan(&x)
				insertionSortPlat(kendaraan, *n, true)
				idx = binarySearchPlat(*kendaraan, *n, x)

				if idx != -1 {

					fmt.Printf("\nKendaraan dengan plat %s ditemukan\n", x)
					fmt.Println("========================================================")
					fmt.Printf("| %-10s | %-12s | %-12s | %-9s |\n",
						"Plat", "Merk", "Model", "Warna")
					fmt.Println("========================================================")

					fmt.Printf("| %-10s | %-12s | %-12s | %-9s |\n", kendaraan[idx].plat, kendaraan[idx].merk, kendaraan[idx].model, kendaraan[idx].warna)
					fmt.Println("========================================================")
				} else {
					fmt.Printf("Kendaraan dengan plat %s tidak ditemukan", x)
				}
			case 2:
				fmt.Print("Merk yang dicari  : ")
				fmt.Scan(&x)

				sequentialSearchMerk(*kendaraan, *n, x)
			case 3:
				fmt.Print("Model yang dicari  : ")
				fmt.Scan(&x)
				sequentialSearchModel(*kendaraan, *n, x)
			case 4:
				fmt.Print("Tahun yang dicari  : ")
				fmt.Scan(&tahun)
				sequentialSearchTahun(*kendaraan, *n, tahun)
			}
		case 4:
			fmt.Print("Plat kendaraan yang diupdate  : ")
			fmt.Scan(&x)
			updateKendaraan(kendaraan, n, x)
		case 5:
			fmt.Print("Plat kendaraan yang ingin dihapus  : ")
			fmt.Scan(&x)
			hapusKendaraan(kendaraan, n, x)
		case 6:
			fmt.Println("[1] Ascending")
			fmt.Println("[2] Descending")
			fmt.Print("Pilih [1/2]? ")
			fmt.Scan(&sorting)
			if sorting == 1 {
				ascending = true
			} else {
				ascending = false
			}

			fmt.Println("[1] Sorting berdasarkan plat kendaraan")
			fmt.Println("[2] Sorting berdasarkan merk kendaraan")
			fmt.Println("[3] Sorting berdasarkan model kendaraan")
			fmt.Println("[4] Sorting berdasarkan tahun kendaraan")
			fmt.Println("[5] Sorting berdasarkan warna kendaraan")
			fmt.Print("Pilih [1/2/3/4/5]? ")
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				insertionSortPlat(kendaraan, *n, ascending)
				daftarKendaraan(*kendaraan, *n, *pemilik, *np)
			case 2:
				selectionSortKendaraan(kendaraan, *n, ascending, "merk")
				daftarKendaraan(*kendaraan, *n, *pemilik, *np)
			case 3:
				selectionSortKendaraan(kendaraan, *n, ascending, "nama")
				daftarKendaraan(*kendaraan, *n, *pemilik, *np)
			case 4:
				selectionSortKendaraan(kendaraan, *n, ascending, "tahun")
				daftarKendaraan(*kendaraan, *n, *pemilik, *np)
			case 5:
				selectionSortKendaraan(kendaraan, *n, ascending, "warna")
				daftarKendaraan(*kendaraan, *n, *pemilik, *np)
			}

		case 0:
			return
		default:
			fmt.Println("Error")
		}
	}
}

func daftarKendaraan(kendaraan tabKendaraan, n int, pemilik tabPemilik, np int) {
	/*
	   IS		: Terdefinisi array kendaraan dengan jumlah data n dan array pemilik dengan jumlah data np.
	   Proses	: Menampilkan daftar kendaraan dengan data pemiliknya
	   FS		: Daftar kendaraan ditampilkan dilayar
	*/
	var i, idxPemilik int

	if n == 0 {
		fmt.Println("Belum ada data yang ditambahkan")
	} else {

		fmt.Println("+------------+----------------+----------------+----------------+----------------+----------------+---------+--------------+")
		fmt.Printf("| %-10s | %-14s | %-14s | %-14s | %-14s | %-14s | %-7s | %-12s |\n",
			"ID Pemilik", "Nama Pemilik", "No Telepon", "Plat", "Merk", "Model", "Tahun", "Warna")
		fmt.Println("+------------+----------------+----------------+----------------+----------------+----------------+---------+--------------+")

		for i = 0; i < n; i++ {
			idxPemilik = dataPemilik(pemilik, np, kendaraan[i].idPemilik)

			if idxPemilik != -1 {
				fmt.Printf("| %-10d | %-14s | %-14s | %-14s | %-14s | %-14s | %-7d | %-12s |\n",
					kendaraan[i].idPemilik,
					pemilik[idxPemilik].namaPemilik,
					pemilik[idxPemilik].no_telp,
					kendaraan[i].plat,
					kendaraan[i].merk,
					kendaraan[i].model,
					kendaraan[i].tahun,
					kendaraan[i].warna)
			}
		}
		fmt.Println("+------------+----------------+----------------+----------------+----------------+----------------+---------+--------------+")
	}
}

func dataPemilik(A tabPemilik, n int, id int) int {
	/*
		Mengembalikan indeks pemilik berdasarkan id yang dicari atau -1 jika data tidak ditemukan.
	*/
	var i int
	for i = 0; i < n; i++ {
		if A[i].idPemilik == id {
			return i
		}
	}
	return -1
}

func tambahKendaraan(kendaraan *tabKendaraan, n *int, pemilik *tabPemilik, np *int) {
	/*
	   IS		: Terdefinisi array kendaraan dengan jumlah data n dan array pemilik dengan jumlah data np.
	   Proses	: Menambahkan data kendaraan dan data pemilik jika belum terdaftar
	   FS		: Data kendaraan tersimpan dan jumlah data bertambah
	*/
	var idPemilik int
	var idx int
	var tambah string
	var valid bool

	tambah = "Ya"
	for tambah == "Ya" && *n < NMAX {
		fmt.Print("Masukkan ID Pemilik: ")
		fmt.Scan(&idPemilik)
		idx = dataPemilik(*pemilik, *np, idPemilik)
		if idx == -1 {
			fmt.Printf("Pemilik tidak ditemukan\n")
			fmt.Println()
			fmt.Print("Tambahkan data Pemilik \n")
			(*pemilik)[*np].idPemilik = idPemilik
			fmt.Print("Nama Pemilik: ")
			fmt.Scan(&(*pemilik)[*np].namaPemilik)
			valid = false
			for !valid {
				fmt.Print("Telepon Pemilik (12 digit): ")
				fmt.Scan(&(*pemilik)[*np].no_telp)

				if len((*pemilik)[*np].no_telp) == 12 {
					valid = true
				} else {
					fmt.Println("Nomor telepon tidak valid! Harus 12 digit.")
				}
			}

			idx = *np
			*np = *np + 1
		} else {
			fmt.Printf("Pemilik dengan ID %d ditemukan: %s\n", idPemilik, (*pemilik)[idx].namaPemilik)
		}

		(*kendaraan)[*n].idPemilik = (*pemilik)[idx].idPemilik

		fmt.Println("Masukkan data kendaraan")

		fmt.Printf("Plat %4s ", ":")
		fmt.Scan(&kendaraan[*n].plat)

		fmt.Printf("Merk %4s ", ":")
		fmt.Scan(&kendaraan[*n].merk)

		fmt.Printf("Model %3s ", ":")
		fmt.Scan(&kendaraan[*n].model)

		fmt.Printf("Tahun %3s ", ":")
		fmt.Scan(&kendaraan[*n].tahun)

		fmt.Printf("Warna %3s ", ":")
		fmt.Scan(&kendaraan[*n].warna)

		*n = *n + 1

		fmt.Println("Data Kendaraan berhasil ditambahkan")
		fmt.Println("Apakah Anda Ingin Menambahkan Kendaraan Lagi?")
		fmt.Print("Pilih [Ya/Tidak] : ")
		if *n < NMAX {
			fmt.Scan(&tambah)
		} else {
			fmt.Println("Anda sudah tidak bisa menambahkan kendaraan lagi")
		}
	}
}

func cariKendaraan(kendaraan tabKendaraan, n int, x string) int {
	/*
		Mencari kendaraan menggunakan sequential search, digunakan pada fungsi booking, hapus, edit, dll
	*/
	var idx, i int
	idx = -1
	i = 0
	for i < n && idx == -1 {
		if kendaraan[i].plat == x {
			idx = i
		}
		i++
	}
	return idx
}

func sequentialSearchMerk(kendaraan tabKendaraan, n int, x string) {
	/*
	   IS		: Terdefinisi data kendaraan sebanyak n data dan merk x yang dicari.
	   Proses	: Mencari kendaraan berdasarkan merk menggunakan sequential search
	   FS		: Menampilkan informasi kendaraan dengan merk yang dicari
	*/
	var i int
	var ditemukan bool
	fmt.Println("========================================================")
	fmt.Printf("| %-8s | %-12s | %-12s | %-10s |\n", "ID", "MERK", "MODEL", "WARNA")
	fmt.Println("========================================================")
	for i = 0; i < n; i++ {
		if kendaraan[i].merk == x {
			fmt.Printf("| %-8s | %-12s | %-12s | %-10s |\n", kendaraan[i].plat, kendaraan[i].merk,
				kendaraan[i].model, kendaraan[i].warna)
			ditemukan = true
		}
	}
	fmt.Println("========================================================")
	if !ditemukan {
		fmt.Println("Kendaraan tidak ditemukan")
	}
}

func sequentialSearchModel(kendaraan tabKendaraan, n int, x string) {
	/*
	    IS		: Terdefinisi data kendaraan sebanyak n data dan model x yang dicari.
	   Proses	: Mencari kendaraan berdasarkan model menggunakan sequential search
	   FS		: Menampilkan informasi kendaraan dengan model yang dicari
	*/
	var i int
	var ditemukan bool
	fmt.Println("========================================================")
	fmt.Printf("| %-8s | %-12s | %-12s | %-10s |\n", "ID", "MERK", "MODEL", "WARNA")
	fmt.Println("========================================================")
	for i = 0; i < n; i++ {
		if kendaraan[i].model == x {
			fmt.Printf("| %-8s | %-12s | %-12s | %-10s |\n", kendaraan[i].plat, kendaraan[i].merk,
				kendaraan[i].model, kendaraan[i].warna)
			ditemukan = true
		}
	}
	fmt.Println("========================================================")
	if !ditemukan {
		fmt.Println("Kendaraan tidak ditemukan")
	}

}

func sequentialSearchTahun(kendaraan tabKendaraan, n int, x int) {
	/*
	   IS		: Terdefinisi data kendaraan sebanyak n data dan tahun x yang dicari.
	   Proses	: Mencari kendaraan berdasarkan tahun menggunakan sequential search
	   FS		: Menampilkan informasi kendaraan dengan tahun yang dicari
	*/
	var i int
	var ditemukan bool
	fmt.Println("========================================================")
	fmt.Printf("| %-8s | %-12s | %-12s | %-10s |\n", "ID", "MERK", "MODEL", "WARNA")
	fmt.Println("========================================================")
	for i = 0; i < n; i++ {
		if kendaraan[i].tahun == x {
			fmt.Printf("| %-8s | %-12s | %-12s | %-10s |\n", kendaraan[i].plat, kendaraan[i].merk,
				kendaraan[i].model, kendaraan[i].warna)
			ditemukan = true
		}
	}
	fmt.Println("========================================================")
	if !ditemukan {
		fmt.Println("Kendaraan tidak ditemukan")
	}
}

func binarySearchPlat(kendaraan tabKendaraan, n int, x string) int {
	/*
		Mengembalikan indeks data kendaraan berdasarkan plat yang dicari mmenggunakan binary search
	*/
	var left, right, mid, idx int

	idx = -1
	left = 0
	right = n - 1
	for left <= right {
		mid = (left + right) / 2

		if kendaraan[mid].plat == x {
			return mid
		} else if kendaraan[mid].plat < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return idx
}

func hapusKendaraan(kendaraan *tabKendaraan, n *int, x string) { // Fungsi untuk menghapus data kendaraan berdasarkan plat
	/*
	   IS		: Terdefinisi data kendaraan sebanyak n data dan plat kendaraan yang akan dihapus
	   Proses	: Menghapus data kendaraan berdasarkan plat yang dicari
	   FS		: Data kendaraan berhasil dihapus atau tidak ditemukan
	*/
	var found, i int
	found = cariKendaraan(*kendaraan, *n, x)
	if found == -1 {
		fmt.Println("Kendaraan tidak ditemukan")
	} else {
		i = found
		for i <= *n-2 {
			(*kendaraan)[i] = (*kendaraan)[i+1]
			i++
		}
		*n = *n - 1
		fmt.Printf("Kendaraan dengan plat %s berhasil dihapus", x)
	}
}

func updateKendaraan(kendaraan *tabKendaraan, n *int, x string) { // Fungsi untuk mengupdate data kendaraan berdasarkan plat
	/*
	   IS		: Terdefinisi data kendaraan sebanyak n data dan plat kendaraan yang akan diubah
	   Proses	: Mengupdate data kendaraan berdasarkan plat yang dicari
	   FS		: Data kendaraan berhasil diupdate atau tidak ditemukan
	*/
	var idx int

	idx = cariKendaraan(*kendaraan, *n, x)

	if idx != -1 {
		fmt.Println("Masukkan Data Baru")
		fmt.Printf("Plat %4s", ":")
		fmt.Scan(&(*kendaraan)[idx].plat)
		fmt.Printf("Warna %3s", ":")
		fmt.Scan(&(*kendaraan)[idx].warna)
		fmt.Println("Data kendaraan berhasil diupdate")
	} else {
		fmt.Println("Kendaraan tidak ditemukan")
	}
}

func min(kendaraan tabKendaraan, n, i int, kategori string) int {
	/*
	   Menghasilkan indeks data kendaraan dengan nilai terkecil berdasarkan kategori yang dipilih.
	*/
	var min, j int
	min = i
	for j = i + 1; j < n; j++ {

		switch kategori {
		case "plat":
			if kendaraan[min].plat > kendaraan[j].plat {
				min = j
			}
		case "merk":
			if kendaraan[min].merk > kendaraan[j].merk {
				min = j
			}
		case "model":
			if kendaraan[min].model > kendaraan[j].model {
				min = j
			}
		case "tahun":
			if kendaraan[min].tahun > kendaraan[j].tahun {
				min = j
			}
		default:
			if kendaraan[min].warna > kendaraan[j].warna {
				min = j
			}
		}
	}
	return min
}

func max(kendaraan tabKendaraan, n, i int, kategori string) int {
	/*
	  Menghasilkan indeks data kendaraan dengan nilai terkecil berdasarkan kategori yang dipilih.
	*/
	var max, j int
	max = i
	for j = i + 1; j < n; j++ {

		switch kategori {
		case "plat":
			if kendaraan[max].plat < kendaraan[j].plat {
				max = j
			}
		case "merk":
			if kendaraan[max].merk < kendaraan[j].merk {
				max = j
			}
		case "model":
			if kendaraan[max].model < kendaraan[j].model {
				max = j
			}
		case "tahun":
			if kendaraan[max].tahun < kendaraan[j].tahun {
				max = j
			}
		case "warna":
			if kendaraan[max].warna < kendaraan[j].warna {
				max = j
			}
		}
	}
	return max
}

func selectionSortKendaraan(k *tabKendaraan, n int, Ascending bool, kategori string) {
	/*
	   IS		: Terdefinisi data kendaraan sebanyak n data
	   Proses	: Mengurutkan data kendaraan menggunakan selection sort berdasarkan kategori yang dipilih
	   FS		: Data kendaraan terurut secara ascending atau descending
	*/
	var temp kendaraan
	var idx int
	var i int

	for i = 0; i < n; i++ {
		if Ascending {
			idx = min(*k, n, i, kategori)
		} else {
			idx = max(*k, n, i, kategori)
		}
		temp = (*k)[i]
		(*k)[i] = (*k)[idx]
		(*k)[idx] = temp
	}
}

func insertionSortPlat(k *tabKendaraan, n int, Ascending bool) {
	/*
	   IS		: Terdefinisi data kendaraan sebanyak n data
	   Proses   : Mengurutkan data kendaraan berdasarkan plat menggunakan insertion sort berdasarkan kategori yang dipilih
	   FS		: Data kendaraan terurut secara ascending atau descending berdasarkan plat
	*/
	var temp kendaraan
	var pass, i int

	pass = 1
	for pass <= n-1 {
		i = pass
		temp = (*k)[pass]
		if Ascending {
			for i > 0 && (*k)[i-1].plat > temp.plat {
				(*k)[i] = (*k)[i-1]
				i--
			}
		} else {
			for i > 0 && (*k)[i-1].plat < temp.plat {
				(*k)[i] = (*k)[i-1]
				i--
			}
		}
		(*k)[i] = temp
		pass = pass + 1
	}
}

func servisKendaraanKM(jenisServis *string, keterangan *string) {
	/*
	   IS		: Jenis dan keterangan servis kendaraan belum diketahui.
	   Proses	: Input kilometer kendaraan kemudian menentukan kategori servis berdasarkan kilometer kendaraan
	   FS		: Jenis dan keterangan servis kendaraan telah diketahui.
	*/
	var km float64
	fmt.Printf("Kilometer Kendaraan %9s ", ":")
	fmt.Scan(&km)

	if km >= 5000 && km <= 10000 {
		*jenisServis = "Servis Ringan"
		*keterangan = "Oli mesin, filter oli"
	} else if km > 10000 && km <= 20000 {
		*jenisServis = "Servis Berkala"
		*keterangan = "Throttle body, aki"
	} else if km > 20000 && km <= 40000 {
		*jenisServis = "Servis Menengah"
		*keterangan = "Rem, filter udara"
	} else if km > 40000 && km <= 50000 {
		*jenisServis = "Servis Besar"
		*keterangan = "Busi, oli transmisi"
	} else {
		*jenisServis = "Servis Lanjutan"
		*keterangan = "Timing belt, suspensi"
	}
}

func servisKendaraanKerusakan(jenisServis, keterangan *string) {
	/*
	   IS		: Jenis dan keterangan servis kendaraan belum diketahui.
	   Proses	: Input kerusakan kendaraan kemudian menentukan kategori servis berdasarkan kerusakan kendaraan
	   FS		: Jenis dan keterangan servis kendaraan telah diketahui.
	*/
	var kerusakan string
	fmt.Println("JENIS KERUSAKAN")
	fmt.Println("1. Mesin")
	fmt.Println("2. Rem")
	fmt.Println("3. Ban")
	fmt.Println("4. Transmisi")
	fmt.Println("5. Servis Lanjutan")
	fmt.Printf("Jenis Kerusakan %10s", ":")
	fmt.Scan(&kerusakan)

	switch kerusakan {
	case "Mesin":
		*jenisServis = "Servis Mesin"
		*keterangan = "Tune Up, cek busi dan aki"

	case "Rem":
		*jenisServis = "Servis Rem"
		*keterangan = "Ganti kampas rem, cek minyak rem"

	case "Ban":
		*jenisServis = "Servis Ban"
		*keterangan = "Tambal atau ganti ban"

	case "Transmisi":
		*jenisServis = "Servis Transmisi"
		*keterangan = "Ganti oli transmisi"

	default:
		*jenisServis = "Servis Lanjutan"
		*keterangan = "Perlu pemeriksaan lebih lanjut"
	}
}

func tambahRiwayatServis(kendaraan tabKendaraan, pemilik tabPemilik, servis *tabServis, nk int, np int, ns *int) {
	/*
	   IS		: Terdefinisi array kendaraan dengan jumlah data nk, array pemilik dengan jumlah data np, dan array servis dengan jumlah data ns.
	   Proses	: Melakukan tambah riwayat servis kendaraan
	   FS		: Data Riwayat servis baru tersimpan
	*/
	var x string
	var idx, pilih int

	fmt.Println("Masukan data kendaraan")
	fmt.Printf("Plat %24s ", ":")
	fmt.Scan(&x)
	idx = cariKendaraan(kendaraan, nk, x)

	if idx != -1 {
		fmt.Println("PILIH SERVIS KENDARAAN")
		fmt.Println("[1] Servis berdasarkan kilometer")
		fmt.Println("[2] Servis berdasarkan kerusakan")
		fmt.Print("Pilih [1/2]? ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			servisKendaraanKM(&(*servis)[*ns].jenisServis, &(*servis)[*ns].keterangan)
		} else if pilih == 2 {
			servisKendaraanKerusakan(&(*servis)[*ns].jenisServis, &(*servis)[*ns].keterangan)
		}
		(*servis)[*ns].idServis = *ns + 1
		(*servis)[*ns].plat = kendaraan[idx].plat
		(*servis)[*ns].merk = kendaraan[idx].merk
		(*servis)[*ns].model = kendaraan[idx].model
		(*servis)[*ns].tahun = kendaraan[idx].tahun
		(*servis)[*ns].warna = kendaraan[idx].warna

		(*servis)[*ns].namaPemilik = pemilik[idx].namaPemilik
		(*servis)[*ns].no_telp = pemilik[idx].no_telp
		fmt.Printf("Tanggal Servis [dd mm yyyy] %1s ", ":")
		fmt.Scan(&(*servis)[*ns].day, &(*servis)[*ns].month, &(*servis)[*ns].year)

		fmt.Println("Tambah riwayat servis berhasil dilakukan")
		*ns = *ns + 1
	} else {
		fmt.Println("Kendaraan tidak ditemukan")
	}
}

func riwayatServis(servis tabServis, ns int) {
	/*
	   IS      : Terdefinisi array servis dengan jumlah data ns
	   Proses  : Menampilkan seluruh data riwayat servis kendaraan
	   FS      : Riwayat servis ditampilkan di layar
	*/
	var i int

	if ns == 0 {
		fmt.Println("Belum ada riwayat servis")
	} else {
		fmt.Println("\n=============================================== RIWAYAT SERVIS ===============================================")
		fmt.Println("+------+----------------+--------------+------------+------------+------------+------------+------------------+")
		fmt.Printf("| %-4s | %-14s | %-12s | %-10s | %-10s | %-10s | %-10s | %-16s |\n",
			"ID", "Nama Pemilik", "No Telepon", "Plat", "Merk", "Model", "Tanggal", "Jenis Servis")
		fmt.Println("+------+----------------+--------------+------------+------------+------------+------------+------------------+")

		for i = 0; i < ns; i++ {
			fmt.Printf("| %-4d | %-14s | %-12s | %-10s | %-10s | %-10s | %02d-%02d-%04d | %-16s |\n",
				servis[i].idServis, servis[i].namaPemilik, servis[i].no_telp, servis[i].plat, servis[i].merk,
				servis[i].model, servis[i].day, servis[i].month, servis[i].year, servis[i].jenisServis)
		}

		fmt.Println("+------+----------------+--------------+------------+------------+------------+------------+------------------+")
	}
}
func statistikServis(servis tabServis, ns int) {
	/*
	   IS : tersedia data servis sebanyak ns
	   FS : menampilkan statistik jumlah servis per bulan
	        pada tahun tertentu dan kategori servis
	        yang paling sering muncul
	*/

	var jumlahBulan [12]int
	var kategori [12][9]int

	var namaBulan = [12]string{
		"Januari", "Februari", "Maret", "April",
		"Mei", "Juni", "Juli", "Agustus",
		"September", "Oktober", "November", "Dessember",
	}
	var namaServis = [9]string{
		"Ringan", "Berkala", "Menengah", "Besar",
		"Lanjutan", "Mesin", "Rem", "Ban", "Transmisi",
	}

	var tahunCari int
	var i, bulan int
	var ada bool

	fmt.Print("Masukkan tahun yang ingin ditampilkan: ")
	fmt.Scan(&tahunCari)

	for i = 0; i < ns; i++ {
		if servis[i].year == tahunCari {
			ada = true
			bulan = servis[i].month - 1
			if bulan >= 0 && bulan < 12 {
				jumlahBulan[bulan]++
				switch servis[i].jenisServis {
				case "Servis Ringan":
					kategori[bulan][0]++
				case "Servis Berkala":
					kategori[bulan][1]++
				case "Servis Menengah":
					kategori[bulan][2]++
				case "Servis Besar":
					kategori[bulan][3]++
				case "Servis Lanjutan":
					kategori[bulan][4]++
				case "Servis Mesin":
					kategori[bulan][5]++
				case "Servis Rem":
					kategori[bulan][6]++
				case "Servis Ban":
					kategori[bulan][7]++
				case "Servis Transmisi":
					kategori[bulan][8]++
				}
			}
		}
	}

	if !ada {
		fmt.Printf("Tidak ada data servis pada tahun %d\n", tahunCari)
		return
	}

	fmt.Println("\n==============================================================")
	fmt.Printf("STATISTIK SERVIS TAHUN %d\n", tahunCari)
	fmt.Println("==============================================================")
	fmt.Printf("%-12s %-15s %-20s\n", "Bulan", "Jumlah Servis", "Kategori Terbanyak")
	fmt.Println("--------------------------------------------------------------")
	maxBulan := 0
	for i = 0; i < 12; i++ {
		if jumlahBulan[i] == 0 {
			fmt.Printf("%-12s %-15d %-20s\n", namaBulan[i], 0, "-")
		} else {
			maxKategori := 0
			for j := 1; j < 9; j++ {
				if kategori[i][j] > kategori[i][maxKategori] {
					maxKategori = j
				}
			}
			fmt.Printf("%-12s %-15d %-20s\n", namaBulan[i], jumlahBulan[i], namaServis[maxKategori])
		}
		if jumlahBulan[i] > jumlahBulan[maxBulan] {
			maxBulan = i
		}
	}
}
