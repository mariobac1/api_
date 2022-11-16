package models

import "time"

type Community struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Communities slice of Community
type Communities []Community
