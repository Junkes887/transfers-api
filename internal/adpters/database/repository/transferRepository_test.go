package repository

import (
	"errors"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Junkes887/transfers-api/internal/adpters/database"
	"github.com/Junkes887/transfers-api/internal/domain/model"
)

func TestCreateTransfer(t *testing.T) {
	type args struct {
		model *model.TransferModel
	}

	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       model.TransferModel
		wantErr    bool
	}{
		{
			name: "success create transfer",
			args: args{
				model: &model.TransferModel{ID: "ID", AccountOriginID: "AccountOriginID", AccountDestinationID: "AccountDestinationID", Amount: float64(10), CreatedAt: "2023-08-15 19:26:21"},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectExec(regexp.QuoteMeta("INSERT INTO TRANSFERS (ID, ACCOUNT_ORIGIN_ID, ACCOUNT_DESTINATION_ID, AMOUNT, CREATED_AT) VALUES(?,?,?,?,?)")).
					WithArgs("ID", "AccountOriginID", "AccountDestinationID", float64(10), "2023-08-15 19:26:21").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "error create transfer",
			args: args{
				model: &model.TransferModel{ID: "ID", AccountOriginID: "AccountOriginID", AccountDestinationID: "AccountDestinationID", Amount: float64(10), CreatedAt: "2023-08-15 19:26:21"},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectExec(regexp.QuoteMeta("INSERT INTO TRANSFERS (ID, ACCOUNT_ORIGIN_ID, ACCOUNT_DESTINATION_ID, AMOUNT, CREATED_AT) VALUES(?,?,?,?,?)")).
					WithArgs("ID", "AccountOriginID", "AccountDestinationID", float64(10), "2023-08-15 19:26:21").
					WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			u := &Repository{
				CFG: &database.ConfigMySql{DB: mockDB},
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			err := u.CreateTransfer(tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.CreateTransfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetTransfer(t *testing.T) {
	type args struct {
		accountOriginID string
	}

	var models []*model.TransferModel
	models = append(models, &model.TransferModel{ID: "123", AccountOriginID: "123", AccountDestinationID: "123", Amount: float64(10), CreatedAt: "2023-08-15 19:26:21"})

	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       []*model.TransferModel
		wantErr    bool
	}{
		{
			name: "success get transfer",
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta("SELECT ID, ACCOUNT_ORIGIN_ID, ACCOUNT_DESTINATION_ID, AMOUNT, CREATED_AT FROM TRANSFERS WHERE ACCOUNT_ORIGIN_ID = ?")).
					WithArgs("123").
					WillReturnRows(
						sqlmock.NewRows([]string{"ID", "ACCOUNT_ORIGIN_ID", "ACCOUNT_DESTINATION_ID", "AMOUNT", "CREATED_AT"}).
							AddRow("123", "123", "123", float64(10), "2023-08-15 19:26:21"))
			},
			args: args{
				accountOriginID: "123",
			},
			want: models,
		},
		{
			name: "error get transfer",
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta("SELECT ID, ACCOUNT_ORIGIN_ID, ACCOUNT_DESTINATION_ID, AMOUNT, CREATED_AT FROM TRANSFERS WHERE ACCOUNT_ORIGIN_ID = ?")).
					WithArgs("123").
					WillReturnError(errors.New("error"))
			},
			args: args{
				accountOriginID: "123",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			u := &Repository{
				CFG: &database.ConfigMySql{DB: mockDB},
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			got, err := u.GetTransfer(tt.args.accountOriginID)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetTransfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(len(got), len(tt.want)) {
				t.Errorf("repository.GetAllAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
