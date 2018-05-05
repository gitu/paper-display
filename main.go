package main

import (
	"github.com/wiless/waveshare"
	"net/http"
	"fmt"
	"golang.org/x/image/bmp"
)

var epd ws.EPD

func main() {
	ws.InitHW()
	epd.Init(true)

	url := "https://paper-display.herokuapp.com/clock"
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

	gray := ws.ConvertToGray(image)
	epd.SetFrame(*gray)
}
