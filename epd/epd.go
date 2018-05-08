package epd

import (
	"github.com/kidoman/embd"
	"time"
	"log"
	"image"
	"errors"
)

type Command int

const (
	PANEL_SETTING                  Command = iota
	POWER_SETTING
	POWER_OFF
	POWER_OFF_SEQUENCE_SETTING
	POWER_ON
	POWER_ON_MEASURE
	BOOSTER_SOFT_START
	DEEP_SLEEP
	DATA_START_TRANSMISSION_1
	DATA_STOP
	DISPLAY_REFRESH
	IMAGE_PROCESS
	LUT_FOR_VCOM
	LUT_BLUE
	LUT_WHITE
	LUT_GRAY_1
	LUT_GRAY_2
	LUT_RED_0
	LUT_RED_1
	LUT_RED_2
	LUT_RED_3
	LUT_XON
	PLL_CONTROL
	TEMPERATURE_SENSOR_COMMAND
	TEMPERATURE_CALIBRATION
	TEMPERATURE_SENSOR_WRITE
	TEMPERATURE_SENSOR_READ
	VCOM_AND_DATA_INTERVAL_SETTING
	LOW_POWER_DETECTION
	TCON_SETTING
	TCON_RESOLUTION
	SPI_FLASH_CONTROL
	REVISION
	GET_STATUS
	AUTO_MEASUREMENT_VCOM
	READ_VCOM_VALUE
	VCM_DC_SETTING
	FLASH_MODE
)

type Display struct {
	Width    int
	Height   int
	Commands map[Command]byte
	Init     func(e *Display) ()
}

func (e *Display) CallFunction(command Command, data ...byte) {
	e.SendCommand(command)
	e.SendData(data...)
}

func (e *Display) SendCommand(cmd Command) {
	c := e.Commands[cmd]
	writeCmd(c)
}

func (e *Display) SendData(data ...byte) {
	writeData(data...)
}

// Ensure to wait before any next command is executed.. monitors the
// BUSY_PIN
func (e *Display) Wait() {
	var busy int
	var err error
	for ; busy == 1; busy, err = embd.DigitalRead(BUSY_PIN) {
		if err != nil {
			log.Println("Error waiting BUSY_PIN", err)
			time.Sleep(500 * time.Millisecond)
		}
		time.Sleep(100 * time.Millisecond) // polling for every 100ms
	}
}

func (e *Display) Reset() {
	embd.DigitalWrite(RST_PIN, embd.Low)
	time.Sleep(200 * time.Millisecond)
	embd.DigitalWrite(RST_PIN, embd.High)
	time.Sleep(200 * time.Millisecond)
}

func (e *Display) DisplayImage(img image.Image) error {
	if img.Bounds().Dx() != e.Width {
		return errors.New("Height doesn't match " + string(img.Bounds().Dx()) + " expected: " + string(e.Width))
	}
	if img.Bounds().Dy() != e.Height {
		return errors.New("Height doesn't match " + string(img.Bounds().Dy()) + " expected: " + string(e.Height))
	}
	e.SendCommand(DATA_START_TRANSMISSION_1)

	val := make([]byte, e.Width*e.Height/2)
	i := 0
	for y := 0; y < e.Height; y++ {
		for x := 0; x < e.Width; x++ {
			at := img.At(x, y)
			newVal := byte(0x00) // black
			r, g, b, _ := at.RGBA()
			if r > 15000 && b > 15000 && g > 15000 {
				newVal = 0x03 // white
			}
			if r > 15000 && b < 15000 {
				newVal = 0x04 // red
			}

			if i%2 == 0 {
				val[i/2] = newVal << 4
			} else {
				val[i/2] = val[i/2] | newVal
				e.SendData(val[i/2])
			}
			i++
		}
	}
	//e.SendData(val...)
	e.SendCommand(DISPLAY_REFRESH)
	time.Sleep(100 * time.Millisecond)
	e.Wait()
	return nil
}

func (e *Display) Sleep() {
	e.SendCommand(POWER_OFF)
	e.Wait()
	e.CallFunction(DEEP_SLEEP, 0xa5)
}
