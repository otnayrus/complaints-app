package model

type GetCategoriesResponse struct {
	Categories []Category `json:"categories"`
}

type GetCategoryByIDResponse struct {
	Category Category `json:"category"`
}

type GetComplaintsResponse struct {
	Complaints []Complaint `json:"complaints"`
}

type GetComplaintByIDResponse struct {
	Complaint Complaint `json:"complaint"`
}
