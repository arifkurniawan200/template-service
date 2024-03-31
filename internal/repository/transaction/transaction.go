package transaction

import (
	"database/sql"
	"template/adapter/cache"
	"template/internal/repository"
)

type TransactionHandler struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB, cache cache.Cache) repository.TransactionRepository {
	return &TransactionHandler{db}
}
