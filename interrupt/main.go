package main

import (
	"machine"
	"runtime/volatile"
)

func pinInterruptChan(pin machine.Pin) <-chan bool {
	var state volatile.Register8
	ch := make(chan bool, 3)

	pin.SetInterrupt(machine.PinFalling, func(p machine.Pin) {
		b := false
		if state.Get() != 1 {
			state.Set(1)
			b = true
		} else {
			state.Set(0)
		}
		select {
		case ch <- b:
		default:
		}
	})

	return ch
}

func main() {
	button := machine.GP0
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	chBtn := pinInterruptChan(button)
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		select {
		case btn := <-chBtn:
			led.Set(btn)
		default:
		}
	}
}
