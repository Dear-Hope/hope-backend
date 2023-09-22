package consultation

import (
	"HOPE-backend/internal/entity/consultation"
	"context"
	"fmt"
)

func (r *repository) GetAllConsultation(ctx context.Context, expertId, userId uint64,
	status string) (consultation.Consultations, error) {
	var (
		consultations consultation.Consultations
		args          []interface{}
	)

	query := `SELECT id, user_id, expert_id, type_id, booking_date, start_time, end_time, status 
			FROM "counsel".consultations WHERE is_deleted = false`

	if expertId > 0 {
		query += ` AND expert_id = ?`
		args = append(args, expertId)
	}

	if userId > 0 {
		query += ` AND user_id = ?`
		args = append(args, userId)
	}

	if status != "" {
		query += ` AND status = ?`
		args = append(args, status)
	}

	err := r.db.SelectContext(ctx, &consultations, r.db.Rebind(query), args...)
	if err != nil {
		return nil, fmt.Errorf("[ConsultationRepo.GetAllConsultation] Failed select: %w", err)
	}

	return consultations, nil
}
