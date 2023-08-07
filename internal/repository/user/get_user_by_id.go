package user

import (
	"HOPE-backend/internal/entity/user"
	"context"
	"fmt"
)

func (r *repository) GetUserById(ctx context.Context, uid uint64) (*user.User, error) {
	var res user.User

	query := r.db.Rebind(
		`SELECT u.id, u.email, u.password, u.name, u.alias, u.is_verified, u.secret_key,
		p.photo, p.total_audio_played, p.total_time_played, p.longest_streak, r.role_name
		FROM "auth".users u, "auth".user_profiles p, "auth".user_roles r
		WHERE u.id = p.user_id AND u.id = r.user_id AND u.id = ?`,
	)

	err := r.db.GetContext(ctx, &res, query, uid)
	if err != nil {
		return nil, fmt.Errorf("[UserRepo.GetUserById][010021] Failed: %v", err)
	}

	return &res, nil
}
