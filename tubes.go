package main

import "fmt"

const max = 100

type langganan struct {
	nama     string
	kategori string
	biaya    int
	metode   string
	status   string
	tanggal  int
	bulan    int
	tahun    int
}

var daftarlangganan [max]langganan
var total int = 0
var i int
var l langganan

func tambahlangganan() {
	if total >= max {
		fmt.Println("Kapasitas langganan penuh!")
		return
	}
	var nama, kategori, metode, status string
	var biaya, tanggal, bulan, tahun int
	fmt.Print("Nama Langganan: ")
	fmt.Scan(&nama)
	fmt.Print("Kategori: ")
	fmt.Scan(&kategori)
	fmt.Print("Biaya: ")
	fmt.Scan(&biaya)
	fmt.Print("Metode Pembayaran: ")
	fmt.Scan(&metode)
	fmt.Print("Status (Aktif/Nonaktif): ")
	fmt.Scan(&status)
	fmt.Print("Tanggal Jatuh Tempo (DD MM YYYY): ")
	fmt.Scan(&tanggal, &bulan, &tahun)
	if tanggal > 31 {
		fmt.Println("tanggal jatuh tempo tidak sesuai")
		fmt.Println("masukan tanggal atau jatuh tempo yang sesuai!")
		fmt.Scan(&tanggal)
	} else if bulan > 12 {
		fmt.Println("bulan jatuh tempo tidak sesuai")
		fmt.Println("masukan bulan jatuh tempo yang sesuai!")
		fmt.Scan(&bulan)
	}

	daftarlangganan[total] = langganan{nama, kategori, biaya, metode, status, tanggal, bulan, tahun}
	total++
	fmt.Println("Berhasil ditambahkan.")

}

func tampilkanlangganan() {
	if total == 0 {
		fmt.Println("data masih kosong, tambahkan data langganan nya dulu ya")
		return
	}
	fmt.Printf("%-4s %-20s %-15s %-10s %-15s %-10s %-15s\n",
		"No.", "Nama", "Kategori", "Biaya", "Metode", "Status", "Jatuh Tempo")

	for i = 0; i < total; i++ {
		l = daftarlangganan[i]
		fmt.Printf("%-4d %-20s %-15s Rp%-9d %-15s %-10s %d %d %d\n", i+1, l.nama, l.kategori, l.biaya, l.metode, l.status, l.tanggal, l.bulan, l.tahun)
	}
	fmt.Printf("Total: %d langganan\n", total)
}

func hapuslangganan() {
	if total == 0 {
		fmt.Println("data masih kosong, tambahkan data langganan nya dulu ya")
		return
	}
	var index int
	tampilkanlangganan()
	fmt.Print("masukan no langganan yang akan dihapus")
	fmt.Scan(&index)
	index--
	if index < 0 || index >= total {
		fmt.Println("Index tidak valid")
		return
	}

	for i = index; i < total-1; i++ {
		daftarlangganan[i] = daftarlangganan[i+1]
	}

	total--
	fmt.Println("Data berhasil dihapus.")

}

func caridatalangganan() {
	if total == 0 {
		fmt.Println("data masih kosong, tambahkan data langganan nya dulu ya")
		return
	}
	var pilihan string
	var index int
	fmt.Println("Cari berdasarkan apa? (nama/kategori): ")
	fmt.Scan(&pilihan)

	if pilihan == "nama" {
		var nama string
		fmt.Print("Masukkan nama langganan yang dicari: ")
		fmt.Scan(&nama)

		urutkanBerdasarkanNama()

		index = binarySearchNama(nama)
		if index != -1 {
			l := daftarlangganan[index]
			fmt.Println("Ditemukan:")
			fmt.Printf("Nama: %s | Kategori: %s | Biaya: Rp%d | Metode: %s | Status: %s | Jatuh Tempo: %d %d %d\n",
				l.nama, l.kategori, l.biaya, l.metode, l.status, l.tanggal, l.bulan, l.tahun)
		} else {
			fmt.Println("Langganan tidak ditemukan.")
		}

	} else if pilihan == "kategori" {
		var kategori string
		fmt.Print("Masukkan kategori yang dicari: ")
		fmt.Scan(&kategori)
		cariKategori(kategori)

	} else {
		fmt.Println("Pilihan tidak valid. Gunakan 'nama' atau 'kategori'.")
	}
}

func urutkanBerdasarkanNama() {
	var temp langganan
	var j int
	for i = 1; i < total; i++ {
		temp = daftarlangganan[i]
		j = i - 1
		for j >= 0 && daftarlangganan[j].nama > temp.nama {
			daftarlangganan[j+1] = daftarlangganan[j]
			j--
		}
		daftarlangganan[j+1] = temp
	}
}

func binarySearchNama(namaCari string) int {
	var low, high, mid int
	low = 0
	high = total - 1

	for low <= high {
		mid = (low + high) / 2
		if daftarlangganan[mid].nama == namaCari {
			return mid
		} else if daftarlangganan[mid].nama < namaCari {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func cariKategori(kategoriCari string) {
	var ketemu bool = false
	for i = 0; i < total; i++ {
		if daftarlangganan[i].kategori == kategoriCari {
			fmt.Printf("Ditemukan: %s | Biaya: Rp%d\n", daftarlangganan[i].nama, daftarlangganan[i].biaya)
			ketemu = true
		}
	}
	if ketemu == false {
		fmt.Println("Kategori tidak ditemukan.")
	}
}

func ubahlangganan() {
	if total == 0 {
		fmt.Println("Belum ada data langganan.")
		return
	}

	tampilkanlangganan()
	var index int
	fmt.Print("Masukkan nomor langganan yang ingin diubah: ")
	fmt.Scan(&index)
	index--

	if index < 0 || index >= total {
		fmt.Println("Nomor tidak valid.")
		return
	}

	for {
		fmt.Println("\nPilih data yang ingin diubah:")
		fmt.Println("1. Nama")
		fmt.Println("2. Kategori")
		fmt.Println("3. Biaya")
		fmt.Println("4. Metode Pembayaran")
		fmt.Println("5. Status")
		fmt.Println("6. Jatuh Tempo")
		fmt.Println("7. Selesai")

		var pilihan int
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Print("Masukkan nama baru: ")
			fmt.Scan(&daftarlangganan[index].nama)
		case 2:
			fmt.Print("Masukkan kategori baru: ")
			fmt.Scan(&daftarlangganan[index].kategori)
		case 3:
			fmt.Print("Masukkan biaya baru: ")
			fmt.Scan(&daftarlangganan[index].biaya)
		case 4:
			fmt.Print("Masukkan metode pembayaran baru: ")
			fmt.Scan(&daftarlangganan[index].metode)
		case 5:
			fmt.Print("Masukkan status baru (Aktif/Nonaktif): ")
			fmt.Scan(&daftarlangganan[index].status)
		case 6:
			fmt.Print("Masukkan jatuh tempo baru (DD-MM-YYYY): ")
			fmt.Scan(&daftarlangganan[index].tanggal, &daftarlangganan[index].bulan, &daftarlangganan[index].tahun)
		case 7:
			fmt.Println("Perubahan selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func hitungbiaya() {
	var totalbiaya int
	if total == 0 {
		fmt.Print("data masih kosong, tambahkan data nya dulu")
	}
	totalbiaya = 0
	for i = 0; i < total; i++ {
		if daftarlangganan[i].status == "aktif" {
			totalbiaya += daftarlangganan[i].biaya
		}
	}
	fmt.Print("total biaya seluruh langgananmu yang masih aktif sebesar RP.", totalbiaya)
}

func pengingatJatuhTempo() {
	var ditemukan bool
	var selisih int
	if total == 0 {
		fmt.Println("Belum ada data langganan.")
		return
	}

	var hariIni, bulanIni, tahunIni int
	fmt.Print("Masukkan tanggal hari ini (DD MM YYYY): ")
	fmt.Scan(&hariIni, &bulanIni, &tahunIni)
	if hariIni > 31 {
		fmt.Println("tanggal jatuh tempo tidak sesuai")
		fmt.Println("masukan tanggal atau jatuh tempo yang sesuai!")
		fmt.Scan(&hariIni)
	} else if bulanIni > 12 {
		fmt.Println("bulan jatuh tempo tidak sesuai")
		fmt.Println("masukan bulan jatuh tempo yang sesuai!")
		fmt.Scan(&bulanIni)
	}

	ditemukan = false
	for i = 0; i < total; i++ {
		if daftarlangganan[i].status == "aktif" {
			l = daftarlangganan[i]
			if l.tahun == tahunIni && l.bulan == bulanIni {
				selisih = l.tanggal - hariIni
				if selisih >= 0 && selisih <= 3 {
					fmt.Printf("- %s | Jatuh Tempo: %02d-%02d-%04d (dalam %d hari)\n", l.nama, l.tanggal, l.bulan, l.tahun, selisih)
					ditemukan = true
				}
			}
		}
	}

	if ditemukan == false {
		fmt.Println("Tidak ada langganan yang jatuh tempo dalam 3 hari.")
	}
}

func urutkanbiayadescending() {
	if total == 0 {
		fmt.Println("data masih kosong, tambahkan data langganan nya dulu ya")
		return
	}

	var maxIdx, i, j int
	var temp langganan

	for i = 0; i < total-1; i++ {
		maxIdx = i
		for j = i + 1; j < total; j++ {
			if daftarlangganan[j].biaya > daftarlangganan[maxIdx].biaya {
				maxIdx = j
			}
		}

		if maxIdx != i {
			temp = daftarlangganan[i]
			daftarlangganan[i] = daftarlangganan[maxIdx]
			daftarlangganan[maxIdx] = temp
		}
	}

	fmt.Println("Data telah diurutkan berdasarkan biaya (tertinggi ke terendah), silahkan pilih opsi menu tampilkan langganan jika ingin melihat tampilan daftar langganan.")
}

func urutkantanggalascending() {
	if total == 0 {
		fmt.Println("data masih kosong, tambahkan data langganan nya dulu ya")
		return
	}

	var j int

	for i = 1; i < total; i++ {
		l = daftarlangganan[i]
		j = i - 1

		for j >= 0 && (daftarlangganan[j].tahun > l.tahun ||
			(daftarlangganan[j].tahun == l.tahun && daftarlangganan[j].bulan > l.bulan) ||
			(daftarlangganan[j].tahun == l.tahun && daftarlangganan[j].bulan == l.bulan && daftarlangganan[j].tanggal > l.tanggal)) {

			daftarlangganan[j+1] = daftarlangganan[j]
			j--
		}
		daftarlangganan[j+1] = l
	}

	fmt.Println("Data telah diurutkan berdasarkan jatuh tempo (terdekat lebih dulu), silahkan pilih opsi menu tampilkan langganan jika ingin melihat tampilan daftar langganan.")
}

func rekomendasipenghematan() {
	if total == 0 {
		fmt.Println("data masih kosong, tambahkan data langganan nya dulu ya")
		return
	}

	var indexTermahal int = 0
	for i = 1; i < total; i++ {
		if daftarlangganan[i].status == "aktif" {
			if daftarlangganan[i].biaya > daftarlangganan[indexTermahal].biaya {
				indexTermahal = i
			}
		}
	}
	fmt.Printf("Langganan termahal: %s (Rp%d)\n", daftarlangganan[indexTermahal].nama, daftarlangganan[indexTermahal].biaya)
	fmt.Println("  Pertimbangkan untuk mengurangi atau menonaktifkan jika tidak terlalu penting.")

}

func main() {
	var pilihan int
	var pass int
	fmt.Println("masukan password untuk mengakses aplikasi")
	fmt.Scan(&pass)
	for pass != 1111 {
		fmt.Println("password salah silahkan coba lagi")
		fmt.Scan(&pass)
	}
	for {
		fmt.Println("\n=== MENU MANAJEMEN SUBSKRIPSI ===")
		fmt.Println("1. Tambah Langganan")
		fmt.Println("2. Tampilkan Semua Langganan")
		fmt.Println("3. Hapus Langganan")
		fmt.Println("4. Cari Langganan")
		fmt.Println("5. ubah data langganan")
		fmt.Println("6. total biaya seluruh langganan")
		fmt.Println("7. Pengingat Jatuh Tempo")
		fmt.Println("8. pengurutan (descending) langganan berdasarkan biaya")
		fmt.Println("9. pengurutan (ascending) langganan berdasarkan tanggal jatuh tempo")
		fmt.Println("10. rekomendasi penghematan")
		fmt.Println("11. Keluar")
		fmt.Println("mohon gunakan huruf kecil pada proses input data langganan\n")
		fmt.Print("Pilih menu (1-11): ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahlangganan()
		case 2:
			tampilkanlangganan()
		case 3:
			hapuslangganan()
		case 4:
			caridatalangganan()
		case 5:
			ubahlangganan()
		case 6:
			hitungbiaya()
		case 7:
			pengingatJatuhTempo()
		case 8:
			urutkanbiayadescending()
		case 9:
			urutkantanggalascending()
		case 10:
			rekomendasipenghematan()
		case 11:
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
