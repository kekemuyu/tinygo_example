package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
)

func main() {

	time.Sleep(time.Millisecond * 100)
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 18000000,
		SCK:       machine.SPI0_SCK_PIN,
		SDO:       machine.SPI0_SDO_PIN,
		SDI:       machine.SPI0_SDI_PIN,
	})
	dis := ssd1306.NewSPI(machine.SPI0, machine.A4, machine.A3, machine.A2)
	config := ssd1306.Config{
		Width:    128,
		Height:   64,
		VccState: ssd1306.SWITCHCAPVCC,

		Address: 0x3C,
	}

	dis.Configure(config)
	dis.ClearBuffer()

	for {
		dis.ClearBuffer()
		strTime := time.Now()

		tinyfont.WriteLine(&dis, &freemono.Oblique9pt7b, 0, 20, strTime.Format("2006.01.02"), color.RGBA{1, 1, 1, 255})
		tinyfont.WriteLine(&dis, &freemono.Oblique9pt7b, 10, 40, strTime.Format("15:04:05"), color.RGBA{1, 1, 1, 255})

		dis.Display()

	}
}
