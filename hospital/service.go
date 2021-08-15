package hospital

import (
	"HOPE-backend/models"
	"strings"
)

type service struct {
	repo models.HospitalRepository
}

func NewHospitalService(repo models.HospitalRepository) models.HospitalService {
	return &service{
		repo: repo,
	}
}

func (ths *service) List(search, location, filter string) ([]*models.Hospital, error) {
	hospitals, err := ths.repo.GetAll(models.Hospital{})

	conditions := map[string]string{
		"search":   search,
		"location": location,
		"filter":   filter,
	}
	hospitals = filterByParams(hospitals, conditions)

	return hospitals, err
}

func filterByParams(hospitals []*models.Hospital, conditions map[string]string) []*models.Hospital {
	filteredHospitals := []*models.Hospital{}
	search := conditions["search"]
	location := conditions["location"]
	filter := conditions["filter"]

	for _, hospital := range hospitals {
		if checkConditions(hospital, search, location, filter) {
			filteredHospitals = append(filteredHospitals, hospital)
		}
	}

	return filteredHospitals
}

func checkConditions(hospital *models.Hospital, search, location, filter string) bool {
	isLocationEmpty := location == ""
	isSearchEmpty := search == ""
	isFilterEmpty := filter == ""

	n := len(hospital.Polyclinics)
	var m int
	if !isFilterEmpty {
		m = len(strings.Split(filter, ","))
	}

	isSearchApplied := strings.Contains(strings.ToLower(hospital.Name), strings.ToLower(search))
	isFilterApplied := checkSubArray(n, m, hospital.Polyclinics, strings.Split(filter, ","))
	isLocationApplied := strings.EqualFold(strings.ToLower(hospital.Location), strings.ToLower(location))

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
