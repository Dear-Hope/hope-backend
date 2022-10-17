package repository

import (
	"HOPE-backend/v2/models"
	"errors"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
)

type postgreSQLRepository struct {
	db *sqlx.DB
}

func NewPostgreSQLRepository(db *sqlx.DB) models.SelfCareRepository {
	return &postgreSQLRepository{
		db: db,
	}
}

func (ths *postgreSQLRepository) Create(newItem models.SelfCareItem) (*models.SelfCareItem, error) {
	rows, err := ths.db.NamedQuery(
		`INSERT INTO "selfcare".items (link, mood, type, title, description) 
		VALUES (:link, :mood, :type, :title, :description) RETURNING id`,
		newItem,
	)
	if err != nil {
		log.Printf("new self care item create failed: %s", err.Error())

		if strings.Contains(err.Error(), "duplicate key") {
			return nil, errors.New("this self care item already exist")
		}

		return nil, errors.New("failed to create new self care item")
	}

	for rows.Next() {
		if err = rows.Scan(&newItem.ID); err != nil {
			log.Printf("new self care item create failed: %s", err.Error())
			return nil, errors.New("failed to create new self care item")
		}
	}

	return &newItem, nil
}

func (ths *postgreSQLRepository) GetItemsByMood(mood string) ([]*models.SelfCareItem, error) {
	items := []*models.SelfCareItem{}
	err := ths.db.Select(
		&items,
		`SELECT id, mood, link, title, type, description 
		FROM "selfcare".items WHERE mood = $1`,
		mood,
	)
	if err != nil {
		log.Printf("self care items get by mood: %s", err.Error())

		err = errors.New("something wrong when get all self care items by mood")
		return nil, err
	}

	return items, nil
}

func (ths *postgreSQLRepository) GetAllItems() ([]*models.SelfCareItem, error) {
	items := []*models.SelfCareItem{}
	err := ths.db.Select(
		&items,
		`SELECT id, mood, link, title, type, description 
		FROM "selfcare".items`,
	)
	if err != nil {
		log.Printf("self care items get all: %s", err.Error())

		err = errors.New("something wrong when get all self care items")
		return nil, err
	}

	return items, nil
}

func (ths *postgreSQLRepository) GetAllTypesWithTotal() ([]models.SelfCareTypeInfo, error) {
	types := []models.SelfCareTypeInfo{}
	err := ths.db.Select(
		&types,
		`SELECT type, COUNT(*) as total 
		FROM "selfcare".items
		GROUP BY type`,
	)
	if err != nil {
		log.Printf("self care items get all types with total: %s", err.Error())

		err = errors.New("something wrong when get all self care types")
		return nil, err
	}

	return types, nil
}
