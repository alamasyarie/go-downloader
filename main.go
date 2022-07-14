package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	filename := "foto.png"
	url := "https://static.republika.co.id/uploads/images/inpicture_slide/film-the-batman-meluncurkan-situs-bertema-tokoh-penjahat-riddler_211229161652-466.jpeg"

	error := download(filename, url)

	if error != nil {
		fmt.Println("Download gagal")
	} else {
		fmt.Println("Download Berhasil")
	}
}

func download(fileName string, url string) error {
	response, error := http.Get(url)

	if error != nil {
		return error
	}
	defer response.Body.Close()

	file, error := os.Create((fileName))

	if error != nil {
		return error
	}
	defer file.Close()

	_, err := io.Copy(file, response.Body)

	return err
}
