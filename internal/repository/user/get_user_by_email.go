package user

import (
	"HOPE-backend/internal/entity/user"
	"context"
	"fmt"
)

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	var res user.User

	query := r.db.Rebind(
		`SELECT u.id, u.email, u.password, u.name, u.alias, u.is_verified, u.secret_key,
		p.photo, p.total_audio_played, p.total_time_played, p.longest_streak
		FROM "auth".users u, "auth".user_profiles p WHERE u.id = p.user_id AND u.email = ?`,
	)

	err := r.db.GetContext(ctx, &res, query, email)
	if err != nil {
		return nil, fmt.Errorf("[UserRepo.GetUserByEmail][010011] Failed: %v", err)
	}

	return &res, nil
}
