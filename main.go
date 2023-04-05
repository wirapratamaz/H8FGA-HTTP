package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//Get request dengan menggunakan function http.Get
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		log.Fatal(err)
	}
	//Cetak response body ke console
	fmt.Println(res.Body)

	//Membaca isi dari hasil HTTP GET request dan simpan ke variabel body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	//defer untuk menjalankan res.Body.Close() secara otomatis pada akhir
	defer res.Body.Close()

	//Konversi isi response body(tipe data slice of byte) ke dalam bentuk string
	sb := string(body)

	log.Println(sb)
}
