package repository

import (
	"fmt"
	"warningfloodsystem/domain/model"

	"gorm.io/gorm"
)

type FloodReport interface {
	CreateReport(report *model.FloodReport) (*model.FloodReport, error)
	GetReportByID(reportID uint) (*model.FloodReport, error)
	GetAllReports() ([]model.FloodReport, error)
	UpdateReport(report *model.FloodReport) (*model.FloodReport, error)
	DeleteReport(id uint) error
}

type FloodReportRepository struct {
	db *gorm.DB
}

func NewFloodReportRepository(db *gorm.DB) *FloodReportRepository {
	return &FloodReportRepository{db: db}
}

func (r *FloodReportRepository) CreateReport(report *model.FloodReport) (*model.FloodReport, error) {
	// Create the report in the database
	if err := r.db.Create(report).Error; err != nil {
		return nil, fmt.Errorf("terjadi kesalahan saat membuat laporan banjir: %w", err)
	}
	// Return the created report along with nil error
	return report, nil
}

func (r *FloodReportRepository) GetReportByID(reportID uint) (*model.FloodReport, error) {
	var report model.FloodReport
	err := r.db.First(&report, reportID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Jika data tidak ditemukan
			return nil, fmt.Errorf("laporan banjir dengan ID %d tidak ditemukan", reportID)
		}
		// Pesan kesalahan umum
		return nil, fmt.Errorf("terjadi kesalahan saat mengambil laporan banjir dengan ID %d: %w", reportID, err)
	}
	return &report, nil
}

func (r *FloodReportRepository) GetAllReports() ([]model.FloodReport, error) {
	var reports []model.FloodReport
	err := r.db.Find(&reports).Error
	if err != nil {
		return nil, fmt.Errorf("terjadi kesalahan saat mengambil semua laporan banjir: %w", err)
	}
	return reports, nil
}

func (r *FloodReportRepository) UpdateReport(report *model.FloodReport) (*model.FloodReport, error) {
	if err := r.db.Model(report).Omit("created_at").Updates(report).Error; err != nil {
		return nil, fmt.Errorf("terjadi kesalahan saat memperbarui laporan banjir: %w", err)
	}

	return report, nil
}

func (r *FloodReportRepository) DeleteReport(id uint) error {
	if err := r.db.Delete(&model.FloodReport{}, id).Error; err != nil {
		return fmt.Errorf("terjadi kesalahan saat menghapus laporan banjir dengan ID %d: %w", id, err)
	}
	return nil
}
