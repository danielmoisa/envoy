package controller

import (
	"net/http"

	"github.com/danielmoisa/envoy/src/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetAllCandidates returns a list of all candidates
// @Summary Get all candidates
// @Description Get a list of all candidates
// @Tags Candidates
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {array} model.Candidate
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 403 {object} map[string]string "Forbidden"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /candidates [get]
func (ctrl *Controller) GetAllCandidates(c *gin.Context) {
	candidates, err := ctrl.Repository.CandidatesRepository.GetAll()
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_CANDIDATE_GET_FAILED, err.Error())
		return
	}

	c.JSON(http.StatusOK, candidates)
}

// GetCandidate returns a single candidate by ID
// @Summary Get candidate by ID
// @Description Get a single candidate by its ID
// @Tags Candidates
// @Accept json
// @Produce json
// @Security Bearer
// @Param candidateId path string true "Candidate ID"
// @Success 200 {object} model.Candidate
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 403 {object} map[string]string "Forbidden"
// @Failure 404 {object} map[string]string "Not Found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /candidates/{candidateId} [get]
func (ctrl *Controller) GetCandidate(c *gin.Context) {
	id, err := uuid.Parse(c.Param("candidateId"))
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_INPUT, "Invalid candidate ID format: "+err.Error())
		return
	}

	candidate, err := ctrl.Repository.CandidatesRepository.GetByID(id)
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_CANDIDATE_NOT_FOUND, "Candidate not found: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, candidate)
}

// CreateCandidate creates a new candidate
// @Summary Create candidate
// @Description Create a new candidate
// @Tags Candidates
// @Accept json
// @Produce json
// @Security Bearer
// @Param candidate body model.Candidate true "Candidate Data"
// @Success 201 {object} model.Candidate
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 403 {object} map[string]string "Forbidden"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /candidates [post]
func (ctrl *Controller) CreateCandidate(c *gin.Context) {
	var candidate model.Candidate
	if err := c.ShouldBindJSON(&candidate); err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_INPUT, err.Error())
		return
	}

	createdCandidate, err := ctrl.Repository.CandidatesRepository.Create(&candidate)
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_CANDIDATE_CREATE_FAILED, err.Error())
		return
	}

	c.JSON(http.StatusCreated, createdCandidate)
}

// UpdateCandidate updates an existing candidate
// @Summary Update candidate
// @Description Update an existing candidate by ID
// @Tags Candidates
// @Accept json
// @Produce json
// @Security Bearer
// @Param candidateId path string true "Candidate ID"
// @Param candidate body model.Candidate true "Candidate Data"
// @Success 200 {object} model.Candidate
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 403 {object} map[string]string "Forbidden"
// @Failure 404 {object} map[string]string "Not Found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /candidates/{candidateId} [put]
func (ctrl *Controller) UpdateCandidate(c *gin.Context) {
	id, err := uuid.Parse(c.Param("candidateId"))
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_INPUT, "Invalid candidate ID format: "+err.Error())
		return
	}

	var candidate model.Candidate
	if err := c.ShouldBindJSON(&candidate); err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_INPUT, err.Error())
		return
	}

	candidate.ID = id

	updatedCandidate, err := ctrl.Repository.CandidatesRepository.Update(&candidate)
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_CANDIDATE_UPDATE_FAILED, err.Error())
		return
	}

	c.JSON(http.StatusOK, updatedCandidate)
}

// DeleteCandidate deletes a candidate by ID
// @Summary Delete candidate
// @Description Delete a candidate by ID
// @Tags Candidates
// @Accept json
// @Produce json
// @Security Bearer
// @Param candidateId path string true "Candidate ID"
// @Success 204 "No Content"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 403 {object} map[string]string "Forbidden"
// @Failure 404 {object} map[string]string "Not Found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /candidates/{candidateId} [delete]
func (ctrl *Controller) DeleteCandidate(c *gin.Context) {
	id, err := uuid.Parse(c.Param("candidateId"))
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_INPUT, "Invalid candidate ID format: "+err.Error())
		return
	}

	if err := ctrl.Repository.CandidatesRepository.Delete(id); err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_CANDIDATE_DELETE_FAILED, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
