// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	models "github.com/grafana/grafana-openapi-client-go/models"
	mock "github.com/stretchr/testify/mock"
)

// OrganizationsApi is an autogenerated mock type for the OrganizationsApi type
type OrganizationsApi struct {
	mock.Mock
}

type OrganizationsApi_Expecter struct {
	mock *mock.Mock
}

func (_m *OrganizationsApi) EXPECT() *OrganizationsApi_Expecter {
	return &OrganizationsApi_Expecter{mock: &_m.Mock}
}

// AddUserToOrg provides a mock function with given fields: role, userId, orgId
func (_m *OrganizationsApi) AddUserToOrg(role string, userId int64, orgId int64) error {
	ret := _m.Called(role, userId, orgId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int64, int64) error); ok {
		r0 = rf(role, userId, orgId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrganizationsApi_AddUserToOrg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddUserToOrg'
type OrganizationsApi_AddUserToOrg_Call struct {
	*mock.Call
}

// AddUserToOrg is a helper method to define mock.On call
//   - role string
//   - userId int64
//   - orgId int64
func (_e *OrganizationsApi_Expecter) AddUserToOrg(role interface{}, userId interface{}, orgId interface{}) *OrganizationsApi_AddUserToOrg_Call {
	return &OrganizationsApi_AddUserToOrg_Call{Call: _e.mock.On("AddUserToOrg", role, userId, orgId)}
}

func (_c *OrganizationsApi_AddUserToOrg_Call) Run(run func(role string, userId int64, orgId int64)) *OrganizationsApi_AddUserToOrg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int64), args[2].(int64))
	})
	return _c
}

func (_c *OrganizationsApi_AddUserToOrg_Call) Return(_a0 error) *OrganizationsApi_AddUserToOrg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OrganizationsApi_AddUserToOrg_Call) RunAndReturn(run func(string, int64, int64) error) *OrganizationsApi_AddUserToOrg_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteUserFromOrg provides a mock function with given fields: userId, orgId
func (_m *OrganizationsApi) DeleteUserFromOrg(userId int64, orgId int64) error {
	ret := _m.Called(userId, orgId)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, int64) error); ok {
		r0 = rf(userId, orgId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrganizationsApi_DeleteUserFromOrg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteUserFromOrg'
type OrganizationsApi_DeleteUserFromOrg_Call struct {
	*mock.Call
}

// DeleteUserFromOrg is a helper method to define mock.On call
//   - userId int64
//   - orgId int64
func (_e *OrganizationsApi_Expecter) DeleteUserFromOrg(userId interface{}, orgId interface{}) *OrganizationsApi_DeleteUserFromOrg_Call {
	return &OrganizationsApi_DeleteUserFromOrg_Call{Call: _e.mock.On("DeleteUserFromOrg", userId, orgId)}
}

func (_c *OrganizationsApi_DeleteUserFromOrg_Call) Run(run func(userId int64, orgId int64)) *OrganizationsApi_DeleteUserFromOrg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(int64))
	})
	return _c
}

func (_c *OrganizationsApi_DeleteUserFromOrg_Call) Return(_a0 error) *OrganizationsApi_DeleteUserFromOrg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OrganizationsApi_DeleteUserFromOrg_Call) RunAndReturn(run func(int64, int64) error) *OrganizationsApi_DeleteUserFromOrg_Call {
	_c.Call.Return(run)
	return _c
}

// DownloadOrganizations provides a mock function with given fields:
func (_m *OrganizationsApi) DownloadOrganizations() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// OrganizationsApi_DownloadOrganizations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DownloadOrganizations'
type OrganizationsApi_DownloadOrganizations_Call struct {
	*mock.Call
}

// DownloadOrganizations is a helper method to define mock.On call
func (_e *OrganizationsApi_Expecter) DownloadOrganizations() *OrganizationsApi_DownloadOrganizations_Call {
	return &OrganizationsApi_DownloadOrganizations_Call{Call: _e.mock.On("DownloadOrganizations")}
}

func (_c *OrganizationsApi_DownloadOrganizations_Call) Run(run func()) *OrganizationsApi_DownloadOrganizations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *OrganizationsApi_DownloadOrganizations_Call) Return(_a0 []string) *OrganizationsApi_DownloadOrganizations_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OrganizationsApi_DownloadOrganizations_Call) RunAndReturn(run func() []string) *OrganizationsApi_DownloadOrganizations_Call {
	_c.Call.Return(run)
	return _c
}

// GetTokenOrganization provides a mock function with given fields:
func (_m *OrganizationsApi) GetTokenOrganization() *models.OrgDetailsDTO {
	ret := _m.Called()

	var r0 *models.OrgDetailsDTO
	if rf, ok := ret.Get(0).(func() *models.OrgDetailsDTO); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.OrgDetailsDTO)
		}
	}

	return r0
}

// OrganizationsApi_GetTokenOrganization_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTokenOrganization'
type OrganizationsApi_GetTokenOrganization_Call struct {
	*mock.Call
}

// GetTokenOrganization is a helper method to define mock.On call
func (_e *OrganizationsApi_Expecter) GetTokenOrganization() *OrganizationsApi_GetTokenOrganization_Call {
	return &OrganizationsApi_GetTokenOrganization_Call{Call: _e.mock.On("GetTokenOrganization")}
}

func (_c *OrganizationsApi_GetTokenOrganization_Call) Run(run func()) *OrganizationsApi_GetTokenOrganization_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *OrganizationsApi_GetTokenOrganization_Call) Return(_a0 *models.OrgDetailsDTO) *OrganizationsApi_GetTokenOrganization_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OrganizationsApi_GetTokenOrganization_Call) RunAndReturn(run func() *models.OrgDetailsDTO) *OrganizationsApi_GetTokenOrganization_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserOrganization provides a mock function with given fields:
func (_m *OrganizationsApi) GetUserOrganization() *models.OrgDetailsDTO {
	ret := _m.Called()

	var r0 *models.OrgDetailsDTO
	if rf, ok := ret.Get(0).(func() *models.OrgDetailsDTO); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.OrgDetailsDTO)
		}
	}

	return r0
}

// OrganizationsApi_GetUserOrganization_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserOrganization'
type OrganizationsApi_GetUserOrganization_Call struct {
	*mock.Call
}

// GetUserOrganization is a helper method to define mock.On call
func (_e *OrganizationsApi_Expecter) GetUserOrganization() *OrganizationsApi_GetUserOrganization_Call {
	return &OrganizationsApi_GetUserOrganization_Call{Call: _e.mock.On("GetUserOrganization")}
}

func (_c *OrganizationsApi_GetUserOrganization_Call) Run(run func()) *OrganizationsApi_GetUserOrganization_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *OrganizationsApi_GetUserOrganization_Call) Return(_a0 *models.OrgDetailsDTO) *OrganizationsApi_GetUserOrganization_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OrganizationsApi_GetUserOrganization_Call) RunAndReturn(run func() *models.OrgDetailsDTO) *OrganizationsApi_GetUserOrganization_Call {
	_c.Call.Return(run)
	return _c
}

// InitOrganizations provides a mock function with given fields:
func (_m *OrganizationsApi) InitOrganizations() {
	_m.Called()
}

// OrganizationsApi_InitOrganizations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InitOrganizations'
type OrganizationsApi_InitOrganizations_Call struct {
	*mock.Call
}

// InitOrganizations is a helper method to define mock.On call
func (_e *OrganizationsApi_Expecter) InitOrganizations() *OrganizationsApi_InitOrganizations_Call {
	return &OrganizationsApi_InitOrganizations_Call{Call: _e.mock.On("InitOrganizations")}
}

func (_c *OrganizationsApi_InitOrganizations_Call) Run(run func()) *OrganizationsApi_InitOrganizations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *OrganizationsApi_InitOrganizations_Call) Return() *OrganizationsApi_InitOrganizations_Call {
	_c.Call.Return()
	return _c
}

func (_c *OrganizationsApi_InitOrganizations_Call) RunAndReturn(run func()) *OrganizationsApi_InitOrganizations_Call {
	_c.Call.Return(run)
	return _c
}

// ListOrgUsers provides a mock function with given fields: orgId
func (_m *OrganizationsApi) ListOrgUsers(orgId int64) []*models.OrgUserDTO {
	ret := _m.Called(orgId)

	var r0 []*models.OrgUserDTO
	if rf, ok := ret.Get(0).(func(int64) []*models.OrgUserDTO); ok {
		r0 = rf(orgId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.OrgUserDTO)
		}
	}

	return r0
}

// OrganizationsApi_ListOrgUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListOrgUsers'
type OrganizationsApi_ListOrgUsers_Call struct {
	*mock.Call
}

// ListOrgUsers is a helper method to define mock.On call
//   - orgId int64
func (_e *OrganizationsApi_Expecter) ListOrgUsers(orgId interface{}) *OrganizationsApi_ListOrgUsers_Call {
	return &OrganizationsApi_ListOrgUsers_Call{Call: _e.mock.On("ListOrgUsers", orgId)}
}

func (_c *OrganizationsApi_ListOrgUsers_Call) Run(run func(orgId int64)) *OrganizationsApi_ListOrgUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *OrganizationsApi_ListOrgUsers_Call) Return(_a0 []*models.OrgUserDTO) *OrganizationsApi_ListOrgUsers_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OrganizationsApi_ListOrgUsers_Call) RunAndReturn(run func(int64) []*models.OrgUserDTO) *OrganizationsApi_ListOrgUsers_Call {
	_c.Call.Return(run)
	return _c
}

// ListOrganizations provides a mock function with given fields:
func (_m *OrganizationsApi) ListOrganizations() []*models.OrgDTO {
	ret := _m.Called()

	var r0 []*models.OrgDTO
	if rf, ok := ret.Get(0).(func() []*models.OrgDTO); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.OrgDTO)
		}
	}

	return r0
}

// OrganizationsApi_ListOrganizations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListOrganizations'
type OrganizationsApi_ListOrganizations_Call struct {
	*mock.Call
}

// ListOrganizations is a helper method to define mock.On call
func (_e *OrganizationsApi_Expecter) ListOrganizations() *OrganizationsApi_ListOrganizations_Call {
	return &OrganizationsApi_ListOrganizations_Call{Call: _e.mock.On("ListOrganizations")}
}

func (_c *OrganizationsApi_ListOrganizations_Call) Run(run func()) *OrganizationsApi_ListOrganizations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *OrganizationsApi_ListOrganizations_Call) Return(_a0 []*models.OrgDTO) *OrganizationsApi_ListOrganizations_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OrganizationsApi_ListOrganizations_Call) RunAndReturn(run func() []*models.OrgDTO) *OrganizationsApi_ListOrganizations_Call {
	_c.Call.Return(run)
	return _c
}

// ListUserOrganizations provides a mock function with given fields:
func (_m *OrganizationsApi) ListUserOrganizations() ([]*models.UserOrgDTO, error) {
	ret := _m.Called()

	var r0 []*models.UserOrgDTO
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*models.UserOrgDTO, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*models.UserOrgDTO); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.UserOrgDTO)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrganizationsApi_ListUserOrganizations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListUserOrganizations'
type OrganizationsApi_ListUserOrganizations_Call struct {
	*mock.Call
}

// ListUserOrganizations is a helper method to define mock.On call
func (_e *OrganizationsApi_Expecter) ListUserOrganizations() *OrganizationsApi_ListUserOrganizations_Call {
	return &OrganizationsApi_ListUserOrganizations_Call{Call: _e.mock.On("ListUserOrganizations")}
}

func (_c *OrganizationsApi_ListUserOrganizations_Call) Run(run func()) *OrganizationsApi_ListUserOrganizations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *OrganizationsApi_ListUserOrganizations_Call) Return(_a0 []*models.UserOrgDTO, _a1 error) *OrganizationsApi_ListUserOrganizations_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrganizationsApi_ListUserOrganizations_Call) RunAndReturn(run func() ([]*models.UserOrgDTO, error)) *OrganizationsApi_ListUserOrganizations_Call {
	_c.Call.Return(run)
	return _c
}

// SetOrganizationByName provides a mock function with given fields: name, useSlug
func (_m *OrganizationsApi) SetOrganizationByName(name string, useSlug bool) error {
	ret := _m.Called(name, useSlug)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, bool) error); ok {
		r0 = rf(name, useSlug)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrganizationsApi_SetOrganizationByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetOrganizationByName'
type OrganizationsApi_SetOrganizationByName_Call struct {
	*mock.Call
}

// SetOrganizationByName is a helper method to define mock.On call
//   - name string
//   - useSlug bool
func (_e *OrganizationsApi_Expecter) SetOrganizationByName(name interface{}, useSlug interface{}) *OrganizationsApi_SetOrganizationByName_Call {
	return &OrganizationsApi_SetOrganizationByName_Call{Call: _e.mock.On("SetOrganizationByName", name, useSlug)}
}

func (_c *OrganizationsApi_SetOrganizationByName_Call) Run(run func(name string, useSlug bool)) *OrganizationsApi_SetOrganizationByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(bool))
	})
	return _c
}

func (_c *OrganizationsApi_SetOrganizationByName_Call) Return(_a0 error) *OrganizationsApi_SetOrganizationByName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OrganizationsApi_SetOrganizationByName_Call) RunAndReturn(run func(string, bool) error) *OrganizationsApi_SetOrganizationByName_Call {
	_c.Call.Return(run)
	return _c
}

// SetUserOrganizations provides a mock function with given fields: id
func (_m *OrganizationsApi) SetUserOrganizations(id int64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrganizationsApi_SetUserOrganizations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetUserOrganizations'
type OrganizationsApi_SetUserOrganizations_Call struct {
	*mock.Call
}

// SetUserOrganizations is a helper method to define mock.On call
//   - id int64
func (_e *OrganizationsApi_Expecter) SetUserOrganizations(id interface{}) *OrganizationsApi_SetUserOrganizations_Call {
	return &OrganizationsApi_SetUserOrganizations_Call{Call: _e.mock.On("SetUserOrganizations", id)}
}

func (_c *OrganizationsApi_SetUserOrganizations_Call) Run(run func(id int64)) *OrganizationsApi_SetUserOrganizations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *OrganizationsApi_SetUserOrganizations_Call) Return(_a0 error) *OrganizationsApi_SetUserOrganizations_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OrganizationsApi_SetUserOrganizations_Call) RunAndReturn(run func(int64) error) *OrganizationsApi_SetUserOrganizations_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUserInOrg provides a mock function with given fields: role, userId, orgId
func (_m *OrganizationsApi) UpdateUserInOrg(role string, userId int64, orgId int64) error {
	ret := _m.Called(role, userId, orgId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int64, int64) error); ok {
		r0 = rf(role, userId, orgId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrganizationsApi_UpdateUserInOrg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUserInOrg'
type OrganizationsApi_UpdateUserInOrg_Call struct {
	*mock.Call
}

// UpdateUserInOrg is a helper method to define mock.On call
//   - role string
//   - userId int64
//   - orgId int64
func (_e *OrganizationsApi_Expecter) UpdateUserInOrg(role interface{}, userId interface{}, orgId interface{}) *OrganizationsApi_UpdateUserInOrg_Call {
	return &OrganizationsApi_UpdateUserInOrg_Call{Call: _e.mock.On("UpdateUserInOrg", role, userId, orgId)}
}

func (_c *OrganizationsApi_UpdateUserInOrg_Call) Run(run func(role string, userId int64, orgId int64)) *OrganizationsApi_UpdateUserInOrg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int64), args[2].(int64))
	})
	return _c
}

func (_c *OrganizationsApi_UpdateUserInOrg_Call) Return(_a0 error) *OrganizationsApi_UpdateUserInOrg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OrganizationsApi_UpdateUserInOrg_Call) RunAndReturn(run func(string, int64, int64) error) *OrganizationsApi_UpdateUserInOrg_Call {
	_c.Call.Return(run)
	return _c
}

// UploadOrganizations provides a mock function with given fields:
func (_m *OrganizationsApi) UploadOrganizations() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// OrganizationsApi_UploadOrganizations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UploadOrganizations'
type OrganizationsApi_UploadOrganizations_Call struct {
	*mock.Call
}

// UploadOrganizations is a helper method to define mock.On call
func (_e *OrganizationsApi_Expecter) UploadOrganizations() *OrganizationsApi_UploadOrganizations_Call {
	return &OrganizationsApi_UploadOrganizations_Call{Call: _e.mock.On("UploadOrganizations")}
}

func (_c *OrganizationsApi_UploadOrganizations_Call) Run(run func()) *OrganizationsApi_UploadOrganizations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *OrganizationsApi_UploadOrganizations_Call) Return(_a0 []string) *OrganizationsApi_UploadOrganizations_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OrganizationsApi_UploadOrganizations_Call) RunAndReturn(run func() []string) *OrganizationsApi_UploadOrganizations_Call {
	_c.Call.Return(run)
	return _c
}

// NewOrganizationsApi creates a new instance of OrganizationsApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrganizationsApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrganizationsApi {
	mock := &OrganizationsApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
