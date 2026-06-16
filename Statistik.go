package main
import "fmt"
//==========================MAIN PROGRAM NO 5 = STATISTIK=========================================================
func Statistik(data Films, banyakfilm int) {
	var dataGenre GENRES
	var i, j int
	var pilihan string
	fmt.Println("+---------------------------------------------+")
	fmt.Printf("| %-43s |\n", "STATISTIK FILM")
	fmt.Println("+----------------------+----------------------+")
	fmt.Printf("| %-20s | %-20s |\n", "GENRE", "JUMLAH FILM")
	fmt.Println("+----------------------+----------------------+")
	BanyakGenre(data, banyakfilm, &dataGenre, &j)
	for i = 0; i < j; i++ {
		fmt.Printf("| %-20s | %-20d |\n", dataGenre[i].Genre, dataGenre[i].Count)
	}
	fmt.Println("+----------------------+----------------------+")
	fmt.Printf("| %-20s | %-20d |\n", "Banyak Film", banyakfilm)
	fmt.Printf("| %-20s | %-20.1f |\n", "Rata-rata Rating", rataRating(data, banyakfilm))
	fmt.Println("+----------------------+----------------------+")
	fmt.Printf("Insert apapun untuk kembali ke menu utama: ")

	fmt.Scan(&pilihan)
	if pilihan == " " {
		UpScreen()
	} else {
		UpScreen()
	}
}
func BanyakGenre(data Films, banyakfilm int, dataGenre *GENRES, genreUnik *int) {
	var GenreNow string
	var found bool
	var i, j int
	*genreUnik = 0
	for i = 0; i < banyakfilm; i++ {
		GenreNow = data[i].Genre
		found = false
		j = 0
		for j < *genreUnik && !found {
			if dataGenre[j].Genre == GenreNow {
				found = true
				dataGenre[j].Count++
			}
			j++
		}
		if !found {
			dataGenre[*genreUnik].Genre = GenreNow
			dataGenre[*genreUnik].Count = 1
			*genreUnik++
		}
	}
}

func rataRating(data Films, banyakfilm int) float64 {
	var total float64
	var i int
	for i = 0; i < banyakfilm; i++ {
		total += data[i].Rating
	}
	return total / float64(banyakfilm)
}