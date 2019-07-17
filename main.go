package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Surat struct {
	Nama       string
	Arti       string
	Asma       string
	Keterangan string
	Audio      string
}

func ambilData() {
	hasil, err := http.Get("https://al-quran-8d642.firebaseio.com/data.json")
	if err != nil {
		fmt.Println("tidak dapat ambil data")
	}

	bacaData, err := ioutil.ReadAll(hasil.Body)
	if err != nil {
		fmt.Println(err)
	}
	var dataJson []Surat

	kesalahan := json.Unmarshal([]byte(bacaData), &dataJson)
	if kesalahan != nil {
		log.Fatal(kesalahan)
	}
	var i int
	for i = 0; i < len(dataJson); i++ {
		fmt.Println("No:", i)
		fmt.Println(dataJson[i].Asma)
		fmt.Println("Surat : " + dataJson[i].Nama)
		fmt.Println("Mp3 (klik untuk mendengarkan): " + dataJson[i].Audio)
		fmt.Println("Keterangan: " + dataJson[i].Keterangan)

	}
}
func main() {
	ambilData()
}
