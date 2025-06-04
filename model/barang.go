package models

import "time"

type Barang struct {
	Id              string
	Nama_barang     string
	Kategori_id     string
	Stok            string
	Kelompok_barang string
	Harga           string
	CreatedAt       time.Time
}
