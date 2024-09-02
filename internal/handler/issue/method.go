package issue

import (
	"net/http"
	"sitemate-challenge-server/internal/entity"
	"sitemate-challenge-server/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateIssue(c *gin.Context) {
	var issue entity.Issue
	if err := c.ShouldBindJSON(&issue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	issue.ID = uuid.New() // Generate a new UUID for the Issue

	newIssue, err := h.repo.Create(&issue)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.SendResponse(c, http.StatusCreated, "Issue Created", newIssue.ToDTO())
}

// GetIssueByID handles retrieving an Issue by ID
func (h *Handler) GetIssueByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, "Invalid UUID", nil)
		return
	}

	issue, err := h.repo.GetByID(id)
	if err != nil {
		utils.SendResponse(c, http.StatusNotFound, "Issue not found", nil)
		return
	}

	utils.SendResponse(c, http.StatusOK, "Issue Found", issue.ToDTO())
}

// UpdateIssue handles updating an Issue
func (h *Handler) UpdateIssue(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, "Invalid UUID", nil)
		return
	}

	var issue entity.Issue
	if err := c.ShouldBindJSON(&issue); err != nil {
		utils.SendResponse(c, http.StatusBadRequest, "Invalid Request Body", err)
		return
	}

	_, err = h.repo.GetByID(id)
	if err != nil {
		utils.SendResponse(c, http.StatusNotFound, "Issue Not Found", nil)
		return
	}

	issue.ID = id // Use the ID from the URL

	if err := h.repo.Update(&issue); err != nil {
		utils.SendResponse(c, http.StatusInternalServerError, "Something went wrong", err)
		return
	}

	utils.SendResponse(c, http.StatusOK, "Issue Updated", issue.ToDTO())
}

// DeleteIssue handles deleting an Issue
func (h *Handler) DeleteIssue(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, "Invalid UUID", nil)
		return
	}

	_, err = h.repo.GetByID(id)
	if err != nil {
		utils.SendResponse(c, http.StatusNotFound, "Issue Not Found", nil)
		return
	}

	if err := h.repo.Delete(id); err != nil {
		utils.SendResponse(c, http.StatusInternalServerError, "Something went wrong", err)
		return
	}

	utils.SendResponse(c, http.StatusOK, "Issue Deleted", nil)
}

// GetAllIssues handles retrieving all Issues
func (h *Handler) GetAllIssues(c *gin.Context) {
	searchQ := c.Query("search")

	issues, err := h.repo.GetAll(searchQ)
	if err != nil {
		utils.SendResponse(c, http.StatusInternalServerError, "Something went wrong", err)
		return
	}

	dtos := []entity.IssueDTO{}
	for _, dao := range issues {
		dtos = append(dtos, dao.ToDTO())
	}

	utils.SendResponse(c, http.StatusOK, "Issue Data", dtos)
}
