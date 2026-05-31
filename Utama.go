package main
import "fmt"

func main() {
	var Data Films
	var Kondisi, X int
	var stop bool = false
	var i int = 5
	Data[0] = Film{"One_Piece", 1994, "Menceritakan_Perjalanan_Luffy_dan_Krunya_Mencari_One_Piece", "Action", 4.0}
	Data[1] = Film{"Naruto", 2002, "Menceritakan_Perjalanan_Naruto", "Action", 4.5}
	Data[2] = Film{"Doraemon", 1973, "Menceritakan_Keseharian_Nobita_dan_Doraemon", "Comedy", 4.8}
	Data[3] = Film{"JunjiIto", 2018, "Menceritakan_Kumpulan_Cerita_Horor_Junji_Ito", "Horror", 4.2}
	Data[4] = Film{"My_Hero_Academia", 2014, "Perjalanan_Deku_Menjadi_Pahlawan_Terbaik", "Action", 4.7}
	X = 1
	fmt.Printf("=================Menu=================\n1. Tampilkan film\n2. Input film\n3. Ubah film\n4. Hapus film\n5. Statistik \n6. Keluar\n")
	for stop == false  {
		if X % 3 == 0 {
			fmt.Printf("\n=================Menu=================\n1. Tampilkan film\n2. Input film\n3. Ubah film\n4. Hapus film\n5. Statistik \n6. Keluar\n")
		}
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
			X++
		}
}
