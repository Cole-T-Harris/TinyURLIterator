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
			time.Sleep(2 * time.Second) // Wait and retry
			continue  
		}

		statusCode := resp.StatusCode
		fmt.Println("Status Code: ", statusCode)
		waitTimer := time.Duration(15)
		if err == nil && statusCode == 200 {
			fmt.Println("Valid URL: ", randomTinyURL)
			err = browser.OpenURL(randomTinyURL)
			if err != nil {
				fmt.Println("Issue opening browser: ", err)
				waitTimer = 2
			}
		} else {
			fmt.Println("Invalid URL: ", randomTinyURL)
			waitTimer = 2
		}

		time.Sleep(waitTimer * time.Second) // Add a delay to prevent spamming the server
	}
}