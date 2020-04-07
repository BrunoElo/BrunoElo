// A bi-directional Unit converter

package main

import (
	"fmt"
	"math"
)

type (
	// Seconds named type
	Seconds float64
	// Minute named type
	Minute float64
	// Feet named type
	Feet float64
	// Centimeter named type
	Centimeter float64
	// Celsius named type
	Celsius float64
	// Fahrenheit named type
	Fahrenheit float64
	// Radian named type
	Radian float64
	// Degree named type
	Degree float64
	// Kilograms named type
	Kilograms float64
	// Pounds named type
	Pounds float64
)

// Converter struct contains field data with named types
type Converter struct {
	Minute     Seconds
	Seconds    Minute
	Feet       Centimeter
	Centimeter Feet
	Celsius    Fahrenheit
	Fahrenheit Celsius
	Radian     Degree
	Degree     Radian
	Kilograms  Pounds
	Pounds     Kilograms
}

// MinuteToSeconds method conversion
func (cvr Converter) MinuteToSeconds(m Minute) Seconds {
	return Seconds(m * 60)
}

// SecondsToMinute method conversion
func (cvr Converter) SecondsToMinute(s Seconds) Minute {
	return Minute(s / 60)
}

// CentimeterToFeet method conversion
func (cvr Converter) CentimeterToFeet(c Centimeter) Feet {
	return Feet(c / 30.48)
}

// FeetToCentimeter method conversion
func (cvr Converter) FeetToCentimeter(f Feet) Centimeter {
	return Centimeter(f * 30.48)
}

// CelsiusToFahrenheit method conversion
func (cvr Converter) CelsiusToFahrenheit(C Celsius) Fahrenheit {
	return Fahrenheit((C * (9 / 5)) + 32)
}

// FahrenheitToCelsius method conversion
func (cvr Converter) FahrenheitToCelsius(F Fahrenheit) Celsius {
	return Celsius((F - 32) * 5 / 9)
}

// RadianToDegree method conversion
func (cvr Converter) RadianToDegree(r Radian) Degree {
	return Degree(r * (180 / math.Pi))
}

// DegreeToRadian method conversion
func (cvr Converter) DegreeToRadian(d Degree) Radian {
	return Radian(d * (math.Pi / 180))
}

// KilogramsToPounds method conversion
func (cvr Converter) KilogramsToPounds(k Kilograms) Pounds {
	return Pounds(k * 2.20462)
}

// PoundsToKilograms method conversion
func (cvr Converter) PoundsToKilograms(p Pounds) Kilograms {
	return Kilograms(p / 2.20462)
}

func main() {
	// Conversion of sample data to corresponding units
	cvr := Converter{Minute: 1, Seconds: 60, Centimeter: 30.48, Feet: 1, Celsius: 0, Fahrenheit: 32, Radian: math.Pi, Degree: 180, Kilograms: 1, Pounds: 2.20462}
	fmt.Println("1 minute =", cvr.MinuteToSeconds(1), "seconds")
	fmt.Println("60 seconds =", cvr.SecondsToMinute(60), "minute")
	fmt.Println("30.48 centimeter =", cvr.CentimeterToFeet(30.48), "foot")
	fmt.Println("1 foot =", cvr.FeetToCentimeter(1), "centimeter")
	fmt.Println("0 Celsius =", cvr.CelsiusToFahrenheit(0), "Fahrenheit")
	fmt.Println("32 Fahrenheit =", cvr.FahrenheitToCelsius(32), "Celsius")
	fmt.Println("Pi radian =", cvr.RadianToDegree(math.Pi), "degree")
	fmt.Println("180 degree =", cvr.DegreeToRadian(180), "radian")
	fmt.Println("1 kilograms =", cvr.KilogramsToPounds(1), "pounds")
	fmt.Println("2.20462 pounds =", cvr.PoundsToKilograms(2.20462), "kilograms")
}
