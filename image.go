package main

import (
	"io"
	"net/http"
	"os"
)

func downloadImage(url string) io.ReadCloser {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return resp.Body
}

func saveImageAssets(img io.ReadCloser, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(file, img)
}
