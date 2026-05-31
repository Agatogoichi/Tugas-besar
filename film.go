package main
import "fmt"

//=====================MAIN PROGRAM NO 1 = TAMPIL =================================================================
//fungsi untuk menampilkan film berdsarkan  Judu;l, Genre, Tahun Rilis, dan Rating
func tampilFilm(data Films, i int) {
	var dataGenre GENRES
	var pilihan string
	var index, j, banyakGenre int
	fmt.Printf("\nMenampilkan film berdasarkan : \n1. Judul \n2. Genre \n3. Berdasarkan Tahun Rilis \n4. Berdasarkan Rating \nMasukan pilihan: ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case "1" :
		MenampilkanJudulFilm(data, i)
		UrutJudulFilm(&data, i)
		index = TampilFilmJudul(data, i)
		if index != -1 {
			fmt.Printf("\nJudul: %s\nTahun Rilis: %d\nDeskripsi: %s\nGenre: %s\nRating: %.1f\n", data[index].judul, data[index].TahunRilis, data[index].Deskripsi, data[index].Genre, data[index].Rating)
		} else {
			fmt.Println("Film tidak ditemukan")
		}
	case "2" :
		fmt.Println("==================================")
		BanyakGenre(data, i, &dataGenre, &banyakGenre)
		fmt.Println("Genre yang tersedia : ")
		for j = 0; j < banyakGenre; j++ {
			fmt.Printf("- %s\n", dataGenre[j].Genre)
		}
		TampilFilmGenre(data)
	case "3" :
		fmt.Println("==================================")
		fmt.Printf("Urutkan secara : \n1. Ascending\n2. Descending\nMasukan pilihan: ")
		fmt.Scan(&pilihan)
		if pilihan == "1" {
		AscendingTahunFilm(&data, i)
		} else if pilihan == "2" {
		DescendingTahunFilm(&data, i)
		}
		fmt.Println("=================================================")
		fmt.Printf("%-10s | %-20s | %-6s\n", "NO", "Judul", "Tahun Rilis")
		fmt.Println("=================================================")
		for j = 0; j < i; j++ {
			fmt.Printf("%-10d | %-20s | %-6d\n", j+1, data[j].judul, data[j].TahunRilis)
		}
	case "4" :
		fmt.Println("==================================")
		fmt.Printf("Urutkan secara : \n1. Ascending\n2. Descending\nMasukan pilihan: ")
		fmt.Scan(&pilihan)
		if pilihan == "1" {
			AscendingRatingFilm(&data, i)
		} else if pilihan == "2" {
			DescendingRatingFilm(&data, i)
		}
		fmt.Println("================================================")
		fmt.Printf("%-10s | %-20s | %-6s\n", "NO", "Judul", "Rating")
		fmt.Println("================================================")
		for j = 0; j < i; j++ {
			fmt.Printf("%-10d | %-20s | %-6.1f\n", j+1, data[j].judul, data[j].Rating)
		}
	}
}


//=====================MAIN PROGRAM NO 2 = INPUT =================================================================
//fungsi untuk melakukan input film baru
func inputFilm(data *Films, i *int) {
	var pilihan string
	var stop bool = false
	for stop == false {
		fmt.Println("Input judul film: ")
		fmt.Scan(&data[*i].judul)
		fmt.Println("Input tahun rilis film: ")
		fmt.Scan(&data[*i].TahunRilis)
		fmt.Println("Input deskripsi film: ")
		fmt.Scan(&data[*i].Deskripsi)
		fmt.Println("Input genre film: ")
		fmt.Scan(&data[*i].Genre)
		fmt.Println("Input rating film(0.0-5.0): ")
		fmt.Scan(&data[*i].Rating)
		*i++
		fmt.Println("Apakah ingin input film lagi? (y/n): ")
		fmt.Scan(&pilihan)
		if pilihan == "n" {
			stop = true
		}
	}
}
//=====================MAIN PROGRAM NO 3 = HAPUS =================================================================
func hapusFilm(data *Films, banyakfilm *int) {
	var JudulDicari string
	var find bool = false
	var j int
	var i int = 0
	fmt.Println("Masukkan judul film yang ingin dihapus: ")
	fmt.Scan(&JudulDicari)
	for i < *banyakfilm && !find {
		if data[i].judul == JudulDicari {
			find = true
		} else {
			i++
		}
	}
	if find {
		for j = i; j < *banyakfilm; j++ {
			data[j] = data[j+1]
		}
		*banyakfilm--
		fmt.Println("Film berhasil dihapus")
	}
}

//=====================MAIN PROGRAM NO 4 = UBAH =================================================================
//prosedur untuk mengubah data film menggunakan teknik sequential search untuk mencari film yang ingin diubah
func ubahFilm(data *Films) {
	var pilihan, masukanJudul, judulBaru, deskripsiBaru, genreBaru string
	var tahunRilisBaru int
	var ratingBaru float64
	var find bool = false
	var jumlahFilm int = len(data)
	var i int = 0

	fmt.Println("\nMasukkan judul film yang ingin diubah: ")
	fmt.Scan(&masukanJudul)
	for i < jumlahFilm && !find {
		if data[i].judul == masukanJudul {
			find = true
		} else {
			i++
		}
	}
	if find {
		fmt.Printf("\nMasukkan informasi baru untuk film %s:\n1. Judul\n2. Tahun Rilis\n3. Deskripsi\n4. Genre\n5. Rating\nMasukan menu: ", data[i].judul)
		fmt.Scan(&pilihan)
		switch pilihan {
			case "1" :
			fmt.Print("Masukkan judul baru: ")
			fmt.Scan(&judulBaru)
			data[i].judul = judulBaru
			case "2" :
			fmt.Print("Masukkan tahun rilis baru: ")
			fmt.Scan(&tahunRilisBaru)
			data[i].TahunRilis = tahunRilisBaru
			case "3" :
			fmt.Print("Masukkan deskripsi baru: ")
			fmt.Scan(&deskripsiBaru)
			data[i].Deskripsi = deskripsiBaru
			case "4" :
			fmt.Print("Masukkan genre baru: ")
			fmt.Scan(&genreBaru)
			data[i].Genre = genreBaru
			case "5" :
			fmt.Print("Masukkan rating baru: ")
			fmt.Scan(&ratingBaru)
			data[i].Rating = ratingBaru
			default :
			fmt.Println("Pilihan tidak valid")
		}
	} else {
		fmt.Println("Film tidak ditemukan")
	}
}
//==========================MAIN PROGRAM NO 5 = STATISTIK=========================================================
func Statistik(data Films, banyakfilm int) {
	var dataGenre GENRES
	var i, j int
	fmt.Println("====================Statistik Cinema====================")
	BanyakGenre(data, banyakfilm, &dataGenre, &j)
	fmt.Println("Banyak Film per Genre: ")
	for i = 0; i < j; i++ {
		fmt.Printf("- %-20s : %d film\n", dataGenre[i].Genre, dataGenre[i].Count)
	}
	fmt.Println("========================================================")
	fmt.Printf("Rata-rata rating film: %.1f\n", rataRating(data, banyakfilm))
}
