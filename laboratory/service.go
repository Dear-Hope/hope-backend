package laboratory

import (
	"HOPE-backend/models"
	"strings"
)

type service struct {
	repo models.LaboratoryRepository
}

func NewLaboratoryService(repo models.LaboratoryRepository) models.LaboratoryService {
	return &service{
		repo: repo,
	}
}

func (ths *service) List(search, location, filter string) ([]*models.Laboratory, error) {
	laboratories, err := ths.repo.GetAll(models.Laboratory{})

	conditions := map[string]string{
		"search":   search,
		"location": location,
		"filter":   filter,
	}
	laboratories = filterByParams(laboratories, conditions)

	return laboratories, err
}

func filterByParams(laboratories []*models.Laboratory, conditions map[string]string) []*models.Laboratory {
	filteredLaboratories := []*models.Laboratory{}
	search := conditions["search"]
	location := conditions["location"]
	filter := conditions["filter"]

	for _, laboratory := range laboratories {
		if checkConditions(laboratory, search, location, filter) {
			filteredLaboratories = append(filteredLaboratories, laboratory)
		}
	}

	return filteredLaboratories
}

func checkConditions(laboratory *models.Laboratory, search, location, filter string) bool {
	isLocationEmpty := location == ""
	isSearchEmpty := search == ""
	isFilterEmpty := filter == ""

	n := len(laboratory.Services)
	var m int
	if !isFilterEmpty {
		m = len(strings.Split(filter, ","))
	}

	isSearchApplied := strings.Contains(strings.ToLower(laboratory.Name), strings.ToLower(search))
	isFilterApplied := checkSubArray(n, m, laboratory.Services, strings.Split(filter, ","))
	isLocationApplied := strings.EqualFold(strings.ToLower(laboratory.Location), strings.ToLower(location))

	if isLocationEmpty {
		if isSearchEmpty && isFilterEmpty {
			return true
		}
		if isSearchEmpty && !isFilterEmpty {
			return isFilterApplied
		}
		if !isSearchEmpty && isFilterEmpty {
			return isSearchApplied
		}
		if !isSearchEmpty && !isFilterEmpty {
			return isSearchApplied && isFilterApplied
		}
	} else {
		if isSearchEmpty && isFilterEmpty {
			return isLocationApplied
		}
		if isSearchEmpty && !isFilterEmpty {
			return isLocationApplied && isFilterApplied
		}
		if !isSearchEmpty && isFilterEmpty {
			return isLocationApplied && isSearchApplied
		}
	}

	return isLocationApplied && isSearchApplied && isFilterApplied
}

func checkSubArray(n, m int, arr, subarr []string) bool {
	if m == 0 {
		return true
	}
	if n == 0 {
		return false
	}

	counter := 0

	for _, x := range arr {
		for _, y := range subarr {
			if x == y {
				counter++
				break
			}
		}
	}
	return counter == m
}
