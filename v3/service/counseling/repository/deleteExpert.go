package repository

import (
	"log"
)

func (ths *repository) DeleteExpert(id uint) error {
	tx, err := ths.db.Beginx()
	if err != nil {
		log.Printf("delete post: %s", err.Error())
		return err
	}

	_, err = tx.Exec(ths.db.Rebind(`UPDATE "counseling".experts SET is_deleted = true, updated_at = now() WHERE id = ?`), id)
	if err != nil {
		log.Printf("delete post: %s", err.Error())
		_ = tx.Rollback()
		return err
	}

	_, err = tx.Exec(ths.db.Rebind(`UPDATE "counseling".expert_topics SET is_deleted = true, updated_at = now() WHERE expert_id = ?`), id)
	if err != nil {
		log.Printf("delete post: %s", err.Error())
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
