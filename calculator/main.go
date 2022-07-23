package main

import (
	"bufio"
	"fmt"
	"github.com/Property-Finder-Patika/week-4-homework-1-erdemcemal/calculator/mathFunction"
	"math"
	"os"
	"strings"
)

type Calculator struct {
	functions []mathFunction.MathFunction
}

// addMathFunction adds a math function to the calculator
func (c *Calculator) addMathFunction(m mathFunction.MathFunction) {
	c.functions = append(c.functions, m)
}

// doCalculation calculates the value of the given function with the given argument
func (c Calculator) doCalculation(name string, arg float64) float64 {
	var result float64
	for _, f := range c.functions {
		// case insensitive comparison
		if name == strings.ToLower(f.GetName()) {
			result = f.Calculate(arg)
		}
	}
	return result
}

func main() {
	//  functions()
	startCalculator()
}

func startCalculator() {
	myCalculator := Calculator{}

	// add the math functions to the calculator
	myCalculator.addMathFunction(mathFunction.Sin{Name: "Sinus"})
	myCalculator.addMathFunction(mathFunction.Cos{Name: "Cosines"})
	myCalculator.addMathFunction(mathFunction.Log{Name: "Log"})

	input := bufio.NewScanner(os.Stdin)
	fmt.Println("\nCalculator started.")
	fmt.Println("You can calculate using following functions")
	for _, f := range myCalculator.functions {
		fmt.Println(f.GetName())
	}

	for input.Scan() {
		var fName string
		var arg float64
		fmt.Println("> Enter name of the calculation:")
		_, err := fmt.Scanf("%s", &fName)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		//
		if fName == "x" {
			os.Exit(0)
		} else {
			fmt.Println("> Enter a value for the calculation:")
			_, err := fmt.Scanf("%f", &arg)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
		}

		// do calculation with case-insensitive fName and arg
		value := myCalculator.doCalculation(strings.ToLower(fName), arg)
		fmt.Printf("Result of %s of %f : %f\n", fName, arg, value)
	}
}

func functions() {
	sin := mathFunction.Sin{Name: "Sinus"}
	fmt.Printf("%v\n", sin)
	sin30 := sin.Calculate(math.Pi / 6)
	fmt.Printf("Sinus of 30 degree is %f\n", sin30)

	cos := mathFunction.Cos{Name: "Cosinus"}
	fmt.Printf("%v\n", sin)
	cos30 := cos.Calculate(math.Pi / 6)
	fmt.Printf("Cosinus of 30 degree is %f\n", cos30)

	log := mathFunction.Log{Name: "Log of base e"}
	fmt.Printf("%v\n", log)
	logE := log.Calculate(2.71828)
	fmt.Printf("Log of Euler constant is %f\n", logE)

	var mf1 mathFunction.MathFunction = sin
	fmt.Printf("%v\n", mf1)

	mf1 = cos
	fmt.Printf("%v\n", mf1)

	mf1 = log
	fmt.Printf("%v\n", mf1)
}

func calculator() {
	myCalculator := Calculator{}
	//myCalculator.functions[0] = Sin{"Sinus"}
	//myCalculator.functions[1] = Cos{"Cosinus"}
	//myCalculator.functions[2] = Log{"Log"}

	myCalculator.addMathFunction(mathFunction.Sin{Name: "Sinus"})
	myCalculator.addMathFunction(mathFunction.Cos{Name: "Cosines"})
	myCalculator.addMathFunction(mathFunction.Log{Name: "Log"})

	fmt.Println(myCalculator.doCalculation("Sinus", math.Pi/6))
	fmt.Println(myCalculator.doCalculation("Cosines", math.Pi/6))
	fmt.Println(myCalculator.doCalculation("Log", math.E))
}
