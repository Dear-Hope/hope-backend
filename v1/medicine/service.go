package medicine

import (
	"HOPE-backend/v1/models"
)

type service struct {
	repo models.MedicineRepository
}

func NewMedicineService(repo models.MedicineRepository) models.MedicineService {
	return &service{
		repo: repo,
	}
}

func (ths *service) List(kind models.Kind) ([]*models.Medicine, error) {
	whereClause := models.Medicine{
		Kind: kind,
	}

	medicines, err := ths.repo.GetAll(whereClause)

	medicines = convertKindToName(medicines)

	return medicines, err
}

func convertKindToName(medicines []*models.Medicine) []*models.Medicine {
	for _, medicine := range medicines {
		medicine.ConvertKindToName()
	}

	return medicines
}
