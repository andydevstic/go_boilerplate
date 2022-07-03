package core

import (
	"gorm.io/gorm"
)

type AppState struct {
	Db *gorm.DB
}

var appState AppState

// Generates singleton app state
func GenerateAppState(store *gorm.DB) {
	appState.Db = store
}

func GetAppState() *AppState {
	return &appState
}
