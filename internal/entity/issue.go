package entity

import (
	"time"

	"github.com/google/uuid"
)

type Issue struct {
	ID          uuid.UUID `gorm:"primaryKey,default:uuid_generate_v4()"`
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
