package service

import (
	"Structure-Project/pkg/database"
	"Structure-Project/internal/core/domain"
)

// GetCustomers: ดึงลูกค้าทั้งหมด
func GetCustomers() ([]domain.CustomerModel, error) {
	rows, err := database.DB.Query("CALL GetCustomers()")
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

// GetCustomerByID: ดึงลูกค้าตาม ID
func GetCustomerByID(id int) (*domain.CustomerModel, error) {
	var cust domain.CustomerModel
	query := "CALL GetCustomerByID(?)"
	err := database.DB.QueryRow(query, id).Scan(&cust.CustomerID, &cust.FirstName, &cust.LastName, &cust.Email)
	if err != nil {
		return nil, err
	}
	return &cust, nil
}

// CreateCustomer: สร้างลูกค้า
func CreateCustomer(cust *domain.CustomerModel) error {
	_, err := database.DB.Exec("CALL InsertCustomer(?, ?, ?)", cust.FirstName, cust.LastName, cust.Email)
	return err
}

// UpdateCustomer: อัปเดตลูกค้า (return bool เพื่อบอกว่าเจอ ID ไหม)
func UpdateCustomer(cust *domain.CustomerModel) (bool, error) {
	result, err := database.DB.Exec("CALL UpdateCustomer(?, ?, ?, ?)", cust.CustomerID, cust.FirstName, cust.LastName, cust.Email)
	if err != nil {
		return false, err
	}
	rowsAffected, err := result.RowsAffected()
	return rowsAffected > 0, nil
}

// DeleteCustomer: ลบลูกค้า (return bool เพื่อบอกว่าเจอ ID ไหม)
func DeleteCustomer(id int) (bool, error) {
	result, err := database.DB.Exec("CALL DeleteCustomer(?)", id)
	if err != nil {
		return false, err
	}
	rowsAffected, _ := result.RowsAffected()
	return rowsAffected > 0, nil
}

// GetVehiclesByCustomerID: ดึงรถของลูกค้าคนนั้น
func GetVehiclesByCustomerID(customerID int) ([]domain.VehicleModel, error) {
	rows, err := database.DB.Query("CALL GetVehiclesByCustomerID(?)", customerID)
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
