package consultation

import (
	"HOPE-backend/internal/entity/consultation"
	"context"
	"fmt"
)

func (r *repository) GetConsultationById(ctx context.Context, consulId uint64) (*consultation.Consultation, error) {
	var consul consultation.Consultation

	err := r.db.GetContext(ctx, &consul, r.db.Rebind(`
			SELECT id, user_id, expert_id, type_id, to_char(booking_date, 'YYYY-MM-DD') AS booking_date, 
		    start_time, end_time, status, user_notes, counsel_notes, document
			FROM "counsel".consultations WHERE is_deleted = false AND id = ?`), consulId)
	if err != nil {
		return nil, fmt.Errorf("[ConsultationRepo.GetConsultationById] Failed get: %w", err)
	}

	return &consul, nil
}
