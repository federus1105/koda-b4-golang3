package main

import (
	"day1golang3/internals"
	"fmt"
)

func main() {
	daftarnama := []string{"federus", "itsna", "sidiq", "fiki", "ari", "anggi", "yoga"}
	var keyword string
	fmt.Println("Masukan Nama yang ingin anda cari....")
	fmt.Scan(&keyword)

	// Pencarian
	found := internals.Search(daftarnama, keyword)

	if len(found) > 0 {
		fmt.Println("Nama yang anda cari:")
		for _, name := range found {
			fmt.Println(name)
		}
	} else {
		fmt.Println(found)
	}
}
