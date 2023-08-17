package mocks

import (
	"github.com/Junkes887/transfers-api/internal/domain/model"
	"github.com/Junkes887/transfers-api/pkg/httperr"
	"github.com/stretchr/testify/mock"
)

type MockUseCase struct {
	mock.Mock
}

type mockConstructorTestingTNewUseCase interface {
	mock.TestingT
	Cleanup(func())
}

func NewMockUseCase(t mockConstructorTestingTNewUseCase) *MockUseCase {
	mock := &MockUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

func (m *MockUseCase) GetAccount(ar1 string) (*model.AccountModel, httperr.RequestError) {
	ret := m.Called(ar1)

	var r0 *model.AccountModel
	if rf, ok := ret.Get(0).(func(string) *model.AccountModel); ok {
		r0 = rf(ar1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AccountModel)
		}
	}

	var r1 httperr.RequestError
	if rf, ok := ret.Get(1).(func(string) httperr.RequestError); ok {
		r1 = rf(ar1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(httperr.RequestError)
		}
	}

	return r0, r1
}

func (m *MockUseCase) GetAllAccount() ([]*model.AccountModel, httperr.RequestError) {
	ret := m.Called()

	var r0 []*model.AccountModel
	if rf, ok := ret.Get(0).(func() []*model.AccountModel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.AccountModel)
		}
	}

	var r1 httperr.RequestError
	if rf, ok := ret.Get(1).(func() httperr.RequestError); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(httperr.RequestError)
		}
	}

	return r0, r1
}

func (m *MockUseCase) CreateAccount(ar1 *model.AccountModel) (*model.AccountModel, httperr.RequestError) {
	ret := m.Called(ar1)

	var r0 *model.AccountModel
	if rf, ok := ret.Get(0).(func(*model.AccountModel) *model.AccountModel); ok {
		r0 = rf(ar1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AccountModel)
		}
	}

	var r1 httperr.RequestError
	if rf, ok := ret.Get(1).(func(*model.AccountModel) httperr.RequestError); ok {
		r1 = rf(ar1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(httperr.RequestError)
		}
	}

	return r0, r1
}

func (m *MockUseCase) Login(ar1 *model.LoginModel) (string, httperr.RequestError) {
	ret := m.Called(ar1)

	var r0 string
	if rf, ok := ret.Get(0).(func(*model.LoginModel) string); ok {
		r0 = rf(ar1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(string)
		}
	}

	var r1 httperr.RequestError
	if rf, ok := ret.Get(1).(func(*model.LoginModel) httperr.RequestError); ok {
		r1 = rf(ar1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(httperr.RequestError)
		}
	}

	return r0, r1
}

func (m *MockUseCase) CreateTransfer(ar1 string, ar2 *model.TransferModel) (*model.TransferModel, httperr.RequestError) {
	ret := m.Called(ar1, ar2)

	var r0 *model.TransferModel
	if rf, ok := ret.Get(0).(func(string, *model.TransferModel) *model.TransferModel); ok {
		r0 = rf(ar1, ar2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.TransferModel)
		}
	}

	var r1 httperr.RequestError
	if rf, ok := ret.Get(1).(func(string, *model.TransferModel) httperr.RequestError); ok {
		r1 = rf(ar1, ar2)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(httperr.RequestError)
		}
	}

	return r0, r1

}
func (m *MockUseCase) GetTransfer(ar1 string) ([]*model.TransferModel, httperr.RequestError) {
	ret := m.Called(ar1)

	var r0 []*model.TransferModel
	if rf, ok := ret.Get(0).(func(string) []*model.TransferModel); ok {
		r0 = rf(ar1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.TransferModel)
		}
	}

	var r1 httperr.RequestError
	if rf, ok := ret.Get(1).(func(string) httperr.RequestError); ok {
		r1 = rf(ar1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(httperr.RequestError)
		}
	}

	return r0, r1
}
