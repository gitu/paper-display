package epd

var epd75b = &Display{
	Width:  640,
	Height: 384,
	Commands: map[Command]byte{
		PANEL_SETTING:                  0x00,
		POWER_SETTING:                  0x01,
		POWER_OFF:                      0x02,
		POWER_OFF_SEQUENCE_SETTING:     0x03,
		POWER_ON:                       0x04,
		POWER_ON_MEASURE:               0x05,
		BOOSTER_SOFT_START:             0x06,
		DEEP_SLEEP:                     0x07,
		DATA_START_TRANSMISSION_1:      0x10,
		DATA_STOP:                      0x11,
		DISPLAY_REFRESH:                0x12,
		IMAGE_PROCESS:                  0x13,
		LUT_FOR_VCOM:                   0x20,
		LUT_BLUE:                       0x21,
		LUT_WHITE:                      0x22,
		LUT_GRAY_1:                     0x23,
		LUT_GRAY_2:                     0x24,
		LUT_RED_0:                      0x25,
		LUT_RED_1:                      0x26,
		LUT_RED_2:                      0x27,
		LUT_RED_3:                      0x28,
		LUT_XON:                        0x29,
		PLL_CONTROL:                    0x30,
		TEMPERATURE_SENSOR_COMMAND:     0x40,
		TEMPERATURE_CALIBRATION:        0x41,
		TEMPERATURE_SENSOR_WRITE:       0x42,
		TEMPERATURE_SENSOR_READ:        0x43,
		VCOM_AND_DATA_INTERVAL_SETTING: 0x50,
		LOW_POWER_DETECTION:            0x51,
		TCON_SETTING:                   0x60,
		TCON_RESOLUTION:                0x61,
		SPI_FLASH_CONTROL:              0x65,
		REVISION:                       0x70,
		GET_STATUS:                     0x71,
		AUTO_MEASUREMENT_VCOM:          0x80,
		READ_VCOM_VALUE:                0x81,
		VCM_DC_SETTING:                 0x82,
		FLASH_MODE:                     0xe5,
	},
	Init: func(e *Display) () {
		e.Reset()
		e.CallFunction(POWER_SETTING, 0x37, 0x00)
		e.CallFunction(PANEL_SETTING, 0xCF, 0x08)
		e.CallFunction(BOOSTER_SOFT_START, 0xc7, 0xcc, 0x28)
		e.CallFunction(POWER_ON)
		e.Wait()
		e.CallFunction(PLL_CONTROL, 0x3c)
		e.CallFunction(TEMPERATURE_CALIBRATION, 0x00)
		e.CallFunction(VCOM_AND_DATA_INTERVAL_SETTING, 0x77)
		e.CallFunction(TCON_SETTING, 0x22)
		e.CallFunction(TCON_RESOLUTION, 0x02, 0x80, 0x01, 0x80)
		e.CallFunction(VCM_DC_SETTING, 0x1E)
		e.CallFunction(FLASH_MODE, 0x03)
	},
}

func Epd75b() (*Display) {
	return epd75b
}

