package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	baseURL := "https://tinyurl.com/"
	fmt.Println("Hello", baseURL)

	resp, err := http.Get(baseURL)
	if err != nil {
		fmt.Println("Invalid URL: ", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error Reading Body: ", err)
		return
	}
	fmt.Println(string(body[:]))
	fmt.Println("Status Code: ", resp.StatusCode)
}