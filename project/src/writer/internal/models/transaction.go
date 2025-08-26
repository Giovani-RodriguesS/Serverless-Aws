package models

type Transaction struct {
    TransactionID string  `json:"TransactionID"`
    AccountID     string  `json:"AccountID"`
    Amount        float64 `json:"Amount"`
    Timestamp     string  `json:"Timestamp"`
    Type          string  `json:"Type"` // "DEBIT" ou "CREDIT"
}