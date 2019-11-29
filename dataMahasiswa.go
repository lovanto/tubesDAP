package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
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
	countData()
	fmt.Println(" ", "Apa yang ingin anda lakukan?")
	fmt.Println(" ", "1. Menampilkan data.")
	fmt.Println(" ", "2. Menambah data.")
	fmt.Println(" ", "3. Mencari data.")
	fmt.Println(" ", "4. Mengubah data.")
	fmt.Println(" ", "5. Menghapus data.")
	fmt.Println(" ", "6. Optional menu.")
	fmt.Println(" ", "7. Keluar.")

	spacingData()
	menuPilih()
}

func menuTampilData() {
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
	spacingData()
	fmt.Println(" ", "Cari berdasarkan?")
	fmt.Println(" ", "1. NIM.")
	fmt.Println(" ", "2. Nama.")
	fmt.Println(" ", "3. Fakultas.")
	fmt.Println(" ", "4. Jurusan.")
	fmt.Println(" ", "5. Kelas.")
	fmt.Println(" ", "6. Kembali.")

	spacingData()
	menuCari()
}

func tampilMenuExport() {
	spacingData()
	createFile()
	fmt.Println(" ", "Apa yang ingin anda lakukan?")
	fmt.Println(" ", "1. Export data ke txt.")
	fmt.Println(" ", "2. Tampilkan data dari txt.")
	fmt.Println(" ", "3. Hapus file txt.")
	fmt.Println(" ", "4. Kembali.")

	spacingData()
	menuTampilExport()
}

// END OF MENU LIST CODE

// START OF ANSWER MENU CODE

func menuPilih() {
	var menu int

	fmt.Print(" ", "Menu -> ")
	fmt.Scanln(&menu)
	switch menu {
	case 1:
		deleteSpace()
		menuTampilData()
	case 2:
		deleteSpace()
		tambahData()
	case 3:
		deleteSpace()
		cariData()
	case 4:
		spacingData()
		updateData()
	case 5:
		spacingData()
		hapusData()
	case 6:
		deleteSpace()
		tampilMenuExport()
	case 7:
		spacingData()
		os.Exit(3)
	default:
		deleteSpace()
		fmt.Println(" ", "Menu yang dipilih tidak ada atau salah. Silakan diulangi!")
		tampilMenu()
	}
}

func menuTampilDataPilih() {
	var menu int

	fmt.Print(" ", "Menu -> ")
	fmt.Scanln(&menu)
	switch menu {
	case 1:
		deleteSpace()
		tampilData()
	case 2:
		deleteSpace()
		tampilMinToMax()
	case 3:
		deleteSpace()
		tampilMaxToMin()
	case 4:
		deleteSpace()
		tampilAToZ()
	case 5:
		deleteSpace()
		tampilZToA()
	case 6:
		deleteSpace()
		tampilMenu()
	default:
		deleteSpace()
		fmt.Println(" ", "Menu yang dipilih tidak ada atau salah. Silakan diulangi!")
		menuTampilData()
	}
}

func menuCari() {
	var menu int

	fmt.Print(" ", "Menu -> ")
	fmt.Scanln(&menu)
	switch menu {
	case 1:
		deleteSpace()
		cariBerdasarkanNim()
	case 2:
		deleteSpace()
		cariBerdasarkanNama()
	case 3:
		deleteSpace()
		cariBerdasarkanFakultas()
	case 4:
		deleteSpace()
		cariBerdasarkanJurusan()
	case 5:
		deleteSpace()
		cariBerdasarkanKelas()
	case 6:
		deleteSpace()
		tampilMenu()
	default:
		deleteSpace()
		fmt.Println(" ", "Kata kunci yang dipilih tidak ada atau salah. Silakan diulangi!")
		spacingData()
		tampilMenuCari()
	}
}

func menuTampilExport() {
	var menu int

	fmt.Print(" ", "Menu -> ")
	fmt.Scanln(&menu)
	switch menu {
	case 1:
		deleteSpace()
		writeFile()
		tampilMenuExport()
	case 2:
		deleteSpace()
		readFile()
		tampilMenuExport()
	case 3:
		deleteSpace()
		deleteFile()
		tampilMenu()
	case 4:
		deleteSpace()
		tampilMenu()
	default:
		deleteSpace()
		fmt.Println(" ", "Menu yang dipilih tidak ada atau salah. Silakan diulangi!")
		spacingData()
		menuTampilData()
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
	for i := 0; i < count; i++ {
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
	var key bool
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
		if mahasiswa[i].nim != 0 {
			key = true
		}

		spacingData()
		fmt.Println("=======================================================")
		spacingData()
		if addMoreOrNo() == true {
			spacingData()
			break
		}
	}
	deleteSpace()
	if key == true {
		fmt.Println(" ", "Data berhasil ditambahkan.")
	} else {
		fmt.Println(" ", "Data gagal ditambahkan.")
	}
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
	fmt.Scanln(&kataKunci) //digunakan untuk kata kunci data yang ingin dicari
	spacingData()

	for i := 0; i < count; i++ { //digunakan untuk mencari data sesuai kata kunci sebanyak count
		if mahasiswa[i].nim == kataKunci { //jika data ada maka akan ditampilkan
			spaceEquals()
			showData(i) //tampil data sesuai kata kunci
			spaceEquals()
			key = true
		}
	}

	if key == false {
		spaceNull()
	}

	if findMoreOrNo() == true { // ini pilihan apakah ingin mencari data lagi atau tidak
		deleteSpace()
		tampilMenuCari()
	} else {
		deleteSpace()
		tampilMenu()
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
			spaceEquals()
			showData(i)
			spaceEquals()
			key = true
		}
	}

	if key == false {
		spaceNull()
	}

	if findMoreOrNo() == true {
		deleteSpace()
		tampilMenuCari()
	} else {
		deleteSpace()
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
			spaceEquals()
			showData(i)
			spaceEquals()
			key = true
		}
	}

	if key == false {
		spaceNull()
	}

	if findMoreOrNo() == true {
		deleteSpace()
		tampilMenuCari()
	} else {
		deleteSpace()
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
			spaceEquals()
			showData(i)
			spaceEquals()
			key = true
		}
	}

	if key == false {
		spaceNull()
	}

	if findMoreOrNo() == true {
		deleteSpace()
		tampilMenuCari()
	} else {
		deleteSpace()
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
			spaceEquals()
			showData(i)
			spaceEquals()
			key = true
		}
	}

	if key == false {
		spaceNull()
	}

	if findMoreOrNo() == true {
		deleteSpace()
		tampilMenuCari()
	} else {
		deleteSpace()
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
	deleteSpace()
	fmt.Print(" ", "NIM: ")
	fmt.Scanln(&nimS) //kata kunci data mana yang ingin diubah

	spacingData()
	fmt.Println(" ", "Ketik '-' jika tidak ingin mengubah data.")
	spacingData()
	fmt.Println("=======================================================")

	for i := 0; i < count; i++ { // digunakan looping pencarian data
		if mahasiswa[i].nim == nimS { // jika ada data yang sesuai kata kunci maka input akan dilakukan

			showData(i) //tampil data sesuai i
			spacingData()
			fmt.Println("=======================================================")
			spacingData()
			pos = i

			namaM = mahasiswa[i].nama         // data sebelumnya akan disimpan di variabel
			fakultasM = mahasiswa[i].fakultas // data sebelumnya akan disimpan di variabel
			jurusanM = mahasiswa[i].jurusan   // data sebelumnya akan disimpan di variabel
			kelasM = mahasiswa[i].kelas       // data sebelumnya akan disimpan di variabel
			nohpM = mahasiswa[i].nohp         // data sebelumnya akan disimpan di variabel
			emailM = mahasiswa[i].email       // data sebelumnya akan disimpan di variabel

			fmt.Print(" ", "Nama: ")
			namaS = inputWithSpacing() //input data dengan bufio agar data dapat menggunakan spasi
			fmt.Print(" ", "Fakultas: ")
			fakultasS = inputWithSpacing() //input data dengan bufio agar data dapat menggunakan spasi
			fmt.Print(" ", "Jurusan: ")
			jurusanS = inputWithSpacing() //input data dengan bufio agar data dapat menggunakan spasi
			fmt.Print(" ", "Kelas: ")
			kelasS = inputWithSpacing() //input data dengan bufio agar data dapat menggunakan spasi
			fmt.Print(" ", "No.HP: ")
			nohpS = inputWithSpacing() //input data dengan bufio agar data dapat menggunakan spasi
			fmt.Print(" ", "Email: ")
			emailS = inputWithSpacing() //input data dengan bufio agar data dapat menggunakan spasi
		}
	}

	spacingData()
	fmt.Println("=======================================================")

	if pos != -1 {
		if updateOrNo() == true { //cek apakah data benar ingin diperbarui
			mahasiswa[pos].nim = nimS

			mahasiswa[pos].nama = namaS
			if namaS == "-" { //jika input == - maka data tidak akan berubah, jika bukan maka akan sesuai yang di inputkan
				mahasiswa[pos].nama = namaM
			}

			mahasiswa[pos].fakultas = fakultasS
			if fakultasS == "-" { //jika input == - maka data tidak akan berubah, jika bukan maka akan sesuai yang di inputkan
				mahasiswa[pos].fakultas = fakultasM
			}

			mahasiswa[pos].jurusan = jurusanS
			if jurusanS == "-" { //jika input == - maka data tidak akan berubah, jika bukan maka akan sesuai yang di inputkan
				mahasiswa[pos].jurusan = jurusanM
			}

			mahasiswa[pos].kelas = kelasS
			if kelasS == "-" { //jika input == - maka data tidak akan berubah, jika bukan maka akan sesuai yang di inputkan
				mahasiswa[pos].kelas = kelasM
			}

			mahasiswa[pos].nohp = nohpS
			if nohpS == "-" { //jika input == - maka data tidak akan berubah, jika bukan maka akan sesuai yang di inputkan
				mahasiswa[pos].nohp = nohpM
			}

			mahasiswa[pos].email = emailS
			if emailS == "-" { //jika input == - maka data tidak akan berubah, jika bukan maka akan sesuai yang di inputkan
				mahasiswa[pos].email = emailM
			}
			deleteSpace()
			fmt.Println(" ", "Data berhasil diperbarui.")
			spacingData()
			tampilMenu()
		} else {
			deleteSpace()
			fmt.Println(" ", "Data gagal diperbarui.")
			tampilMenu()
		}
	} else {
		deleteSpace()
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

	deleteSpace()
	fmt.Print(" ", "NIM: ")
	fmt.Scanln(&nimS)            // kata kunci pencarian
	for i := 0; i < count; i++ { // melakukan looping untuk menentukan apakah data yang dicari ada
		if mahasiswa[i].nim == nimS {

			showData(i) // tampil data yang ingin dihapus
			key = true
		}
	}

	if key == true {
		if deleteOrNo() == true { // jika jawaban y maka data akan dihapus
			for j := 0; j < count; j++ {
				if mahasiswa[j].nim == nimS {
					mahasiswa[j] = mahasiswa[j+1]
				}
				if key == true {
					mahasiswa[j+1] = mahasiswa[j+2]
				}
			}
			deleteSpace()
			fmt.Println(" ", "Data berhasil dihapus.")
			spacingData()
			tampilMenu()
		} else {
			deleteSpace()
			fmt.Println(" ", "Data gagal diperbarui.")
			tampilMenu()
		}
	} else {
		deleteSpace()
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
		fmt.Print(" ", "Apakah anda ingin mencari data lagi? Y/N. ")
		fmt.Scanln(&yesOrNo)
		spacingData()
		kebenaran = yesOrNo == "Y" || yesOrNo == "y"

		if yesOrNo == "N" || yesOrNo == "n" {
			deleteSpace()
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
			deleteSpace()
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
			deleteSpace()
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

func countData() {
	for i := 0; i < n; i++ {
		if mahasiswa[i].nim == 0 {
			count = i
			break
		}
	}
}

func spaceNull() {
	spacingData()
	fmt.Println(" ", "Data tidak ditemukan.")
	spacingData()
}

func spaceEquals() {
	spacingData()
	fmt.Println("=======================================================")
}

func deleteSpace() {
	clsCode()
	spacingData()
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

// START OF EXPORT CODE

var path = "hasilExport.txt" //memberi lokasi untuk file txt

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func createFile() {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) { //jika file pada path tidak ada, maka akan dibuat file tsb
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}
}

func writeFile() {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644) //membaca file ada atau tidak
	if isError(err) {
		return
	}
	defer file.Close()
	for i := 0; i < count; i++ { //mengulang penulisan data sesuai count(banyak data)
		_, err = file.WriteString(strconv.Itoa(mahasiswa[i].nim)) //tulis data
		_, err = file.WriteString("\n")                           //spasi
		_, err = file.WriteString(mahasiswa[i].nama)              //tulis data
		_, err = file.WriteString("\n")                           //spasi
		_, err = file.WriteString(mahasiswa[i].fakultas)          //tulis data
		_, err = file.WriteString("\n")                           //spasi
		_, err = file.WriteString(mahasiswa[i].jurusan)           //tulis data
		_, err = file.WriteString("\n")                           //spasi
		_, err = file.WriteString(mahasiswa[i].kelas)             //tulis data
		_, err = file.WriteString("\n")                           //spasi
		_, err = file.WriteString(mahasiswa[i].nohp)              //tulis data
		_, err = file.WriteString("\n")                           //spasi
		_, err = file.WriteString(mahasiswa[i].email)             //tulis data
		_, err = file.WriteString("\n")                           //spasi
		_, err = file.WriteString("\n")                           //spasi
	}

	err = file.Sync()
	if isError(err) {
		return
	}
	fmt.Println(" ", "Berhasil membuat file txt.")
	spacingData()
	fmt.Println(" ", "Berhasil menulis data ke txt.")
}

func readFile() {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644) //membaca file ada atau tidak
	if isError(err) {
		return
	}
	defer file.Close()

	var text = make([]byte, 1024)
	for { //membaca isi file yang ada pada path
		n, err := file.Read(text)
		if err != io.EOF {
			if isError(err) {
				break
			}
		}
		if n == 0 {
			break
		}
	}
	if isError(err) {
		return
	}
	fmt.Println(string(text)) //menampilkan data pada path ke cmd
}

func deleteFile() {
	var err = os.Remove(path) //menghapus file txt
	if isError(err) {
		return
	}

	fmt.Println(" ", "File txt berhasil dihapus.")
}

// END OF EXPORT CODE

func main() {
	clsCode()
	countData()
	tampilMenu()
}
