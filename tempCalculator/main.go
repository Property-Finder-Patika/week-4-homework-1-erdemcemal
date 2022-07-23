package main

import (
	"fmt"
	temp "github.com/Property-Finder-Patika/week-4-homework-1-erdemcemal/tempCalculator/temperatures"
	"os"
	"strings"
	"unicode/utf8"
)

type Converter struct {
	functions []temp.Temperature
}

// addTemperature adds a temperature to the list of temperatures
func (c *Converter) addTemperature(t temp.Temperature) {
	c.functions = append(c.functions, t)
}

// Display the list of temperatures
func (c Converter) displayTemperature() {
	fmt.Println("1. Celsius (C)")
	fmt.Println("2. Fahrenheit (F)")
	fmt.Println("3. Kelvin (K)")
}

// Do the conversion from one temperature to another
func (c Converter) doConversion(from rune, to rune, value float64) float64 {
	var result float64
	for _, f := range c.functions {
		if from == f.GetSymbol() {
			result = f.ConvertTo(value, to)
		}
	}
	return result
}

// Check if the temperature is valid
func (c Converter) validTemperatureSymbol(s rune) bool {
	for _, f := range c.functions {
		if s == f.GetSymbol() {
			return true
		}
	}
	return false
}

// Initialize the temperatures
func (c *Converter) initTemperatures() {
	c.functions = []temp.Temperature{}
	c.addTemperature(temp.Celsius{})
	c.addTemperature(temp.Fahrenheit{})
	c.addTemperature(temp.Kelvin{})
}
func main() {
	startConverter()
}
func startConverter() {
	converter := Converter{}
	converter.initTemperatures()

	fmt.Println("\nConverter started.")
	fmt.Println("Please enter the temperature symbol you want to convert from:")
	converter.displayTemperature()

	// Get the temperature to convert from
	var from string
	_, err := fmt.Scanf("%s", &from)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	// Check if the temperature is valid
	if !converter.validTemperatureSymbol(stringToRune(from)) {
		fmt.Println("Invalid temperature symbol")
		os.Exit(0)
	}

	fmt.Println("Please enter the temperature symbol you want to convert to:")
	converter.displayTemperature()

	// Get the temperature to convert to
	var to string
	_, err = fmt.Scanf("%s", &to)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	// Check if the temperature is valid
	if !converter.validTemperatureSymbol(stringToRune(to)) {
		fmt.Println("Invalid temperature symbol")
		os.Exit(0)
	}

	fmt.Println("Please enter the temperature value:")

	// Get the temperature value
	var value float64
	_, err = fmt.Scanf("%f", &value)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	result := converter.doConversion(stringToRune(from), stringToRune(to), value)
	fmt.Println("Result:", result)

	// Check if user wants to convert again
	fmt.Println("Do you want to convert again? (y/n)")
	var answer string
	_, err = fmt.Scanf("%s", &answer)
	if err != nil {
		fmt.Println(err)
		return
	}
	if strings.ToLower(answer) == "y" {
		startConverter()
	} else {
		fmt.Println("Exiting...")
		os.Exit(0)
	}

}

// Convert string to rune
func stringToRune(s string) rune {
	r, _ := utf8.DecodeRuneInString(s)
	return r
}
