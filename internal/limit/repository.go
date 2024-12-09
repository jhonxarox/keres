package limit

import (
	"database/sql"
	"fmt"
)

type Limit struct {
	LimitID     int     `json:"limit_id"`
	CustomerID  int     `json:"customer_id"`
	Tenor       int     `json:"tenor"`
	LimitAmount float64 `json:"limit_amount"`
}

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetLimitsByCustomerID(customerID int) ([]Limit, error) {
	query := "SELECT limit_id, customer_id, tenor, limit_amount FROM limits WHERE customer_id = $1"
	rows, err := r.DB.Query(query, customerID)
	if err != nil {
		return nil, fmt.Errorf("failed to query limits: %w", err)
	}
	defer rows.Close()

	var limits []Limit
	for rows.Next() {
		var l Limit
		if err := rows.Scan(&l.LimitID, &l.CustomerID, &l.Tenor, &l.LimitAmount); err != nil {
			return nil, fmt.Errorf("failed to scan limit: %w", err)
		}
		limits = append(limits, l)
	}
	return limits, nil
}

func (r *Repository) CreateLimit(limit *Limit) error {
	query := `
	INSERT INTO limits (customer_id, tenor, limit_amount)
	VALUES ($1, $2, $3)
	RETURNING limit_id`
	err := r.DB.QueryRow(query, limit.CustomerID, limit.Tenor, limit.LimitAmount).Scan(&limit.LimitID)
	if err != nil {
		return fmt.Errorf("failed to create limit: %w", err)
	}
	return nil
}
