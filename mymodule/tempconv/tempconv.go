package tempconv

import "fmt"

// Celsius Type
type Celsius float64

// Fahrenheit Type
type Fahrenheit float64

const (
	// AbsoluteZeroC tmp
	AbsoluteZeroC Celsius = -273.15
	// FreezingC tmp
	FreezingC Celsius = 0
	// BoilingC tmp
	BoilingC Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
