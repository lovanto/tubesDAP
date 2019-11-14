package main

import (
	"bufio"
	"fmt"
	"os"
)

const n = 100

type typeMahasiswa struct {
	nim      int
	nama     string
	fakultas string
	jurusan  string
	kelas    string
	nohp     string
	email    string
}
type arrType [n]typeMahasiswa

var mahasiswa arrType
var count int

func tampilMenu() {
	fmt.Println("Apa yang ingin anda lakukan?")
	fmt.Println("1. Menampilkan data.")
	fmt.Println("2. Menambah data.")
	fmt.Println("3. Mencari data.")
	fmt.Println("4. Mengubah data.")
	fmt.Println("5. Keluar.")

	menuPilih()
}

func menuPilih() {
	var menu int

	fmt.Scanln(&menu)
	switch menu {
	case 1:
		spasiData()
		menuTampilData()
	case 2:
		spasiData()
		tambahData()
	case 3:
		spasiData()
		cariData()
	case 4:
		spasiData()
		ubahData()
	case 5:
		spasiData()
		os.Exit(3)
	default:
		fmt.Println(" ")
		fmt.Println("Menu yang dipilih tidak ada. Silakan diulangi!")
		spasiData()
		tampilMenu()
	}
}

func menuTampilData() {
	fmt.Println("Tampilkan dengan?")
	fmt.Println("1. Apa adanya.")
	fmt.Println("2. Berdasarkan nim terkecil ke terbesar.")
	fmt.Println("3. Berdasarkan nim terbesar ke terkecil.")
	fmt.Println("4. Berdasarkan nama (A-Z).")
	fmt.Println("5. Berdasarkan nama (Z-A).")
	fmt.Println("6. Kembali.")

	menuTampilDataPilih()
}

func menuTampilDataPilih() {
	var menu int

	fmt.Scanln(&menu)
	switch menu {
	case 1:
		spasiData()
		tampilData()
	case 2:
		spasiData()
		tampilMinToMax()
	case 3:
		spasiData()
		tampilMaxToMin()
	case 4:
		spasiData()
		tampilAToZ()
	case 5:
		spasiData()
		tampilZToA()
	case 6:
		spasiData()
		tampilMenu()
	default:
		fmt.Println(" ")
		fmt.Println("Menu yang dipilih tidak ada. Silakan diulangi!")
		spasiData()
		tampilMenu()
	}
}

func tampilMinToMax() {
	// INSERTION SORT
	for i := 1; i < count; i++ {
		min := mahasiswa[i].nim
		nama := mahasiswa[i].nama
		fakultas := mahasiswa[i].fakultas
		jurusan := mahasiswa[i].jurusan
		kelas := mahasiswa[i].kelas
		nohp := mahasiswa[i].nohp
		email := mahasiswa[i].email
		j := i
		for j > 0 && mahasiswa[j-1].nim > min {
			mahasiswa[j].nim = mahasiswa[j-1].nim
			mahasiswa[j].nama = mahasiswa[j-1].nama
			mahasiswa[j].fakultas = mahasiswa[j-1].fakultas
			mahasiswa[j].jurusan = mahasiswa[j-1].jurusan
			mahasiswa[j].kelas = mahasiswa[j-1].kelas
			mahasiswa[j].nohp = mahasiswa[j-1].nohp
			mahasiswa[j].email = mahasiswa[j-1].email
			j--
		}
		mahasiswa[j].nim = min
		mahasiswa[j].nama = nama
		mahasiswa[j].fakultas = fakultas
		mahasiswa[j].jurusan = jurusan
		mahasiswa[j].kelas = kelas
		mahasiswa[j].nohp = nohp
		mahasiswa[j].email = email
	}
	tampilData()
}

func tampilMaxToMin() {
	// INSERTION SORT
	for i := 1; i < count; i++ {
		min := mahasiswa[i].nim
		nama := mahasiswa[i].nama
		fakultas := mahasiswa[i].fakultas
		jurusan := mahasiswa[i].jurusan
		kelas := mahasiswa[i].kelas
		nohp := mahasiswa[i].nohp
		email := mahasiswa[i].email
		j := i
		for j > 0 && mahasiswa[j-1].nim < min {
			mahasiswa[j].nim = mahasiswa[j-1].nim
			mahasiswa[j].nama = mahasiswa[j-1].nama
			mahasiswa[j].fakultas = mahasiswa[j-1].fakultas
			mahasiswa[j].jurusan = mahasiswa[j-1].jurusan
			mahasiswa[j].kelas = mahasiswa[j-1].kelas
			mahasiswa[j].nohp = mahasiswa[j-1].nohp
			mahasiswa[j].email = mahasiswa[j-1].email
			j--
		}
		mahasiswa[j].nim = min
		mahasiswa[j].nama = nama
		mahasiswa[j].fakultas = fakultas
		mahasiswa[j].jurusan = jurusan
		mahasiswa[j].kelas = kelas
		mahasiswa[j].nohp = nohp
		mahasiswa[j].email = email
	}
	tampilData()
}

func tampilAToZ() {
	// INSERTION SORT
	for i := 1; i < count; i++ {
		min := mahasiswa[i].nama
		nim := mahasiswa[i].nim
		fakultas := mahasiswa[i].fakultas
		jurusan := mahasiswa[i].jurusan
		kelas := mahasiswa[i].kelas
		nohp := mahasiswa[i].nohp
		email := mahasiswa[i].email
		j := i
		for j > 0 && mahasiswa[j-1].nama > min {
			mahasiswa[j].nim = mahasiswa[j-1].nim
			mahasiswa[j].nama = mahasiswa[j-1].nama
			mahasiswa[j].fakultas = mahasiswa[j-1].fakultas
			mahasiswa[j].jurusan = mahasiswa[j-1].jurusan
			mahasiswa[j].kelas = mahasiswa[j-1].kelas
			mahasiswa[j].nohp = mahasiswa[j-1].nohp
			mahasiswa[j].email = mahasiswa[j-1].email
			j--
		}
		mahasiswa[j].nim = nim
		mahasiswa[j].nama = min
		mahasiswa[j].fakultas = fakultas
		mahasiswa[j].jurusan = jurusan
		mahasiswa[j].kelas = kelas
		mahasiswa[j].nohp = nohp
		mahasiswa[j].email = email
	}
	tampilData()
}

func tampilZToA() {
	// INSERTION SORT
	for i := 1; i < count; i++ {
		min := mahasiswa[i].nama
		nim := mahasiswa[i].nim
		fakultas := mahasiswa[i].fakultas
		jurusan := mahasiswa[i].jurusan
		kelas := mahasiswa[i].kelas
		nohp := mahasiswa[i].nohp
		email := mahasiswa[i].email
		j := i
		for j > 0 && mahasiswa[j-1].nama < min {
			mahasiswa[j].nim = mahasiswa[j-1].nim
			mahasiswa[j].nama = mahasiswa[j-1].nama
			mahasiswa[j].fakultas = mahasiswa[j-1].fakultas
			mahasiswa[j].jurusan = mahasiswa[j-1].jurusan
			mahasiswa[j].kelas = mahasiswa[j-1].kelas
			mahasiswa[j].nohp = mahasiswa[j-1].nohp
			mahasiswa[j].email = mahasiswa[j-1].email
			j--
		}
		mahasiswa[j].nim = nim
		mahasiswa[j].nama = min
		mahasiswa[j].fakultas = fakultas
		mahasiswa[j].jurusan = jurusan
		mahasiswa[j].kelas = kelas
		mahasiswa[j].nohp = nohp
		mahasiswa[j].email = email
	}
	tampilData()
}

func tampilData() {
	for i := 0; i < n; i++ {
		if mahasiswa[i].nim == 0 {
			count = i
			break
		}
		showData(i)
		fmt.Println("=======================================================")
		spasiData()
	}
	tampilMenu()
	menuPilih()
}

func tambahData() {
	for i := count; i < n; i++ {
		fmt.Print("NIM: ")
		fmt.Scanln(&mahasiswa[i].nim)
		fmt.Print("Nama: ")
		mahasiswa[i].nama = inputWithSpacing()
		fmt.Print("Fakultas: ")
		mahasiswa[i].fakultas = inputWithSpacing()
		fmt.Print("Jurusan: ")
		mahasiswa[i].jurusan = inputWithSpacing()
		fmt.Print("Kelas: ")
		mahasiswa[i].kelas = inputWithSpacing()
		fmt.Print("No.HP: ")
		mahasiswa[i].nohp = inputWithSpacing()
		fmt.Print("Email: ")
		mahasiswa[i].email = inputWithSpacing()
		fmt.Println("Data berhasil ditambahkan.")

		if cekTambahLagiOrNo() == true {
			spasiData()
			break
		}
	}
	spasiData()
	main()
}

func cariData() {
	tampilMenuCari()
}

func tampilMenuCari() {
	fmt.Println("Metode apa yang ingin digunakan?")
	fmt.Println("1. NIM.")
	fmt.Println("2. Nama.")
	fmt.Println("3. Fakultas.")
	fmt.Println("4. Jurusan.")
	fmt.Println("5. Kelas.")
	fmt.Println("6. Kembali.")

	menuCari()
}

func menuCari() {
	var menu int

	fmt.Scanln(&menu)
	switch menu {
	case 1:
		spasiData()
		cariBerdasarkanNim()
	case 2:
		spasiData()
		cariBerdasarkanNama()
	case 3:
		spasiData()
		cariBerdasarkanFakultas()
	case 4:
		spasiData()
		cariBerdasarkanJurusan()
	case 5:
		spasiData()
		cariBerdasarkanKelas()
	case 6:
		spasiData()
		tampilMenu()
	default:
		fmt.Println(" ")
		fmt.Println("Metode yang dipilih tidak ada. Silakan diulangi!")
		spasiData()
		tampilMenuCari()
	}
}

func cariBerdasarkanNim() {
	var kataKunci int
	var key bool

	fmt.Print("NIM: ")
	fmt.Scanln(&kataKunci)
	spasiData()

	for i := 0; i < count; i++ {
		if mahasiswa[i].nim == kataKunci {
			showData(i)
			key = true
		}
	}

	if key == false {
		fmt.Println("Data Tidak Ditemukan.")
		spasiData()
	}

	if cekCariLagiOrNo() == true {
		tampilMenuCari()
	} else {
		tampilMenu()
	}
}

func cariBerdasarkanNama() {
	var kataKunci string
	var key bool

	fmt.Print("Nama: ")
	kataKunci = inputWithSpacing()
	spasiData()

	for i := 0; i < count; i++ {
		if mahasiswa[i].nama == kataKunci {
			showData(i)
			key = true
		}
	}

	if key == false {
		fmt.Println("Data tidak ditemukan.")
		spasiData()
	}

	if cekCariLagiOrNo() == true {
		tampilMenuCari()
	} else {
		tampilMenu()
	}
}

func cariBerdasarkanFakultas() {
	var kataKunci string
	var key bool

	fmt.Print("Fakultas: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	kataKunci = scanner.Text()
	spasiData()

	for i := 0; i < count; i++ {
		if mahasiswa[i].fakultas == kataKunci {
			showData(i)
			key = true
		}
	}

	if key == false {
		fmt.Println("Data tidak ditemukan.")
		spasiData()
	}

	if cekCariLagiOrNo() == true {
		tampilMenuCari()
	} else {
		tampilMenu()
	}
}

func cariBerdasarkanJurusan() {
	var kataKunci string
	var key bool

	fmt.Print("Jurusan: ")
	kataKunci = inputWithSpacing()
	spasiData()

	for i := 0; i < count; i++ {
		if mahasiswa[i].jurusan == kataKunci {
			showData(i)
			key = true
		}
	}

	if key == false {
		fmt.Println("Data tidak ditemukan.")
		spasiData()
	}

	if cekCariLagiOrNo() == true {
		tampilMenuCari()
	} else {
		tampilMenu()
	}
}

func cariBerdasarkanKelas() {
	var kataKunci string
	var key bool

	fmt.Print("Kelas: ")
	kataKunci = inputWithSpacing()
	spasiData()

	for i := 0; i < count; i++ {
		if mahasiswa[i].kelas == kataKunci {
			showData(i)
			key = true
		}
	}

	if key == false {
		fmt.Println("Data tidak ditemukan.")
		spasiData()
	}

	if cekCariLagiOrNo() == true {
		tampilMenuCari()
	} else {
		tampilMenu()
	}
}

func showData(i int) {
	spasiData()
	fmt.Println("Data ditemukan")
	spasiData()
	fmt.Println("NIM: ", mahasiswa[i].nim)
	fmt.Println("Nama: ", mahasiswa[i].nama)
	fmt.Println("Fakultas: ", mahasiswa[i].fakultas)
	fmt.Println("Jurusan: ", mahasiswa[i].jurusan)
	fmt.Println("Kelas: ", mahasiswa[i].kelas)
	fmt.Println("No.HP: ", mahasiswa[i].nohp)
	fmt.Println("Email: ", mahasiswa[i].email)
	spasiData()
}

func ubahData() {
	var nimS int
	var namaS, fakultasS, jurusanS, kelasS, nohpS, emailS string
	var pos int

	fmt.Print("NIM: ")
	fmt.Scanln(&nimS)
	for i := 0; i < count; i++ {
		if mahasiswa[i].nim == nimS {
			showData(i)
			pos = i
		}
	}

	fmt.Print("Nama: ")
	namaS = inputWithSpacing()
	fmt.Print("Fakultas: ")
	fakultasS = inputWithSpacing()
	fmt.Print("Jurusan: ")
	jurusanS = inputWithSpacing()
	fmt.Print("Kelas: ")
	kelasS = inputWithSpacing()
	fmt.Print("No.HP: ")
	nohpS = inputWithSpacing()
	fmt.Print("Email: ")
	emailS = inputWithSpacing()

	if cekUpdateOrNo() == true {
		mahasiswa[pos].nim = nimS
		mahasiswa[pos].nama = namaS
		mahasiswa[pos].fakultas = fakultasS
		mahasiswa[pos].jurusan = jurusanS
		mahasiswa[pos].kelas = kelasS
		mahasiswa[pos].nohp = nohpS
		mahasiswa[pos].email = emailS
		fmt.Println("Data berhasil diperbarui")
		spasiData()
		tampilMenu()
	} else {
		fmt.Println("Data gagal diperbarui")
		tampilMenu()
	}
}

func tempData() {
	mahasiswa[0].nim = 1302194068
	mahasiswa[0].nama = "Rifky Lovanto"
	mahasiswa[0].fakultas = "Fakultas Informatika"
	mahasiswa[0].jurusan = "S1 Rekayasa Perangkat Lunak"
	mahasiswa[0].kelas = "SE-43-02"
	mahasiswa[0].nohp = "+62-878-2383-7566"
	mahasiswa[0].email = "lovanto@student.telkomuniversity.ac.id"

	mahasiswa[1].nim = 1302194070
	mahasiswa[1].nama = "Rizal Maidan Firdaus"
	mahasiswa[1].fakultas = "Fakultas Rekayasa Industri"
	mahasiswa[1].jurusan = "S1 Sistem Informasi"
	mahasiswa[1].kelas = "SE-43-02"
	mahasiswa[1].nohp = "+62-878-2383-7555"
	mahasiswa[1].email = "rizalmf@student.telkomuniversity.ac.id"

	mahasiswa[2].nim = 1302194048
	mahasiswa[2].nama = "Aku Bukan Ya"
	mahasiswa[2].fakultas = "Fakultas Informatika"
	mahasiswa[2].jurusan = "S1 Rekayasa Perangkat Lunak"
	mahasiswa[2].kelas = "SE-43-02"
	mahasiswa[2].nohp = "+62-878-2383-7566"
	mahasiswa[2].email = "lovanto@student.telkomuniversity.ac.id"
}

func hitungData() {
	for i := 0; i < n; i++ {
		if mahasiswa[i].nim == 0 {
			count = i
			break
		}
	}
}

func spasiData() {
	fmt.Println(" ")
}

func cekTambahLagiOrNo() bool {
	var yesOrNo string
	var kebenaran bool

	fmt.Println("Apakah anda ingin menambah data lagi? Y/N")
	fmt.Scanln(&yesOrNo)
	spasiData()
	if yesOrNo == "N" || yesOrNo == "n" {
		kebenaran = true
	} else if yesOrNo == "Y" || yesOrNo == "y" {
		kebenaran = false
	} else {
		fmt.Println("Jawaban anda tidak sesuai. Silakan diulangi!")
		cekTambahLagiOrNo()
	}
	return kebenaran
}

func cekCariLagiOrNo() bool {
	var yesOrNo string
	var kebenaran bool

	fmt.Println("Apakah anda ingin mencari data lagi? Y/N")
	fmt.Scanln(&yesOrNo)
	spasiData()
	if yesOrNo == "Y" || yesOrNo == "y" {
		kebenaran = true
	} else if yesOrNo == "N" || yesOrNo == "n" {
		kebenaran = false
	} else {
		fmt.Println("Jawaban anda tidak sesuai. Silakan diulangi!")
		cekCariLagiOrNo()
	}
	return kebenaran
}

func cekUpdateOrNo() bool {
	var yesOrNo string
	var kebenaran bool

	fmt.Println("Apakah anda yakin ingin memperbarui data? Y/N")
	fmt.Scanln(&yesOrNo)
	spasiData()
	if yesOrNo == "Y" || yesOrNo == "y" {
		kebenaran = true
	} else if yesOrNo == "N" || yesOrNo == "n" {
		kebenaran = false
	} else {
		fmt.Println("Jawaban anda tidak sesuai. Silakan diulangi!")
		cekUpdateOrNo()
	}
	return kebenaran
}

func inputWithSpacing() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	tempData()
	hitungData()
	tampilMenu()
}
