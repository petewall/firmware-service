// Code generated by counterfeiter. DO NOT EDIT.
package internalfakes

import (
	"sync"

	"github.com/petewall/firmware-service/v2/internal"
	"github.com/petewall/firmware-service/v2/lib"
)

type FakeFirmwareStore struct {
	AddFirmwareStub        func(string, string, []byte) error
	addFirmwareMutex       sync.RWMutex
	addFirmwareArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 []byte
	}
	addFirmwareReturns struct {
		result1 error
	}
	addFirmwareReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteFirmwareStub        func(string, string) error
	deleteFirmwareMutex       sync.RWMutex
	deleteFirmwareArgsForCall []struct {
		arg1 string
		arg2 string
	}
	deleteFirmwareReturns struct {
		result1 error
	}
	deleteFirmwareReturnsOnCall map[int]struct {
		result1 error
	}
	GetAllFirmwareStub        func() (lib.FirmwareList, error)
	getAllFirmwareMutex       sync.RWMutex
	getAllFirmwareArgsForCall []struct {
	}
	getAllFirmwareReturns struct {
		result1 lib.FirmwareList
		result2 error
	}
	getAllFirmwareReturnsOnCall map[int]struct {
		result1 lib.FirmwareList
		result2 error
	}
	GetAllFirmwareByTypeStub        func(string) (lib.FirmwareList, error)
	getAllFirmwareByTypeMutex       sync.RWMutex
	getAllFirmwareByTypeArgsForCall []struct {
		arg1 string
	}
	getAllFirmwareByTypeReturns struct {
		result1 lib.FirmwareList
		result2 error
	}
	getAllFirmwareByTypeReturnsOnCall map[int]struct {
		result1 lib.FirmwareList
		result2 error
	}
	GetAllTypesStub        func() ([]string, error)
	getAllTypesMutex       sync.RWMutex
	getAllTypesArgsForCall []struct {
	}
	getAllTypesReturns struct {
		result1 []string
		result2 error
	}
	getAllTypesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	GetFirmwareStub        func(string, string) (*lib.Firmware, error)
	getFirmwareMutex       sync.RWMutex
	getFirmwareArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getFirmwareReturns struct {
		result1 *lib.Firmware
		result2 error
	}
	getFirmwareReturnsOnCall map[int]struct {
		result1 *lib.Firmware
		result2 error
	}
	GetFirmwareDataStub        func(string, string) ([]byte, error)
	getFirmwareDataMutex       sync.RWMutex
	getFirmwareDataArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getFirmwareDataReturns struct {
		result1 []byte
		result2 error
	}
	getFirmwareDataReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFirmwareStore) AddFirmware(arg1 string, arg2 string, arg3 []byte) error {
	var arg3Copy []byte
	if arg3 != nil {
		arg3Copy = make([]byte, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.addFirmwareMutex.Lock()
	ret, specificReturn := fake.addFirmwareReturnsOnCall[len(fake.addFirmwareArgsForCall)]
	fake.addFirmwareArgsForCall = append(fake.addFirmwareArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 []byte
	}{arg1, arg2, arg3Copy})
	stub := fake.AddFirmwareStub
	fakeReturns := fake.addFirmwareReturns
	fake.recordInvocation("AddFirmware", []interface{}{arg1, arg2, arg3Copy})
	fake.addFirmwareMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeFirmwareStore) AddFirmwareCallCount() int {
	fake.addFirmwareMutex.RLock()
	defer fake.addFirmwareMutex.RUnlock()
	return len(fake.addFirmwareArgsForCall)
}

func (fake *FakeFirmwareStore) AddFirmwareCalls(stub func(string, string, []byte) error) {
	fake.addFirmwareMutex.Lock()
	defer fake.addFirmwareMutex.Unlock()
	fake.AddFirmwareStub = stub
}

func (fake *FakeFirmwareStore) AddFirmwareArgsForCall(i int) (string, string, []byte) {
	fake.addFirmwareMutex.RLock()
	defer fake.addFirmwareMutex.RUnlock()
	argsForCall := fake.addFirmwareArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeFirmwareStore) AddFirmwareReturns(result1 error) {
	fake.addFirmwareMutex.Lock()
	defer fake.addFirmwareMutex.Unlock()
	fake.AddFirmwareStub = nil
	fake.addFirmwareReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFirmwareStore) AddFirmwareReturnsOnCall(i int, result1 error) {
	fake.addFirmwareMutex.Lock()
	defer fake.addFirmwareMutex.Unlock()
	fake.AddFirmwareStub = nil
	if fake.addFirmwareReturnsOnCall == nil {
		fake.addFirmwareReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addFirmwareReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFirmwareStore) DeleteFirmware(arg1 string, arg2 string) error {
	fake.deleteFirmwareMutex.Lock()
	ret, specificReturn := fake.deleteFirmwareReturnsOnCall[len(fake.deleteFirmwareArgsForCall)]
	fake.deleteFirmwareArgsForCall = append(fake.deleteFirmwareArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.DeleteFirmwareStub
	fakeReturns := fake.deleteFirmwareReturns
	fake.recordInvocation("DeleteFirmware", []interface{}{arg1, arg2})
	fake.deleteFirmwareMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeFirmwareStore) DeleteFirmwareCallCount() int {
	fake.deleteFirmwareMutex.RLock()
	defer fake.deleteFirmwareMutex.RUnlock()
	return len(fake.deleteFirmwareArgsForCall)
}

func (fake *FakeFirmwareStore) DeleteFirmwareCalls(stub func(string, string) error) {
	fake.deleteFirmwareMutex.Lock()
	defer fake.deleteFirmwareMutex.Unlock()
	fake.DeleteFirmwareStub = stub
}

func (fake *FakeFirmwareStore) DeleteFirmwareArgsForCall(i int) (string, string) {
	fake.deleteFirmwareMutex.RLock()
	defer fake.deleteFirmwareMutex.RUnlock()
	argsForCall := fake.deleteFirmwareArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeFirmwareStore) DeleteFirmwareReturns(result1 error) {
	fake.deleteFirmwareMutex.Lock()
	defer fake.deleteFirmwareMutex.Unlock()
	fake.DeleteFirmwareStub = nil
	fake.deleteFirmwareReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFirmwareStore) DeleteFirmwareReturnsOnCall(i int, result1 error) {
	fake.deleteFirmwareMutex.Lock()
	defer fake.deleteFirmwareMutex.Unlock()
	fake.DeleteFirmwareStub = nil
	if fake.deleteFirmwareReturnsOnCall == nil {
		fake.deleteFirmwareReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteFirmwareReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFirmwareStore) GetAllFirmware() (lib.FirmwareList, error) {
	fake.getAllFirmwareMutex.Lock()
	ret, specificReturn := fake.getAllFirmwareReturnsOnCall[len(fake.getAllFirmwareArgsForCall)]
	fake.getAllFirmwareArgsForCall = append(fake.getAllFirmwareArgsForCall, struct {
	}{})
	stub := fake.GetAllFirmwareStub
	fakeReturns := fake.getAllFirmwareReturns
	fake.recordInvocation("GetAllFirmware", []interface{}{})
	fake.getAllFirmwareMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFirmwareStore) GetAllFirmwareCallCount() int {
	fake.getAllFirmwareMutex.RLock()
	defer fake.getAllFirmwareMutex.RUnlock()
	return len(fake.getAllFirmwareArgsForCall)
}

func (fake *FakeFirmwareStore) GetAllFirmwareCalls(stub func() (lib.FirmwareList, error)) {
	fake.getAllFirmwareMutex.Lock()
	defer fake.getAllFirmwareMutex.Unlock()
	fake.GetAllFirmwareStub = stub
}

func (fake *FakeFirmwareStore) GetAllFirmwareReturns(result1 lib.FirmwareList, result2 error) {
	fake.getAllFirmwareMutex.Lock()
	defer fake.getAllFirmwareMutex.Unlock()
	fake.GetAllFirmwareStub = nil
	fake.getAllFirmwareReturns = struct {
		result1 lib.FirmwareList
		result2 error
	}{result1, result2}
}

func (fake *FakeFirmwareStore) GetAllFirmwareReturnsOnCall(i int, result1 lib.FirmwareList, result2 error) {
	fake.getAllFirmwareMutex.Lock()
	defer fake.getAllFirmwareMutex.Unlock()
	fake.GetAllFirmwareStub = nil
	if fake.getAllFirmwareReturnsOnCall == nil {
		fake.getAllFirmwareReturnsOnCall = make(map[int]struct {
			result1 lib.FirmwareList
			result2 error
		})
	}
	fake.getAllFirmwareReturnsOnCall[i] = struct {
		result1 lib.FirmwareList
		result2 error
	}{result1, result2}
}

func (fake *FakeFirmwareStore) GetAllFirmwareByType(arg1 string) (lib.FirmwareList, error) {
	fake.getAllFirmwareByTypeMutex.Lock()
	ret, specificReturn := fake.getAllFirmwareByTypeReturnsOnCall[len(fake.getAllFirmwareByTypeArgsForCall)]
	fake.getAllFirmwareByTypeArgsForCall = append(fake.getAllFirmwareByTypeArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetAllFirmwareByTypeStub
	fakeReturns := fake.getAllFirmwareByTypeReturns
	fake.recordInvocation("GetAllFirmwareByType", []interface{}{arg1})
	fake.getAllFirmwareByTypeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFirmwareStore) GetAllFirmwareByTypeCallCount() int {
	fake.getAllFirmwareByTypeMutex.RLock()
	defer fake.getAllFirmwareByTypeMutex.RUnlock()
	return len(fake.getAllFirmwareByTypeArgsForCall)
}

func (fake *FakeFirmwareStore) GetAllFirmwareByTypeCalls(stub func(string) (lib.FirmwareList, error)) {
	fake.getAllFirmwareByTypeMutex.Lock()
	defer fake.getAllFirmwareByTypeMutex.Unlock()
	fake.GetAllFirmwareByTypeStub = stub
}

func (fake *FakeFirmwareStore) GetAllFirmwareByTypeArgsForCall(i int) string {
	fake.getAllFirmwareByTypeMutex.RLock()
	defer fake.getAllFirmwareByTypeMutex.RUnlock()
	argsForCall := fake.getAllFirmwareByTypeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFirmwareStore) GetAllFirmwareByTypeReturns(result1 lib.FirmwareList, result2 error) {
	fake.getAllFirmwareByTypeMutex.Lock()
	defer fake.getAllFirmwareByTypeMutex.Unlock()
	fake.GetAllFirmwareByTypeStub = nil
	fake.getAllFirmwareByTypeReturns = struct {
		result1 lib.FirmwareList
		result2 error
	}{result1, result2}
}

func (fake *FakeFirmwareStore) GetAllFirmwareByTypeReturnsOnCall(i int, result1 lib.FirmwareList, result2 error) {
	fake.getAllFirmwareByTypeMutex.Lock()
	defer fake.getAllFirmwareByTypeMutex.Unlock()
	fake.GetAllFirmwareByTypeStub = nil
	if fake.getAllFirmwareByTypeReturnsOnCall == nil {
		fake.getAllFirmwareByTypeReturnsOnCall = make(map[int]struct {
			result1 lib.FirmwareList
			result2 error
		})
	}
	fake.getAllFirmwareByTypeReturnsOnCall[i] = struct {
		result1 lib.FirmwareList
		result2 error
	}{result1, result2}
}

func (fake *FakeFirmwareStore) GetAllTypes() ([]string, error) {
	fake.getAllTypesMutex.Lock()
	ret, specificReturn := fake.getAllTypesReturnsOnCall[len(fake.getAllTypesArgsForCall)]
	fake.getAllTypesArgsForCall = append(fake.getAllTypesArgsForCall, struct {
	}{})
	stub := fake.GetAllTypesStub
	fakeReturns := fake.getAllTypesReturns
	fake.recordInvocation("GetAllTypes", []interface{}{})
	fake.getAllTypesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFirmwareStore) GetAllTypesCallCount() int {
	fake.getAllTypesMutex.RLock()
	defer fake.getAllTypesMutex.RUnlock()
	return len(fake.getAllTypesArgsForCall)
}

func (fake *FakeFirmwareStore) GetAllTypesCalls(stub func() ([]string, error)) {
	fake.getAllTypesMutex.Lock()
	defer fake.getAllTypesMutex.Unlock()
	fake.GetAllTypesStub = stub
}

func (fake *FakeFirmwareStore) GetAllTypesReturns(result1 []string, result2 error) {
	fake.getAllTypesMutex.Lock()
	defer fake.getAllTypesMutex.Unlock()
	fake.GetAllTypesStub = nil
	fake.getAllTypesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeFirmwareStore) GetAllTypesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.getAllTypesMutex.Lock()
	defer fake.getAllTypesMutex.Unlock()
	fake.GetAllTypesStub = nil
	if fake.getAllTypesReturnsOnCall == nil {
		fake.getAllTypesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.getAllTypesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeFirmwareStore) GetFirmware(arg1 string, arg2 string) (*lib.Firmware, error) {
	fake.getFirmwareMutex.Lock()
	ret, specificReturn := fake.getFirmwareReturnsOnCall[len(fake.getFirmwareArgsForCall)]
	fake.getFirmwareArgsForCall = append(fake.getFirmwareArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.GetFirmwareStub
	fakeReturns := fake.getFirmwareReturns
	fake.recordInvocation("GetFirmware", []interface{}{arg1, arg2})
	fake.getFirmwareMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFirmwareStore) GetFirmwareCallCount() int {
	fake.getFirmwareMutex.RLock()
	defer fake.getFirmwareMutex.RUnlock()
	return len(fake.getFirmwareArgsForCall)
}

func (fake *FakeFirmwareStore) GetFirmwareCalls(stub func(string, string) (*lib.Firmware, error)) {
	fake.getFirmwareMutex.Lock()
	defer fake.getFirmwareMutex.Unlock()
	fake.GetFirmwareStub = stub
}

func (fake *FakeFirmwareStore) GetFirmwareArgsForCall(i int) (string, string) {
	fake.getFirmwareMutex.RLock()
	defer fake.getFirmwareMutex.RUnlock()
	argsForCall := fake.getFirmwareArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeFirmwareStore) GetFirmwareReturns(result1 *lib.Firmware, result2 error) {
	fake.getFirmwareMutex.Lock()
	defer fake.getFirmwareMutex.Unlock()
	fake.GetFirmwareStub = nil
	fake.getFirmwareReturns = struct {
		result1 *lib.Firmware
		result2 error
	}{result1, result2}
}

func (fake *FakeFirmwareStore) GetFirmwareReturnsOnCall(i int, result1 *lib.Firmware, result2 error) {
	fake.getFirmwareMutex.Lock()
	defer fake.getFirmwareMutex.Unlock()
	fake.GetFirmwareStub = nil
	if fake.getFirmwareReturnsOnCall == nil {
		fake.getFirmwareReturnsOnCall = make(map[int]struct {
			result1 *lib.Firmware
			result2 error
		})
	}
	fake.getFirmwareReturnsOnCall[i] = struct {
		result1 *lib.Firmware
		result2 error
	}{result1, result2}
}

func (fake *FakeFirmwareStore) GetFirmwareData(arg1 string, arg2 string) ([]byte, error) {
	fake.getFirmwareDataMutex.Lock()
	ret, specificReturn := fake.getFirmwareDataReturnsOnCall[len(fake.getFirmwareDataArgsForCall)]
	fake.getFirmwareDataArgsForCall = append(fake.getFirmwareDataArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.GetFirmwareDataStub
	fakeReturns := fake.getFirmwareDataReturns
	fake.recordInvocation("GetFirmwareData", []interface{}{arg1, arg2})
	fake.getFirmwareDataMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFirmwareStore) GetFirmwareDataCallCount() int {
	fake.getFirmwareDataMutex.RLock()
	defer fake.getFirmwareDataMutex.RUnlock()
	return len(fake.getFirmwareDataArgsForCall)
}

func (fake *FakeFirmwareStore) GetFirmwareDataCalls(stub func(string, string) ([]byte, error)) {
	fake.getFirmwareDataMutex.Lock()
	defer fake.getFirmwareDataMutex.Unlock()
	fake.GetFirmwareDataStub = stub
}

func (fake *FakeFirmwareStore) GetFirmwareDataArgsForCall(i int) (string, string) {
	fake.getFirmwareDataMutex.RLock()
	defer fake.getFirmwareDataMutex.RUnlock()
	argsForCall := fake.getFirmwareDataArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeFirmwareStore) GetFirmwareDataReturns(result1 []byte, result2 error) {
	fake.getFirmwareDataMutex.Lock()
	defer fake.getFirmwareDataMutex.Unlock()
	fake.GetFirmwareDataStub = nil
	fake.getFirmwareDataReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeFirmwareStore) GetFirmwareDataReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.getFirmwareDataMutex.Lock()
	defer fake.getFirmwareDataMutex.Unlock()
	fake.GetFirmwareDataStub = nil
	if fake.getFirmwareDataReturnsOnCall == nil {
		fake.getFirmwareDataReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getFirmwareDataReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeFirmwareStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addFirmwareMutex.RLock()
	defer fake.addFirmwareMutex.RUnlock()
	fake.deleteFirmwareMutex.RLock()
	defer fake.deleteFirmwareMutex.RUnlock()
	fake.getAllFirmwareMutex.RLock()
	defer fake.getAllFirmwareMutex.RUnlock()
	fake.getAllFirmwareByTypeMutex.RLock()
	defer fake.getAllFirmwareByTypeMutex.RUnlock()
	fake.getAllTypesMutex.RLock()
	defer fake.getAllTypesMutex.RUnlock()
	fake.getFirmwareMutex.RLock()
	defer fake.getFirmwareMutex.RUnlock()
	fake.getFirmwareDataMutex.RLock()
	defer fake.getFirmwareDataMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFirmwareStore) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ internal.FirmwareStore = new(FakeFirmwareStore)