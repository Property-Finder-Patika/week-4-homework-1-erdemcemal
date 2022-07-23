package temperatures

type (
	// Celsius is a scale and unit of measurement for temperature.
	Celsius struct {
		Value float64
	}
	// Fahrenheit is a scale and unit of measurement for temperature.
	Fahrenheit struct {
		Value float64
	}
	// Kelvin is a scale and unit of measurement for temperature.
	Kelvin struct {
		Value float64
	}
)

const (
	AbsoluteZeroC = -273.15
	AbsoluteZeroF = 459.67
)

type Temperature interface {
	GetSymbol() rune
	ConvertTo(float64, rune) float64
}

// GetSymbol returns the symbol of the temperature.
func (c Celsius) GetSymbol() rune {
	return 'C'
}

// ConvertTo converts a temperature to a different scale.
func (c Celsius) ConvertTo(value float64, s rune) float64 {
	c.Value = value
	switch s {
	case 'F':
		return CToF(c)
	case 'K':
		return CToK(c)
	}
	return 0
}

// GetSymbol returns the symbol of the temperature.
func (f Fahrenheit) GetSymbol() rune {
	return 'F'
}

// ConvertTo converts a temperature to a different scale.
func (f Fahrenheit) ConvertTo(value float64, s rune) float64 {
	f.Value = value
	switch s {
	case 'C':
		return FToC(f)
	case 'K':
		return FToK(f)
	}
	return 0
}

// GetSymbol returns the symbol of the temperature.
func (k Kelvin) GetSymbol() rune {
	return 'K'
}

// ConvertTo converts a temperature to a different scale.

func (k Kelvin) ConvertTo(value float64, s rune) float64 {
	k.Value = value
	switch s {
	case 'C':
		return KToC(k)
	case 'F':
		return KToF(k)
	}
	return 0
}
