package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CelciusToFarenheit(value float32) float32 {
	return (value*1.8 + 32)
}

func FarenheitToCelcius(value float32) float32 {
	return ((value - 32) * 5 / 9)
}

func GetValueToConvert() float64 {
	fmt.Printf("Type your value: ")
	get_scan := bufio.NewScanner(os.Stdin)
	get_scan.Scan()

	string_input := strings.TrimSpace(get_scan.Text())
	float_output := StrToFloat(string_input)

	return float64(float_output)
}

func StrToFloat(data string) float32 {
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

	choice := int8(StrToFloat(user_input))

	switch choice {
	case 1:
		buffer := GetValueToConvert()
		fmt.Printf("Farenheit: %v --> Celcius: %v\n", buffer, FarenheitToCelcius(float32(buffer)))

	case 2:
		buffer := GetValueToConvert()
		fmt.Printf("Celcius: %v --> Farenheit: %v\n", buffer, CelciusToFarenheit(float32(buffer)))

	default:
		fmt.Println("Incorrect option bruh")
	}

}
