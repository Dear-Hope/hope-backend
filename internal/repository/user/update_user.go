package user

import (
	"HOPE-backend/internal/entity/user"
	"context"
	"fmt"
	"time"
)

func (r *repository) UpdateUser(ctx context.Context, req user.User) (*user.User, error) {
	req.UpdatedAt = time.Now().UTC()

	queryUser := `WITH query_user AS (
		UPDATE "auth".users 
		SET email = COALESCE(NULLIF(:email, ''), email),
		    password = COALESCE(NULLIF(:password, ''), password),
		    name = COALESCE(NULLIF(:name, ''), name),
		    alias = COALESCE(NULLIF(:alias, ''), alias),
		    updated_at = :updated_at
		WHERE id = :id
	)`

	queryProfile := `UPDATE "auth".user_profiles
		SET photo = COALESCE(NULLIF(:photo, ''), photo),
		    total_audio_played = COALESCE(NULLIF(:total_audio_played, 0), total_audio_played),
		    total_time_played = COALESCE(NULLIF(:total_time_played, 0), total_time_played),
		    longest_streak = COALESCE(NULLIF(:longest_streak, 0), longest_streak),
		    updated_at = (
		        CASE WHEN COALESCE(
		            NULLIF(:photo, ''),
		            TEXT(NULLIF(:total_audio_played, 0)),
		            TEXT(NULLIF(:total_time_played, 0)),
		            TEXT(NULLIF(:longest_streak, 0))
				) IS NULL THEN updated_at ELSE :updated_at END)
		WHERE user_id = :id`

	_, err := r.db.NamedQueryContext(ctx, queryUser+queryProfile, req)
	if err != nil {
		return nil, fmt.Errorf("[UserRepo.UpdateUser][010019] Failed: %v", err)
	}

	return &req, nil
}
