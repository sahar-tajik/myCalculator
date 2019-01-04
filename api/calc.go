package api

import "strconv"

func add(num1, num2 float64) float64 {
	return num1 + num2
}

func subtract(num1, num2 float64) float64 {
	return num1 - num2
}

func multiply(num1, num2 float64) float64 {
	return num1 * num2
}

func divide(num1, num2 float64) float64 {
	return num1 / num2
}

func calculator(input Input) Output{
	var output Output

	switch input.Operation{
	case "ADD":
		output = Output{floatToString(add(input.NumberA, input.NumberB))}

	case "SUBTRACT":
		output = Output{floatToString(subtract(input.NumberA, input.NumberB))}

	case "MULTIPLY":
		output = Output{floatToString(multiply(input.NumberA, input.NumberB))}

	case "DIVIDE":
		output = Output{floatToString(divide(input.NumberA, input.NumberB))}

	}
	return output
}


func floatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 2, 64)
}