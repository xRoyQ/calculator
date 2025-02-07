package calculator

// local
var logMessage = "[LOG]"

// global
var Version = "1.0"

// local
func internalSum(number int) int {
	return number - 1
}

// global
func Sum(number1, number2 int) int {
	return number1 + number2
}
