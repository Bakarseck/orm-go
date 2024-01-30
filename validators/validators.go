package validators

import (
	_ "github.com/mattn/go-sqlite3"
)

type ValidationResult struct {
	Field  string
	Tag    string
	Valid  bool
	Reason string
}
