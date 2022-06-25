package main

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }

func main() {
	var k = Kelvin(1000)
	c := Celsius(1000)
	var f Fahrenheit = 1000

	fmt.Println("Celsius to Fahrenheit:", CToF(c).String())
	fmt.Println("Kelvin to Celsius:", KToC(k).String())
	fmt.Println("Fahrenheit to Celsius:", FToC(f).String())

}
