package issue

import (
	"sitemate-challenge-server/internal/entity"

	"github.com/google/uuid"
)

type IssueRepository interface {
	Create(issue *entity.Issue) (*entity.Issue, error)
	GetAll() ([]entity.Issue, error)
	GetByID(id uuid.UUID) (*entity.Issue, error)
	Update(issue *entity.Issue) error
	Delete(id uuid.UUID) error
}
