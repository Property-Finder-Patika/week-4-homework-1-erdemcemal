package temperatures

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) float64 {
	return c.Value*9/5 + 32
}

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) float64 {
	return c.Value + AbsoluteZeroC
}

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) float64 {
	return (f.Value - 32) * 5 / 9
}

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) float64 {
	return (f.Value + AbsoluteZeroF) * 5 / 9
}

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) float64 {
	return k.Value - AbsoluteZeroC
}

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) float64 {
	return (k.Value * 9 / 5) - AbsoluteZeroF
}
