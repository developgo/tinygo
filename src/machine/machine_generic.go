// +build !avr,!nrf,!sam,!stm32

package machine

// Dummy machine package, filled with no-ops.

var (
	SPI0 = SPI{0}
)

type GPIOMode uint8

const (
	GPIO_INPUT = iota
	GPIO_OUTPUT
)

func (p GPIO) Configure(config GPIOConfig) {
	gpioConfigure(p.Pin, config)
}

func (p GPIO) Set(value bool) {
	gpioSet(p.Pin, value)
}

func (p GPIO) Get() bool {
	return gpioGet(p.Pin)
}

//go:export __tinygo_gpio_configure
func gpioConfigure(pin uint8, config GPIOConfig)

//go:export __tinygo_gpio_set
func gpioSet(pin uint8, value bool)

//go:export __tinygo_gpio_get
func gpioGet(pin uint8) bool

type SPI struct {
	Bus uint8
}

type SPIConfig struct {
	Frequency uint32
	SCK       uint8
	MOSI      uint8
	MISO      uint8
	Mode      uint8
}

func (spi SPI) Configure(config SPIConfig) {
	if config.SCK == 0 {
		config.SCK = SPI0_SCK_PIN
	}
	if config.MOSI == 0 {
		config.MOSI = SPI0_MOSI_PIN
	}
	if config.MISO == 0 {
		config.MISO = SPI0_MISO_PIN
	}
	spiConfigure(spi.Bus, config.SCK, config.MOSI, config.MISO)
}

// Transfer writes/reads a single byte using the SPI interface.
func (spi SPI) Transfer(w byte) (byte, error) {
	return spiTransfer(spi.Bus, w), nil
}

//go:export __tinygo_spi_configure
func spiConfigure(bus uint8, sck uint8, mosi uint8, miso uint8)

//go:export __tinygo_spi_transfer
func spiTransfer(bus uint8, w uint8) uint8
