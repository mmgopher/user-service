package dao

import (
	"fmt"
)

var (
	// Asc - ascending sort order
	Asc = PagingSortOrder{
		orderNext:        "ASC",
		operatorNext:     ">",
		orderPrevious:    "DESC",
		operatorPrevious: "<",
	}
	// Desc - descending sort order
	Desc = PagingSortOrder{
		orderNext:        "DESC",
		operatorNext:     "<",
		orderPrevious:    "ASC",
		operatorPrevious: ">",
	}
)

// PagingSortOrder represents sort order used in ORDER BY clause
// and operator used in WHERE clause
type PagingSortOrder struct {
	orderNext        string
	operatorNext     string
	orderPrevious    string
	operatorPrevious string
}

// GetOperator returns comparison operator based on direction
// Direction is the bool value which defines if next or previous page is queried.
// It returns '>' or '<'
func (p PagingSortOrder) GetOperator(nextPage bool) string {
	if nextPage {
		return p.operatorNext
	}
	return p.operatorPrevious
}

// GetOrder returns order used in Group By clause
func (p PagingSortOrder) GetOrder(nextPage bool) string {
	if nextPage {
		return p.orderNext
	}
	return p.orderPrevious
}

// PagingSearchBuilder represents input parameters used to search using paging.
type PagingSearchBuilder struct {
	PrimaryKey string
	Limit      int
	StartID    int
	SortColumn string
	SortOrder  PagingSortOrder
	NextPage   bool
}

// NewPagingSearchBuilder creates new instance of Paging Search Builder.
func NewPagingSearchBuilder(limit, afterID, beforeID int, sortColumn, sortOrdeName, primaryKey string,
) PagingSearchBuilder {

	startID := 0
	next := true

	if afterID > 0 {
		startID = afterID
	} else if beforeID > 0 {
		startID = beforeID
		next = false
	}

	if sortColumn == "" {
		sortColumn = primaryKey
	}

	var sortOrder PagingSortOrder

	switch sortOrdeName {
	case "desc":
		sortOrder = Desc
	case "asc":
		sortOrder = Asc
	default:
		sortOrder = Asc
	}

	return PagingSearchBuilder{
		PrimaryKey: primaryKey,
		Limit:      limit,
		StartID:    startID,
		SortColumn: sortColumn,
		SortOrder:  sortOrder,
		NextPage:   next,
	}
}

func (b PagingSearchBuilder) getSortOrder() string {
	return b.SortOrder.GetOrder(b.NextPage)
}

func (b PagingSearchBuilder) getWhereOperator() string {
	return b.SortOrder.GetOperator(b.NextPage)
}

func (b PagingSearchBuilder) isSortByPrimaryKey() bool {
	return b.SortColumn == b.PrimaryKey
}

// GetOrderByCriteria returns order criteria.
// If results is sorted by primary key it returns primary_key ASC/DESC
// In query it is used like this: ORDER BY primary_key ASC/DESC
// if it is sorted by another column it returns  column_name ASC/DESC, primary_key
// In query it is used like this: column_name ASC/DESC, primary_key
func (b PagingSearchBuilder) GetOrderByCriteria() string {
	if b.isSortByPrimaryKey() {
		return fmt.Sprintf("%s %s", b.PrimaryKey, b.getSortOrder())
	}

	return fmt.Sprintf("%s %s,%s", b.SortColumn, b.getSortOrder(), b.PrimaryKey)

}

// GetWhereCriteria returns "where" criteria used to know the place where to start quering next or prevoius page
func (b PagingSearchBuilder) GetWhereCriteria() (whereCondition string, args []interface{}) {

	if b.StartID > 0 {
		whereCondition = fmt.Sprintf("%s %s", b.SortColumn, b.getWhereOperator())
		if b.isSortByPrimaryKey() {
			args = []interface{}{b.StartID}
		}
	} else {
		whereCondition = "1 ="
		args = []interface{}{1}
	}

	return
}
