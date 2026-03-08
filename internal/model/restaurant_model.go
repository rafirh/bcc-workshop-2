package model

type CreateRestaurant struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}
