package transaction

import (
	"database/sql"
	"fmt"
)

type Transaction struct {
	TransactionID  int     `json:"transaction_id"`
	CustomerID     int     `json:"customer_id"`
	ContractNumber string  `json:"contract_number"`
	OTR            float64 `json:"otr"`
	AdminFee       float64 `json:"admin_fee"`
	Installment    float64 `json:"installment"`
	InterestAmount float64 `json:"interest_amount"`
	AssetName      string  `json:"asset_name"`
	CreatedAt      string  `json:"created_at"`
}

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) CreateTransaction(t *Transaction) error {
	query := `
		INSERT INTO transactions (customer_id, contract_number, otr, admin_fee, installment_amount, interest_amount, asset_name)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING transaction_id, created_at`
	err := r.DB.QueryRow(query, t.CustomerID, t.ContractNumber, t.OTR, t.AdminFee, t.Installment, t.InterestAmount, t.AssetName).Scan(&t.TransactionID, &t.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to create transaction: %w", err)
	}
	return nil
}
