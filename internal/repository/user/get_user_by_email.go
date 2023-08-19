package user

import (
	"HOPE-backend/internal/entity/user"
	"context"
	"fmt"
)

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	var res user.User

	query := r.db.Rebind(
		`SELECT id, email, password, name, alias, is_verified, secret_key,
		photo, total_audio_played, total_time_played, longest_streak
		FROM "user".users WHERE email = ?`,
	)

	err := r.db.GetContext(ctx, &res, query, email)
	if err != nil {
		return nil, fmt.Errorf("[UserRepo.GetUserByEmail] Failed: %w", err)
	}

	return &res, nil
}
