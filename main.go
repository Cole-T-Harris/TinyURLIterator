package main

import (
	"fmt"
	"github.com/pkg/browser"
	"math/rand"
	"net/http"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
var baseURL = "https://tinyurl.com/"

func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func main() {

	for {
		randomTinyURL := baseURL + randSeq(6)
		fmt.Println("Random URL: ", randomTinyURL)
		resp, err := http.Get(randomTinyURL)
		if err != nil {
			fmt.Println("Issue getting response: ", err)
		}
		defer resp.Body.Close()

		statusCode := resp.StatusCode
		fmt.Println("Status Code: ", statusCode)

		if err == nil && statusCode == 200 {
			fmt.Println("Valid URL: ", randomTinyURL)
			err = browser.OpenURL(randomTinyURL)
			if err != nil {
				fmt.Println("Issue opening browser: ", err)
			}
			return // Exit the loop and program when URL is valid
		} else {
			fmt.Println("Invalid URL: ", randomTinyURL)
		}

		time.Sleep(2 * time.Second) // Add a delay to prevent spamming the server
	}
}