package storage

import (
	"context"
	"database/sql"

	"github.com/242617/pace/model"
)

func GetUserByPhone(ctx context.Context, phone string) (*model.User, error) {

	var user model.User
	var personID, name, alias sql.NullString
	err := db.QueryRowContext(ctx, queryGetUserByPhone, phone).Scan(
		&user.Phone,
		&personID,
		&name,
		&alias,
	)
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	if ctx.Err() != nil {
		err = ctx.Err()
		return nil, err
	}
	if personID.Valid {
		user.PersonID = personID.String
	}
	if name.Valid {
		user.Name = name.String
	}
	if alias.Valid {
		user.Alias = alias.String
	}
	return &user, nil

}

func GetUserByPersonID(ctx context.Context, personID string) (*model.User, error) {

	var user model.User
	var nullPersonID, name, alias sql.NullString
	err := db.QueryRowContext(ctx, queryGetUserByPersonID, personID).Scan(
		&user.Phone,
		&nullPersonID,
		&name,
		&alias,
	)
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	if ctx.Err() != nil {
		err = ctx.Err()
		return nil, err
	}
	if nullPersonID.Valid {
		user.PersonID = nullPersonID.String
	}
	if name.Valid {
		user.Name = name.String
	}
	if alias.Valid {
		user.Alias = alias.String
	}
	return &user, nil

}

func CreateUser(ctx context.Context, phone string) error {

	_, err := db.ExecContext(ctx, queryCreateUser, phone)
	if err != nil {
		return err
	}
	if ctx.Err() != nil {
		err = ctx.Err()
		return err
	}
	return nil

}

func UpdateUserName(ctx context.Context, phone string, name string) error {

	_, err := db.ExecContext(ctx, queryUpdateUserName, name, phone)
	if err != nil {
		return err
	}
	if ctx.Err() != nil {
		err = ctx.Err()
		return err
	}
	return nil

}

func UpdateUserAlias(ctx context.Context, phone string, alias string) error {

	_, err := db.ExecContext(ctx, queryUpdateUserAlias, alias, phone)
	if err != nil {
		return err
	}
	if ctx.Err() != nil {
		err = ctx.Err()
		return err
	}
	return nil

}

func UpdateUserPersonID(ctx context.Context, phone string, personID string) error {

	_, err := db.ExecContext(ctx, queryUpdateUserPersonID, personID, phone)
	if err != nil {
		return err
	}
	if ctx.Err() != nil {
		err = ctx.Err()
		return err
	}
	return nil

}
