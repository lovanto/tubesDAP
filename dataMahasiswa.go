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
	deleteSpace()
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

var path = "hasilExport.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func createFile() {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}
}

func writeFile() {
	// buka file dengan level akses READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()
	for i := 0; i < 3; i++ {
		_, err = file.WriteString(strconv.Itoa(mahasiswa[i].nim))
		_, err = file.WriteString("\n")
		_, err = file.WriteString(mahasiswa[i].nama)
		_, err = file.WriteString("\n")
		_, err = file.WriteString(mahasiswa[i].fakultas)
		_, err = file.WriteString("\n")
		_, err = file.WriteString(mahasiswa[i].jurusan)
		_, err = file.WriteString("\n")
		_, err = file.WriteString(mahasiswa[i].kelas)
		_, err = file.WriteString("\n")
		_, err = file.WriteString(mahasiswa[i].nohp)
		_, err = file.WriteString("\n")
		_, err = file.WriteString(mahasiswa[i].email)
		_, err = file.WriteString("\n")
		_, err = file.WriteString("\n")
	}

	// simpan perubahan
	err = file.Sync()
	if isError(err) {
		return
	}
	fmt.Println(" ", "Berhasil membuat file txt.")
	spacingData()
	fmt.Println(" ", "Berhasil menulis data ke txt.")
}

func readFile() {
	// buka file
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// baca file
	var text = make([]byte, 1024)
	for {
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
	fmt.Println(string(text))
}

func deleteFile() {
	var err = os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println(" ", "File txt berhasil dihapus.")
}

// END OF EXPORT CODE

func tempData() {
	mahasiswa[0].nim = 1302194068
	mahasiswa[0].nama = "Sir. Rifky Lovanto"
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
	mahasiswa[2].kelas = "SE-43-01"
	mahasiswa[2].nohp = "+62-878-2383-7556"
	mahasiswa[2].email = "akuBukanDia@student.telkomuniversity.ac.id"
}

func main() {
	clsCode()
	tempData()
	countData()
	tampilMenu()
}
