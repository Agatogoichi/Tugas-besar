package main

//Prosedur untuk mengurutkan film berdasarkan nama menggunakan teknik selection sort
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
