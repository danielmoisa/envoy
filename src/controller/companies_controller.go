package controller

import (
	"net/http"

	"github.com/danielmoisa/envoy/src/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetAllCompanies returns a list of all companies
// @Summary Get all companies
// @Description Get a list of all companies
// @Tags Companies
// @Accept json
// @Produce json
// @Success 200 {array} model.Company
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /companies [get]
func (ctrl *Controller) GetAllCompanies(c *gin.Context) {
	companies, err := ctrl.Repository.CompaniesRepository.GetAll()
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_COMPANY_GET_FAILED, err.Error())
		return
	}

	c.JSON(http.StatusOK, companies)
}

// GetCompany returns a single company by ID
// @Summary Get company by ID
// @Description Get a single company by its ID
// @Tags Companies
// @Accept json
// @Produce json
// @Param companyId path string true "Company ID"
// @Success 200 {object} model.Company
// @Failure 404 {object} map[string]string "Not Found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /companies/{companyId} [get]
func (ctrl *Controller) GetCompany(c *gin.Context) {
	id, err := uuid.Parse(c.Param("companyId"))
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_INPUT, "Invalid company ID format: "+err.Error())
		return
	}

	company, err := ctrl.Repository.CompaniesRepository.GetByID(id)
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_COMPANY_NOT_FOUND, "Company not found: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, company)
}

// CreateCompany creates a new company
// @Summary Create company
// @Description Create a new company
// @Tags Companies
// @Accept json
// @Produce json
// @Param company body model.Company true "Company Data"
// @Success 201 {object} model.Company
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /companies [post]
func (ctrl *Controller) CreateCompany(c *gin.Context) {
	var company model.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_INPUT, err.Error())
		return
	}

	createdCompany, err := ctrl.Repository.CompaniesRepository.Create(&company)
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_COMPANY_CREATE_FAILED, err.Error())
		return
	}

	c.JSON(http.StatusCreated, createdCompany)
}

// UpdateCompany updates an existing company
// @Summary Update company
// @Description Update an existing company by ID
// @Tags Companies
// @Accept json
// @Produce json
// @Param companyId path string true "Company ID"
// @Param company body model.Company true "Company Data"
// @Success 200 {object} model.Company
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 404 {object} map[string]string "Not Found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /companies/{companyId} [put]
func (ctrl *Controller) UpdateCompany(c *gin.Context) {
	id, err := uuid.Parse(c.Param("companyId"))
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_INPUT, "Invalid company ID format: "+err.Error())
		return
	}

	var company model.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_INPUT, err.Error())
		return
	}

	company.ID = id
	updateCompany, err := ctrl.Repository.CompaniesRepository.Update(&company)
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_COMPANY_UPDATE_FAILED, err.Error())
		return
	}

	c.JSON(http.StatusOK, updateCompany)
}

// DeleteCompany deletes a company by ID
// @Summary Delete company
// @Description Delete a company by ID
// @Tags Companies
// @Accept json
// @Produce json
// @Param companyId path string true "Company ID"
// @Success 204 "No Content"
// @Failure 404 {object} map[string]string "Not Found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /companies/{companyId} [delete]
func (ctrl *Controller) DeleteCompany(c *gin.Context) {
	id, err := uuid.Parse(c.Param("companyId"))
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_INPUT, "Invalid company ID format: "+err.Error())
		return
	}

	if err := ctrl.Repository.CompaniesRepository.Delete(id); err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_COMPANY_DELETE_FAILED, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
