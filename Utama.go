package main
import "fmt"

const NMAX = 9999
type Film struct {
	judul string
	TahunRilis int
	Deskripsi string
	Genre string
	Rating float64
}
type Films [NMAX]Film

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
		fmt.Println("=====================================")
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
			hapusFilm(&Data, i-1)
			case 5 :
			fmt.Println("Statistik Cinema")
			case 6 :
			stop = true
			default :
			fmt.Println("TIdak ada pilihan")
			}
		}
		fmt.Println(i)
}

//INPUT & DELETE & UPDATE////////////////////////////////////////////////////////////////////////////////
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
//Sistem untuk menghapus film menggunakan teknik sequential search untuk mencari film yang ingin dihapus, jika ditemukan maka data film tersebut akan dihapus dengan cara menggeser semua data setelahnya ke kiri
func hapusFilm(data *Films, banyakfilm int) {
	var masukanJudul string
	var find bool = false
	var j int
	var i int = 0
	fmt.Println("Masukkan judul film yang ingin dihapus: ")
	fmt.Scan(&masukanJudul)
	for i < banyakfilm && !find {
		if data[i].judul == masukanJudul {
			find = true
		} else {
			i++
		}
	}
	if find {
		for j = i; j < banyakfilm; j++ {
			data[j] = data[j+1]
		}
	}
}

//fungsi untuk menampilkan film berdasarkan genre yang diinputkan
func tampilFilm(data Films, i int) {
	var pilihan string
	var index, j int
	fmt.Printf("\nMenampilkan film berdasarkan : \n1. Judul \n2. Genre \n3.Berdasarkan Tahun Rilis \n4. Berdasarkan Rating \nMasukan pilihan: ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case "1" :
		UrutJudulFilm(&data, i)
		index = TampilFilmJudul(data, i)
		if index != -1 {
			fmt.Printf("\nJudul: %s\nTahun Rilis: %d\nDeskripsi: %s\nGenre: %s\nRating: %.1f\n", data[index].judul, data[index].TahunRilis, data[index].Deskripsi, data[index].Genre, data[index].Rating)
		} else {
			fmt.Println("Film tidak ditemukan")
		}
	case "2" :
		TampilFilmGenre(data)
	case "3" :
		tahunRilisFilm(&data, i)
		fmt.Println("=======================================")
		fmt.Printf("%-10s | %-20s | %-6s\n", "NO", "Judul", "Tahun Rilis")
		fmt.Println("=======================================")
		for j = 0; j < i; j++ {
			fmt.Printf("%-10d | %-20s | %-6d\n", j+1, data[j].judul, data[j].TahunRilis)
		}
	case "4" :
		ratingFilm(&data, i)
		fmt.Println("=======================================")
		fmt.Printf("%-10s | %-20s | %-6s\n", "NO", "Judul", "Rating")
		fmt.Println("=======================================")
		for j = 0; j < i; j++ {
			fmt.Printf("%-10d | %-20s | %-6.1f\n", j+1, data[j].judul, data[j].Rating)
		}
	}
}
//prosedur untuk mengubah data film menggunakan teknik sequential search untuk mencari film yang ingin diubah
func ubahFilm(data *Films) {
	var pilihan, masukanJudul, judulBaru, deskripsiBaru, genreBaru string
	var tahunRilisBaru int
	var ratingBaru float64
	var find bool = false
	var jumlahFilm int = len(data)
	var i int = 0

	fmt.Println("Masukkan judul film yang ingin diubah: ")
	fmt.Scan(&masukanJudul)
	for i < jumlahFilm && !find {
		if data[i].judul == masukanJudul {
			find = true
		} else {
			i++
		}
	}
	if find {
		fmt.Println("Masukkan informasi baru untuk film ini:\n1. Judul\n2. Tahun Rilis\n3. Deskripsi\n4. Genre\n5. Rating")
		fmt.Scan(&pilihan)
		switch pilihan {
			case "1" :
			fmt.Println("Masukkan judul baru: ")
			fmt.Scan(&judulBaru)
			data[i].judul = judulBaru
			case "2" :
			fmt.Println("Masukkan tahun rilis baru: ")
			fmt.Scan(&tahunRilisBaru)
			data[i].TahunRilis = tahunRilisBaru
			case "3" :
			fmt.Println("Masukkan deskripsi baru: ")
			fmt.Scan(&deskripsiBaru)
			data[i].Deskripsi = deskripsiBaru
			case "4" :
			fmt.Println("Masukkan genre baru: ")
			fmt.Scan(&genreBaru)
			data[i].Genre = genreBaru
			case "5" :
			fmt.Println("Masukkan rating baru: ")
			fmt.Scan(&ratingBaru)
			data[i].Rating = ratingBaru
			default :
			fmt.Println("Pilihan tidak valid")
		}
	} else {
		fmt.Println("Film tidak ditemukan")
	}
}
//SEARCH///////////////////////////////////////////////////////////////////////////
//Fungsi untuk Menampilkan film berdasarkan genre yang diinputkan menggunakan teknik sequential search
//akan menampilkan semua film yang memiliki genre yang dicari oleh pengguna atau Use
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
		fmt.Println("tidak ada film dengan genre tersebut")
	}
}

//function mencari sebuah judul menggunakanan teknik binary search, akan menampilkan informasi film jika ditemukan dan menampilkan pesan jika tidak ditemukan
func TampilFilmJudul(data Films, i int) int {	
var judul string
var low, high, mid int
	fmt.Printf("\nMasukan judul film :  ")
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

///SORTING////////////////////////////////////////////////////////////////////////////////
//funsi untuk mengurutkan film berdasarkan nama menggunakan teknik selection sort
func UrutJudulFilm(data *Films, banyakfilm int) {
	var temp Film
	var Mindex, i, j int
	for i = 0; i < banyakfilm-1; i++ {
		Mindex = i
		for j = i + 1; j < banyakfilm; j++ {
			if data[j].judul < data[Mindex].judul {
				Mindex = j
			}
		}
		temp = data[i]
		data[i] = data[Mindex]
		data[Mindex] = temp
	}
}

//prosedur untuk mengurutkan rating film menggunakan teknik insertion sort
func ratingFilm(data *Films, banyakfilm int) {
	var temp Film
	var i, j int
	for i = 1; i < banyakfilm; i++ {
		temp = data[i]
		j = i - 1
		for j >= 0 && data[j].Rating < temp.Rating {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}
//prosedur untuk mengurutkan tahun rilis film menggunakan teknik insertion sort
func tahunRilisFilm(data *Films, banyakfilm int) {
	var temp Film
	var i, j int
	for i = 1; i < banyakfilm; i++ {
		temp = data[i]
		j = i - 1
		for j >= 0 && data[j].TahunRilis < temp.TahunRilis {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}

//STATISTIK////////////////////////////////////////////////////////////////////////////////
func BanyakGenre(data Films, banyakfilm int) {
	var genres [NMAX]string
	var count [NMAX]int
	var GenreUnik int = 0
	var GenreNow string
	var found bool
	var i, j int
	for i = 0; i < banyakfilm; i++ {
		GenreNow = data[i].Genre
		found = false
		j = 0
		for j < GenreUnik && !found {
			if genres[j] == GenreNow {
				found = true
				count[j]++
			}
		}
		if !found {
			genres[GenreUnik] = GenreNow
			count[GenreUnik] = 1
			GenreUnik++
		}
	}
	fmt.Println("Statistik Genre Film:")
	for i = 0; i < GenreUnik; i++ {
		fmt.Printf("- %s: %d film\n", genres[i], count[i])
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
