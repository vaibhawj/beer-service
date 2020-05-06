package services

import (
	"github.com/vaibhawj/beer-service/app/models"
)

func ListBeers() []models.Beer {

	return []models.Beer{
		models.Beer{1, "Kingfisher", "Lager"},
		models.Beer{2, "Tuborg", "Dark"},
		models.Beer{3, "Budweiser", "Ale"}}
}
