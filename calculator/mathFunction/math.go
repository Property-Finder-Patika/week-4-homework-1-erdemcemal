package mathFunction

import (
	"math"
	"math/rand"
	"time"
)

// MathFunction is an interface for math functions
type MathFunction interface {
	Calculate(arg float64) float64 // Calculate returns the result of the math function
	GetName() string               // GetName returns the name of the math function
}

type Sin struct {
	Name string
}

func (s Sin) Calculate(arg float64) float64 {
	return math.Sin(arg)
}

func (s Sin) GetName() string {
	return s.Name
}

type Cos struct {
	Name string
}

func (c Cos) Calculate(arg float64) float64 {
	return math.Cos(arg)
}

func (c Cos) GetName() string {
	return c.Name
}

type Log struct {
	Name string
}

func (l Log) Calculate(arg float64) float64 {
	return math.Log(arg)
}

func (l Log) GetName() string {
	return l.Name
}

// MathFunctionFactory is a factory to create math function randomly
func MathFunctionFactory() MathFunction {
	var mf MathFunction
	rand.Seed(time.Now().UnixNano())
	i := rand.Int31n(3)

	switch i {
	case 0:
		mf = Sin{"Sinus"}
	case 1:
		mf = Cos{"Cosines"}
	case 2:
		mf = Log{"Log"}
	}
	return mf
}
