package shared

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type ICrudService[Model any] interface {
	FindOne(context.Context, map[string]any) (Model, error)
	Find(ctx context.Context, model map[string]any, limit int, offset int) (result []Model, err error)
	Create(context.Context, map[string]any) error
	Update(context.Context, map[string]any) error
	Delete(context.Context, map[string]any) error
}

type ColumnFilter struct {
	Column   string
	Operator string
	Value    string
}

type FilterConditions = []map[string]ColumnFilter
