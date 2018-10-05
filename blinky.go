// Example blinky project from https://gobot.io/documentation/platforms/arduino/.
package blinky

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
	"time"
)

func Run(verbose bool, serial string, rate time.Duration) {
	firmataAdaptor := firmata.NewAdaptor(serial)
	led := gpio.NewLedDriver(firmataAdaptor, "13")

	work := func() {
		gobot.Every(rate, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led},
		work,
	)

	robot.Start()
}
