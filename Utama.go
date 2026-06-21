package main
import "fmt"

func main() {
	var Data Films
	var Kondisi int
	var stop bool = false
	var i int = 11 //ini adalah deklarsi awal dari banyak film yang akan terus dibawa di setiap codenya
	//data awal yang sudah dimasukan sejak
	Data[0] = Film{"One_Piece", 1994, "Sebuah_Kisah_Perjalanan_Luffy_dan_Krunya_Mencari_One_Piece", "Adventure", 4.0}
	Data[1] = Film{"Naruto", 2002, "Kisah_Perjalanan_Naruto", "Action", 4.5}
	Data[2] = Film{"Doraemon", 1973, "Keseharian_Nobita_dan_Doraemon", "Comedy", 4.8}
	Data[3] = Film{"JunjiIto", 2018, "Kumpulan_Cerita_Horor_Junji_Ito", "Horror", 4.2}
	Data[4] = Film{"My_Hero_Academia", 2014, "Perjalanan_Deku_Menjadi_Pahlawan_Terbaik", "Action", 4.7}
	Data[5] = Film{"Anohana", 2011, "Lupa sama jalan ceritanya", "Drama", 4.6}
	Data[6] = Film{"Attack_on_Titan", 2013, "Menceritakan_Perjuangan_Manusia_Melawan_Titan", "Action", 4.9}
	Data[7] = Film{"Death_Note", 2006, "Menceritakan_Perjuangan_Light_Yagami_Melawan_Justice", "Thriller", 4.4}
	Data[8] = Film{"Tokyo_Ghoul", 2014, "Menceritakan_Perjuangan_Kaneki_Melawan_Ghoul", "Horror", 4.3}
	Data[9] = Film{"Demon_Slayer", 2019, "Menceritakan_Perjuangan_Tanjiro_Melawan_Demon", "Action", 4.8}
	Data[10] = Film{"Dragonest", 2007, "Perjalanan fantasy seorang pahlawan", "Adventure", 4.0}
	UpScreen()
	
	for stop == false  {
		cover1TextLong("SELAMAT DATANG DI CINEMA GWEH!")
		fmt.Printf("| %-50s |\n| %-50s |\n| %-50s |\n| %-50s |\n", "1. Tampilkan Film", "2. Input/Update Film", "3. Statistik", "4. Exit")
		fmt.Println("+----------------------------------------------------+")
		fmt.Print("|Pilih Menu (1-4):")
		fmt.Scan(&Kondisi)
		UpScreen()
		switch Kondisi {
			case 1 :
			tampilFilm(Data, i)
		    case 2 :
			UpdateArrayFilm(&Data, &i)
			case 3 :
			Statistik(Data, i)
			case 4 :
			stop = true
			default :
			fmt.Println("TIdak ada pilihan")
		}
	}
	UpScreen()
	CoverExit() //ketika masuk ke case 4, stop menjadi true dan while selesai.
}