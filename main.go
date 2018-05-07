package main

import (
	"net/http"
	"fmt"
	"golang.org/x/image/bmp"
	"github.com/gitu/paper-display/epd"
	"time"
	"image"
)

func main() {
	epd.InitHW()
	display := epd.Epd75b()
	display.Init(display)

	url := "https://paper-display.herokuapp.com/clock"

	previous := image.Image(image.NewRGBA(image.Rect(0, 0, 1, 1)))
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

		if !sameImage(previous, image) {
			fmt.Println("new Image")
			previous = image
			display.DisplayImage(image)
		} else {
			fmt.Println("same image - skip update")
		}

		time.Sleep(5 * time.Second)
	}
}
func sameImage(i1 image.Image, i2 image.Image) bool {
	if i1.Bounds() != i2.Bounds() {
		return false
	}
	for y := 0; y < i1.Bounds().Dy(); y++ {
		for x := 0; x < i1.Bounds().Dx(); x++ {
			r1, g1, b1, _ := i1.At(x, y).RGBA()
			r2, g2, b2, _ := i2.At(x, y).RGBA()
			if r1 != r2 || g1 != g2 || b1 != b2 {
				return false
			}
		}
	}
	return true
}
