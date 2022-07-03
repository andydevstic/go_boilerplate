package filter

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/andydevstic/boilerplate-backend/shared"
	"github.com/andydevstic/boilerplate-backend/shared/custom"
	"github.com/doug-martin/goqu/v9"
)

func ParseScalarValueType(rawValue string) interface{} {
	if value, err := strconv.Atoi(rawValue); err == nil {
		return value
	}

	if value, err := strconv.ParseFloat(rawValue, 64); err == nil {
		return value
	}

	if value, err := strconv.ParseBool(rawValue); err == nil {
		return value
	}

	return rawValue
}

func ParseFilterString(columnName string, filterString string) (parsed goqu.Expression, err error) {
	if len(filterString) == 0 {
		return nil, custom.NewError(http.StatusBadGateway, fmt.Errorf("invalid filter string: %s", filterString))
	}

	splitted := strings.Split(filterString, ":")
	if len(splitted) != 2 {
		return nil, custom.NewError(http.StatusBadGateway, fmt.Errorf("invalid filter string: %s", filterString))
	}

	operator := splitted[0]
	filterValue := splitted[1]

	columnFilter := shared.ColumnFilter{
		Column:   columnName,
		Operator: operator,
		Value:    filterValue,
	}

	return parseFilterCondition(&columnFilter)
}

func parseFilterCondition(rawFilter *shared.ColumnFilter) (parsed goqu.Expression, err error) {
	column := goqu.C(rawFilter.Column)
	value := ParseScalarValueType(rawFilter.Value)
	parsed, err = nil, nil

	switch rawFilter.Operator {
	case "gte":
		parsed = column.Gte(value)
	case "lte":
		parsed = column.Lt(value)
	case "eq":
		parsed = column.Eq(value)
	case "in":
		parsed = goqu.L(fmt.Sprintf("%s IN (%s)", rawFilter.Column, rawFilter.Value))
	default:
		err = errors.New("operator not supported")
	}

	return
}
