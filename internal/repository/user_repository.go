package repository

import (
    "context"
	"task-management/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(ctx context.Context, u *model.User) error {
	query := `
		INSERT INTO users (full_name, phone, national_id, id_image_path, is_active)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`
	return r.DB.QueryRow(ctx, query,
		u.FullName,
		u.Phone,
		u.NationalID,
		u.IDImagePath,
		u.IsActive,
	).Scan(&u.ID, &u.CreatedAt)
}
func (r *UserRepository) GetByNationalID(ctx context.Context, nationalID string) (*model.User, error) {
	var u model.User
	err := r.DB.QueryRow(ctx, `
		SELECT id, full_name, phone, national_id, id_image_path, is_active, created_at
		FROM users
		WHERE national_id = $1
	`, nationalID).Scan(
		&u.ID,
		&u.FullName,
		&u.Phone,
		&u.NationalID,
		&u.IDImagePath,
		&u.IsActive,
		&u.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &u, nil
}