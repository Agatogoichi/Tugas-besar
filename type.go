package main

const NMAX = 9999
type Film struct {
	judul string
	TahunRilis int
	Deskripsi string
	Genre string
	Rating float64
}
type Films [NMAX]Film

type CountGenre struct {
	Genre string
	Count int
}
type GENRES [NMAX]CountGenre
