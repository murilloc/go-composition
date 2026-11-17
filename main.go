package main

import (
	"fmt"
	"os"

	"github.com/murilloc/go-composition/car"
)

func main() {
	myCar, err := car.New(50, 3)
	if err != nil {
		fmt.Fprintln(os.Stderr, "erro ao criar o carro:", err)
		os.Exit(1)
	}
	fmt.Printf("combustível: %.0f%%, bateria: %.0f%%\n", myCar.FuelQuantity(), myCar.BatteryCharge())
	if err := myCar.Start(); err != nil {
		fmt.Fprintln(os.Stderr, "erro ao ligar o carro:", err)
		os.Exit(1)
	}
}
