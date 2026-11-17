// Package car modela um carro híbrido, composto por um motor (Engine).
package car

import (
	"errors"
	"fmt"
)

type Engine interface {
	Start()
}

type CombustionEngine struct{}

type ElectricEngine struct{}

func (e *ElectricEngine) Start() {
	fmt.Println("Electric engine started....")
}

func (c *CombustionEngine) Start() {
	fmt.Println("Combustion engine started....")
}

type Car struct {
	engine        Engine
	fuelQuantity  float64
	batteryCharge float64
}

// minBatteryChargePercent é a carga mínima (em %, de 0 a 100) para usar o motor elétrico.
const minBatteryChargePercent = 10.0

// validatePercent garante que o valor está entre 0 e 100.
// A condição negada também rejeita NaN, que falha em qualquer comparação.
func validatePercent(name string, value float64) error {
	if !(value >= 0 && value <= 100) {
		return fmt.Errorf("%s deve estar entre 0 e 100, recebido: %v", name, value)
	}
	return nil
}

// New recebe o combustível e a carga da bateria em percentual (0 a 100).
// Retorna erro se algum dos valores estiver fora desse intervalo.
func New(fuelQuantity float64, batteryCharge float64) (*Car, error) {
	c := &Car{}
	if err := c.setFuelQuantity(fuelQuantity); err != nil {
		return nil, err
	}
	if err := c.setBatteryCharge(batteryCharge); err != nil {
		return nil, err
	}

	if c.batteryCharge >= minBatteryChargePercent {
		c.engine = &ElectricEngine{}
	} else {
		c.engine = &CombustionEngine{}
	}
	return c, nil
}

func (c *Car) setBatteryCharge(charge float64) error {
	if err := validatePercent("batteryCharge", charge); err != nil {
		return err
	}
	c.batteryCharge = charge
	return nil
}

func (c *Car) setFuelQuantity(fuelQuantity float64) error {
	if err := validatePercent("fuelQuantity", fuelQuantity); err != nil {
		return err
	}
	c.fuelQuantity = fuelQuantity
	return nil
}

// ErrNoEngine é retornado por Start quando o carro não tem motor,
// o que acontece se ele for criado sem passar por New (ex.: car.Car{}).
var ErrNoEngine = errors.New("carro sem motor: crie o carro com car.New")

// Start liga o carro, delegando para o motor escolhido em New.
// Retorna ErrNoEngine se o carro não tiver motor.
func (c *Car) Start() error {
	if c.engine == nil {
		return ErrNoEngine
	}
	c.engine.Start()
	return nil
}

// FuelQuantity retorna o combustível em percentual (0 a 100).
func (c *Car) FuelQuantity() float64 {
	return c.fuelQuantity
}

// BatteryCharge retorna a carga da bateria em percentual (0 a 100).
func (c *Car) BatteryCharge() float64 {
	return c.batteryCharge
}
