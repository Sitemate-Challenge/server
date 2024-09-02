package issue

import (
	"sitemate-challenge-server/internal/entity"

	"github.com/google/uuid"
)

func (r *Repository) Create(issue *entity.Issue) (*entity.Issue, error) {
	if err := r.db.Create(issue).Error; err != nil {
		return nil, err
	}
	return issue, nil
}

func (r *Repository) GetAll(search string) ([]entity.Issue, error) {
	var issues []entity.Issue
	query := r.db.Model(&entity.Issue{})

	if search != "" {
		query = query.Where("title ILIKE ?", "%"+search+"%")
	}

	err := query.Find(&issues).Error
	return issues, err
}

func (r *Repository) GetByID(id uuid.UUID) (*entity.Issue, error) {
	var issue entity.Issue
	err := r.db.First(&issue, "id = ?", id).Error
	return &issue, err
}

func (r *Repository) Update(issue *entity.Issue) error {
	return r.db.Save(issue).Error
}

func (r *Repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entity.Issue{}, "id = ?", id).Error
}
