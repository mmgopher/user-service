package db

import (
	"errors"
	"reflect"
)

// FindColumnNames returns set of db columns names for provided struct
// It searches for tag `db:column-name` in every struc field to gather column names.
func FindColumnNames(model interface{}) (map[string]struct{}, error) {
	columnNames := map[string]struct{}{}
	v := reflect.ValueOf(model)
	if v.Kind() != reflect.Struct {
		return columnNames, errors.New("only struct type allowed as input parameter")
	}
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		tag := typeOfS.Field(i).Tag
		if v, ok := tag.Lookup("db"); ok {
			columnNames[v] = struct{}{}
		}

	}
	return columnNames, nil
}
