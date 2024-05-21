package users

import (
	"context"
	"time"
	"work/internal/database"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

func New(userDB *pgx.Conn, timeout time.Duration) *Repository {
	return &Repository{userDB: userDB, timeout: timeout}
}

type Repository struct {
	userDB  *pgx.Conn
	timeout time.Duration
}

func (r *Repository) Create(ctx context.Context, req CreateUserReq) (database.User, error) {
	var u database.User

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	// implement me
	//changes
	err := r.userDB.QueryRow(ctx, "INSERT INTO users (Username, Password) VALUES ($1, $2) RETURNING id, username, pasword", req.Username, req.Password).Scan(&u.ID, &u.Username, &u.Password)
	if err != nil {
		return u, err
	}

	return u, nil
}

func (r *Repository) FindByID(ctx context.Context, userID uuid.UUID) (database.User, error) {
	var u database.User

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	// implement me
	//changes
	row := r.userDB.QueryRow(ctx, "SELECT ID, Username FROM users WHERE id = $1", userID)

	err := row.Scan(&u.ID, &u.Username)
	if err != nil {
		if err == pgx.ErrNoRows {
			return u, nil
		}
		return u, err
	}
	return u, nil
}

func (r *Repository) FindByUsername(ctx context.Context, username string) (database.User, error) {
	var u database.User

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	// implement me
	//changes
	row := r.userDB.QueryRow(ctx, "SELECT id, username FROM users WHERE username = $1", username)

	err := row.Scan(&u.ID, &u.Username)
	if err != nil {
		if err == pgx.ErrNoRows {
			return u, nil
		}
		return u, err
	}
	return u, nil
}
