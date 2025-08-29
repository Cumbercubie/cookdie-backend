package restaurants

import "time"

type Restaurant struct {
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	Star        byte      `json:"star"`
	PostalCode  string    `json:"postal_code"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type Recipe struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title,omitempty"`
	Images      []string  `json:"images,omitempty"`
	Ingredients []string  `json:"ingredients,omitempty"`
	Instruction []string  `json:"instructions,omitempty"`
	CreatorID   int64     `json:"creator_id"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	Rating      int64     `json:"rating,omitempty"`
}

type RecipeInput struct {
	CreatorID   int64    `json:"creator_id"`
	Title       string   `json:"title,omitempty"`
	Images      []string `json:"images,omitempty"`
	Ingredients []string `json:"ingredients,omitempty"`
	Instruction []string `json:"instructions,omitempty"`
}


