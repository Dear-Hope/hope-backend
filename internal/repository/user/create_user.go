package user

import (
	"HOPE-backend/internal/entity/user"
	"context"
	"fmt"
)

func (r *repository) CreateUser(ctx context.Context, req user.User) (*user.User, error) {
	rows, err := r.db.NamedQueryContext(ctx,
		`INSERT INTO "user".users (email, password, name, alias, is_verified, secret_key, 
				photo, total_audio_played, total_time_played, longest_streak) 
			VALUES (:email, :password, :name, :alias, :is_verified, :secret_key, :photo, 
				:total_audio_played, :total_time_played, :longest_streak) RETURNING id`,
		req,
	)
	if err != nil {
		return nil, fmt.Errorf("[UserRepo.Create] Failed exec: %v", err)
	}

	defer func() {
		_ = rows.Close()
	}()
	for rows.Next() {
		err = rows.Scan(&req.Id)
		if err != nil {
			return nil, fmt.Errorf("[UserRepo.Create] Failed scan: %v", err)
		}
	}
	return &req, nil
}
