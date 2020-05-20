package dao

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	dbhelper "github.com/mmgopher/user-service/app/db"
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
	// GetByNameAndSurname returns User object by name and surname
	GetByNameAndSurname(name, surname string) (*model.User, error)
	// FindUsers finds users in database using pagination, sorting and filtering.
	FindUsers(sb *UserSearchBuilder,
	) ([]model.User, int, int, error)
}

// UserRepository represents object to work with  database User entity
type UserRepository struct {
	db               *sqlx.DB
	setOfUserColumns map[string]struct{}
}

// NewUserRepository creates new instance of Repository.
func NewUserRepository(db *sqlx.DB) *UserRepository {
	//I am not checking an error, because this function is called internally and I send Struct type as input parameter
	columns, _ := dbhelper.FindColumnNames(model.User{})
	return &UserRepository{
		db:               db,
		setOfUserColumns: columns,
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

// GetByNameAndSurname returns User object by name and surname
func (r UserRepository) GetByNameAndSurname(name, surname string) (*model.User, error) {
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
		WHERE name = $1
		AND surname = $2`, name, surname,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, errors.Wrap(err, "impossible to get user")
	}

	return &user, nil
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

// FindUsers finds users in database using pagination, sorting and filtering.
func (r UserRepository) FindUsers(sb *UserSearchBuilder,
) ([]model.User, int, int, error) {

	//Input sanitization for sort column name
	//Check if column provided as sort parameter is one of the columns on User entity
	if _, ok := r.setOfUserColumns[sb.SortColumn]; !ok {
		return nil, 0, 0, errors.Errorf("column used to sort does not exist in user table, sortColumn=%s",
			sb.SortColumn)
	}

	orderByCriteria := sb.GetOrderByCriteria()
	filterCriteria, filterArgs := sb.GetFilterCriteria()
	whereCriteria, whereArgs := sb.GetWhereCriteria()

	// Used when querying next page to find row to start db searching
	if len(whereArgs) == 0 {
		var whereValue interface{}
		if err := r.db.Get(&whereValue,
			// nolint
			fmt.Sprintf(`
				SELECT %s
				FROM user_sch.user
				WHERE id = $1`,
				sb.SortColumn,
			), sb.StartID); err != nil {
			if err == sql.ErrNoRows {
				return nil, 0, 0, errors.Errorf("row with start ID for next page not found, creator_profile_id=%d",
					sb.StartID)
			}
			return nil, 0, 0, err
		}
		whereArgs = append(whereArgs, whereValue)
	}

	args := append(whereArgs, filterArgs...)
	args = append(args, sb.Limit+1)

	query := sb.BuildSearchQuery(whereCriteria, filterCriteria, orderByCriteria)
	query = r.db.Rebind(query)

	var users []model.User
	if err := r.db.Select(&users, query, args...); err != nil {
		return nil, 0, 0, err
	}

	rowsReturnedCount := len(users)
	if rowsReturnedCount == 0 {
		return []model.User{}, 0, 0, nil
	}

	var beforeID int
	var afterID int
	var usersToReturn []model.User
	if sb.NextPage {
		if sb.StartID > 0 {
			beforeID = users[0].ID
		}
		if rowsReturnedCount > sb.Limit {
			afterID = users[rowsReturnedCount-2].ID
			usersToReturn = users[:rowsReturnedCount-1]
		} else {
			usersToReturn = users
		}
	} else {
		afterID = users[rowsReturnedCount-1].ID
		if rowsReturnedCount > sb.Limit {
			usersToReturn = users[1:rowsReturnedCount]
			beforeID = users[1].ID
		} else {
			usersToReturn = users
		}
	}

	return usersToReturn, beforeID, afterID, nil
}
