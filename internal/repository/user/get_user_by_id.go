package user

import (
	"HOPE-backend/internal/entity/user"
	"context"
	"fmt"
)

func (r *repository) GetUserById(ctx context.Context, uid uint64) (*user.User, error) {
	var res user.User

	query := r.db.Rebind(
		`SELECT id, email, password, name, alias, is_verified, secret_key,
		photo, total_audio_played, total_time_played, longest_streak
		FROM "user".users WHERE id = ?`,
	)

	err := r.db.GetContext(ctx, &res, query, uid)
	if err != nil {
		return nil, fmt.Errorf("[UserRepo.GetUserById] Failed: %v", err)
	}

	return &res, nil
}
