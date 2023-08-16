package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Junkes887/transfers-api/internal/adpters/database"
	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {
	mockDB, _, _ := sqlmock.New()
	rep := NewRepository(&database.ConfigMySql{DB: mockDB})

	assert.NotEmpty(t, rep)
}
