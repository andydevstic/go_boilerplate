package core

import (
	"database/sql"
)

type AppState struct {
	Store *sql.DB
}

var appState AppState

// Generates singleton app state
func GenerateAppState(store *sql.DB) {
	appState.Store = store
}

func GetAppState() *AppState {
	return &appState
}
