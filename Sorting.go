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

//sorting rating
//selection sort ascending
func AscendingRatingFilm(data *Films, banyakfilm int) {
	var temp Film
	var Mindex, i, j int
	for i = 0; i < banyakfilm-1; i++ {
		Mindex = i
		for j = i + 1; j < banyakfilm; j++ {
			if data[j].Rating < data[Mindex].Rating {
				Mindex = j
			}
		}
		temp = data[i]
		data[i] = data[Mindex]
		data[Mindex] = temp
	}
}

//insertion sort descending
func DescendingRatingFilm(data *Films, banyakfilm int) {
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

//sorting tahun film
//insertion sort ascending
func AscendingTahunFilm(data *Films, banyakfilm int) {
var temp Film
	var i, j int
	for i = 1; i < banyakfilm; i++ {
		temp = data[i]
		j = i - 1
		for j >= 0 && data[j].TahunRilis > temp.TahunRilis {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}

//selection sort descending
func DescendingTahunFilm(data *Films, banyakfilm int) {
var temp Film
	var Mindex, i, j int
	for i = 0; i < banyakfilm-1; i++ {
		Mindex = i
		for j = i + 1; j < banyakfilm; j++ {
			if data[j].TahunRilis > data[Mindex].TahunRilis {
				Mindex = j
			}
		}
		temp = data[i]
		data[i] = data[Mindex]
		data[Mindex] = temp
	}
}
