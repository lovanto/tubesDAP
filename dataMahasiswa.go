package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

const n = 1000

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

// START OF MENU LIST CODE

func tampilMenu() {
	spacingData()
	fmt.Println(" ", "Apa yang ingin anda lakukan?")
	fmt.Println(" ", "1. Menampilkan data.")
	fmt.Println(" ", "2. Menambah data.")
	fmt.Println(" ", "3. Mencari data.")
	fmt.Println(" ", "4. Mengubah data.")
	fmt.Println(" ", "5. Menghapus data.")
	fmt.Println(" ", "6. Keluar.")

	spacingData()
	menuPilih()
}

func menuTampilData() {
	clsCode()
	spacingData()
	fmt.Println(" ", "Tampilkan dengan?")
	fmt.Println(" ", "1. Apa adanya.")
	fmt.Println(" ", "2. Berdasarkan nim terkecil ke terbesar.")
	fmt.Println(" ", "3. Berdasarkan nim terbesar ke terkecil.")
	fmt.Println(" ", "4. Berdasarkan nama (A-Z).")
	fmt.Println(" ", "5. Berdasarkan nama (Z-A).")
	fmt.Println(" ", "6. Kembali.")

	spacingData()
	menuTampilDataPilih()
}

func tampilMenuCari() {
	clsCode()
	spacingData()
	fmt.Println(" ", "Metode apa yang ingin digunakan?")
	fmt.Println(" ", "1. NIM.")
	fmt.Println(" ", "2. Nama.")
	fmt.Println(" ", "3. Fakultas.")
	fmt.Println(" ", "4. Jurusan.")
	fmt.Println(" ", "5. Kelas.")
	fmt.Println(" ", "6. Kembali.")

	spacingData()
	menuCari()
}

// END OF MENU LIST CODE

// START OF ANSWER MENU CODE

func menuPilih() {
	var menu int

	fmt.Print(" ", "Menu -> ")
	fmt.Scanln(&menu)
	switch menu {
	case 1:
		spacingData()
		menuTampilData()
	case 2:
		clsCode()
		spacingData()
		tambahData()
	case 3:
		spacingData()
		cariData()
	case 4:
		spacingData()
		updateData()
	case 5:
		spacingData()
		hapusData()
	case 6:
		spacingData()
		os.Exit(3)
	default:
		clsCode()
		fmt.Println("Menu yang dipilih tidak ada. Silakan diulangi!")
		spacingData()
		tampilMenu()
	}
}

func menuTampilDataPilih() {
	var menu int

	fmt.Print(" ", "Menu -> ")
	fmt.Scanln(&menu)
	switch menu {
	case 1:
		spacingData()
		tampilData()
	case 2:
		spacingData()
		tampilMinToMax()
	case 3:
		spacingData()
		tampilMaxToMin()
	case 4:
		spacingData()
		tampilAToZ()
	case 5:
		spacingData()
		tampilZToA()
	case 6:
		spacingData()
		clsCode()
		tampilMenu()
	default:
		clsCode()
		fmt.Println("Menu yang dipilih tidak ada. Silakan diulangi!")
		spacingData()
		menuTampilData()
	}
}

func menuCari() {
	var menu int

	fmt.Print(" ", "Menu -> ")
	fmt.Scanln(&menu)
	switch menu {
	case 1:
		clsCode()
		spacingData()
		cariBerdasarkanNim()
	case 2:
		clsCode()
		spacingData()
		cariBerdasarkanNama()
	case 3:
		clsCode()
		spacingData()
		cariBerdasarkanFakultas()
	case 4:
		clsCode()
		spacingData()
		cariBerdasarkanJurusan()
	case 5:
		clsCode()
		spacingData()
		cariBerdasarkanKelas()
	case 6:
		clsCode()
		tampilMenu()
	default:
		clsCode()
		fmt.Println("Metode yang dipilih tidak ada. Silakan diulangi!")
		spacingData()
		tampilMenuCari()
	}
}

// END OF ANSWER MENU CODE

// START OF SORT DATA CODE
// INSERTION SORT
func tampilMinToMax() {
	clsCode()
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
	clsCode()
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
	clsCode()
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
	clsCode()
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

// END OF SORT DATA CODE

// START OF SHOW DATA CODE

func tampilData() {
	clsCode()
	for i := 0; i < n; i++ {
		if mahasiswa[i].nim == 0 {
			count = i
			break
		}
		fmt.Println("=======================================================")
		showData(i)
	}
	fmt.Println("=======================================================")
	spacingData()
	tampilMenu()
	menuPilih()
}

func showData(i int) {
	spacingData()
	fmt.Println(" ", "NIM: ", mahasiswa[i].nim)
	fmt.Println(" ", "Nama: ", mahasiswa[i].nama)
	fmt.Println(" ", "Fakultas: ", mahasiswa[i].fakultas)
	fmt.Println(" ", "Jurusan: ", mahasiswa[i].jurusan)
	fmt.Println(" ", "Kelas: ", mahasiswa[i].kelas)
	fmt.Println(" ", "No.HP: ", mahasiswa[i].nohp)
	fmt.Println(" ", "Email: ", mahasiswa[i].email)
	spacingData()
}

// END OF SHOW DATA CODE

// START OF ADD NEW DATA CODE

func tambahData() {
	fmt.Println("=======================================================")
	spacingData()
	for i := count; i < n; i++ {
		fmt.Print(" ", "NIM: ")
		fmt.Scanln(&mahasiswa[i].nim)
		fmt.Print(" ", "Nama: ")
		mahasiswa[i].nama = inputWithSpacing()
		fmt.Print(" ", "Fakultas: ")
		mahasiswa[i].fakultas = inputWithSpacing()
		fmt.Print(" ", "Jurusan: ")
		mahasiswa[i].jurusan = inputWithSpacing()
		fmt.Print(" ", "Kelas: ")
		mahasiswa[i].kelas = inputWithSpacing()
		fmt.Print(" ", "No.HP: ")
		mahasiswa[i].nohp = inputWithSpacing()
		fmt.Print(" ", "Email: ")
		mahasiswa[i].email = inputWithSpacing()

		spacingData()
		fmt.Println("=======================================================")
		spacingData()
		if addMoreOrNo() == true {
			spacingData()
			break
		}
	}
	fmt.Println(" ", "Data berhasil ditambahkan.")
	tampilMenu()
}

// END OF ADD NEW DATA CODE

// START OF FIND DATA CODE

func cariData() {
	tampilMenuCari()
}

func cariBerdasarkanNim() {
	var kataKunci int
	var key bool

	fmt.Print(" ", "NIM: ")
	fmt.Scanln(&kataKunci)
	spacingData()

	for i := 0; i < count; i++ {
		if mahasiswa[i].nim == kataKunci {
			spacingData()
			fmt.Println("=======================================================")
			showData(i)
			spacingData()
			fmt.Println("=======================================================")
			key = true
		}
	}

	if key == false {
		spacingData()
		fmt.Println(" ", "Data Tidak Ditemukan.")
		spacingData()
	}

	spacingData()
	if findMoreOrNo() == true {
		tampilMenuCari()
	}
}

func cariBerdasarkanNama() {
	var kataKunci string
	var key bool

	fmt.Print(" ", "Nama: ")
	kataKunci = inputWithSpacing()
	spacingData()

	for i := 0; i < count; i++ {
		if mahasiswa[i].nama == kataKunci {
			spacingData()
			fmt.Println("=======================================================")
			showData(i)
			spacingData()
			fmt.Println("=======================================================")
			key = true
		}
	}

	if key == false {
		fmt.Println(" ", "Data tidak ditemukan.")
		spacingData()
	}

	if findMoreOrNo() == true {
		tampilMenuCari()
	} else {
		tampilMenu()
	}
}

func cariBerdasarkanFakultas() {
	var kataKunci string
	var key bool

	fmt.Print(" ", "Fakultas: ")
	kataKunci = inputWithSpacing()
	spacingData()

	for i := 0; i < count; i++ {
		if mahasiswa[i].fakultas == kataKunci {
			spacingData()
			fmt.Println("=======================================================")
			showData(i)
			spacingData()
			fmt.Println("=======================================================")
			key = true
		}
	}

	if key == false {
		fmt.Println(" ", "Data tidak ditemukan.")
		spacingData()
	}

	if findMoreOrNo() == true {
		tampilMenuCari()
	} else {
		tampilMenu()
	}
}

func cariBerdasarkanJurusan() {
	var kataKunci string
	var key bool

	fmt.Print(" ", "Jurusan: ")
	kataKunci = inputWithSpacing()
	spacingData()

	for i := 0; i < count; i++ {
		if mahasiswa[i].jurusan == kataKunci {
			spacingData()
			fmt.Println("=======================================================")
			showData(i)
			spacingData()
			fmt.Println("=======================================================")
			key = true
		}
	}

	if key == false {
		fmt.Println(" ", "Data tidak ditemukan.")
		spacingData()
	}

	if findMoreOrNo() == true {
		tampilMenuCari()
	} else {
		tampilMenu()
	}
}

func cariBerdasarkanKelas() {
	var kataKunci string
	var key bool

	fmt.Print(" ", "Kelas: ")
	kataKunci = inputWithSpacing()
	spacingData()

	for i := 0; i < count; i++ {
		if mahasiswa[i].kelas == kataKunci {
			spacingData()
			fmt.Println("=======================================================")
			showData(i)
			spacingData()
			fmt.Println("=======================================================")
			key = true
		}
	}

	if key == false {
		fmt.Println(" ", "Data tidak ditemukan.")
		spacingData()
	}

	if findMoreOrNo() == true {
		tampilMenuCari()
	} else {
		tampilMenu()
	}
}

// END OF FIND DATA CODE

// START OF UPDATE DATA CODE

func updateData() {
	var nimS int
	var namaS, fakultasS, jurusanS, kelasS, nohpS, emailS string
	var namaM, fakultasM, jurusanM, kelasM, nohpM, emailM string
	var pos int

	pos = -1
	clsCode()
	spacingData()
	fmt.Print(" ", "NIM: ")
	fmt.Scanln(&nimS)

	spacingData()
	fmt.Println(" ", "Ketik '-' jika tidak ingin mengubah data.")
	spacingData()
	fmt.Println("=======================================================")

	for i := 0; i < count; i++ {
		if mahasiswa[i].nim == nimS {

			showData(i)
			spacingData()
			fmt.Println("=======================================================")
			spacingData()
			pos = i

			namaM = mahasiswa[i].nama
			fakultasM = mahasiswa[i].fakultas
			jurusanM = mahasiswa[i].jurusan
			kelasM = mahasiswa[i].kelas
			nohpM = mahasiswa[i].nohp
			emailM = mahasiswa[i].email

			fmt.Print(" ", "Nama: ")
			namaS = inputWithSpacing()
			fmt.Print(" ", "Fakultas: ")
			fakultasS = inputWithSpacing()
			fmt.Print(" ", "Jurusan: ")
			jurusanS = inputWithSpacing()
			fmt.Print(" ", "Kelas: ")
			kelasS = inputWithSpacing()
			fmt.Print(" ", "No.HP: ")
			nohpS = inputWithSpacing()
			fmt.Print(" ", "Email: ")
			emailS = inputWithSpacing()
		}
	}

	spacingData()
	fmt.Println("=======================================================")

	if pos != -1 {
		if updateOrNo() == true {
			mahasiswa[pos].nim = nimS

			mahasiswa[pos].nama = namaS
			if namaS == "-" {
				mahasiswa[pos].nama = namaM
			}

			mahasiswa[pos].fakultas = fakultasS
			if fakultasS == "-" {
				mahasiswa[pos].fakultas = fakultasM
			}

			mahasiswa[pos].jurusan = jurusanS
			if jurusanS == "-" {
				mahasiswa[pos].jurusan = jurusanM
			}

			mahasiswa[pos].kelas = kelasS
			if kelasS == "-" {
				mahasiswa[pos].kelas = kelasM
			}

			mahasiswa[pos].nohp = nohpS
			if nohpS == "-" {
				mahasiswa[pos].nohp = nohpM
			}

			mahasiswa[pos].email = emailS
			if emailS == "-" {
				mahasiswa[pos].email = emailM
			}
			clsCode()

			spacingData()
			fmt.Println(" ", "Data berhasil diperbarui.")
			spacingData()
			tampilMenu()
		} else {
			clsCode()

			spacingData()
			fmt.Println(" ", "Data gagal diperbarui.")
			tampilMenu()
		}
	} else {
		clsCode()

		spacingData()
		fmt.Println(" ", "Data tidak ditemukan.")
		spacingData()
		tampilMenu()
	}
}

// END OF UPDATE DATA CODE

// START OF DELETE DATA CODE

func hapusData() {
	var nimS int
	var key bool

	clsCode()
	spacingData()
	fmt.Print(" ", "NIM: ")
	fmt.Scanln(&nimS)
	for i := 0; i < count; i++ {
		if mahasiswa[i].nim == nimS {

			showData(i)
			key = true
		}
	}

	if key == true {
		if deleteOrNo() == true {
			for j := 0; j < count; j++ {
				if mahasiswa[j].nim == nimS {
					mahasiswa[j] = mahasiswa[j+1]
				}
				if key == true {
					mahasiswa[j+1] = mahasiswa[j+2]
				}
			}
			clsCode()
			spacingData()
			fmt.Println(" ", "Data berhasil dihapus.")
			spacingData()
			tampilMenu()
		} else {
			clsCode()
			spacingData()
			fmt.Println(" ", "Data gagal diperbarui.")
			tampilMenu()
		}
	} else {
		clsCode()
		spacingData()
		fmt.Println(" ", "Data tidak ditemukan.")
		spacingData()
		tampilMenu()
	}
}

// END OF DELETE DATA CODE

// START OF BOOLEAN CODE

func addMoreOrNo() bool {
	var yesOrNo string
	var kebenaran bool

	for kebenaran = false; !kebenaran; {
		fmt.Print(" ", "Apakah anda ingin menambah data lagi? Y/N. ")
		fmt.Scanln(&yesOrNo)
		spacingData()
		kebenaran = yesOrNo == "N" || yesOrNo == "n"

		if yesOrNo == "Y" || yesOrNo == "y" {
			countData()
			tambahData()
		}
		if yesOrNo != "N" || yesOrNo != "n" || yesOrNo != "Y" || yesOrNo != "y" {
			fmt.Println(" ", "Jawaban anda tidak sesuai. Silakan diulangi!")
		}
	}

	return kebenaran
}

func findMoreOrNo() bool {
	var yesOrNo string
	var kebenaran bool

	for kebenaran = false; !kebenaran; {
		fmt.Print(" ", "Apakah anda ingin menambah data lagi? Y/N. ")
		fmt.Scanln(&yesOrNo)
		spacingData()
		kebenaran = yesOrNo == "Y" || yesOrNo == "y"

		if yesOrNo == "N" || yesOrNo == "n" {
			tampilMenu()
		}
		if yesOrNo != "N" || yesOrNo != "n" || yesOrNo != "Y" || yesOrNo != "y" {
			fmt.Println(" ", "Jawaban anda tidak sesuai. Silakan diulangi!")
		}
	}

	return kebenaran
}

func updateOrNo() bool {
	var yesOrNo string
	var kebenaran bool

	for kebenaran = false; !kebenaran; {
		fmt.Print(" ", "Apakah anda ingin memperbarui data? Y/N. ")
		fmt.Scanln(&yesOrNo)
		spacingData()
		kebenaran = yesOrNo == "Y" || yesOrNo == "y"

		if yesOrNo == "N" || yesOrNo == "n" {
			tampilMenu()
		}
		if yesOrNo != "N" || yesOrNo != "n" || yesOrNo != "Y" || yesOrNo != "y" {
			fmt.Println(" ", "Jawaban anda tidak sesuai. Silakan diulangi!")
		}
	}

	return kebenaran
}

func deleteOrNo() bool {
	var yesOrNo string
	var kebenaran bool

	for kebenaran = false; !kebenaran; {
		fmt.Print(" ", "Apakah anda ingin menghapus data? Y/N. ")
		fmt.Scanln(&yesOrNo)
		spacingData()
		kebenaran = yesOrNo == "Y" || yesOrNo == "y"

		if yesOrNo == "N" || yesOrNo == "n" {
			tampilMenu()
		}
		if yesOrNo != "N" || yesOrNo != "n" || yesOrNo != "Y" || yesOrNo != "y" {
			fmt.Println(" ", "Jawaban anda tidak sesuai. Silakan diulangi!")
		}
	}

	return kebenaran
}

// END OF TRUTH CODE

// START OF ANOTHER CODE

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

func countData() {
	for i := 0; i < n; i++ {
		if mahasiswa[i].nim == 0 {
			count = i
			break
		}
	}
}

func spacingData() {
	fmt.Println(" ")
}

func inputWithSpacing() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func clsCode() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// END OF ANOTHER CODE

func main() {
	clsCode()
	tempData()
	countData()
	tampilMenu()
}
