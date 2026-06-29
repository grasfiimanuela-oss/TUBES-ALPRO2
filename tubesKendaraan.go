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
    -Mengupdate Data Kendaraan
    -Menghapus Data Kendaraan
	-Mencari Data Kendaraan berdaasarkan Plat, Merk, Model, Tahun
    -Menyortir Data Kendaraan berdasarkan Plat, Merk, Model,Tahun, Warna secara ascending maupun descending

 2. Tambah Riwayat Servis Kendaraan
    -Tambah riwayat servis kendaraan yang sudah dilakukan

 3. Riwayat Servis
    Menampilkan riwayat servis kendaraan
*/
package main

import "fmt"

const NMAX int = 99 // Konstanta untuk jumlah maksimal array

type kendaraan struct { // Struct menyimpan informasi data kendaraan
	idPemilik int // untuk mengetahui pemilik dari kendaraan
	plat      string
	merk      string
	model     string
	warna     string
	tahun     int // tahun produksi kendaraan
}
type tabKendaraan [NMAX]kendaraan // Tipe data array untuk menyimpan data kendaraan

type servis struct { // Struct untuk menyimpan informasi data servis kendaraan
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
}

type tabServis [NMAX]servis // Tipe data array untuk menyimpan data booking dan riwayat servis kendaraan

type pemilik struct { //Struct untuk menyimpan informasi data pemilik kendaraan
	namaPemilik string
	no_telp     string
	idPemilik   int
}
type tabPemilik [NMAX]pemilik // Tipe data array untuk menyimpan data pemilik kendaraan

func dataDummy(kendaraans *tabKendaraan, pemiliks *tabPemilik, nk, np *int) {
	/* IS : kendaraan, pemilik, nk, np terdefinisi
	   FS :
	   - Array pemilik berisi data awal pemilik.
	   - Array kendaraan berisi data awal kendaraan.
	   - np dan nk berisi jumlah data awal.
	*/
	// DATA PEMILIK
	(*pemiliks)[0] = pemilik{"Andi", "081234567890", 101}
	(*pemiliks)[1] = pemilik{"Budi", "082345678901", 102}
	(*pemiliks)[2] = pemilik{"Citra", "083456789012", 103}
	(*pemiliks)[3] = pemilik{"Dina", "084567890123", 104}
	(*pemiliks)[4] = pemilik{"Eko", "085678901234", 105}

	*np = 5

	// DATA KENDARAAN
	(*kendaraans)[0] = kendaraan{101, "B1234ABC", "Toyota", "Avanza", "Hitam", 2020}
	(*kendaraans)[1] = kendaraan{102, "D5678DEF", "Honda", "Brio", "Putih", 2021}
	(*kendaraans)[2] = kendaraan{103, "F9876XYZ", "Suzuki", "Ertiga", "Merah", 2019}
	(*kendaraans)[3] = kendaraan{104, "B4321AAA", "Toyota", "Innova", "Silver", 2022}
	(*kendaraans)[4] = kendaraan{105, "Z1111ZZZ", "Honda", "HRV", "Hitam", 2023}

	*nk = 5
}

func main() {
	var kendaraan tabKendaraan //Data kendaraan
	var servis tabServis       // Data servis
	var pemilik tabPemilik     // Data pemilik kendaraan
	var nk, ns, np int         //nk : jumlah data kendaraan, ns : jumlah data servis, np : jumlah data pemilik
	var pilih int              // Input integer

	dataDummy(&kendaraan, &pemilik, &nk, &np)
	pilih = -1
	for pilih != 0 {
		menu_utama()
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			manajemenKendaraan(&kendaraan, &nk, &pemilik, &np)
		case 2:
			tambahRiwayatServis(kendaraan, pemilik, &servis, nk, np, &ns)
		case 3:
			riwayatServis(servis, ns)
		case 0:
			fmt.Println()
			fmt.Println("----Terima kasih telah menggunakan aplikasi ini----")
		default:
			fmt.Println("Pilihan Tidak Valid")
		}
	}
}

func menu_utama() {
	/*
	   Menampilkan pilihan menu utama aplikasi manajemen kendaraan
	*/
	fmt.Println()
	fmt.Println("==============================================================")
	fmt.Printf("|%35s%25s|\n", "AUTOCARE 0.5", "")
	fmt.Printf("|%44s%16s|\n", "APLIKASI MANAJEMEN KENDARAAN", "")
	fmt.Println("==============================================================")
	fmt.Printf("| %-58s |\n", "[1] Manajemen Kendaraan")
	fmt.Printf("| %-58s |\n", "[2] Tambah Riwayat Servis")
	fmt.Printf("| %-58s |\n", "[3] Riwayat Servis")
	fmt.Printf("| %-58s |\n", "[0] Exit")
	fmt.Println("==============================================================")
	fmt.Print("Pilih [1/2/3/0]? ")
}

func manajemenKendaraan(kendaraan *tabKendaraan, n *int, pemilik *tabPemilik, np *int) {
	/*
	   IS 	  : Terdefinisi array kendaraan dengan jumlah data n dan array pemilik dengan jumlah data np.
	   Proses : Menampilkan menu untuk melihat, menambah, mencari,mengupdate, menghapus, dan mengurutkan data kendaraan.
	   FS 	  : Data kendaraan diproses sesuai pilihan pengguna atau keluar dari menu
	*/
	var pilih, idx int // pilih : input integer untuk memilih menu manajemen kendaraan, idx : indeks yang dicari
	var x string       // input string yang akan dicari
	var sorting int    //input integer untuk memilih mengurutkan berdasarkan ascending atau descening
	var terurut bool   // untuk sorting ascending atau descending
	var cariTahun int  // untuk mencari berdasarkan tahun produksi

	pilih = -1
	for pilih != 0 {
		fmt.Println()
		fmt.Println("================================")
		fmt.Printf("| %-28s |\n", "MANAJEMEN KENDARAAN")
		fmt.Println("================================")
		fmt.Printf("| %-28s |\n", "[1] Daftar Kendaraan")
		fmt.Printf("| %-28s |\n", "[2] Tambah Kendaraan")
		fmt.Printf("| %-28s |\n", "[3] Update Kendaraan")
		fmt.Printf("| %-28s |\n", "[4] Hapus Kendaraan")
		fmt.Printf("| %-28s |\n", "[5] Cari Kendaraan")
		fmt.Printf("| %-28s |\n", "[6] Sorting Kendaraan")
		fmt.Printf("| %-28s |\n", "[0] Exit")
		fmt.Println("================================")
		fmt.Print("Pilih [1/2/3/4/5/6/0]? ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			daftarKendaraan(*kendaraan, *n, *pemilik, *np) // memanggil fungsi daftar kendaraan
		case 2:
			tambahKendaraan(kendaraan, n, pemilik, np) // memanggil fungsi tambah kendaraan
		case 3:
			fmt.Print("Plat kendaraan yang ingin diupdate  : ")
			fmt.Scan(&x)                     // input plat kendaraan yang diupdate
			updateKendaraan(kendaraan, n, x) // memanggil fungsi update kendaraan
		case 4:
			fmt.Print("Plat kendaraan yang ingin dihapus  : ")
			fmt.Scan(&x) // input plat kendaraan yang ingin dihapus
			hapusKendaraan(kendaraan, n, x)
		case 5:
			fmt.Println("[1] Search berdasarkan plat kendaraan")
			fmt.Println("[2] Search berdasarkan merk kendaraan")
			fmt.Println("[3] Search berdasarkan model kendaraan")
			fmt.Println("[4] Search berdasarkan tahun kendaraan")
			fmt.Print("Pilih [1/2/3/4]? ")
			fmt.Scan(&pilih)
			switch pilih {
			case 1: // binary search berdasarkan plat
				fmt.Print("Plat kendaraan yang dicari  : ")
				fmt.Scan(&x)
				insertionSortPlat(kendaraan, *n, true) // mengurutkan data kendaraan sebelum dicari menggunakan binary search
				idx = binarySearchPlat(*kendaraan, *n, x)

				if idx != -1 {
					fmt.Printf("\nKendaraan dengan plat %s ditemukan\n", x)
					fmt.Println("========================================================")
					fmt.Printf("| %-10s | %-12s | %-12s | %-9s |\n", "Plat", "Merk", "Model", "Warna")
					fmt.Println("========================================================")
					fmt.Printf("| %-10s | %-12s | %-12s | %-9s |\n", kendaraan[idx].plat, kendaraan[idx].merk, kendaraan[idx].model, kendaraan[idx].warna)
					fmt.Println("========================================================")
				} else {
					fmt.Printf("Kendaraan dengan plat %s tidak ditemukan", x)
				}
			case 2: // sequential search berdasarkan merk kendaraan
				fmt.Print("Merk yang dicari  : ")
				fmt.Scan(&x) // input merk yang akan dicari
				sequentialSearchMerk(*kendaraan, *n, x)
			case 3: // sequential search berdasarkan model kendaraan
				fmt.Print("Model yang dicari  : ")
				fmt.Scan(&x) // input model yang akan dicari
				sequentialSearchModel(*kendaraan, *n, x)
			case 4: // sequential search berdasarkan tahun produksi kendaraan
				fmt.Print("Tahun yang dicari  : ")
				fmt.Scan(&cariTahun) // input tahun produksi yang akan dicari
				sequentialSearchTahun(*kendaraan, *n, cariTahun)
			}
		case 6:
			fmt.Println("[1] Ascending")
			fmt.Println("[2] Descending")
			fmt.Print("Pilih [1/2]? ")
			fmt.Scan(&sorting)
			if sorting == 1 {
				terurut = true // ascending
			} else {
				terurut = false // descending
			}

			fmt.Println("[1] Sorting berdasarkan plat kendaraan")
			fmt.Println("[2] Sorting berdasarkan merk kendaraan")
			fmt.Println("[3] Sorting berdasarkan model kendaraan")
			fmt.Println("[4] Sorting berdasarkan tahun kendaraan")
			fmt.Println("[5] Sorting berdasarkan warna kendaraan")
			fmt.Print("Pilih [1/2/3/4/5]? ")
			fmt.Scan(&pilih)
			switch pilih {
			case 1: // mengurutkan kendaraan berdasarkan plat kendaraan
				insertionSortPlat(kendaraan, *n, terurut)
				daftarKendaraan(*kendaraan, *n, *pemilik, *np) // menampilkan data kendaraan yang sudah diurutkan
			case 2: // mengurutkan kendaraan berdasarkan merk kendaraan
				selectionSortKendaraan(kendaraan, *n, terurut, "merk")
				daftarKendaraan(*kendaraan, *n, *pemilik, *np)
			case 3: // mengurutkan kendaraan berdasarkan model kendaraan
				selectionSortKendaraan(kendaraan, *n, terurut, "model")
				daftarKendaraan(*kendaraan, *n, *pemilik, *np)
			case 4: // mengurutkan kendaraan berdasarkan tahun produksi kendaraan
				selectionSortKendaraan(kendaraan, *n, terurut, "tahun")
				daftarKendaraan(*kendaraan, *n, *pemilik, *np)
			case 5: // mengurutkan kendaraan berdasarkan warna kendaraan
				selectionSortKendaraan(kendaraan, *n, terurut, "warna")
				daftarKendaraan(*kendaraan, *n, *pemilik, *np)
			}
		case 0:
			fmt.Println("Kembali ke menu utama")
		}
	}
}

func daftarKendaraan(kendaraan tabKendaraan, n int, pemilik tabPemilik, np int) {
	/*  IS		: Terdefinisi array kendaraan dengan jumlah data n dan array pemilik dengan jumlah data np.
	Proses	: Menampilkan daftar kendaraan dengan data pemiliknya
	FS		: Daftar kendaraan ditampilkan dilayar */
	var i, idxPemilik int

	if n == 0 {
		fmt.Println("Belum ada data yang ditambahkan")
	} else {

		fmt.Println("+------------+----------------+----------------+----------------+----------------+----------------+---------+--------------+")
		fmt.Printf("| %-10s | %-14s | %-14s | %-14s | %-14s | %-14s | %-7s | %-12s |\n", "ID Pemilik", "Nama Pemilik", "No Telepon", "Plat", "Merk", "Model", "Tahun", "Warna")
		fmt.Println("+------------+----------------+----------------+----------------+----------------+----------------+---------+--------------+")

		for i = 0; i < n; i++ {
			idxPemilik = dataPemilik(pemilik, np, kendaraan[i].idPemilik, 0) // untuk memeriksa pemilik dari kendaraan, dengan idPemilik mencocokan idPemilik di tab kendaraan dan idPemilik di tab pemilik

			if idxPemilik != -1 {
				fmt.Printf("| %-10d | %-14s | %-14s | %-14s | %-14s | %-14s | %-7d | %-12s |\n",
					kendaraan[i].idPemilik, pemilik[idxPemilik].namaPemilik, pemilik[idxPemilik].no_telp,
					kendaraan[i].plat, kendaraan[i].merk, kendaraan[i].model, kendaraan[i].tahun, kendaraan[i].warna)
			}
		}
		fmt.Println("+------------+----------------+----------------+----------------+----------------+----------------+---------+--------------+")
	}
}

func dataPemilik(A tabPemilik, n, id, i int) int {
	/* Mengembalikan indeks pemilik jika ditemukan, atau -1 jika tidak ditemukan.*/
	if i >= n { // jika indeks sudah melewati data terakhir, berarti id tidak ditemukan
		return -1
	}
	if A[i].idPemilik == id { // jika id pada indeks ke-i sesuai dengan id yang dicari
		return i
	}
	return dataPemilik(A, n, id, i+1) // lanjut mencari ke indeks berikutnya
}

func tambahKendaraan(kendaraan *tabKendaraan, n *int, pemilik *tabPemilik, np *int) {
	/*  IS		: Terdefinisi array kendaraan dengan jumlah data n dan array pemilik dengan jumlah data np.
	Proses	: Menambahkan data kendaraan dan data pemilik jika belum terdaftar
	FS		: Data kendaraan tersimpan dan jumlah data bertambah */
	var idPemilik int
	var idx int
	var tambah string // untuk konfirmasi menambah data kendaraan lagi atau tidak
	var valid bool    // untuk cek nomor telepon sudah valid berisi 12 atau belum

	tambah = "Ya"
	if *n >= NMAX {
		fmt.Println("Data sudah penuh, anda tidak dapat menambah data lagi")
	} else {
		for tambah == "Ya" {
			fmt.Print("Masukkan ID Pemilik: ")
			fmt.Scan(&idPemilik)
			idx = dataPemilik(*pemilik, *np, idPemilik, 0) // untuk mencari apakah idPemilik sudah terdaftar
			// jika belum terdaftar maka harus menambahkan pemilik
			if idx == -1 {
				fmt.Printf("Pemilik tidak ditemukan\n")
				fmt.Println()
				fmt.Print("Tambahkan data Pemilik \n")
				(*pemilik)[*np].idPemilik = idPemilik // jika idPemilik tidak ditemukan, maka idPemilik yang diinput akan menjadi idPemilik baru
				fmt.Print("Nama Pemilik: ")
				fmt.Scan(&(*pemilik)[*np].namaPemilik)
				valid = false
				for !valid {
					fmt.Print("Telepon Pemilik (12 digit): ")
					fmt.Scan(&(*pemilik)[*np].no_telp)

					if len((*pemilik)[*np].no_telp) == 12 { // memeriksa panjang no telepon yang diinput
						valid = true
					} else {
						fmt.Println("Nomor telepon tidak valid! Harus 12 digit.")
					}
				}

				idx = *np     // Indeks pemilik baru disimpan agar dapat digunakan untuk menghubungkan kendaraan dengan pemilik tersebut.
				*np = *np + 1 // mengupdate jumlah pemilik
			} else {
				fmt.Printf("Pemilik dengan ID %d ditemukan: %s\n", idPemilik, (*pemilik)[idx].namaPemilik) // jika idPemilik sudah terdaftar atau ditemukan
			}

			(*kendaraan)[*n].idPemilik = (*pemilik)[idx].idPemilik // menghubungkan pemilik dengan kendaraan

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

			*n = *n + 1 // mengupdate jumlah kendaraan

			fmt.Println("Data Kendaraan berhasil ditambahkan")
			fmt.Println("Apakah Anda Ingin Menambahkan Kendaraan Lagi?")
			fmt.Print("Pilih [Ya/Tidak] : ")
			if *n < NMAX {
				fmt.Scan(&tambah) // konfirmasi tambah kendaraan atau tidak selama belum melebihi NMAX
			} else {
				fmt.Println("Data sudah penuh, anda tidak dapat menambah data lagi") // menunjukkan data sudah penuh
			}
		}
	}

}

func updateKendaraan(kendaraan *tabKendaraan, n *int, x string) { // Fungsi untuk mengupdate data kendaraan berdasarkan plat
	/* IS		: Terdefinisi data kendaraan sebanyak n data dan plat kendaraan yang akan diubah
	   Proses	: Mengupdate data kendaraan berdasarkan plat yang dicari
	   FS		: Data kendaraan berhasil diupdate atau tidak ditemukan */
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
		fmt.Printf("Kendaraan dengan plat %s tidak ditemukan", x)
	}
}

func hapusKendaraan(kendaraan *tabKendaraan, n *int, x string) { // Fungsi untuk menghapus data kendaraan berdasarkan plat
	/* IS		: Terdefinisi data kendaraan sebanyak n data dan plat kendaraan yang akan dihapus
	   Proses	: Menghapus data kendaraan berdasarkan plat yang dicari
	   FS		: Data kendaraan berhasil dihapus atau tidak ditemukan */
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

func cariKendaraan(kendaraan tabKendaraan, n int, x string) int {
	/* Mencari kendaraan menggunakan sequential search, digunakan pada fungsi booking, hapus, edit, dll */
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
	/* IS		: Terdefinisi data kendaraan sebanyak n data dan merk x yang dicari.
	   Proses	: Mencari kendaraan berdasarkan merk menggunakan sequential search
	   FS		: Menampilkan informasi kendaraan dengan merk yang dicari */
	var i int
	var ditemukan bool

	fmt.Println("========================================================")
	fmt.Printf("| %-8s | %-12s | %-12s | %-10s |\n", "Plat", "Merk", "Model", "Warna")
	fmt.Println("========================================================")
	for i = 0; i < n; i++ {
		if kendaraan[i].merk == x {
			fmt.Printf("| %-8s | %-12s | %-12s | %-10s |\n", kendaraan[i].plat, kendaraan[i].merk, kendaraan[i].model, kendaraan[i].warna)
			ditemukan = true
		}
	}
	fmt.Println("========================================================")
	if !ditemukan {
		fmt.Printf("Kendaraan dengan merk %s tidak ditemukan", x)
	}
}

func sequentialSearchModel(kendaraan tabKendaraan, n int, x string) {
	/*  IS		: Terdefinisi data kendaraan sebanyak n data dan model x yang dicari.
	Proses	: Mencari kendaraan berdasarkan model menggunakan sequential search
	FS		: Menampilkan informasi kendaraan dengan model yang dicari */
	var i int
	var ditemukan bool

	fmt.Println("========================================================")
	fmt.Printf("| %-8s | %-12s | %-12s | %-10s |\n", "Plat", "Merk", "Model", "Warna")
	fmt.Println("========================================================")
	for i = 0; i < n; i++ {
		if kendaraan[i].model == x {
			fmt.Printf("| %-8s | %-12s | %-12s | %-10s |\n", kendaraan[i].plat, kendaraan[i].merk, kendaraan[i].model, kendaraan[i].warna)
			ditemukan = true
		}
	}
	fmt.Println("========================================================")
	if !ditemukan {
		fmt.Printf("Kendaraan dengan model %s tidak ditemukan", x)
	}

}

func sequentialSearchTahun(kendaraan tabKendaraan, n int, x int) {
	/* IS		: Terdefinisi data kendaraan sebanyak n data dan tahun x yang dicari.
	   Proses	: Mencari kendaraan berdasarkan tahun menggunakan sequential search
	   FS		: Menampilkan informasi kendaraan dengan tahun yang dicari */
	var i int
	var ditemukan bool
	fmt.Println("=======================================================")
	fmt.Printf("| %-8s | %-12s | %-12s | %-10s |\n", "Plat", "Merk", "Model", "Warna")
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
		fmt.Printf("Kendaraan dengan tahun %d tidak ditemukan", x)
	}
}

func binarySearchPlat(kendaraan tabKendaraan, n int, x string) int {
	/* Mengembalikan indeks data kendaraan berdasarkan plat yang dicari mmenggunakan binary search */
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

func min(kendaraan tabKendaraan, n, i int, kategori string) int { // digunakan untuk selectionSort secara ascending
	/* Mengembalikan indeks data kendaraan dengan nilai terkecil berdasarkan kategori yang dipilih. */
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
		case "warna":
			if kendaraan[min].warna > kendaraan[j].warna {
				min = j
			}
		}
	}
	return min
}

func max(kendaraan tabKendaraan, n, i int, kategori string) int { // digunakan untuk selectionSort secara descending
	/* Mengembalikan indeks data kendaraan dengan nilai terkecil berdasarkan kategori yang dipilih. */
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
	/* IS		: Terdefinisi data kendaraan sebanyak n data
	   Proses	: Mengurutkan data kendaraan menggunakan selection sort berdasarkan kategori yang dipilih
	   FS		: Data kendaraan terurut secara ascending atau descending */
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
	/* IS		: Terdefinisi data kendaraan sebanyak n data
	   Proses   : Mengurutkan data kendaraan berdasarkan plat menggunakan insertion sort berdasarkan kategori yang dipilih
	   FS		: Data kendaraan terurut secara ascending atau descending berdasarkan plat */
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
	/* IS		: Jenis dan keterangan servis kendaraan belum diketahui.
	   Proses	: Input kilometer kendaraan kemudian menentukan kategori servis berdasarkan kilometer kendaraan
	   FS		: Jenis dan keterangan servis kendaraan telah diketahui. */
	var km float64
	fmt.Printf("Kilometer Kendaraan %9s ", ":")
	fmt.Scan(&km)
	fmt.Println()

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
	/* IS		: Jenis dan keterangan servis kendaraan belum diketahui.
	   Proses	: Input kerusakan kendaraan kemudian menentukan kategori servis berdasarkan kerusakan kendaraan
	   FS		: Jenis dan keterangan servis kendaraan telah diketahui. */
	var kerusakan string
	fmt.Println("JENIS KERUSAKAN")
	fmt.Println("1. Mesin")
	fmt.Println("2. Rem")
	fmt.Println("3. Ban")
	fmt.Println("4. Transmisi")
	fmt.Println("5. Servis Lanjutan")
	fmt.Printf("Jenis Kerusakan %10s", ":")
	fmt.Scan(&kerusakan)
	fmt.Println()

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
	/* IS		: Terdefinisi array kendaraan dengan jumlah data nk, array pemilik dengan jumlah data np, dan array servis dengan jumlah data ns.
	   Proses	: Melakukan tambah riwayat servis kendaraan
	   FS		: Data Riwayat servis baru tersimpan */
	var x string
	var idx, pilih, idxpemilik int

	if *ns >= NMAX {
		fmt.Println("Data servis sudah penuh, tidak dapat menambah riwayat servis.")
	} else {
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
			fmt.Println()

			for pilih != 1 && pilih != 2 {
				fmt.Println("Pilihan tidak valid!")
				fmt.Print("Silahkan pilih ulang: ")
				fmt.Scan(&pilih)
			}

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

			idxpemilik = dataPemilik(pemilik, np, kendaraan[idx].idPemilik, 0)
			(*servis)[*ns].namaPemilik = pemilik[idxpemilik].namaPemilik // Disambungin dulu kendaraan sama pemiliknya pake binary
			(*servis)[*ns].no_telp = pemilik[idxpemilik].no_telp
			fmt.Printf("Tanggal Servis [dd mm yyyy] %1s ", ":")
			fmt.Scan(&(*servis)[*ns].day, &(*servis)[*ns].month, &(*servis)[*ns].year)

			fmt.Println()
			fmt.Println("Tambah riwayat servis berhasil dilakukan")
			fmt.Println()
			*ns = *ns + 1
		} else {
			fmt.Println("Kendaraan tidak ditemukan")
		}
	}

}
func riwayatServis(servis tabServis, ns int) {
	/* IS      : Terdefinisi array servis dengan jumlah data ns
	   Proses  : Menampilkan seluruh data riwayat servis kendaraan
	   FS      : Riwayat servis ditampilkan di layar */
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
