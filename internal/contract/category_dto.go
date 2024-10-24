package contract

type NewCategoryDTO struct {
	Name string `json:"name" validate:"required,min=2,max=30"`
}
