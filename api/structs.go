package api

type RecipeJSON struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Link        string   `json:"link"`
	Ingredients []string `json:"ingredients"`
}

type RequestJSON struct {
	GoodIngredients []string
	BadIngredients  []string
}
