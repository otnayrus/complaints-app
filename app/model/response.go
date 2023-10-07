package model

type GetCategoriesResponse struct {
	Categories []Category `json:"categories"`
}

type GetCategoryByIDResponse struct {
	Category Category `json:"category"`
}
