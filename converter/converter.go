package main

import (
	"flag"
	"fmt"
)

func main() {
	// Флаги
	celsius := flag.Float64("c", 0.0, "Температура в Цельсиях")
	fahrenheit := flag.Float64("f", 0.0, "Температура в Фаренгейтах")
	flag.Parse()

	// Конвертация
	if *celsius != 0 {
		fmt.Printf("%.2f°C = %.2f°F\n", *celsius, celsiusToFahrenheit(*celsius))
	} else if *fahrenheit != 0 {
		fmt.Printf("%.2f°F = %.2f°C\n", *fahrenheit, fahrenheitToCelsius(*fahrenheit))
	} else {
		fmt.Println("Укажите флаг -c или -f")
	}
}

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}
