package repository

import (
	"errors"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Junkes887/transfers-api/internal/adpters/database"
	"github.com/Junkes887/transfers-api/internal/domain/model"
	"github.com/Junkes887/transfers-api/pkg/crypt"
)

func TestCreateAccount(t *testing.T) {
	type args struct {
		model *model.AccountModel
	}

	t.Setenv("CRYPT_KEY", "0123456789abcdef")

	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       model.AccountModel
		wantErr    bool
	}{
		{
			name: "success create account",
			args: args{
				model: &model.AccountModel{ID: "ID", Name: "Name", CPF: "CPF", Secret: "Secret", Balance: float64(10), CreatedAt: "2023-08-15 19:26:21"},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectExec(regexp.QuoteMeta("INSERT INTO ACCOUNTS (ID, NAME, CPF, SECRET, BALANCE, CREATED_AT) VALUES(?,?,?,?,?,?)")).
					WithArgs("ID", "Name", "CPF", crypt.Encrypt("Secret"), float64(10), "2023-08-15 19:26:21").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "error create account",
			args: args{
				model: &model.AccountModel{ID: "ID", Name: "Name", CPF: "CPF", Secret: "Secret", Balance: float64(10), CreatedAt: "2023-08-15 19:26:21"},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectExec(regexp.QuoteMeta("INSERT INTO ACCOUNTS (ID, NAME, CPF, SECRET, BALANCE, CREATED_AT) VALUES(?,?,?,?,?,?)")).
					WithArgs("ID", "Name", "CPF", crypt.Encrypt("Secret"), float64(10), "2023-08-15 19:26:21").
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

			err := u.CreateAccount(tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetAccount(t *testing.T) {
	type args struct {
		id string
	}

	t.Setenv("CRYPT_KEY", "0123456789abcdef")
	secret := crypt.Encrypt("secret")

	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       *model.AccountModel
		wantErr    bool
	}{
		{
			name: "success get account",
			args: args{
				id: "123",
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta("SELECT ID, NAME, CPF, SECRET, BALANCE, CREATED_AT FROM ACCOUNTS WHERE ID = ?")).
					WithArgs("123").
					WillReturnRows(
						sqlmock.NewRows([]string{"ID", "NAME", "CPF", "SECRET", "BALANCE", "CREATED_AT"}).
							AddRow("123", "Name", "CPF", secret, float64(10), "2023-08-15 19:26:21"))
			},
			want: &model.AccountModel{ID: "123", Name: "Name", CPF: "CPF", Secret: "secret", Balance: float64(10), CreatedAt: "2023-08-15 19:26:21"},
		},
		{
			name: "error get account",
			args: args{
				id: "123",
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta("SELECT ID, NAME, CPF, SECRET, BALANCE, CREATED_AT FROM ACCOUNTS WHERE ID = ?")).
					WithArgs("123").
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

			got, err := u.GetAccount(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAccountByCpf(t *testing.T) {
	type args struct {
		cpf string
	}

	t.Setenv("CRYPT_KEY", "0123456789abcdef")
	secret := crypt.Encrypt("secret")

	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       *model.AccountModel
		wantErr    bool
	}{
		{
			name: "success get account by cpf",
			args: args{
				cpf: "123",
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta("SELECT ID, NAME, CPF, SECRET, BALANCE, CREATED_AT FROM ACCOUNTS WHERE CPF = ?")).
					WithArgs("123").
					WillReturnRows(
						sqlmock.NewRows([]string{"ID", "NAME", "CPF", "SECRET", "BALANCE", "CREATED_AT"}).
							AddRow("123", "Name", "CPF", secret, float64(10), "2023-08-15 19:26:21"))
			},
			want: &model.AccountModel{ID: "123", Name: "Name", CPF: "CPF", Secret: "secret", Balance: float64(10), CreatedAt: "2023-08-15 19:26:21"},
		},
		{
			name: "error get account by cpf",
			args: args{
				cpf: "123",
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta("SELECT ID, NAME, CPF, SECRET, BALANCE, CREATED_AT FROM ACCOUNTS WHERE CPF = ?")).
					WithArgs("123").
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

			got, err := u.GetAccountByCpf(tt.args.cpf)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetAccountByCpf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetAccountByCpf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllAccount(t *testing.T) {

	t.Setenv("CRYPT_KEY", "0123456789abcdef")
	secret := crypt.Encrypt("secret")

	var models []*model.AccountModel
	models = append(models, &model.AccountModel{ID: "123", Name: "Name", CPF: "CPF", Secret: "secret", Balance: float64(10), CreatedAt: "2023-08-15 19:26:21"})

	tests := []struct {
		name       string
		beforeTest func(sqlmock.Sqlmock)
		want       []*model.AccountModel
		wantErr    bool
	}{
		{
			name: "success get account by cpf",
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta("SELECT ID, NAME, CPF, SECRET, BALANCE, CREATED_AT FROM ACCOUNTS")).
					WillReturnRows(
						sqlmock.NewRows([]string{"ID", "NAME", "CPF", "SECRET", "BALANCE", "CREATED_AT"}).
							AddRow("123", "Name", "CPF", secret, float64(10), "2023-08-15 19:26:21"))
			},
			want: models,
		},
		{
			name: "error get account by cpf",
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta("SELECT ID, NAME, CPF, SECRET, BALANCE, CREATED_AT FROM ACCOUNTS")).
					WillReturnError(errors.New("Error"))
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

			got, err := u.GetAllAccount()
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetAllAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetAllAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateAccount(t *testing.T) {
	type args struct {
		id      string
		balance float64
	}

	balance := float64(10)

	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       *model.AccountModel
		wantErr    bool
	}{
		{
			name: "success update account",
			args: args{
				id:      "123",
				balance: balance,
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectExec(regexp.QuoteMeta("UPDATE ACCOUNTS SET BALANCE = ? WHERE ID = ?")).
					WithArgs(balance, "123").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "error update account",
			args: args{
				id:      "123",
				balance: balance,
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectExec(regexp.QuoteMeta("UPDATE ACCOUNTS SET BALANCE = ? WHERE ID = ?")).
					WithArgs(balance, "123").
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

			err := u.UpdateAccount(tt.args.id, tt.args.balance)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.UpdateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
