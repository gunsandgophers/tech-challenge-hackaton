// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	multipart "mime/multipart"

	mock "github.com/stretchr/testify/mock"
)

// StorageServiceInterfaceMock is an autogenerated mock type for the StorageServiceInterface type
type StorageServiceInterfaceMock struct {
	mock.Mock
}

type StorageServiceInterfaceMock_Expecter struct {
	mock *mock.Mock
}

func (_m *StorageServiceInterfaceMock) EXPECT() *StorageServiceInterfaceMock_Expecter {
	return &StorageServiceInterfaceMock_Expecter{mock: &_m.Mock}
}

// DownloadVideo provides a mock function with given fields: videoID, filename
func (_m *StorageServiceInterfaceMock) DownloadVideo(videoID string, filename string) (string, error) {
	ret := _m.Called(videoID, filename)

	if len(ret) == 0 {
		panic("no return value specified for DownloadVideo")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(videoID, filename)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(videoID, filename)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(videoID, filename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StorageServiceInterfaceMock_DownloadVideo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DownloadVideo'
type StorageServiceInterfaceMock_DownloadVideo_Call struct {
	*mock.Call
}

// DownloadVideo is a helper method to define mock.On call
//   - videoID string
//   - filename string
func (_e *StorageServiceInterfaceMock_Expecter) DownloadVideo(videoID interface{}, filename interface{}) *StorageServiceInterfaceMock_DownloadVideo_Call {
	return &StorageServiceInterfaceMock_DownloadVideo_Call{Call: _e.mock.On("DownloadVideo", videoID, filename)}
}

func (_c *StorageServiceInterfaceMock_DownloadVideo_Call) Run(run func(videoID string, filename string)) *StorageServiceInterfaceMock_DownloadVideo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *StorageServiceInterfaceMock_DownloadVideo_Call) Return(_a0 string, _a1 error) *StorageServiceInterfaceMock_DownloadVideo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StorageServiceInterfaceMock_DownloadVideo_Call) RunAndReturn(run func(string, string) (string, error)) *StorageServiceInterfaceMock_DownloadVideo_Call {
	_c.Call.Return(run)
	return _c
}

// DownloadZipFrames provides a mock function with given fields: videoID
func (_m *StorageServiceInterfaceMock) DownloadZipFrames(videoID string) ([]byte, error) {
	ret := _m.Called(videoID)

	if len(ret) == 0 {
		panic("no return value specified for DownloadZipFrames")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]byte, error)); ok {
		return rf(videoID)
	}
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(videoID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(videoID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StorageServiceInterfaceMock_DownloadZipFrames_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DownloadZipFrames'
type StorageServiceInterfaceMock_DownloadZipFrames_Call struct {
	*mock.Call
}

// DownloadZipFrames is a helper method to define mock.On call
//   - videoID string
func (_e *StorageServiceInterfaceMock_Expecter) DownloadZipFrames(videoID interface{}) *StorageServiceInterfaceMock_DownloadZipFrames_Call {
	return &StorageServiceInterfaceMock_DownloadZipFrames_Call{Call: _e.mock.On("DownloadZipFrames", videoID)}
}

func (_c *StorageServiceInterfaceMock_DownloadZipFrames_Call) Run(run func(videoID string)) *StorageServiceInterfaceMock_DownloadZipFrames_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *StorageServiceInterfaceMock_DownloadZipFrames_Call) Return(_a0 []byte, _a1 error) *StorageServiceInterfaceMock_DownloadZipFrames_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StorageServiceInterfaceMock_DownloadZipFrames_Call) RunAndReturn(run func(string) ([]byte, error)) *StorageServiceInterfaceMock_DownloadZipFrames_Call {
	_c.Call.Return(run)
	return _c
}

// GetExternalFramesDir provides a mock function with no fields
func (_m *StorageServiceInterfaceMock) GetExternalFramesDir() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetExternalFramesDir")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// StorageServiceInterfaceMock_GetExternalFramesDir_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExternalFramesDir'
type StorageServiceInterfaceMock_GetExternalFramesDir_Call struct {
	*mock.Call
}

// GetExternalFramesDir is a helper method to define mock.On call
func (_e *StorageServiceInterfaceMock_Expecter) GetExternalFramesDir() *StorageServiceInterfaceMock_GetExternalFramesDir_Call {
	return &StorageServiceInterfaceMock_GetExternalFramesDir_Call{Call: _e.mock.On("GetExternalFramesDir")}
}

func (_c *StorageServiceInterfaceMock_GetExternalFramesDir_Call) Run(run func()) *StorageServiceInterfaceMock_GetExternalFramesDir_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *StorageServiceInterfaceMock_GetExternalFramesDir_Call) Return(_a0 string) *StorageServiceInterfaceMock_GetExternalFramesDir_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StorageServiceInterfaceMock_GetExternalFramesDir_Call) RunAndReturn(run func() string) *StorageServiceInterfaceMock_GetExternalFramesDir_Call {
	_c.Call.Return(run)
	return _c
}

// GetExternalVideoDir provides a mock function with no fields
func (_m *StorageServiceInterfaceMock) GetExternalVideoDir() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetExternalVideoDir")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// StorageServiceInterfaceMock_GetExternalVideoDir_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExternalVideoDir'
type StorageServiceInterfaceMock_GetExternalVideoDir_Call struct {
	*mock.Call
}

// GetExternalVideoDir is a helper method to define mock.On call
func (_e *StorageServiceInterfaceMock_Expecter) GetExternalVideoDir() *StorageServiceInterfaceMock_GetExternalVideoDir_Call {
	return &StorageServiceInterfaceMock_GetExternalVideoDir_Call{Call: _e.mock.On("GetExternalVideoDir")}
}

func (_c *StorageServiceInterfaceMock_GetExternalVideoDir_Call) Run(run func()) *StorageServiceInterfaceMock_GetExternalVideoDir_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *StorageServiceInterfaceMock_GetExternalVideoDir_Call) Return(_a0 string) *StorageServiceInterfaceMock_GetExternalVideoDir_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StorageServiceInterfaceMock_GetExternalVideoDir_Call) RunAndReturn(run func() string) *StorageServiceInterfaceMock_GetExternalVideoDir_Call {
	_c.Call.Return(run)
	return _c
}

// GetLocalVideoDir provides a mock function with given fields: videoID
func (_m *StorageServiceInterfaceMock) GetLocalVideoDir(videoID string) string {
	ret := _m.Called(videoID)

	if len(ret) == 0 {
		panic("no return value specified for GetLocalVideoDir")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(videoID)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// StorageServiceInterfaceMock_GetLocalVideoDir_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLocalVideoDir'
type StorageServiceInterfaceMock_GetLocalVideoDir_Call struct {
	*mock.Call
}

// GetLocalVideoDir is a helper method to define mock.On call
//   - videoID string
func (_e *StorageServiceInterfaceMock_Expecter) GetLocalVideoDir(videoID interface{}) *StorageServiceInterfaceMock_GetLocalVideoDir_Call {
	return &StorageServiceInterfaceMock_GetLocalVideoDir_Call{Call: _e.mock.On("GetLocalVideoDir", videoID)}
}

func (_c *StorageServiceInterfaceMock_GetLocalVideoDir_Call) Run(run func(videoID string)) *StorageServiceInterfaceMock_GetLocalVideoDir_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *StorageServiceInterfaceMock_GetLocalVideoDir_Call) Return(_a0 string) *StorageServiceInterfaceMock_GetLocalVideoDir_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StorageServiceInterfaceMock_GetLocalVideoDir_Call) RunAndReturn(run func(string) string) *StorageServiceInterfaceMock_GetLocalVideoDir_Call {
	_c.Call.Return(run)
	return _c
}

// UploadVideo provides a mock function with given fields: filename, file
func (_m *StorageServiceInterfaceMock) UploadVideo(filename string, file multipart.File) (string, error) {
	ret := _m.Called(filename, file)

	if len(ret) == 0 {
		panic("no return value specified for UploadVideo")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, multipart.File) (string, error)); ok {
		return rf(filename, file)
	}
	if rf, ok := ret.Get(0).(func(string, multipart.File) string); ok {
		r0 = rf(filename, file)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, multipart.File) error); ok {
		r1 = rf(filename, file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StorageServiceInterfaceMock_UploadVideo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UploadVideo'
type StorageServiceInterfaceMock_UploadVideo_Call struct {
	*mock.Call
}

// UploadVideo is a helper method to define mock.On call
//   - filename string
//   - file multipart.File
func (_e *StorageServiceInterfaceMock_Expecter) UploadVideo(filename interface{}, file interface{}) *StorageServiceInterfaceMock_UploadVideo_Call {
	return &StorageServiceInterfaceMock_UploadVideo_Call{Call: _e.mock.On("UploadVideo", filename, file)}
}

func (_c *StorageServiceInterfaceMock_UploadVideo_Call) Run(run func(filename string, file multipart.File)) *StorageServiceInterfaceMock_UploadVideo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(multipart.File))
	})
	return _c
}

func (_c *StorageServiceInterfaceMock_UploadVideo_Call) Return(_a0 string, _a1 error) *StorageServiceInterfaceMock_UploadVideo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StorageServiceInterfaceMock_UploadVideo_Call) RunAndReturn(run func(string, multipart.File) (string, error)) *StorageServiceInterfaceMock_UploadVideo_Call {
	_c.Call.Return(run)
	return _c
}

// UploadZipFrames provides a mock function with given fields: filename, file
func (_m *StorageServiceInterfaceMock) UploadZipFrames(filename string, file multipart.File) (string, error) {
	ret := _m.Called(filename, file)

	if len(ret) == 0 {
		panic("no return value specified for UploadZipFrames")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, multipart.File) (string, error)); ok {
		return rf(filename, file)
	}
	if rf, ok := ret.Get(0).(func(string, multipart.File) string); ok {
		r0 = rf(filename, file)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, multipart.File) error); ok {
		r1 = rf(filename, file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StorageServiceInterfaceMock_UploadZipFrames_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UploadZipFrames'
type StorageServiceInterfaceMock_UploadZipFrames_Call struct {
	*mock.Call
}

// UploadZipFrames is a helper method to define mock.On call
//   - filename string
//   - file multipart.File
func (_e *StorageServiceInterfaceMock_Expecter) UploadZipFrames(filename interface{}, file interface{}) *StorageServiceInterfaceMock_UploadZipFrames_Call {
	return &StorageServiceInterfaceMock_UploadZipFrames_Call{Call: _e.mock.On("UploadZipFrames", filename, file)}
}

func (_c *StorageServiceInterfaceMock_UploadZipFrames_Call) Run(run func(filename string, file multipart.File)) *StorageServiceInterfaceMock_UploadZipFrames_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(multipart.File))
	})
	return _c
}

func (_c *StorageServiceInterfaceMock_UploadZipFrames_Call) Return(_a0 string, _a1 error) *StorageServiceInterfaceMock_UploadZipFrames_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StorageServiceInterfaceMock_UploadZipFrames_Call) RunAndReturn(run func(string, multipart.File) (string, error)) *StorageServiceInterfaceMock_UploadZipFrames_Call {
	_c.Call.Return(run)
	return _c
}

// NewStorageServiceInterfaceMock creates a new instance of StorageServiceInterfaceMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStorageServiceInterfaceMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *StorageServiceInterfaceMock {
	mock := &StorageServiceInterfaceMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
