package main
import "fmt"


//fungsi untuk menampilkan film berdsarkan  Judu;l, Genre, Tahun Rilis, dan Rating
func tampilFilm(data Films, i int) {
	var dataGenre GENRES
	var pilihan,y_n string
	var index, j, banyakGenre int
	var pages int
	UpScreen()
	coverTampilFilm()
	fmt.Scan(&pilihan)
	UpScreen()
	
	switch pilihan {
	case "1" :
		UpScreen()
		UrutJudulFilm(&data, i)
		ListJudulFilm(data, i)
		index = TampilFilmJudul(data, i)
		if index != -1 {
			UpScreen()
			fmt.Println("+-------------------------------+-----------------+-----------------+------------+")
			fmt.Printf("| %-29s | %-15s | %-15s | %-10s |\n", "JUDUL", "GENRE", "TAHUN RILIS", "RATING")
			fmt.Println("+-------------------------------+-----------------+-----------------+------------+")
			fmt.Printf("| %-29s | %-15s | %-15d | %-10.1f |\n", data[index].judul, data[index].Genre, data[index].TahunRilis, data[index].Rating)
			fmt.Println("+--------------------------------------------------------------------------------+")
			fmt.Printf("|Deskripsi: %-69s|\n| %-78s |\n", "", data[index].Deskripsi)
			fmt.Println("+--------------------------------------------------------------------------------+")
			fmt.Printf("sudah selesai mencari filmnya? (y/n): ")
			fmt.Scan(&pilihan)
			switch pilihan {
			case "y" :
				UpScreen()
			case "n" :
				UpScreen()
				tampilFilm(data, i)
			default :
				fmt.Println("Pilihan tidak valid")
				UpScreen()
				tampilFilm(data, i)
			}
		} else {
			UpScreen()
			fmt.Println("+---------------------------------------+")
			fmt.Println("| Film yang kamu cari belum ada nihh :( |")
			fmt.Println("| Mau coba film lain?                   |")
			fmt.Println("+---------------------------------------+")
			fmt.Printf("1. Cari Film lain || 2. Kembali ke Menu Utama\nMasukan pilihan: ")
			fmt.Scan(&pilihan)
			switch pilihan {
			case "1" :
				UpScreen()
				tampilFilm(data, i)
			case "2" :
				UpScreen()
			default :
				fmt.Println("Pilihan tidak valid")
				tampilFilm(data, i)
			}
		}

	case "2" :
		var PGS int = 1
		UpScreen()
		coverGenre()
		BanyakGenre(data, i, &dataGenre, &banyakGenre)
		pages = page(banyakGenre, 5)
		for j = 0; j < banyakGenre; j++ {
			fmt.Printf("| %d. %-29s |\n", j+1, dataGenre[j].Genre)
			if (j % 5 == 4) || (j == banyakGenre-1) {
				fmt.Println("+----------------------------------+")
				fmt.Printf("[%d/%d]Input 'y' untuk melanjutkan: ", PGS, pages)
				fmt.Scan(&y_n)
				if y_n == "y" && j != banyakGenre-1 {
					UpScreen()
					coverGenre()
					PGS++
				} else if y_n == "y" && j == banyakGenre-1 {
					UpScreen()
					coverGenre()
					j = -1
					PGS = 1
				} else {
					j = banyakGenre
				}
			}
		}
		fmt.Println("+----------------------------------+")
		TampilFilmGenre(data)
		fmt.Println("+---------------------------------------------+")
		fmt.Printf("sudah selesai mencari filmnya? (y/n): ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case "y" :
			UpScreen()
		case "n" :
			UpScreen()
			tampilFilm(data, i)
		default :
			fmt.Println("Pilihan tidak valid")
			UpScreen()
			tampilFilm(data, i)
		}
	case "3" :
		var S1 string
		var PGS int = 1
		pages = page(i, 5)
		UpScreen()
		coverAscDesc()
		fmt.Scan(&pilihan)
		if pilihan == "1" {
			AscendingTahunFilm(&data, i)
			S1 = "Ascending"
		} else if pilihan == "2" {
			DescendingTahunFilm(&data, i)
			S1 = "Descending"
		}
		UpScreen()
		CoverRatingTahun(("Judul " + S1), "Tahun Rilis")
		for j = 0; j < i; j++ {
			fmt.Printf("| %-3d | %-29s | %-12d |\n", j+1, data[j].judul, data[j].TahunRilis)
				if (j % 5 == 4)	|| (j == i-1) {
				fmt.Println("+----------------------------------------------------+")
				fmt.Printf("[%d/%d]Input 'y' untuk melanjutkan: ", PGS, pages)
				fmt.Scan(&y_n)
				if y_n == "y" && j != i-1 {
					UpScreen()
					CoverRatingTahun("Judul " + S1, "Tahun Rilis")
					PGS++
				}  else if y_n == "y" && j == i-1 {
					UpScreen()
					CoverRatingTahun("Judul " + S1, "Tahun Rilis")
					j = -1
					PGS = 1
				} else {
					UpScreen()
					j = i
				}
			}
		}
	case "4" :
		var S1 string
		var PGS int = 1
		pages = page(i, 5)
		UpScreen()
		coverAscDesc()
		fmt.Scan(&pilihan)
		if pilihan == "1" {
			AscendingRatingFilm(&data, i)
			S1 = "Ascending"
		} else if pilihan == "2" {
			DescendingRatingFilm(&data, i)
			S1 = "Descending"
		}
		UpScreen()
		CoverRatingTahun(("Judul " + S1), "Rating")
		for j = 0; j < i; j++ {
			fmt.Printf("| %-3d | %-29s | %-12.1f |\n", j+1, data[j].judul, data[j].Rating)
			if (j % 5 == 4)	|| (j == i-1) {
				fmt.Println("+----------------------------------------------------+")
				fmt.Printf("[%d/%d]Input 'y' untuk melanjutkan: ", PGS, pages)
				fmt.Scan(&y_n)
				if y_n == "y" && j != i-1 {
					UpScreen()
					CoverRatingTahun(("Judul " + S1) , "Rating")
					PGS++
				} else if y_n == "y" && j == i-1 {
					UpScreen()
					CoverRatingTahun(("Judul " + S1), "Rating")
					j = -1
					PGS = 1
				} else {
					UpScreen()
					j = i
				}
			}
		}
	case "5" :
		UpScreen()
	default :
		fmt.Println("Pilihan tidak valid")
		tampilFilm(data, i)
	}
}

func ListJudulFilm(data Films, banyakfilm int) {
	var pilihan string
	var i, pages,PGS int
	PGS = 1
	pages = page(banyakfilm, 5)
	CoverListFilm("JUDUL FILM", "DESKRIPSI")
	for i = 0; i < banyakfilm; i++ {
		fmt.Printf("| %-3d| %-29s| %-65s|\n", i+1, data[i].judul, data[i].Deskripsi)
		if (i % 5 == 4) || (i == banyakfilm-1) {
			fmt.Println("+----+------------------------------+------------------------------------------------------------------+")
			fmt.Printf("|Kamu belum dapet yang dicari?!%-67s[%d/%d]|\n", "", PGS, pages)
			fmt.Printf("|Insert apapun jika sudah dapat yang kamu cari :3%-54s|\n", "")
			fmt.Println("+------------------------------------------------------------------------------------------------------+")
			fmt.Print("NEXT PAGES INSERT 'y' : ")
			fmt.Scan(&pilihan)
			if pilihan == "y" && i != banyakfilm-1 {
				UpScreen()
				CoverListFilm("JUDUL FILM", "DESKRIPSI")
				PGS++
			} else if pilihan == "y" && i == banyakfilm-1 {
				UpScreen()
				CoverListFilm("JUDUL FILM", "DESKRIPSI")
				i = -1
				PGS = 1
			} else {
				i = banyakfilm
			}
		}
	}
}
