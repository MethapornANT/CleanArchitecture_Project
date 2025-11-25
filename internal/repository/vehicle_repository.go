package repository

import (
	"database/sql"
	"Structure-Project/internal/core/domain"
)

// * ลองทำแบบมี Repository

type VehicleRepository struct {
	db *sql.DB
}

func NewVehicleRepository(db *sql.DB) *VehicleRepository {
	return &VehicleRepository{db: db}
}

// 1. GetVehicles
func (r *VehicleRepository) GetVehicles() ([]domain.VehicleModel, error) {
	rows, err := r.db.Query("CALL Get_Vehicle()")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vehicles []domain.VehicleModel
	for rows.Next() {
		var v domain.VehicleModel
		if err := rows.Scan(&v.VehicleID, &v.CustomerID, &v.LicensePlate, &v.Model); err != nil {
			return nil, err
		}
		vehicles = append(vehicles, v)
	}
	return vehicles, nil
}

// 2. GetVehicleByID
func (r *VehicleRepository) GetVehicleByID(id int) (*domain.VehicleModel, error) {
	var v domain.VehicleModel
	query := "CALL GetVehicleByID(?)"
	err := r.db.QueryRow(query, id).Scan(&v.VehicleID, &v.CustomerID, &v.LicensePlate, &v.Model)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// 3. GetCustomersByVehicleID
func (r *VehicleRepository) GetCustomersByVehicleID(id int) ([]domain.CustomerModel, error) {
	rows, err := r.db.Query("CALL GetCustomersByVehicleID(?)", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []domain.CustomerModel
	for rows.Next() {
		var cust domain.CustomerModel
		if err := rows.Scan(&cust.CustomerID, &cust.FirstName, &cust.LastName, &cust.Email); err != nil {
			return nil, err
		}
		customers = append(customers, cust)
	}
	return customers, nil
}

// 4. GetVehicleAndCustomerByVehicleID
func (r *VehicleRepository) GetVehicleAndCustomerByVehicleID(id int) (*domain.VehicleWithCustomerModel, error) {
	var data domain.VehicleWithCustomerModel
	query := "CALL GetVehicleAndCustomerByVehicleID(?)"
	err := r.db.QueryRow(query, id).Scan(
		&data.VehicleID, &data.LicensePlate, &data.Model,
		&data.CustomerID, &data.FirstName, &data.LastName,
	)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// 5. CreateVehicle
func (r *VehicleRepository) CreateVehicle(v *domain.VehicleModel) error {
	_, err := r.db.Exec("CALL InsertVehicle(?, ?, ?)", v.CustomerID, v.LicensePlate, v.Model)
	return err
}