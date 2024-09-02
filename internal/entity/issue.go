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

type IssueDTO struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (d *Issue) ToDTO() IssueDTO {
	return IssueDTO{
		ID:          d.ID,
		Title:       d.Title,
		Description: d.Description,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}
