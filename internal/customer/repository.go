package customer

import (
	"database/sql"
	"fmt"
)

type Customer struct {
	CustomerID   int     `json:"customer_id"`
	NIK          string  `json:"nik"`
	FullName     string  `json:"full_name"`
	LegalName    string  `json:"legal_name"`
	PlaceOfBirth string  `json:"place_of_birth"`
	DateOfBirth  string  `json:"date_of_birth"`
	Salary       float64 `json:"salary"`
}

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetAllCustomers() ([]Customer, error) {
	query := "SELECT customer_id, nik, full_name, legal_name, place_of_birth, date_of_birth, salary FROM customers"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query customers: %w", err)
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var c Customer
		if err := rows.Scan(&c.CustomerID, &c.NIK, &c.FullName, &c.LegalName, &c.PlaceOfBirth, &c.DateOfBirth, &c.Salary); err != nil {
			return nil, fmt.Errorf("failed to scan customer: %w", err)
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (r *Repository) CreateCustomer(c *Customer) error {
	query := `
		INSERT INTO customers (nik, full_name, legal_name, place_of_birth, date_of_birth, salary)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING customer_id`
	err := r.DB.QueryRow(query, c.NIK, c.FullName, c.LegalName, c.PlaceOfBirth, c.DateOfBirth, c.Salary).Scan(&c.CustomerID)
	if err != nil {
		return fmt.Errorf("failed to create customer: %w", err)
	}
	return nil
}

func (r *Repository) GetPaginatedCustomers(limit, offset int) ([]Customer, error) {
	query := "SELECT customer_id, nik, full_name, legal_name, place_of_birth, date_of_birth, salary FROM customers LIMIT $1 OFFSET $2"
	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query paginated customers: %w", err)
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var c Customer
		if err := rows.Scan(&c.CustomerID, &c.NIK, &c.FullName, &c.LegalName, &c.PlaceOfBirth, &c.DateOfBirth, &c.Salary); err != nil {
			return nil, fmt.Errorf("failed to scan customer: %w", err)
		}
		customers = append(customers, c)
	}
	return customers, nil
}
