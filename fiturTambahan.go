package main
import "fmt"


func page(banyakfilm int, N int) int {
	var pages int
	if banyakfilm % N == 0 {
		pages = banyakfilm / N
	} else {
		pages = banyakfilm / N + 1
	}
	return pages
}	

func UpScreen() {
    fmt.Print("\033[H\033[2J")
}


func CoverRatingTahun(S1, S2 string) {
	fmt.Println("+----------------------------------------------------+")
		fmt.Printf("| %-3s | %-29s | %-12s |\n", "NO", S1, S2)
		fmt.Println("+----------------------------------------------------+")
}
func coverAscDesc() {
		fmt.Println("+--------------------------------+")
		fmt.Println("|      ASCENDING/DESCENDING      |")
		fmt.Printf("+--------------pilihan-----------+\n| 1. ASCENDING	 	 	 |\n| 2. DESCENDING 		 |\n")
		fmt.Println("+--------------------------------+")
		fmt.Print("Pilih Menu (1-2): ")
}

func coverTampilFilm() {
	fmt.Println("+----------------------------------+")
	fmt.Println("|         MENAMPILKAN FILM         |")
	fmt.Println("+----------------------------------+")
	fmt.Printf("| 1.  Judul		 	   |\n| 2.  Genre		 	   |\n| 3.  Berdasarkan Tahun Rilis      |\n| 4.  Berdasarkan Rating	   |\n| 5.  Kembali ke Menu Utama  	   |\n")
	fmt.Println("+----------------------------------+")
	fmt.Print("Pilih Menu (1-5): ")
}

func coverGenre() {
	fmt.Println("+----------------------------------+")
	fmt.Println("|         GENRE TERSEDIA           |")
	fmt.Println("+----------------------------------+")
}

func CoverListFilm(S1, S2 string) {
	fmt.Println("+----+------------------------------+------------------------------------------------------------------+")
	fmt.Printf("| %-3s| %-29s| %-65s|\n", "NO", S1, S2)
	fmt.Println("+----+------------------------------+------------------------------------------------------------------+")
}


func cover1Text(S1 string) {
	fmt.Println("+----------------------------------+")
	fmt.Printf("| %-32s |\n", S1)
	fmt.Println("+----------------------------------+")
}

func cover1TextLong(S1 string) {
	fmt.Println("+----------------------------------------------------+")
	fmt.Printf("| %-50s |\n", S1)
	fmt.Println("+----------------------------------------------------+")
}

func CoverExit() {
	fmt.Println("+--------------------------------+")
	fmt.Println("| TERIMA KASIH SUDAH             |")
	fmt.Println("|       MENGGUNAKAN CINEMA GWEH! |")
	fmt.Println("+--------------------------------+")
}