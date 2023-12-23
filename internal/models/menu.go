package models

type Menu struct {
	ID           string    `json:"id"`
	Content      string    `json:"content"`
	ShoppingList string    `json:"shoppingList"`
	Specs        MenuSpecs `json:"specs"`
	UserID       string    `json:"userID"`
}

type MenuSpecs struct {
	MaxCalories string   `json:"maxCalories"`
	MaxCarbs    string   `json:"maxCarbs"`
	MaxProteins string   `json:"maxProteins"`
	MaxFats     string   `json:"maxFats"`
	Allergies   []string `json:"allergies"`
	Lang        string   `json:"lang"`
}

type MenuSaveRequest struct {
	Content      string    `json:"content"`
	ShoppingList string    `json:"shoppingList"`
	Specs        MenuSpecs `json:"specs"`
	UserID       string    `json:"userID"`
}
