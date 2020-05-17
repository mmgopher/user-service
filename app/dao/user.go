package dao

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/mmgopher/user-service/app/model"
)

// UserRepositoryProvider provides an interface to work with database User entity
type UserRepositoryProvider interface {
	// GetByID returns User object by ID
	GetByID(id int) (*model.User, error)
	// Create creates new User record
	Create(user *model.User) (int, error)
	// Update updates user record
	Update(user *model.User) (bool, error)
	// Delete deletes user record
	Delete(userID int) (bool, error)
	// CheckIfExistWithNameAndSurname checks if user with provided name and surname exists in DB
	// Returns TRUE if user already exist
	CheckIfExistWithNameAndSurname(name, surname string) (bool, error)
}

// UserRepository represents object to work with  database User entity
type UserRepository struct {
	db *sqlx.DB
}

// NewUserRepository creates new instance of Repository.
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// CheckIfExistWithNameAndSurname checks if user with provided name and surname exists in DB
// Returns TRUE if user already exist
func (r UserRepository) CheckIfExistWithNameAndSurname(name, surname string) (bool, error) {
	var total int
	if err := r.db.Get(&total, `
		SELECT count(*)
		FROM user_sch.user
		WHERE name = $1
		AND surname = $2`, name, surname,
	); err != nil {
		return false, errors.Wrap(err, "impossible to get count of users")
	}

	return total > 0, nil
}

// GetByID returns User object by ID
func (r UserRepository) GetByID(userID int) (*model.User, error) {
	var user model.User
	if err := r.db.Get(&user, `
		SELECT id,
		       name,
		       surname,
		       gender,
		       age,
			   address,
			   created_at
		FROM user_sch.user
		WHERE id = $1`, userID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, errors.Wrapf(err, "impossible to get user, userID=%d", userID)
	}

	return &user, nil
}

// Create creates new User record
func (r UserRepository) Create(user *model.User) (int, error) {

	err := r.db.QueryRow(`
	INSERT INTO user_sch.user(
		name,                                   
		surname,
		gender,
		age,
		address
	) VALUES (
		 $1, $2, $3, $4, $5
	) RETURNING id`,
		user.Name,
		user.Surname,
		user.Gender,
		user.Age,
		user.Address,
	).Scan(&user.ID)

	if err != nil {
		return 0, errors.Wrap(err, "impossible to create user record")
	}

	return user.ID, nil
}

// Update updates user record
func (r UserRepository) Update(user *model.User) (bool, error) {
	res, err := r.db.Exec(`
	UPDATE user_sch.user
	SET  name = $1,
		 surname = $2,
		 gender = $3,
		 age = $4,
		 address = $5
	WHERE id = $6`,
		user.Name,
		user.Surname,
		user.Gender,
		user.Age,
		user.Address,
		user.ID,
	)

	if err != nil {
		return false, errors.Wrap(err, "impossible to update user record")
	}

	count, err := res.RowsAffected()
	if err != nil {
		return false, errors.Wrap(err, "impossible to check if user  was updated")
	}

	return count == 1, nil
}

// Delete deletes user record
func (r UserRepository) Delete(userID int) (bool, error) {
	res, err := r.db.Exec(`
	DELETE from  user_sch.user
	WHERE id = $1`,
		userID,
	)

	if err != nil {
		return false, errors.Wrap(err, "impossible to delete user record")
	}

	count, err := res.RowsAffected()
	if err != nil {
		return false, errors.Wrap(err, "impossible to check if user  was deleted")
	}
	return count == 1, nil
}
