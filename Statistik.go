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
//program untuk memasukan genre
func BanyakGenre(data Films, banyakfilm int, dataGenre *GENRES, genreUnik *int) { 
	//dataGenre adalah untuk menyimpan genre yang ada didalam array genre serta melakukan penjumlahan untuk setiap genre unik
	//GenreUnik adalah banyaknya genre unik yang ada pada array utama (fungsi mirip seperti banyakfilm)
	var GenreNow string
	var found bool
	var i, j int
	*genreUnik = 0 //genre unik dimulai dari 0 
	for i = 0; i < banyakfilm; i++ {
		GenreNow = data[i].Genre //memasukan genre yang ingin dicari duplkatnya (pada awalnya mulai akan false di percabbangan jadi otomatis masuk genre unik) 
		found = false 
		j = 0
		for j < *genreUnik && !found { //membandingkan di array genre
			if dataGenre[j].Genre == GenreNow {
				found = true //ketika masuk found akan true dan count pada genre akan bertambah
				dataGenre[j].Count++ 
			}
			j++
		}
		if !found { //ketika data genreNow tidak ditemukan duplikatnya maka akan masuk ke genre Unik dan dimasukan ke array genre
			dataGenre[*genreUnik].Genre = GenreNow
			dataGenre[*genreUnik].Count = 1 //melakuan initial bahwa Count dimulai dari 1 karena sudah ada genre itu sendiri
			*genreUnik++ //banyak genre unik bertambah
		}
	}
}
//program menghitung rating keseluruh film
func rataRating(data Films, banyakfilm int) float64 {
	var total float64
	var i int
	for i = 0; i < banyakfilm; i++ {
		total += data[i].Rating
	}
	return total / float64(banyakfilm)
}