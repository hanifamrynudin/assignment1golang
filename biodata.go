package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)


type Murid struct {
	ID        int16  `json:"ID"`
	Nama      string `json:"Nama"`
	Alamat    string `json:"Alamat"`
	Pekerjaan string `json:"Pekerjaan"`
	Alasan    string `json:"Alasan"`
}

type Murid_List struct {
	Murid_List []Murid `json:"Murid_List"`
}


func main() {
	if len(os.Args) > 1 {
		kode, _ := strconv.Atoi(os.Args[1])
		murid := GetMurid(int16(kode))

		if (Murid{}) == murid {
			fmt.Println("Tidak ditemukan data")
		} else {
			fmt.Println(murid)
		}
	} else {
		fmt.Println("No user arguments")
	}
}


func GetMurid(kode int16) (data Murid) {
	murid_list := GetAllData("./murid_list.json").Murid_List

	for _, v := range murid_list {
		if v.ID == kode {
			return v
		}
	}
	return Murid{}
}


func GetAllData(path string) (data Murid_List) {
	file, _ := ioutil.ReadFile(path)

	_ = json.Unmarshal([]byte(file), &data)

	return
}