package storage

import (
	"context"

	"github.com/242617/pace/model"
)

func GetUser(ctx context.Context, phone string) (*model.User, error) {

	var user model.User
	err := db.QueryRowContext(ctx, queryGetUser, phone).Scan(
		&user.FaceID,
	)
	if err != nil {
		return nil, err
	}
	if ctx.Err() != nil {
		err = ctx.Err()
		return nil, err
	}
	return &user, nil

}
