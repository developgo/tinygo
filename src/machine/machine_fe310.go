// +build fe310

package machine

import (
	"device/sifive"
)

type PinMode uint8

// Set the pin to high or low.
func (p Pin) Set(high bool) {
}

type UART struct {
	Bus *sifive.UART_Type
}

var (
	UART0 = UART{sifive.UART0}
)

type UARTConfig struct {
}

func (uart UART) Configure(config UARTConfig) {
	sifive.GPIO0.IO_FUNC_SEL.ClearBits(0x00030000)
	sifive.GPIO0.IO_FUNC_EN.SetBits(0x00030000)

	// Assuming a 16Mhz Crystal (which is Y1 on the HiFive1), the divisor for a
	// 115200 baud rate is 138.
	sifive.UART0.DIV.Set(138)
	sifive.UART0.TXCTRL.Set(sifive.UART_TXCTRL_ENABLE)

	// busy loop until the line is asserted...
	var n sifive.Register32
	for n.Get() < 1000000 {
		n.Set(n.Get() + 1)
	}
}
