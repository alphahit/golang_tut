package main

import "fmt"

type gasEngine struct {
	gallons float32
	mpg     float32
}

type electricEngine struct {
	kwh   float32
	mpkwh float32
}

type engine interface {
	calculateRange() float32
}

func (ge gasEngine) calculateRange() float32 {
	return ge.gallons * ge.mpg
}

func (ee electricEngine) calculateRange() float32 {
	return ee.kwh * ee.mpkwh
}

type car[T engine] struct {
	carMake  string
	carModel string
	engine   T
}

func displayCarDetails[T engine](c car[T]) {
	fmt.Printf("Make: %s, Model: %s, Range: %.2f miles\n", c.carMake, c.carModel, c.engine.calculateRange())
}

func hasEnoughRange[T engine](c car[T], distance float32) (bool, error) {
	if rangeVal := c.engine.calculateRange(); rangeVal < distance {
		return false, fmt.Errorf("insufficient range. Car can only travel %.2f miles", rangeVal)
	}
	return true, nil
}

func main() {
	var gasCar = car[gasEngine]{
		carMake:  "Honda",
		carModel: "Civic",
		engine: gasEngine{
			gallons: 12.4,
			mpg:     40,
		},
	}

	var electricCar = car[electricEngine]{
		carMake:  "Tesla",
		carModel: "Model 3",
		engine: electricEngine{
			kwh:   67.5,
			mpkwh: 4.17,
		},
	}

	// Directly handle each car
	displayCarDetails(gasCar)
	displayCarDetails(electricCar)

	distance := float32(200.0)

	hasRange, err := hasEnoughRange(gasCar, distance)
	if err != nil {
		fmt.Println(err)
	} else if hasRange {
		fmt.Printf("Honda Civic has enough range for %.2f miles.\n", distance)
	} else {
		fmt.Println("Honda Civic does not have enough range.")
	}
}
