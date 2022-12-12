package repository

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/selfcare/filter"
	"fmt"
	"log"
	"strings"
)

func (ths *repository) GetAll(f filter.ListCategory) (model.Categories, error) {
	var categories model.Categories

	query := `SELECT id, name, image_url, description FROM ` + model.Category{}.TableWithSchemaName()

	if f.ExcludeIDs != nil && len(f.ExcludeIDs) > 0 {
		var ids []string
		for _, id := range f.ExcludeIDs {
			ids = append(ids, fmt.Sprintf("%d", id))
		}

		query += fmt.Sprintf(" WHERE id NOT IN (%s)", strings.Join(ids, ","))
	}

	err := ths.db.Select(&categories, query)
	if err != nil {
		log.Printf("get all categories: %s", err.Error())
		return nil, err
	}

	return categories, nil
}
