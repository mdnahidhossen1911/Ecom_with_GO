package repo

import (
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"`
	IsOwner   bool      `json:"is_owner" db:"is_owner"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Get(userID string) (*User, error)
	Find(email string, password string) (*User, error)

	List() ([]*User, error)
	Delete(userID string) error
	Update(user User) (*User, error)
}

type userRepo struct {
	dbCon *sqlx.DB
}

func NewUserRepo(dbCon *sqlx.DB) UserRepo {
	return &userRepo{
		dbCon: dbCon,
	}
}

func (u *userRepo) Create(user User) (*User, error) {

	query := `INSERT INTO users (name, email, password, is_owner) 
	VALUES (:name, :email, :password, :is_owner)
	 RETURNING id`

	rows, err := u.dbCon.NamedQuery(query, user)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&user.ID)
	}

	if user.ID != "" {
		return &user, nil
	}

	return &user, nil
}

func (u *userRepo) Find(email string, password string) (*User, error) {

	query := `
		SELECT id, name, email, password, is_owner, created_at, updated_at
		FROM users
		WHERE email = $1 AND password = $2
		LIMIT 1;
	`

	var user User

	err := u.dbCon.Get(&user, query, email, password)
	if err != nil {
		return nil, errors.New("user not found " + err.Error())
	}

	return &user, nil
}

func (u *userRepo) Get(userID string) (*User, error) {
	query := `
		SELECT id, name, email, password, is_owner, created_at, updated_at
		FROM users
		WHERE id = $1
		LIMIT 1;
	`

	var user User

	err := u.dbCon.Get(&user, query, userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (u *userRepo) List() ([]*User, error) {
	query := `
		SELECT id, name, email, password, is_owner, created_at, updated_at
		FROM users;
	`

	var users []*User

	err := u.dbCon.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepo) Update(user User) (*User, error) {
	query := `
		UPDATE users
		SET 
			name = $1,
			email = $2,
			password = $3,
			is_owner = $4,
			updated_at = NOW()
		WHERE id = $5
		RETURNING id, name, email, password, is_owner, created_at, updated_at;
	`

	var updatedUser User

	err := u.dbCon.QueryRow(
		query,
		user.Name,
		user.Email,
		user.Password,
		user.IsOwner,
		user.ID,
	).Scan(
		&updatedUser.ID,
		&updatedUser.Name,
		&updatedUser.Email,
		&updatedUser.Password,
		&updatedUser.IsOwner,
	)

	if err != nil {
		return nil, errors.New("user not found")
	}

	return &updatedUser, nil
}

func (u *userRepo) Delete(userID string) error {
	query := `
		DELETE FROM users
		WHERE id = $1;
	`

	result, err := u.dbCon.Exec(query, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}
