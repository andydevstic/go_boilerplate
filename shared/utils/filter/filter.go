package filter

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/andydevstic/boilerplate-backend/shared/custom"
	"gorm.io/gorm"
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

func AddFilterToStatement(query *gorm.DB, columnName string, filterString string) error {
	if len(filterString) == 0 {
		return custom.NewError(http.StatusBadGateway, fmt.Errorf("invalid filter string: %s", filterString))
	}

	splitted := strings.Split(filterString, ":")
	if len(splitted) != 2 {
		return custom.NewError(http.StatusBadGateway, fmt.Errorf("invalid filter string: %s", filterString))
	}

	operator := splitted[0]
	filterValue := ParseScalarValueType(splitted[1])

	switch operator {
	case "gte":
		query.Where(fmt.Sprintf("%s >= ?", columnName), filterValue)
	case "lte":
		query.Where(fmt.Sprintf("%s <= ?", columnName), filterValue)
	case "eq":
		query.Where(fmt.Sprintf("%s = ?", columnName), filterValue)
	case "in":
		query.Where(fmt.Sprintf("%s IN (?)", columnName), filterValue)
	default:
		return errors.New("operator not supported")
	}

	return nil
}
