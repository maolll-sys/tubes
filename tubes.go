package main

import "fmt"

const NMAX int = 100

type komponen struct {
	noSeri string
	nama   string
	suhu   float64
	beban  int
	status string
}

type tabKomponen [NMAX]komponen

func inputData(T *tabKomponen, n *int) {
	fmt.Println("INPUT DATA ")
	fmt.Print("Nomor Seri : ")
	fmt.Scan(&T[*n].noSeri)
	for T[*n].noSeri != "SELESAI" {
		fmt.Print("Nama Komponen : ")
		fmt.Scan(&T[*n].nama)

		fmt.Print("Suhu : ")
		fmt.Scan(&T[*n].suhu)

		fmt.Print("Beban Kerja : ")
		fmt.Scan(&T[*n].beban)

		if T[*n].suhu > 80 {
			T[*n].status = "OVERHEAT"
		} else if T[*n].beban > 70 {
			T[*n].status = "LAG"
		} else {
			T[*n].status = "NORMAL"
		}
		fmt.Println()
		*n = *n + 1
		fmt.Print("Nomor Seri : ")
		fmt.Scan(&T[*n].noSeri)
	}
	fmt.Println("Data berhasil ditambahkan")
}

func printData(T tabKomponen, n int) {
	var i int
	fmt.Println("DATA KOMPONEN :")
	for i = 0; i < n; i++ {
		fmt.Println("----------------------")
		fmt.Println("Nomor Seri :", T[i].noSeri)
		fmt.Println("Nama       :", T[i].nama)
		fmt.Println("Suhu       :", T[i].suhu)
		fmt.Println("Beban      :", T[i].beban)
		fmt.Println("Status     :", T[i].status)
	}
}

func cari(T tabKomponen, n int, x string) int {
	var i, found int
	found = -1
	i = 0
	for i < n && found == -1 {
		if T[i].nama == x || T[i].status == x {
			found = i
		}
		i = i + 1
	}
	return found
}

func ubah(T *tabKomponen, n int, x string) {
	var found int
	found = cari(*T, n, x)

	if found == -1 {
		fmt.Println("Data tidak ditemukan")
	} else {
		fmt.Print("Nama Baru : ")
		fmt.Scan(&T[found].nama)

		fmt.Print("Suhu Baru : ")
		fmt.Scan(&T[found].suhu)

		fmt.Print("Beban Baru : ")
		fmt.Scan(&T[found].beban)

		if T[found].suhu > 80 {
			T[found].status = "OVERHEAT"
		} else if T[found].beban > 70 {
			T[found].status = "LAG"
		} else {
			T[found].status = "NORMAL"
		}
		fmt.Println("Data berhasil diubah")
	}
}

func hapus(T *tabKomponen, n *int, x string) {
	var found, i int
	found = cari(*T, *n, x)
	if found == -1 {
		fmt.Println("Data tidak ditemukan")
	} else {
		for i = found; i < *n-1; i++ {
			T[i] = T[i+1]
		}
		*n = *n - 1
		fmt.Println("Data berhasil dihapus")
	}
}

func selectionSort(T *tabKomponen, n int) {
	var pass, idx, i int
	var temp komponen

	for pass = 1; pass < n; pass++ {
		idx = pass - 1
		for i = pass; i < n; i++ {
			if T[i].noSeri < T[idx].noSeri {
				idx = i
			}
		}
		temp = T[pass-1]
		T[pass-1] = T[idx]
		T[idx] = temp
	}
	fmt.Println("Data berhasil diurutkan (Selection Sort)")
}

func insertionSort(T *tabKomponen, n int) {
	var pass, i int
	var temp komponen

	for pass = 1; pass < n; pass++ {
		temp = T[pass]
		i = pass
		for i > 0 && temp.noSeri < T[i-1].noSeri {
			T[i] = T[i-1]
			i = i - 1
		}
		T[i] = temp
	}
	fmt.Println("Data berhasil diurutkan (Insertion Sort)")
}

func binarySearch(T tabKomponen, n int, x string) int {
	var left, right, mid int
	var found int

	left = 0
	right = n - 1
	found = -1

	for left <= right && found == -1 {
		mid = (left + right) / 2
		if T[mid].noSeri == x {
			found = mid
		} else if x < T[mid].noSeri {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return found
}

func statistik(T tabKomponen, n int) {
	var i, jumlahMasalah int
	var totalSuhu, rata float64
	totalSuhu = 0
	jumlahMasalah = 0
	for i = 0; i < n; i++ {
		totalSuhu = totalSuhu + T[i].suhu
		if T[i].status == "LAG" || T[i].status == "OVERHEAT" {
			jumlahMasalah = jumlahMasalah + 1
		}
	}

	if n > 0 {
		rata = totalSuhu / float64(n)
	}

	fmt.Println("STATISTIK")
	fmt.Println("Jumlah komponen bermasalah :", jumlahMasalah)
	fmt.Println("Rata-rata suhu :", rata)
}

func main() {
	var T tabKomponen
	var n, menu, idx int
	var x string

	n = 0
	menu = -1

	for menu != 0 {
		fmt.Println("\nSEHATIN PC ")
		fmt.Println("1. Input Data")
		fmt.Println("2. Print Data")
		fmt.Println("3. Ubah Data")
		fmt.Println("4. Hapus Data")
		fmt.Println("5. Sequential Search")
		fmt.Println("6. Selection Sort")
		fmt.Println("7. Insertion Sort")
		fmt.Println("8. Binary Search")
		fmt.Println("9. Statistik")
		fmt.Println("0. Keluar")

		fmt.Print("Pilih menu : ")
		fmt.Scan(&menu)

		switch menu {
		case 1:
			inputData(&T, &n)
		case 2:
			printData(T, n)
		case 3:
			fmt.Print("Cari nama/status : ")
			fmt.Scan(&x)
			ubah(&T, n, x)
		case 4:
			fmt.Print("Cari nama/status : ")
			fmt.Scan(&x)
			hapus(&T, &n, x)
		case 5:
			fmt.Print("Cari nama/status : ")
			fmt.Scan(&x)
			idx = cari(T, n, x)
			if idx == -1 {
				fmt.Println("Data tidak ditemukan")
			} else {
				fmt.Println("Data ditemukan di indeks", idx)
				fmt.Println(T[idx])
			}
		case 6:
			selectionSort(&T, n)
		case 7:
			insertionSort(&T, n)
		case 8:
			fmt.Print("Cari nomor seri : ")
			fmt.Scan(&x)
			idx = binarySearch(T, n, x)
			if idx == -1 {
				fmt.Println("Data tidak ditemukan")
			} else {
				fmt.Println("Data ditemukan")
				fmt.Println(T[idx])
			}
		case 9:
			statistik(T, n)
		case 0:
			fmt.Println("Program selesai")
		default:
			fmt.Println("Menu tidak tersedia")
		}
	}
}
