package controller

import (
	"net/http"

	"github.com/danielmoisa/envoy/src/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetAllApplications returns a list of all applications
// @Summary Get all applications
// @Description Get a list of all applications
// @Tags Applications
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {array} model.Application
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 403 {object} map[string]string "Forbidden"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /applications [get]
func (ctrl *Controller) GetAllApplications(c *gin.Context) {
	applications, err := ctrl.Repository.ApplicationsRepository.GetAll()
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_APPLICATION_GET_FAILED, err.Error())
		return
	}

	c.JSON(http.StatusOK, applications)
}

// GetApplication returns a single application by ID
// @Summary Get application by ID
// @Description Get a single application by its ID
// @Tags Applications
// @Accept json
// @Produce json
// @Security Bearer
// @Param applicationId path string true "Application ID"
// @Success 200 {object} model.Application
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 403 {object} map[string]string "Forbidden"
// @Failure 404 {object} map[string]string "Not Found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /applications/{applicationId} [get]
func (ctrl *Controller) GetApplication(c *gin.Context) {
	id, err := uuid.Parse(c.Param("applicationId"))
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_INPUT, "Invalid application ID format: "+err.Error())
		return
	}

	application, err := ctrl.Repository.ApplicationsRepository.GetByID(id)
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_APPLICATION_NOT_FOUND, "Application not found: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, application)
}

// CreateApplication creates a new application
// @Summary Create application
// @Description Create a new application
// @Tags Applications
// @Accept json
// @Produce json
// @Security Bearer
// @Param application body model.Application true "Application Data"
// @Success 201 {object} model.Application
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 403 {object} map[string]string "Forbidden"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /applications [post]
func (ctrl *Controller) CreateApplication(c *gin.Context) {
	var application model.Application
	if err := c.ShouldBindJSON(&application); err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_INPUT, err.Error())
		return
	}

	createdApplication, err := ctrl.Repository.ApplicationsRepository.Create(&application)
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_APPLICATION_CREATE_FAILED, err.Error())
		return
	}

	c.JSON(http.StatusCreated, createdApplication)
}

// UpdateApplication updates an existing application
// @Summary Update application
// @Description Update an existing application by ID
// @Tags Applications
// @Accept json
// @Produce json
// @Security Bearer
// @Param applicationId path string true "Application ID"
// @Param application body model.Application true "Application Data"
// @Success 200 {object} model.Application
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 403 {object} map[string]string "Forbidden"
// @Failure 404 {object} map[string]string "Not Found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /applications/{applicationId} [put]
func (ctrl *Controller) UpdateApplication(c *gin.Context) {
	id, err := uuid.Parse(c.Param("applicationId"))
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_INPUT, "Invalid application ID format: "+err.Error())
		return
	}

	var application model.Application
	if err := c.ShouldBindJSON(&application); err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_INPUT, err.Error())
		return
	}

	application.ID = id
	updatedApplication, err := ctrl.Repository.ApplicationsRepository.Update(&application)
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_APPLICATION_UPDATE_FAILED, err.Error())
		return
	}

	c.JSON(http.StatusOK, updatedApplication)
}

// DeleteApplication deletes a application by ID
// @Summary Delete application
// @Description Delete a application by ID
// @Tags Applications
// @Accept json
// @Produce json
// @Security Bearer
// @Param applicationId path string true "Application ID"
// @Success 204 "No Content"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 403 {object} map[string]string "Forbidden"
// @Failure 404 {object} map[string]string "Not Found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /applications/{applicationId} [delete]
func (ctrl *Controller) DeleteApplication(c *gin.Context) {
	id, err := uuid.Parse(c.Param("applicationId"))
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_INPUT, "Invalid application ID format: "+err.Error())
		return
	}

	if err := ctrl.Repository.ApplicationsRepository.Delete(id); err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_APPLICATION_DELETE_FAILED, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
