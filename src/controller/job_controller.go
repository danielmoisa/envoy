package controller

import (
	"net/http"

	"github.com/danielmoisa/envoy/src/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetAllJobs returns a list of all jobs
// @Summary Get all jobs
// @Description Get a list of all jobs
// @Tags Jobs
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {array} model.Job
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 403 {object} map[string]string "Forbidden"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /jobs [get]
func (ctrl *Controller) GetAllJobs(c *gin.Context) {
	jobs, err := ctrl.Repository.JobsRepository.GetAll()
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_JOB_GET_FAILED, err.Error())
		return
	}

	c.JSON(http.StatusOK, jobs)
}

// GetJob returns a single job by ID
// @Summary Get job by ID
// @Description Get a single job by its ID
// @Tags Jobs
// @Accept json
// @Produce json
// @Security Bearer
// @Param jobId path string true "Job ID"
// @Success 200 {object} model.Job
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 403 {object} map[string]string "Forbidden"
// @Failure 404 {object} map[string]string "Not Found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /jobs/{jobId} [get]
func (ctrl *Controller) GetJob(c *gin.Context) {
	id, err := uuid.Parse(c.Param("jobId"))
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_INPUT, "Invalid job ID format: "+err.Error())
		return
	}

	job, err := ctrl.Repository.JobsRepository.GetByID(id)
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_JOB_NOT_FOUND, "Job not found: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, job)
}

// CreateJob creates a new job
// @Summary Create job
// @Description Create a new job
// @Tags Jobs
// @Accept json
// @Produce json
// @Security Bearer
// @Param job body model.Job true "Job Data"
// @Success 201 {object} model.Job
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 403 {object} map[string]string "Forbidden"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /jobs [post]
func (ctrl *Controller) CreateJob(c *gin.Context) {
	var job model.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_INPUT, err.Error())
		return
	}

	createdJob, err := ctrl.Repository.JobsRepository.Create(job.Title, job.Description, job.Location, job.JobType, job.SalaryMin, job.SalaryMax, job.CompanyID)
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_JOB_CREATE_FAILED, err.Error())
		return
	}

	c.JSON(http.StatusCreated, createdJob)
}

// UpdateJob updates an existing job
// @Summary Update job
// @Description Update an existing job by ID
// @Tags Jobs
// @Accept json
// @Produce json
// @Security Bearer
// @Param jobId path string true "Job ID"
// @Param job body model.Job true "Job Data"
// @Success 200 {object} model.Job
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 403 {object} map[string]string "Forbidden"
// @Failure 404 {object} map[string]string "Not Found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /jobs/{jobId} [put]
func (ctrl *Controller) UpdateJob(c *gin.Context) {
	id, err := uuid.Parse(c.Param("jobId"))
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_INPUT, "Invalid job ID format: "+err.Error())
		return
	}

	var job model.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_INPUT, err.Error())
		return
	}

	updatedJob, err := ctrl.Repository.JobsRepository.Update(id, job.Title, job.Description, job.Location, job.JobType, job.SalaryMin, job.SalaryMax)
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_JOB_UPDATE_FAILED, err.Error())
		return
	}

	c.JSON(http.StatusOK, updatedJob)
}

// DeleteJob deletes a job by ID
// @Summary Delete job
// @Description Delete a job by ID
// @Tags Jobs
// @Accept json
// @Produce json
// @Security Bearer
// @Param jobId path string true "Job ID"
// @Success 204 "No Content"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 403 {object} map[string]string "Forbidden"
// @Failure 404 {object} map[string]string "Not Found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /jobs/{jobId} [delete]
func (ctrl *Controller) DeleteJob(c *gin.Context) {
	id, err := uuid.Parse(c.Param("jobId"))
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_INPUT, "Invalid job ID format: "+err.Error())
		return
	}

	if err := ctrl.Repository.JobsRepository.Delete(id); err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_JOB_DELETE_FAILED, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
