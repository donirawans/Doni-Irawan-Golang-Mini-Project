package controllers

import (
	"net/http"
	"strconv"
	"warningfloodsystem/domain/model"
	"warningfloodsystem/domain/usecase"
	"warningfloodsystem/utils"

	"github.com/labstack/echo/v4"
)

type FloodReportController struct {
	service usecase.FloodReportUsecase
}

func NewFloodReportController(s usecase.FloodReportUsecase) *FloodReportController {
	return &FloodReportController{
		service: s,
	}
}

func (ctrl *FloodReportController) Create(c echo.Context) error {
	var report model.FloodReport
	if err := c.Bind(&report); err != nil {
		return utils.JSONErrorResponse(c, http.StatusBadRequest, "Invalid input")
	}

	createdReport, err := ctrl.service.CreateReport(&report)
	if err != nil {
		return utils.JSONErrorResponse(c, http.StatusInternalServerError, "Failed to create report")
	}

	return utils.JSONSuccessResponse(c, createdReport)
}

func (ctrl *FloodReportController) GetAll(c echo.Context) error {
	reports, err := ctrl.service.GetAllReports()
	if err != nil {
		return utils.JSONErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve reports")
	}

	return utils.JSONSuccessResponse(c, reports)
}

func (ctrl *FloodReportController) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.JSONErrorResponse(c, http.StatusBadRequest, "Invalid ID format")
	}

	report, err := ctrl.service.GetReportByID(uint(id))
	if err != nil {
		return utils.JSONErrorResponse(c, http.StatusNotFound, "Report not found")
	}

	return utils.JSONSuccessResponse(c, report)
}

func (ctrl *FloodReportController) Update(c echo.Context) error {
	var report model.FloodReport
	if err := c.Bind(&report); err != nil {
		return utils.JSONErrorResponse(c, http.StatusBadRequest, "Invalid input")
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.JSONErrorResponse(c, http.StatusBadRequest, "Invalid ID format")
	}
	report.ID = uint(id)

	updatedReport, err := ctrl.service.UpdateReport(&report)
	if err != nil {
		return utils.JSONErrorResponse(c, http.StatusInternalServerError, "Failed to update report")
	}

	return utils.JSONSuccessResponse(c, updatedReport)
}

func (ctrl *FloodReportController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.JSONErrorResponse(c, http.StatusBadRequest, "Invalid ID format")
	}

	err = ctrl.service.DeleteReport(uint(id))
	if err != nil {
		return utils.JSONErrorResponse(c, http.StatusInternalServerError, "Failed to delete report")
	}

	return utils.JSONSuccessResponse(c, "Report deleted successfully")
}
