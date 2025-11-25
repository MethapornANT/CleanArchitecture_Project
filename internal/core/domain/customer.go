package domain

type CustomerModel struct {
	CustomerID     int    `json:"customer_id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}
