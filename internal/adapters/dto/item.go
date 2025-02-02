package dto

type ItemDto struct {
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float32 `json:"price"`
	ImageUrl string  `json:"image_url"`
} //@name ItemDto
