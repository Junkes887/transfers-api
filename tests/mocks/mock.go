package mocks

import (
	"github.com/Junkes887/transfers-api/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

func NewMockRepository(t mockConstructorTestingTNewUserRepository) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

func (m *MockRepository) GetAccount(ar1 string) (*model.AccountModel, error) {
	ret := m.Called(ar1)

	var r0 *model.AccountModel
	if rf, ok := ret.Get(0).(func(string) *model.AccountModel); ok {
		r0 = rf(ar1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AccountModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ar1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *MockRepository) GetAccountByCpf(ar1 string) (*model.AccountModel, error) {
	ret := m.Called(ar1)

	var r0 *model.AccountModel
	if rf, ok := ret.Get(0).(func(string) *model.AccountModel); ok {
		r0 = rf(ar1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AccountModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ar1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *MockRepository) GetAllAccount() ([]*model.AccountModel, error) {
	ret := m.Called()

	var r0 []*model.AccountModel
	if rf, ok := ret.Get(0).(func() []*model.AccountModel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.AccountModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *MockRepository) CreateAccount(ar1 *model.AccountModel) error {
	ret := m.Called(ar1)

	var r0 error
	if rf, ok := ret.Get(1).(func(*model.AccountModel) error); ok {
		r0 = rf(ar1)
	} else {
		r0 = ret.Error(1)
	}

	return r0
}

func (m *MockRepository) UpdateAccount(ar1 string, ar2 float64) error {
	ret := m.Called(ar1, ar2)
	var r0 error

	if rf, ok := ret.Get(0).(func(string, float64) error); ok {
		r0 = rf(ar1, ar2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (m *MockRepository) CreateTransfer(ar1 *model.TransferModel) error {
	ret := m.Called(ar1)

	var r0 error
	if rf, ok := ret.Get(1).(func(*model.TransferModel) error); ok {
		r0 = rf(ar1)
	} else {
		r0 = ret.Error(1)
	}

	return r0
}

func (m *MockRepository) GetTransfer(ar1 string) ([]*model.TransferModel, error) {
	ret := m.Called(ar1)

	var r0 []*model.TransferModel
	if rf, ok := ret.Get(1).(func(string) []*model.TransferModel); ok {
		r0 = rf(ar1)
	} else {
		r0 = ret.Get(0).([]*model.TransferModel)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ar1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
