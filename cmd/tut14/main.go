package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type contactInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}

func main() {
	contacts, err := loadJSON[contactInfo]("./contactInfo.json")
	if err != nil {
		log.Fatalf("Error loading contacts: %v", err)
	}
	fmt.Printf("\n%+v", contacts)

	purchases, err := loadJSON[purchaseInfo]("./purchaseInfo.json")
	if err != nil {
		log.Fatalf("Error loading purchases: %v", err)
	}
	fmt.Printf("\n%+v", purchases)
}

func loadJSON[T any](filePath string) ([]T, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}
	var loaded []T
	err = json.Unmarshal(data, &loaded)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling JSON: %w", err)
	}
	return loaded, nil
}
