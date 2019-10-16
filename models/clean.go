package models

import "time"

type Clean struct {
	ID       int       `json:"id"`
	Optional time.Time `json:"optional, omitempty"` // klo di vscode ini dapat warning gara" si omitempty tp bisa diabaikan
}
