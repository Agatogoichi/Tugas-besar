package main
import "fmt"

func UpdateArrayFilm(data *Films, banyakfilm *int) {
	var pilihan string
	UpScreen()
	fmt.Println("+----------------------------------+")
	fmt.Println("|         INPUT/UPDATE FILM        |")
	fmt.Println("+----------------------------------+")
	fmt.Printf("| 1.  Input film baru		   |\n| 2.  Update film		   |\n| 3.  Hapus film	  	   |\n| 4.  Kembali ke Menu Utama  	   |\n")
	fmt.Println("+----------------------------------+")
	fmt.Print("Pilih Menu (1-4): ")
	fmt.Scan(&pilihan)
	UpScreen()
	switch pilihan {
	case "1" :
		inputFilm(data, banyakfilm)
	case "2" :
		ubahFilm(data)
	case "3" :
		hapusFilm(data, banyakfilm)
	case "4" :
		UpScreen()
	default :
		fmt.Println("Pilihan tidak valid")
	}
}

//fungsi untuk melakukan input film baru
func inputFilm(data *Films, i *int) {
	var pilihan string
	var stop bool = false
	for stop == false {
		UpScreen()
		cover1Text("INPUT FILM")
		fmt.Print("Judul: ")
		fmt.Scan(&data[*i].judul)
		fmt.Print("Genre: ")
		fmt.Scan(&data[*i].Genre)
		fmt.Print("Tahun rilis: ")
		fmt.Scan(&data[*i].TahunRilis)
		fmt.Println("Deskripsi Singkat: ")
		fmt.Scan(&data[*i].Deskripsi)
		fmt.Print("Rating film(0.0-5.0): ")
		fmt.Scan(&data[*i].Rating)
		for data[*i].Rating < 0.0 || data[*i].Rating > 5.0 {
			fmt.Println("Rating tidak valid")
			fmt.Print("Input rating film(0.0-5.0): ")
			fmt.Scan(&data[*i].Rating)
		} 
		*i++
		fmt.Println("+----------------------------------+")
		fmt.Println("Apakah ingin input film lagi? (y/n): ")
		fmt.Scan(&pilihan)
		UpScreen()
		if pilihan == "n" {
			stop = true
		}
	}
}

//Program untuk menghapus film
func hapusFilm(data *Films, banyakfilm *int) {
	var JudulDicari, pilihan string
	var find bool = false
	var j int
	var i int = 0
	UpScreen()
	cover1Text("HAPUS FILM")
	fmt.Println("Judul film yang ingin dihapus: ")
	fmt.Scan(&JudulDicari)
	for i < *banyakfilm && !find {
		if data[i].judul == JudulDicari {
			find = true
		} else {
			i++
		}
	}
	if find {
		for j = i; j < *banyakfilm; j++ { // ketika film di temukan film bergeser ke kiri, dimulai dari index ditemukan
			data[j] = data[j+1] //lalu index yang dicari di timpa dengan yang baru
		}
		*banyakfilm-- //banyak film di kurangi
		cover1Text("FILM BERHASIL DI HAPUS")
	} else {
		cover1Text("FILM YANG DICARI TDK ADA")
	}
	fmt.Printf("Masih ingin menghapus film??\nInsert Y untuk menghapus, apapun untuk kembali: ")
	fmt.Scan(&pilihan)
	if pilihan == "Y" || pilihan == "y" {
		hapusFilm(data, banyakfilm)
	} else {
		UpScreen()
	}
}

//program untuk mengubah data film menggunakan teknik sequential search untuk mencari film yang ingin diubah
func ubahFilm(data *Films) {
	var pilihan, masukanJudul, judulBaru, deskripsiBaru, genreBaru string
	var tahunRilisBaru int
	var ratingBaru float64
	var find bool = false
	var jumlahFilm int = len(data)
	var i int = 0
	UpScreen()
	cover1Text("UPDATE DATA FILM")
	fmt.Print("| Film yang ingin diubah: ")
	fmt.Scan(&masukanJudul)
	for i < jumlahFilm && !find {
		if data[i].judul == masukanJudul {
			find = true
		} else {
			i++
		}
	}
	if find {
		UpScreen()
		cover1TextLong("Data baru untuk film " + data[i].judul)
		fmt.Printf("| %-50s |", "Informasi yang dapat diubah")
		fmt.Printf("\n| %-50s |\n| %-50s |\n| %-50s |\n| %-50s |\n| %-50s |\n| %-50s |\n", "1. Judul", "2. Tahun Rilis", "3. Deskripsi", "4. Genre","5. Rating", "6. Tidak Jadi")
		fmt.Println("+----------------------------------------------------+")
		fmt.Print("Insert 1-6 :")
		fmt.Scan(&pilihan)
		switch pilihan {
			case "1" :
				UpScreen()
				cover1Text("JUDUL")
				fmt.Print("Masukkan judul baru: ")
				fmt.Scan(&judulBaru)
				data[i].judul = judulBaru
				cover1Text("Judul Berhasil Diubah")
			case "2" :
				UpScreen()
				cover1Text("TAHUN RILIS")
				fmt.Print("Masukkan tahun rilis baru: ")
				fmt.Scan(&tahunRilisBaru)
				data[i].TahunRilis = tahunRilisBaru
				cover1Text("Tahun Rilis Berhasil DiUbah")
			case "3" :
				UpScreen()
				cover1Text("DESKRIPSI")
				fmt.Println("Masukkan deskripsi baru: ")
				fmt.Scan(&deskripsiBaru)
				data[i].Deskripsi = deskripsiBaru
				cover1Text("Deskripsi Berhasil Diubah")
			case "4" :
				UpScreen()
				cover1Text("GENRE")
				fmt.Print("Masukkan genre baru: ")
				fmt.Scan(&genreBaru)
				data[i].Genre = genreBaru
				cover1Text("Genre Berhasil Diubah")
			case "5" :
				UpScreen()
				cover1Text("RATING")
				fmt.Print("Masukkan rating baru: ")
				fmt.Scan(&ratingBaru)
				data[i].Rating = ratingBaru
				cover1Text("Rating Berhasil Di Ubah")
			case "6" :
				UpScreen()
			default :
			fmt.Println("Pilihan tidak valid")
		}
	} else {
		cover1Text("Data/film tidak ditemukan")
	}
}
