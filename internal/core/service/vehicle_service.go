package service

import (
	"Structure-Project/internal/core/domain"
	"Structure-Project/internal/repository"
	"Structure-Project/pkg/utils"
	"errors"
)

type VehicleService struct {
	repo *repository.VehicleRepository
}

func NewVehicleService(repo *repository.VehicleRepository) *VehicleService {
	return &VehicleService{repo: repo}
}

// 1. GetVehicles
func (s *VehicleService) GetVehicles() ([]domain.VehicleModel, error) {
	return s.repo.GetVehicles()
}

// 2. GetVehicleByID
func (s *VehicleService) GetVehicleByID(id int) (*domain.VehicleModel, error) {
	return s.repo.GetVehicleByID(id)
}

// 3. GetCustomersByVehicleID
func (s *VehicleService) GetCustomersByVehicleID(id int) ([]domain.CustomerModel, error) {
	customers, err := s.repo.GetCustomersByVehicleID(id)
	if err != nil {
		return nil, err
	}
	// Logic เพิ่มเติม: ถ้าไม่เจอข้อมูล ให้ส่ง Error กลับไป (เพื่อให้ Handler รู้ว่า 404)
	if len(customers) == 0 {
		return nil, errors.New("data_not_found")
	}
	return customers, nil
}

// 4. GetVehicleAndCustomerByVehicleID
func (s *VehicleService) GetVehicleAndCustomerByVehicleID(id int) (*domain.VehicleWithCustomerModel, error) {
	return s.repo.GetVehicleAndCustomerByVehicleID(id)
}

// 5. CreateVehicle
func (s *VehicleService) CreateVehicle(v *domain.VehicleModel) error {

	if !utils.IsPlateValid(v.LicensePlate) {
		return errors.New("invalid_license_plate")
	}
	return s.repo.CreateVehicle(v)
}
