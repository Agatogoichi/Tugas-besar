package main
import "fmt"

func main() {
	var Data Films
	var Kondisi int
	var stop bool = false
	var i int = 5
	Data[0] = Film{"One_Piece", 1994, "Menceritakan_Perjalanan_Luffy_dan_Krunya_Mencari_One_Piece", "Action", 4.0}
	Data[1] = Film{"Naruto", 2002, "Menceritakan_Perjalanan_Naruto", "Action", 4.5}
	Data[2] = Film{"Doraemon", 1973, "Menceritakan_Keseharian_Nobita_dan_Doraemon", "Comedy", 4.8}
	Data[3] = Film{"JunjiIto", 2018, "Menceritakan_Kumpulan_Cerita_Horor_Junji_Ito", "Horror", 4.2}
	Data[4] = Film{"My_Hero_Academia", 2014, "Perjalanan_Deku_Menjadi_Pahlawan_Terbaik", "Action", 4.7}

	fmt.Printf("=================Menu=================\n1. Tampilan film\n2. Input film\n3. Ubah film\n4. Hapus film\n5. Statistik \n6. Keluar\n")
	for stop == false  {
		fmt.Println("\n=====================================")
		fmt.Println("Pilih Menu (1-6): ")
		fmt.Scan(&Kondisi)
		switch Kondisi {
			case 1 :
			tampilFilm(Data, i)
		    case 2 :
			inputFilm(&Data, &i)
			case 3 :
			ubahFilm(&Data)
			case 4 :
			hapusFilm(&Data, &i)
			case 5 :
			Statistik(Data, i)
			case 6 :
			stop = true
			default :
			fmt.Println("TIdak ada pilihan")
			}
		}
}

//STATISTIK////////////////////////////////////////////////////////////////////////////////
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