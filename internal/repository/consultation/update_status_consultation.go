package consultation

import (
	"context"
	"fmt"
)

func (r *repository) UpdateStatusConsultation(ctx context.Context, id uint64, status string) (bool, error) {
	query := `UPDATE "counsel".consultations SET status = ? WHERE id = ?`

	result, err := r.db.ExecContext(ctx, r.db.Rebind(query), status, id)
	if err != nil {
		return false, fmt.Errorf("[ConsultationRepo.UpdateStatusConsultation] Failed exec: %w", err)
	}

	if affected, err := result.RowsAffected(); affected == 0 || err != nil {
		return false, fmt.Errorf("[ConsultationRepo.UpdateStatusConsultation] no rows affected: %w", err)
	}

	return true, nil
}
