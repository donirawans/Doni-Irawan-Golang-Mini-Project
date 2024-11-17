package usecase

import (
	"errors"
	"warningfloodsystem/domain/model"
	"warningfloodsystem/domain/repository"
)

type FloodReportUsecase interface {
	CreateReport(report *model.FloodReport) (*model.FloodReport, error)
	GetReportByID(reportID uint) (*model.FloodReport, error)
	GetAllReports() ([]model.FloodReport, error)
	UpdateReport(report *model.FloodReport) (*model.FloodReport, error)
	DeleteReport(reportID uint) error
}

type FloodReportUsecaseImpl struct {
	repo repository.FloodReportRepository
}

func NewFloodReportUsecase(repo repository.FloodReportRepository) FloodReportUsecase {
	return &FloodReportUsecaseImpl{repo}
}

func (s *FloodReportUsecaseImpl) CreateReport(report *model.FloodReport) (*model.FloodReport, error) {
	if report == nil {
		return nil, errors.New("data laporan banjir tidak valid")
	}
	// Assuming the CreateReport method in the repository returns the created report
	createdReport, err := s.repo.CreateReport(report)
	if err != nil {
		return nil, err
	}
	return createdReport, nil
}

func (s *FloodReportUsecaseImpl) GetReportByID(reportID uint) (*model.FloodReport, error) {
	report, err := s.repo.GetReportByID(reportID)
	if err != nil {
		return nil, err
	}
	return report, nil
}

func (s *FloodReportUsecaseImpl) GetAllReports() ([]model.FloodReport, error) {
	reports, err := s.repo.GetAllReports()
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func (s *FloodReportUsecaseImpl) UpdateReport(report *model.FloodReport) (*model.FloodReport, error) {
	if report == nil {
		return nil, errors.New("data laporan banjir tidak valid")
	}
	updatedReport, err := s.repo.UpdateReport(report)
	if err != nil {
		return nil, err
	}
	return updatedReport, nil
}

func (s *FloodReportUsecaseImpl) DeleteReport(reportID uint) error {
	return s.repo.DeleteReport(reportID)
}
