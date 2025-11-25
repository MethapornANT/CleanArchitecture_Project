package domain

type VehicleModel struct {
	VehicleID    int    `json:"vehicle_id"`
	CustomerID   int    `json:"customer_id"`
	LicensePlate string `json:"license_plate"`
	Model        string `json:"model"`
}

type VehicleWithCustomerModel struct {
	VehicleID    int    `json:"vehicle_id"`
	LicensePlate string `json:"license_plate"`
	Model        string `json:"model"`
	CustomerID   int    `json:"customer_id"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
}
