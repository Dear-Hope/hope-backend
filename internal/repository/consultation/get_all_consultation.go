package consultation

import (
	"HOPE-backend/internal/entity/consultation"
	"context"
	"fmt"
)

func (r *repository) GetAllConsultation(ctx context.Context, req consultation.Consultation, isFilterPerMonth bool) (
	consultation.Consultations, error) {
	var (
		consultations consultation.Consultations
		args          []interface{}
	)

	query := `SELECT id, user_id, expert_id, type_id, to_char(booking_date, 'YYYY-MM-DD') AS booking_date, 
       		start_time, end_time, status FROM "counsel".consultations WHERE is_deleted = false`

	if req.ExpertId > 0 {
		query += ` AND expert_id = ?`
		args = append(args, req.ExpertId)
	}

	if req.UserId > 0 {
		query += ` AND user_id = ?`
		args = append(args, req.UserId)
	}

	if req.Status != "" {
		query += ` AND status = ?`
		args = append(args, req.Status)
	}

	if req.BookingDate != "" {
		if isFilterPerMonth {
			query += ` AND to_char(booking_date, 'YYYY-MM') = ?`
		} else {
			query += ` AND booking_date = ?`
		}
		args = append(args, req.BookingDate)
	}

	query += ` ORDER BY booking_date`
	err := r.db.SelectContext(ctx, &consultations, r.db.Rebind(query), args...)
	if err != nil {
		return nil, fmt.Errorf("[ConsultationRepo.GetAllConsultation] Failed select: %w", err)
	}

	return consultations, nil
}
