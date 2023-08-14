package repository

import "github.com/Junkes887/transfers-api/internal/adpters/database"

type Repository struct {
	CFG *database.ConfigMySql
}

func NewRepository(cfg *database.ConfigMySql) *Repository {
	return &Repository{
		CFG: cfg,
	}
}
