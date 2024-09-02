package issue

type Handler struct {
	repo IssueRepository
}

func New(issueRepo IssueRepository) *Handler {
	return &Handler{
		repo: issueRepo,
	}
}
