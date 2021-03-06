// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	sync "sync"

	v2v3action "code.cloudfoundry.org/cli/actor/v2v3action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeShareServiceActor struct {
	CloudControllerV3APIVersionStub        func() string
	cloudControllerV3APIVersionMutex       sync.RWMutex
	cloudControllerV3APIVersionArgsForCall []struct {
	}
	cloudControllerV3APIVersionReturns struct {
		result1 string
	}
	cloudControllerV3APIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationStub        func(string, string, string, string) (v2v3action.Warnings, error)
	shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationMutex       sync.RWMutex
	shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}
	shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationReturns struct {
		result1 v2v3action.Warnings
		result2 error
	}
	shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationReturnsOnCall map[int]struct {
		result1 v2v3action.Warnings
		result2 error
	}
	ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameStub        func(string, string, string, string) (v2v3action.Warnings, error)
	shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameMutex       sync.RWMutex
	shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}
	shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameReturns struct {
		result1 v2v3action.Warnings
		result2 error
	}
	shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameReturnsOnCall map[int]struct {
		result1 v2v3action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeShareServiceActor) CloudControllerV3APIVersion() string {
	fake.cloudControllerV3APIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerV3APIVersionReturnsOnCall[len(fake.cloudControllerV3APIVersionArgsForCall)]
	fake.cloudControllerV3APIVersionArgsForCall = append(fake.cloudControllerV3APIVersionArgsForCall, struct {
	}{})
	fake.recordInvocation("CloudControllerV3APIVersion", []interface{}{})
	fake.cloudControllerV3APIVersionMutex.Unlock()
	if fake.CloudControllerV3APIVersionStub != nil {
		return fake.CloudControllerV3APIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.cloudControllerV3APIVersionReturns
	return fakeReturns.result1
}

func (fake *FakeShareServiceActor) CloudControllerV3APIVersionCallCount() int {
	fake.cloudControllerV3APIVersionMutex.RLock()
	defer fake.cloudControllerV3APIVersionMutex.RUnlock()
	return len(fake.cloudControllerV3APIVersionArgsForCall)
}

func (fake *FakeShareServiceActor) CloudControllerV3APIVersionCalls(stub func() string) {
	fake.cloudControllerV3APIVersionMutex.Lock()
	defer fake.cloudControllerV3APIVersionMutex.Unlock()
	fake.CloudControllerV3APIVersionStub = stub
}

func (fake *FakeShareServiceActor) CloudControllerV3APIVersionReturns(result1 string) {
	fake.cloudControllerV3APIVersionMutex.Lock()
	defer fake.cloudControllerV3APIVersionMutex.Unlock()
	fake.CloudControllerV3APIVersionStub = nil
	fake.cloudControllerV3APIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeShareServiceActor) CloudControllerV3APIVersionReturnsOnCall(i int, result1 string) {
	fake.cloudControllerV3APIVersionMutex.Lock()
	defer fake.cloudControllerV3APIVersionMutex.Unlock()
	fake.CloudControllerV3APIVersionStub = nil
	if fake.cloudControllerV3APIVersionReturnsOnCall == nil {
		fake.cloudControllerV3APIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerV3APIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeShareServiceActor) ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganization(arg1 string, arg2 string, arg3 string, arg4 string) (v2v3action.Warnings, error) {
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationMutex.Lock()
	ret, specificReturn := fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationReturnsOnCall[len(fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationArgsForCall)]
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationArgsForCall = append(fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganization", []interface{}{arg1, arg2, arg3, arg4})
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationMutex.Unlock()
	if fake.ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationStub != nil {
		return fake.ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeShareServiceActor) ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationCallCount() int {
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationMutex.RLock()
	defer fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationMutex.RUnlock()
	return len(fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationArgsForCall)
}

func (fake *FakeShareServiceActor) ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationCalls(stub func(string, string, string, string) (v2v3action.Warnings, error)) {
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationMutex.Lock()
	defer fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationMutex.Unlock()
	fake.ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationStub = stub
}

func (fake *FakeShareServiceActor) ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationArgsForCall(i int) (string, string, string, string) {
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationMutex.RLock()
	defer fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationMutex.RUnlock()
	argsForCall := fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeShareServiceActor) ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationReturns(result1 v2v3action.Warnings, result2 error) {
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationMutex.Lock()
	defer fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationMutex.Unlock()
	fake.ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationStub = nil
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationReturns = struct {
		result1 v2v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeShareServiceActor) ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationReturnsOnCall(i int, result1 v2v3action.Warnings, result2 error) {
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationMutex.Lock()
	defer fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationMutex.Unlock()
	fake.ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationStub = nil
	if fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationReturnsOnCall == nil {
		fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationReturnsOnCall = make(map[int]struct {
			result1 v2v3action.Warnings
			result2 error
		})
	}
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationReturnsOnCall[i] = struct {
		result1 v2v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeShareServiceActor) ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationName(arg1 string, arg2 string, arg3 string, arg4 string) (v2v3action.Warnings, error) {
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameMutex.Lock()
	ret, specificReturn := fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameReturnsOnCall[len(fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameArgsForCall)]
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameArgsForCall = append(fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationName", []interface{}{arg1, arg2, arg3, arg4})
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameMutex.Unlock()
	if fake.ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameStub != nil {
		return fake.ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeShareServiceActor) ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameCallCount() int {
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameMutex.RLock()
	defer fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameMutex.RUnlock()
	return len(fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameArgsForCall)
}

func (fake *FakeShareServiceActor) ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameCalls(stub func(string, string, string, string) (v2v3action.Warnings, error)) {
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameMutex.Lock()
	defer fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameMutex.Unlock()
	fake.ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameStub = stub
}

func (fake *FakeShareServiceActor) ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameArgsForCall(i int) (string, string, string, string) {
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameMutex.RLock()
	defer fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameMutex.RUnlock()
	argsForCall := fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeShareServiceActor) ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameReturns(result1 v2v3action.Warnings, result2 error) {
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameMutex.Lock()
	defer fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameMutex.Unlock()
	fake.ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameStub = nil
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameReturns = struct {
		result1 v2v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeShareServiceActor) ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameReturnsOnCall(i int, result1 v2v3action.Warnings, result2 error) {
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameMutex.Lock()
	defer fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameMutex.Unlock()
	fake.ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameStub = nil
	if fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameReturnsOnCall == nil {
		fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameReturnsOnCall = make(map[int]struct {
			result1 v2v3action.Warnings
			result2 error
		})
	}
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameReturnsOnCall[i] = struct {
		result1 v2v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeShareServiceActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerV3APIVersionMutex.RLock()
	defer fake.cloudControllerV3APIVersionMutex.RUnlock()
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationMutex.RLock()
	defer fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationMutex.RUnlock()
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameMutex.RLock()
	defer fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeShareServiceActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.ShareServiceActor = new(FakeShareServiceActor)
