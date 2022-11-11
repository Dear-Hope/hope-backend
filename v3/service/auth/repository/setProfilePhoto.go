package repository

import "log"

func (ths *repository) SetProfilePhoto(id uint, link string) error {
	_, err := ths.db.Queryx(
		`UPDATE "auth".profiles SET photo = $1 WHERE id = $2`,
		link,
		id,
	)
	if err != nil {
		log.Printf("user set profile photo: %s", err.Error())
		return err
	}

	return nil
}
