package database

const (
	DbAddress = "host=localhost port=5432 user=postgres password=password dbname=go_resto_app sslmode=disable"
	Secret    = "AES256Key-32Characters1092384756"
	Time      = 1
	Memory    = 64 * 1024
	Threads   = 4
	KeyLen    = 32
)
