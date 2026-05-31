package main

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