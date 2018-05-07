package main

import (
	"net/http"
	"fmt"
	"golang.org/x/image/bmp"
	"github.com/gitu/paper-display/epd"
	"time"
)

func main() {
	epd.InitHW()
	display := epd.Epd75b()
	display.Init(display)

	url := "https://paper-display.herokuapp.com/clock"

	for {
		response, err := http.Get(url)
		if err != nil {
			fmt.Println("Error while downloading", url, "-", err)
			return
		}
		defer response.Body.Close()

		image, err := bmp.Decode(response.Body)
		if err != nil {
			fmt.Println("Error while parsing", url, "-", err)
			return
		}

		display.DisplayImage(image)

		time.Sleep(10 * time.Second)
	}
}
