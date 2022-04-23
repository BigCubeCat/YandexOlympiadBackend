package api

type RecipeJSON struct {
	Title                string   `json:"title"`
	Description          string   `json:"description"`
	Link                 string   `json:"link"`
	Ingredients          []string `json:"ingredients"`
	IngredientsSetCounts string   `json:"counts"`
}

type RequestJSON struct {
	GoodIngredients []string `json:"good"`
	BadIngredients  []string `json:"bad"`
}
