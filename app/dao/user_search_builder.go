package dao

import (
	"fmt"
	"strings"

	"github.com/mmgopher/user-service/app/api/request"
)

const (
	userListDefaultPageSize = 30
	userListMaxPageSize     = 200
)

// UserSearchBuilder represents input parameters used to find creator activity data.
type UserSearchBuilder struct {
	filter userFilter
	PagingSearchBuilder
}

// userFilter collects filter criteria for user search.
type userFilter struct {
	name    string
	surname string
	gender  string
	address string
	minAge  int
	maxAge  int
}

// NewUserSearchBuilder creates new instance of User Search Builder.
func NewUserSearchBuilder(request *request.FindUsers) *UserSearchBuilder {

	rowsToReturn := request.Limit

	if request.Limit == 0 || request.Limit > userListMaxPageSize {
		rowsToReturn = userListDefaultPageSize
	}

	var sortColumn string
	var sortOrder string
	if len(request.Sort) > 0 {
		s := strings.Split(request.Sort, ":")
		sortColumn = s[0]
		sortOrder = s[1]
	}

	pagingSearchBuilder := NewPagingSearchBuilder(
		rowsToReturn,
		request.AfterID,
		request.BeforeID,
		sortColumn,
		sortOrder,
		"id")

	return &UserSearchBuilder{
		filter: userFilter{
			name:    request.Name,
			surname: request.Surname,
			gender:  request.Gender,
			address: request.Address,
			minAge:  request.MinAge,
			maxAge:  request.MaxAge,
		},
		PagingSearchBuilder: pagingSearchBuilder,
	}
}

// GetFilterCriteria returns "filter" criteria to filter results
func (usb UserSearchBuilder) GetFilterCriteria() (string, []interface{}) {

	var sb strings.Builder
	sb.WriteString("1 = ?")
	args := []interface{}{1}

	if strings.TrimSpace(usb.filter.name) != "" {
		sb.WriteString(" AND name ilike ?")
		args = append(args, fmt.Sprintf("%%%s%%", usb.filter.name))
	}

	if strings.TrimSpace(usb.filter.surname) != "" {
		sb.WriteString(" AND surname ilike ?")
		args = append(args, fmt.Sprintf("%%%s%%", usb.filter.surname))
	}

	if strings.TrimSpace(usb.filter.gender) != "" {
		sb.WriteString(" AND gender ilike ?")
		args = append(args, fmt.Sprintf("%s", usb.filter.gender))
	}

	if strings.TrimSpace(usb.filter.surname) != "" {
		sb.WriteString(" AND address ilike ?")
		args = append(args, fmt.Sprintf("%%%s%%", usb.filter.address))
	}

	if usb.filter.minAge > 0 {
		sb.WriteString(" AND age >= ?")
		args = append(args, usb.filter.minAge)
	}

	if usb.filter.maxAge > 0 {
		sb.WriteString(" AND age <= ?")
		args = append(args, usb.filter.maxAge)
	}

	return sb.String(), args
}

// BuildSearchQuery builds final query with all criteria
func (usb UserSearchBuilder) BuildSearchQuery(whereCriteria, filterCriteria, orderByCriteria string) string {

	// nolint
	basicQuery := fmt.Sprintf(`
	SELECT
		id,
		name,
		surname,
		gender,
		age,
		address,
		created_at
	FROM user_sch.user
	WHERE %s ?
	AND %s
	ORDER BY %s
	LIMIT ?`,
		whereCriteria,
		filterCriteria,
		orderByCriteria,
	)
	if usb.NextPage {
		return basicQuery
	}

	return fmt.Sprintf(`
	SELECT * 
	FROM (%s) as alias
	ORDER BY alias.%s %s`,
		basicQuery,
		usb.SortColumn,
		usb.SortOrder.GetOrder(!usb.NextPage),
	)
}
