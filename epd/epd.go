package epd

import (
	"github.com/kidoman/embd"
	"time"
	"log"
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

/*

def get_frame_buffer(self, image):
buf = [0xFF] * (self.width * self.height / 8)
# Set buffer to value of Python Imaging Library image.
# Image must be in mode L.
image_grayscale = image.convert('1')
imwidth, imheight = image_grayscale.size
if imwidth != self.width or imheight != self.height:
raise ValueError('Image must be same dimensions as display \
({0}x{1}).' .format(self.width, self.height))

pixels = image_grayscale.load()
for y in range (self.height):
for x in range (self.width):
# Set the bits for the column of pixels at the current position.
if pixels[x, y] == 0:
buf[(x + y * self.width) / 8] &= ~(0x80 >> (x % 8))
return buf

def display_frame(self, frame_buffer_black, frame_buffer_red):
self.send_command(DATA_START_TRANSMISSION_1)
for i in range (0, self.width / 8 * self.height):
temp1 = frame_buffer_black[i]
temp2 = frame_buffer_red[i]
j = 0
while (j < 8):
if ((temp2 & 0x80) == 0x00):
temp3 = 0x04                #red
elif ((temp1 & 0x80) == 0x00):
temp3 = 0x00                #black
else:
temp3 = 0x03                #white

temp3 = (temp3 << 4) & 0xFF
temp1 = (temp1 << 1) & 0xFF
temp2 = (temp2 << 1) & 0xFF
j += 1
if ((temp2 & 0x80) == 0x00):
temp3 |= 0x04              #red
elif ((temp1 & 0x80) == 0x00):
temp3 |= 0x00              #black else:
temp3 |= 0x03              #white
temp1 = (temp1 << 1) & 0xFF
temp2 = (temp2 << 1) & 0xFF
self.send_data(temp3)
j += 1
self.send_command(DISPLAY_REFRESH)
self.delay_ms(100)
self.wait_until_idle()

def sleep(self):
self.send_command(POWER_OFF)
self.wait_until_idle()
self.send_command(DEEP_SLEEP)
self.send_data(0xa5)

### END OF FILE ###

*/
