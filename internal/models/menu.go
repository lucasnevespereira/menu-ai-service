package models

import "encoding/json"

type Menu struct {
	ID           string          `json:"id"`
	Content      json.RawMessage `json:"content"`
	ShoppingList json.RawMessage `json:"shoppingList"`
	Specs        MenuSpecs       `json:"specs"`
}

type MenuSpecs struct {
	MaxCalories string   `json:"maxCalories"`
	MaxCarbs    string   `json:"maxCarbs"`
	MaxProteins string   `json:"maxProteins"`
	MaxFats     string   `json:"maxFats"`
	Allergies   []string `json:"allergies"`
}

type MenuRequest struct {
	Content      json.RawMessage `json:"content"`
	ShoppingList json.RawMessage `json:"shoppingList"`
	Specs        MenuSpecs       `json:"specs"`
}
