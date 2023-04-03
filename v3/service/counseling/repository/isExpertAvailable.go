package repository

import "log"

func (ths *repository) IsExpertAvailable(id uint) (bool, error) {
	var isAvailable bool
	query := `SELECT is_available FROM "counseling".experts WHERE is_deleted = false AND id = ?`

	err := ths.db.Get(&isAvailable, ths.db.Rebind(query), id)
	if err != nil {
		log.Printf("get expert availability: %s", err.Error())
		return false, err
	}

	return isAvailable, nil
}
