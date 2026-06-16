package main
import "fmt"


func TampilFilmGenre(data Films) {
var Genre string
var i int
var Find bool = false
	fmt.Printf("Mw cari film genre apa nih? :  ")
	fmt.Scan(&Genre)
	UpScreen()
	for i = 0; i < NMAX; i++ {
		if data[i].Genre == Genre {Find = true}
	}
	if Find {
		fmt.Println("+---------------------------------------------+")
		fmt.Printf("| %-43s |\n", "Film dengan Genre: "+Genre)
		fmt.Println("+---------------------------------------------+")
	}
	for i = 0; i < NMAX; i++ {
		if data[i].Genre == Genre {
			fmt.Printf("| %-29s | Rating: %.1f |\n", data[i].judul, data[i].Rating)
		}
	}
	if !Find {
		fmt.Printf("Genre yang kamu cari belum ada :3\n")
	}
}

//function mencari sebuah judul menggunakanan teknik binary search, akan menampilkan informasi film jika ditemukan dan menampilkan pesan jika tidak ditemukan
func TampilFilmJudul(data Films, i int) int {	
var judul string
var low, high, mid int
	fmt.Printf("Masukan judul film yang mau ditampilkan :  ")
	fmt.Scan(&judul)
	low = 0
	high = i - 1
	for low <= high {
		mid = (low + high) / 2
		if data[mid].judul == judul {
			return mid
		} else if data[mid].judul < judul {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}