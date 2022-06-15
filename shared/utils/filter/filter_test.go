package filter_test

import (
	"testing"

	"github.com/andydevstic/boilerplate-backend/shared/utils/filter"
	"github.com/doug-martin/goqu/v9"
	"github.com/stretchr/testify/assert"
)

func TestParseInt(t *testing.T) {
	intString := "123"

	parsed := filter.ParseScalarValueType(intString)
	parsed, ok := parsed.(int)

	assert.Equal(t, ok, true, "should be able to convert to int")
	assert.Equal(t, parsed, 123, "should get parsed int correctly")
}

func TestParseIntFailWhenProvideNaN(t *testing.T) {
	invalidIntString := "12s"

	parsed := filter.ParseScalarValueType(invalidIntString)
	parsed, ok := parsed.(int)

	assert.Equal(t, ok, false, "should not be able to convert to int")
}

func TestParseFloat(t *testing.T) {
	floatString := "123.5"

	parsed := filter.ParseScalarValueType(floatString)
	parsed, ok := parsed.(float64)

	assert.Equal(t, ok, true, "should be able to convert to float64")
	assert.Equal(t, parsed, 123.5, "should get parsed float64 correctly")
}

func TestParseFloatFailWhenProvideNaN(t *testing.T) {
	invalidFloatString := "12.5s"

	parsed := filter.ParseScalarValueType(invalidFloatString)
	parsed, ok := parsed.(int)

	assert.Equal(t, ok, false, "should not be able to convert to float")
}

func TestParseEqFilterString(t *testing.T) {
	filterString1 := "eq:andy"
	filterString2 := "gte:1"

	testSql := goqu.Dialect("postgres").From("users")

	parsed1, err := filter.ParseFilterString("name", filterString1)
	assert.Nil(t, err, "Should be able to parse name filter string")

	parsed2, err := filter.ParseFilterString("status", filterString2)
	assert.Nil(t, err, "Should be able to parse status filter string")

	parsedSQL, _, err := testSql.Where(parsed1).Where(parsed2).ToSQL()
	assert.Nil(t, err, "Should be able to combine filter string into sql")

	assert.Equal(t, parsedSQL, "SELECT * FROM \"users\" WHERE ((\"name\" = 'andy') AND (\"status\" >= 1))")
}
