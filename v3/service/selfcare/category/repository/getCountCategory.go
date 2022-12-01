package repository

import (
	"log"
)

func (ths *repository) GetCountCategory(tableName string) (int, error) {
	var total int
	err := ths.db.Get(&total, `SELECT COUNT(*) FROM `+tableName)
	if err != nil {
		log.Printf("get count per category: %s", err.Error())
		return 0, err
	}

	return total, nil
}
