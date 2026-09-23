package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func celcius_to_farenheit(value float32) float32 {
	return (value*1.8 + 32)
}

func farenheit_to_celcius(value float32) float32 {
	return ((value - 32) * 5 / 9)
}

func get_value_to_convert() float64 {
	fmt.Printf("Type your value: ")
	get_scan := bufio.NewScanner(os.Stdin)
	get_scan.Scan()

	string_input := strings.TrimSpace(get_scan.Text())
	float_output := str_to_float(string_input)

	return float64(float_output)
}

func str_to_float(data string) float32 {
	conv_data, err := strconv.ParseFloat(data, 32)
	if err != nil {
		fmt.Println("Incorrect input type")
		os.Exit(0)
	}

	return float32(conv_data)
}

func main() {
	fmt.Println("Choose an available option bruh")
	fmt.Println("1 > to Celcius\n2 > to Farnheit")
	fmt.Printf("Enter your choice: ")
	conversion_type := bufio.NewScanner(os.Stdin)
	conversion_type.Scan()

	user_input := strings.TrimSpace(conversion_type.Text())

	choice := int8(str_to_float(user_input))

	switch choice {
	case 1:
		buffer := get_value_to_convert()
		fmt.Printf("Farenheit: %v --> Celcius: %v\n", buffer, farenheit_to_celcius(float32(buffer)))

	case 2:
		buffer := get_value_to_convert()
		fmt.Printf("Celcius: %v --> Farenheit: %v\n", buffer, celcius_to_farenheit(float32(buffer)))

	default:
		fmt.Println("Incorrect option bruh")
	}

}
