package repository

import (
	"HOPE-backend/v2/models"
	"database/sql"
	"errors"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
)

type postgreSQLRepository struct {
	db *sqlx.DB
}

func NewPostgreSQLRepository(db *sqlx.DB) models.AuthRepository {
	return &postgreSQLRepository{
		db: db,
	}
}

func (ths *postgreSQLRepository) CreateUserWithProfile(user models.DBUserWithProfile) (uint, uint, error) {
	rows, err := ths.db.NamedQuery(
		`WITH new_user AS (INSERT INTO "auth".users (email, password, first_name, last_name, profile_photo, is_active) 
		VALUES (:email, :password, :first_name, :last_name, :profile_photo, :is_active) RETURNING id)
		INSERT INTO "auth".profiles (weight, height, job, activities, user_id)
		VALUES (:weight, :height, :job, :activities, (SELECT id from new_user))
		RETURNING user_id, id as profile_id`,
		user,
	)
	if err != nil {
		log.Printf("user create failed: %s", err.Error())

		if strings.Contains(err.Error(), "duplicate key") {
			return 0, 0, errors.New("user with this email address already exists")
		}
		return 0, 0, errors.New("failed to create user")
	}

	var userID, profileID uint
	for rows.Next() {
		err = rows.Scan(&userID, &profileID)
		if err != nil {
			log.Printf("user create failed: %s", err.Error())
			return 0, 0, errors.New("failed to create user")
		}
	}
	return userID, profileID, nil
}

func (ths *postgreSQLRepository) GetUserWithProfileByEmail(email string) (*models.DBUserWithProfile, error) {
	var dbUser models.DBUserWithProfile
	err := ths.db.Get(
		&dbUser,
		`SELECT u.id AS user_id, password, p.id AS profile_id, is_active
		FROM "auth".users AS u, "auth".profiles AS p 
		WHERE u.id = p.user_id AND email=$1`,
		email,
	)
	if err != nil {
		log.Printf("user get by email: %s", err.Error())

		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found with given email")
		}
		return nil, errors.New("failed to get user")
	}

	return &dbUser, nil
}

func (ths *postgreSQLRepository) GetUserWithProfileByID(id uint) (*models.DBUserWithProfile, error) {
	var dbUser models.DBUserWithProfile
	err := ths.db.Get(
		&dbUser,
		`SELECT u.id AS user_id, email, password, first_name, last_name, profile_photo, is_active
		weight, height, job, activities, p.id AS profile_id
		FROM "auth".users AS u, "auth".profiles AS p 
		WHERE u.id = p.user_id AND u.id=$1`,
		id,
	)
	if err != nil {
		log.Printf("user get by id: %s", err.Error())

		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("failed to get user")
	}

	return &dbUser, nil
}

func (ths *postgreSQLRepository) UpdateUserWithProfile(user models.DBUserWithProfile) (*models.DBUserWithProfile, error) {
	_, err := ths.db.NamedQuery(
		`WITH updated_query AS (UPDATE "auth".users 
			SET email = :email, password = :password, first_name = :first_name, 
			last_name = :last_name, profile_photo = :profile_photo, is_active = :is_active 
			WHERE id = :user_id RETURNING id
		) 
		UPDATE "auth".profiles SET weight = :weight, height = :height, job = :job,
		activities = :activities, user_id = (SELECT id from updated_query)
		WHERE id = :profile_id RETURNING user_id, id AS profile_id`,
		user,
	)
	if err != nil {
		log.Printf("user update: %s", err.Error())

		return nil, errors.New("failed to update user")
	}

	return &user, nil
}
