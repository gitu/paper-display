package main

import (
	"flag"
	"fmt"
	"github.com/denisbrodbeck/machineid"
	"github.com/gitu/paper-display/epd"
	"github.com/golang/glog"
	"golang.org/x/image/bmp"
	"image"
	"net/http"
	"time"
)

var url = flag.String("url", "https://paper-display.herokuapp.com/clock?display="+getMid(), "url to fetch")

func getMid() string {
	mid, err := machineid.ProtectedID("paper-display")
	if err != nil {
		fmt.Println("Error while getting machine id ", err)
	} else {
		fmt.Println("MachineId", mid)
	}
	return mid
}

func main() {
	flag.Parse()

	epd.InitHW()
	display := epd.Epd75b()

	previous := image.Image(image.NewRGBA(image.Rect(0, 0, 1, 1)))
	for {
		img := fetchImage(*url)
		if img == nil {
			glog.Info("no image - skip update")
		} else if !sameImage(previous, img) {
			glog.Info("new Image")
			previous = img
			display.Init(display)
			display.DisplayImage(img)
			display.Sleep()
		} else {
			glog.Info("same image - skip update")
		}

		time.Sleep(5 * time.Second)
	}
}
func fetchImage(url string) image.Image {
	response, err := http.Get(url)
	if err != nil {
		glog.Error("Error while downloading", url, "-", err)
		return nil
	}
	defer response.Body.Close()

	img, err := bmp.Decode(response.Body)
	if err != nil {
		glog.Error("Error while parsing", url, "-", err)
		return nil
	}
	return img
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
