package ambulance

import (
	"HOPE-backend/models"
	"strings"
)

type service struct {
	repo models.AmbulanceRepository
}

func NewAmbulanceService(repo models.AmbulanceRepository) models.AmbulanceService {
	return &service{
		repo: repo,
	}
}

func (ths *service) List(search, location string) ([]*models.Ambulance, error) {
	ambulances, err := ths.repo.GetAll(models.Ambulance{})

	conditions := map[string]string{
		"search":   search,
		"location": location,
	}
	ambulances = filterByParams(ambulances, conditions)

	return ambulances, err
}

func filterByParams(ambulances []*models.Ambulance, conditions map[string]string) []*models.Ambulance {
	var filteredAmbulances []*models.Ambulance
	search := conditions["search"]
	location := conditions["location"]

	for _, ambulance := range ambulances {
		if checkConditions(ambulance, search, location) {
			filteredAmbulances = append(filteredAmbulances, ambulance)
		}
	}

	return filteredAmbulances
}

func checkConditions(ambulance *models.Ambulance, search, location string) bool {
	isLocationEmpty := location == ""
	isSearchEmpty := search == ""

	if isLocationEmpty {
		if isSearchEmpty {
			return true
		}
		return strings.Contains(strings.ToLower(ambulance.Name), strings.ToLower(search))
	} else {
		if isSearchEmpty {
			return strings.EqualFold(strings.ToLower(ambulance.Location), strings.ToLower(location))
		}
	}

	return strings.Contains(strings.ToLower(ambulance.Name), strings.ToLower(search)) &&
		strings.EqualFold(strings.ToLower(ambulance.Location), strings.ToLower(location))
}
