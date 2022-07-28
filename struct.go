package main

import "fmt"

// type Student struct {
// 	Name string
// 	Alamat string
// 	Nim int
// }

// func (student Student) newstudent(){
// 	fmt.Printf("Halo aku %s aku adalah murid baru \n", student.Name)
// }

// func main(){
// 	var student Student
// 	student.Name = "Doni"
// 	student.Nim = 101010
// 	student.Alamat = "Jakarta"
// 	fmt.Println(student)
// 	fmt.Printf("Nama saya %s, nim saya %d, alamat saya %s \n", student.Name, student.Nim, student.Alamat)

// 	doni := Student{Name: "doni"}
// 	doni.newstudent()
// }

// KUIS 1 (KUIS DATA PENDUDUK)
// Tugas kali ini adalah kamu diminta untuk membuat sebuah struct bernama “Penduduk” yang mempunya field sebagai berikut :
// Nama (string)
// Umur (int)
// Alamat (string)

// Setelah itu, tampilkanlah pada console struct yang sudah dibuat dengan format:
// “Hello, my name is (nama). Im (umur) years old. I live in (alamat)”

//Tulis kode struct disini
// type Penduduk struct {
//     nama string
//     umur int
//     alamat string
// }

// func main() {
//     var penduduk1 Penduduk
//     scanner := bufio.NewScanner(os.Stdin)
//     scanner.Scan()
//     penduduk1.nama = scanner.Text()
//     scanner.Scan()
//     penduduk1.umur, _ = strconv.Atoi(scanner.Text())
//     scanner.Scan()
//     penduduk1.alamat = scanner.Text()

//     //Tulis kode format disini
//     fmt.Printf("Hello, my name is %s. Im %d years old. I live in %s\n", penduduk1.nama, penduduk1.umur, penduduk1.alamat)
// }

// KUIS 2 (KUIS PERKENALKAN DIRI)
// Buatlah struct method dari struct student yang sudah disediakan dimana struct method tersebut bernama "Introduction" dan jika dipanggil maka akan menampilkan pada console berupa perkenalan diri dengan format :
// "Hello, my name is (name). Im (age) years old"

type Student struct {
 // Isi struct ini
    name string
    age int
}

func (s *Student) Introduction() {
   // Tulis kode disini
    fmt.Printf("Hello, my name is %s. Im %d years old\n", s.name, s.age)
}

func main() {
 // Tulis kode disini
    var student1 Student
    student1.name = "Sammy"
    student1.age = 17
  	student1.Introduction()
}