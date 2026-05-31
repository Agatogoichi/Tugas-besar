package main
import "fmt"


func TampilFilmGenre(data Films) {
var Genre string
var i int
var Find bool = false
	fmt.Printf("\nMenampilkan film berdasarkan genre: ")
	fmt.Scan(&Genre)
	for i = 0; i < NMAX; i++ {
		if data[i].Genre == Genre {
			fmt.Printf("\nJudul: %s\nTahun Rilis: %d\nDeskripsi: %s\nGenre: %s\nRating: %.1f\n", data[i].judul, data[i].TahunRilis, data[i].Deskripsi, data[i].Genre, data[i].Rating)
			Find = true
		}
	}
	if !Find {
		fmt.Printf("tidak ada film dengan genre tersebut\n")
	}
}

//function mencari sebuah judul menggunakanan teknik binary search, akan menampilkan informasi film jika ditemukan dan menampilkan pesan jika tidak ditemukan
func TampilFilmJudul(data Films, i int) int {	
var judul string
var low, high, mid int
	fmt.Printf("\nMasukan judul film yang mau ditampilkan :  ")
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