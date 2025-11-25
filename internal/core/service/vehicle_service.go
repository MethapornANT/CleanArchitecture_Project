package service

import (
	"Structure-Project/internal/core/domain"
	"Structure-Project/pkg/database"
)

// GetVehicles: ดึงข้อมูลรถทั้งหมด
func GetVehicles() ([]domain.VehicleModel, error) {
	rows, err := database.DB.Query("CALL Get_Vehicle()")
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

// GetVehicleByID: ดึงรถตาม ID
func GetVehicleByID(id int) (*domain.VehicleModel, error) {
	var v domain.VehicleModel
	query := "CALL GetVehicleByID(?)"
	err := database.DB.QueryRow(query, id).Scan(&v.VehicleID, &v.CustomerID, &v.LicensePlate, &v.Model)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// GetCustomersByVehicleID: ดึงลูกค้าตามรถ
func GetCustomersByVehicleID(id int) ([]domain.CustomerModel, error) {
	rows, err := database.DB.Query("CALL GetCustomersByVehicleID(?)", id)
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

// GetVehicleAndCustomerByVehicleID: ดึงข้อมูล Join
func GetVehicleAndCustomerByVehicleID(id int) (*domain.VehicleWithCustomerModel, error) {
	var data domain.VehicleWithCustomerModel
	query := "CALL GetVehicleAndCustomerByVehicleID(?)"
	err := database.DB.QueryRow(query, id).Scan(
		&data.VehicleID, &data.LicensePlate, &data.Model,
		&data.CustomerID, &data.FirstName, &data.LastName,
		// ตรวจสอบว่า SP Return กี่ column แน่นะครับ ใส่ให้ครบตาม struct
	)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// CreateVehicle: เพิ่มรถ
func CreateVehicle(v *domain.VehicleModel) error {
	_, err := database.DB.Exec("CALL InsertVehicle(?, ?, ?)", v.CustomerID, v.LicensePlate, v.Model)
	return err
}
