package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//* GET
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

	//* POST
	//map data
	data := map[string]interface{}{
		"title":  "Airell",
		"body":   "Jordan",
		"userId": 1,
	}

	//ubah map data menjadi string JSON -> simpan ke variable requestJson
	requestJson, err := json.Marshal(data)
	//instance dari http.Client untuk melakukan request.
	client := &http.Client{}
	if err != nil {
		log.Fatalln(err)
	}

	//HTTP request baru dengan menggunakan http.NewRequest()

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts",
		//payload data ditampung dalam request body menggunakan bytes.NewBuffer()
		bytes.NewBuffer(requestJson))
	//mengatur header
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	//menjalankan request dengan client.Do(req)
	res, err = client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	//Menutup body dari responsen
	defer res.Body.Close()

	//Membaca isi dari responsen
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//cetak argumen yang sudah dikonversi menjadi string
	log.Println(string(body))
}
