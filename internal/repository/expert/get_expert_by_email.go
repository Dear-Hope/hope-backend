package expert

import (
	"HOPE-backend/internal/entity/expert"
	"context"
	"fmt"
)

func (r *repository) GetExpertByEmail(ctx context.Context, email string) (*expert.Expert, error) {
	var res expert.Expert

	query := r.db.Rebind(
		`SELECT id, email, password, name, expertise, price,
		photo, is_available, title, education, experience, bio
		FROM "expert".experts WHERE email = ? AND is_deleted = false`,
	)

	err := r.db.GetContext(ctx, &res, query, email)
	if err != nil {
		return nil, fmt.Errorf("[ExpertRepo.GetExpertByEmail] Failed: %w", err)
	}

	return &res, nil
}
