// Code generated by counterfeiter. DO NOT EDIT.
package fluxfakes

import (
	"sync"

	"github.com/weaveworks/weave-gitops/pkg/flux"
)

type FakeFlux struct {
	CreateHelmReleaseGitRepositoryStub        func(string, string, string, string) ([]byte, error)
	createHelmReleaseGitRepositoryMutex       sync.RWMutex
	createHelmReleaseGitRepositoryArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}
	createHelmReleaseGitRepositoryReturns struct {
		result1 []byte
		result2 error
	}
	createHelmReleaseGitRepositoryReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	CreateHelmReleaseHelmRepositoryStub        func(string, string, string) ([]byte, error)
	createHelmReleaseHelmRepositoryMutex       sync.RWMutex
	createHelmReleaseHelmRepositoryArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	createHelmReleaseHelmRepositoryReturns struct {
		result1 []byte
		result2 error
	}
	createHelmReleaseHelmRepositoryReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	CreateKustomizationStub        func(string, string, string, string) ([]byte, error)
	createKustomizationMutex       sync.RWMutex
	createKustomizationArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}
	createKustomizationReturns struct {
		result1 []byte
		result2 error
	}
	createKustomizationReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	CreateSecretGitStub        func(string, string, string) ([]byte, error)
	createSecretGitMutex       sync.RWMutex
	createSecretGitArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	createSecretGitReturns struct {
		result1 []byte
		result2 error
	}
	createSecretGitReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	CreateSourceGitStub        func(string, string, string, string, string) ([]byte, error)
	createSourceGitMutex       sync.RWMutex
	createSourceGitArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}
	createSourceGitReturns struct {
		result1 []byte
		result2 error
	}
	createSourceGitReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	CreateSourceHelmStub        func(string, string, string) ([]byte, error)
	createSourceHelmMutex       sync.RWMutex
	createSourceHelmArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	createSourceHelmReturns struct {
		result1 []byte
		result2 error
	}
	createSourceHelmReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	InstallStub        func(string, bool) ([]byte, error)
	installMutex       sync.RWMutex
	installArgsForCall []struct {
		arg1 string
		arg2 bool
	}
	installReturns struct {
		result1 []byte
		result2 error
	}
	installReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	UninstallStub        func(string, bool) error
	uninstallMutex       sync.RWMutex
	uninstallArgsForCall []struct {
		arg1 string
		arg2 bool
	}
	uninstallReturns struct {
		result1 error
	}
	uninstallReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFlux) CreateHelmReleaseGitRepository(arg1 string, arg2 string, arg3 string, arg4 string) ([]byte, error) {
	fake.createHelmReleaseGitRepositoryMutex.Lock()
	ret, specificReturn := fake.createHelmReleaseGitRepositoryReturnsOnCall[len(fake.createHelmReleaseGitRepositoryArgsForCall)]
	fake.createHelmReleaseGitRepositoryArgsForCall = append(fake.createHelmReleaseGitRepositoryArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	stub := fake.CreateHelmReleaseGitRepositoryStub
	fakeReturns := fake.createHelmReleaseGitRepositoryReturns
	fake.recordInvocation("CreateHelmReleaseGitRepository", []interface{}{arg1, arg2, arg3, arg4})
	fake.createHelmReleaseGitRepositoryMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFlux) CreateHelmReleaseGitRepositoryCallCount() int {
	fake.createHelmReleaseGitRepositoryMutex.RLock()
	defer fake.createHelmReleaseGitRepositoryMutex.RUnlock()
	return len(fake.createHelmReleaseGitRepositoryArgsForCall)
}

func (fake *FakeFlux) CreateHelmReleaseGitRepositoryCalls(stub func(string, string, string, string) ([]byte, error)) {
	fake.createHelmReleaseGitRepositoryMutex.Lock()
	defer fake.createHelmReleaseGitRepositoryMutex.Unlock()
	fake.CreateHelmReleaseGitRepositoryStub = stub
}

func (fake *FakeFlux) CreateHelmReleaseGitRepositoryArgsForCall(i int) (string, string, string, string) {
	fake.createHelmReleaseGitRepositoryMutex.RLock()
	defer fake.createHelmReleaseGitRepositoryMutex.RUnlock()
	argsForCall := fake.createHelmReleaseGitRepositoryArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeFlux) CreateHelmReleaseGitRepositoryReturns(result1 []byte, result2 error) {
	fake.createHelmReleaseGitRepositoryMutex.Lock()
	defer fake.createHelmReleaseGitRepositoryMutex.Unlock()
	fake.CreateHelmReleaseGitRepositoryStub = nil
	fake.createHelmReleaseGitRepositoryReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeFlux) CreateHelmReleaseGitRepositoryReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.createHelmReleaseGitRepositoryMutex.Lock()
	defer fake.createHelmReleaseGitRepositoryMutex.Unlock()
	fake.CreateHelmReleaseGitRepositoryStub = nil
	if fake.createHelmReleaseGitRepositoryReturnsOnCall == nil {
		fake.createHelmReleaseGitRepositoryReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.createHelmReleaseGitRepositoryReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeFlux) CreateHelmReleaseHelmRepository(arg1 string, arg2 string, arg3 string) ([]byte, error) {
	fake.createHelmReleaseHelmRepositoryMutex.Lock()
	ret, specificReturn := fake.createHelmReleaseHelmRepositoryReturnsOnCall[len(fake.createHelmReleaseHelmRepositoryArgsForCall)]
	fake.createHelmReleaseHelmRepositoryArgsForCall = append(fake.createHelmReleaseHelmRepositoryArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.CreateHelmReleaseHelmRepositoryStub
	fakeReturns := fake.createHelmReleaseHelmRepositoryReturns
	fake.recordInvocation("CreateHelmReleaseHelmRepository", []interface{}{arg1, arg2, arg3})
	fake.createHelmReleaseHelmRepositoryMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFlux) CreateHelmReleaseHelmRepositoryCallCount() int {
	fake.createHelmReleaseHelmRepositoryMutex.RLock()
	defer fake.createHelmReleaseHelmRepositoryMutex.RUnlock()
	return len(fake.createHelmReleaseHelmRepositoryArgsForCall)
}

func (fake *FakeFlux) CreateHelmReleaseHelmRepositoryCalls(stub func(string, string, string) ([]byte, error)) {
	fake.createHelmReleaseHelmRepositoryMutex.Lock()
	defer fake.createHelmReleaseHelmRepositoryMutex.Unlock()
	fake.CreateHelmReleaseHelmRepositoryStub = stub
}

func (fake *FakeFlux) CreateHelmReleaseHelmRepositoryArgsForCall(i int) (string, string, string) {
	fake.createHelmReleaseHelmRepositoryMutex.RLock()
	defer fake.createHelmReleaseHelmRepositoryMutex.RUnlock()
	argsForCall := fake.createHelmReleaseHelmRepositoryArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeFlux) CreateHelmReleaseHelmRepositoryReturns(result1 []byte, result2 error) {
	fake.createHelmReleaseHelmRepositoryMutex.Lock()
	defer fake.createHelmReleaseHelmRepositoryMutex.Unlock()
	fake.CreateHelmReleaseHelmRepositoryStub = nil
	fake.createHelmReleaseHelmRepositoryReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeFlux) CreateHelmReleaseHelmRepositoryReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.createHelmReleaseHelmRepositoryMutex.Lock()
	defer fake.createHelmReleaseHelmRepositoryMutex.Unlock()
	fake.CreateHelmReleaseHelmRepositoryStub = nil
	if fake.createHelmReleaseHelmRepositoryReturnsOnCall == nil {
		fake.createHelmReleaseHelmRepositoryReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.createHelmReleaseHelmRepositoryReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeFlux) CreateKustomization(arg1 string, arg2 string, arg3 string, arg4 string) ([]byte, error) {
	fake.createKustomizationMutex.Lock()
	ret, specificReturn := fake.createKustomizationReturnsOnCall[len(fake.createKustomizationArgsForCall)]
	fake.createKustomizationArgsForCall = append(fake.createKustomizationArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	stub := fake.CreateKustomizationStub
	fakeReturns := fake.createKustomizationReturns
	fake.recordInvocation("CreateKustomization", []interface{}{arg1, arg2, arg3, arg4})
	fake.createKustomizationMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFlux) CreateKustomizationCallCount() int {
	fake.createKustomizationMutex.RLock()
	defer fake.createKustomizationMutex.RUnlock()
	return len(fake.createKustomizationArgsForCall)
}

func (fake *FakeFlux) CreateKustomizationCalls(stub func(string, string, string, string) ([]byte, error)) {
	fake.createKustomizationMutex.Lock()
	defer fake.createKustomizationMutex.Unlock()
	fake.CreateKustomizationStub = stub
}

func (fake *FakeFlux) CreateKustomizationArgsForCall(i int) (string, string, string, string) {
	fake.createKustomizationMutex.RLock()
	defer fake.createKustomizationMutex.RUnlock()
	argsForCall := fake.createKustomizationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeFlux) CreateKustomizationReturns(result1 []byte, result2 error) {
	fake.createKustomizationMutex.Lock()
	defer fake.createKustomizationMutex.Unlock()
	fake.CreateKustomizationStub = nil
	fake.createKustomizationReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeFlux) CreateKustomizationReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.createKustomizationMutex.Lock()
	defer fake.createKustomizationMutex.Unlock()
	fake.CreateKustomizationStub = nil
	if fake.createKustomizationReturnsOnCall == nil {
		fake.createKustomizationReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.createKustomizationReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeFlux) CreateSecretGit(arg1 string, arg2 string, arg3 string) ([]byte, error) {
	fake.createSecretGitMutex.Lock()
	ret, specificReturn := fake.createSecretGitReturnsOnCall[len(fake.createSecretGitArgsForCall)]
	fake.createSecretGitArgsForCall = append(fake.createSecretGitArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.CreateSecretGitStub
	fakeReturns := fake.createSecretGitReturns
	fake.recordInvocation("CreateSecretGit", []interface{}{arg1, arg2, arg3})
	fake.createSecretGitMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFlux) CreateSecretGitCallCount() int {
	fake.createSecretGitMutex.RLock()
	defer fake.createSecretGitMutex.RUnlock()
	return len(fake.createSecretGitArgsForCall)
}

func (fake *FakeFlux) CreateSecretGitCalls(stub func(string, string, string) ([]byte, error)) {
	fake.createSecretGitMutex.Lock()
	defer fake.createSecretGitMutex.Unlock()
	fake.CreateSecretGitStub = stub
}

func (fake *FakeFlux) CreateSecretGitArgsForCall(i int) (string, string, string) {
	fake.createSecretGitMutex.RLock()
	defer fake.createSecretGitMutex.RUnlock()
	argsForCall := fake.createSecretGitArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeFlux) CreateSecretGitReturns(result1 []byte, result2 error) {
	fake.createSecretGitMutex.Lock()
	defer fake.createSecretGitMutex.Unlock()
	fake.CreateSecretGitStub = nil
	fake.createSecretGitReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeFlux) CreateSecretGitReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.createSecretGitMutex.Lock()
	defer fake.createSecretGitMutex.Unlock()
	fake.CreateSecretGitStub = nil
	if fake.createSecretGitReturnsOnCall == nil {
		fake.createSecretGitReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.createSecretGitReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeFlux) CreateSourceGit(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string) ([]byte, error) {
	fake.createSourceGitMutex.Lock()
	ret, specificReturn := fake.createSourceGitReturnsOnCall[len(fake.createSourceGitArgsForCall)]
	fake.createSourceGitArgsForCall = append(fake.createSourceGitArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.CreateSourceGitStub
	fakeReturns := fake.createSourceGitReturns
	fake.recordInvocation("CreateSourceGit", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.createSourceGitMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFlux) CreateSourceGitCallCount() int {
	fake.createSourceGitMutex.RLock()
	defer fake.createSourceGitMutex.RUnlock()
	return len(fake.createSourceGitArgsForCall)
}

func (fake *FakeFlux) CreateSourceGitCalls(stub func(string, string, string, string, string) ([]byte, error)) {
	fake.createSourceGitMutex.Lock()
	defer fake.createSourceGitMutex.Unlock()
	fake.CreateSourceGitStub = stub
}

func (fake *FakeFlux) CreateSourceGitArgsForCall(i int) (string, string, string, string, string) {
	fake.createSourceGitMutex.RLock()
	defer fake.createSourceGitMutex.RUnlock()
	argsForCall := fake.createSourceGitArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeFlux) CreateSourceGitReturns(result1 []byte, result2 error) {
	fake.createSourceGitMutex.Lock()
	defer fake.createSourceGitMutex.Unlock()
	fake.CreateSourceGitStub = nil
	fake.createSourceGitReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeFlux) CreateSourceGitReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.createSourceGitMutex.Lock()
	defer fake.createSourceGitMutex.Unlock()
	fake.CreateSourceGitStub = nil
	if fake.createSourceGitReturnsOnCall == nil {
		fake.createSourceGitReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.createSourceGitReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeFlux) CreateSourceHelm(arg1 string, arg2 string, arg3 string) ([]byte, error) {
	fake.createSourceHelmMutex.Lock()
	ret, specificReturn := fake.createSourceHelmReturnsOnCall[len(fake.createSourceHelmArgsForCall)]
	fake.createSourceHelmArgsForCall = append(fake.createSourceHelmArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.CreateSourceHelmStub
	fakeReturns := fake.createSourceHelmReturns
	fake.recordInvocation("CreateSourceHelm", []interface{}{arg1, arg2, arg3})
	fake.createSourceHelmMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFlux) CreateSourceHelmCallCount() int {
	fake.createSourceHelmMutex.RLock()
	defer fake.createSourceHelmMutex.RUnlock()
	return len(fake.createSourceHelmArgsForCall)
}

func (fake *FakeFlux) CreateSourceHelmCalls(stub func(string, string, string) ([]byte, error)) {
	fake.createSourceHelmMutex.Lock()
	defer fake.createSourceHelmMutex.Unlock()
	fake.CreateSourceHelmStub = stub
}

func (fake *FakeFlux) CreateSourceHelmArgsForCall(i int) (string, string, string) {
	fake.createSourceHelmMutex.RLock()
	defer fake.createSourceHelmMutex.RUnlock()
	argsForCall := fake.createSourceHelmArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeFlux) CreateSourceHelmReturns(result1 []byte, result2 error) {
	fake.createSourceHelmMutex.Lock()
	defer fake.createSourceHelmMutex.Unlock()
	fake.CreateSourceHelmStub = nil
	fake.createSourceHelmReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeFlux) CreateSourceHelmReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.createSourceHelmMutex.Lock()
	defer fake.createSourceHelmMutex.Unlock()
	fake.CreateSourceHelmStub = nil
	if fake.createSourceHelmReturnsOnCall == nil {
		fake.createSourceHelmReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.createSourceHelmReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeFlux) Install(arg1 string, arg2 bool) ([]byte, error) {
	fake.installMutex.Lock()
	ret, specificReturn := fake.installReturnsOnCall[len(fake.installArgsForCall)]
	fake.installArgsForCall = append(fake.installArgsForCall, struct {
		arg1 string
		arg2 bool
	}{arg1, arg2})
	stub := fake.InstallStub
	fakeReturns := fake.installReturns
	fake.recordInvocation("Install", []interface{}{arg1, arg2})
	fake.installMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFlux) InstallCallCount() int {
	fake.installMutex.RLock()
	defer fake.installMutex.RUnlock()
	return len(fake.installArgsForCall)
}

func (fake *FakeFlux) InstallCalls(stub func(string, bool) ([]byte, error)) {
	fake.installMutex.Lock()
	defer fake.installMutex.Unlock()
	fake.InstallStub = stub
}

func (fake *FakeFlux) InstallArgsForCall(i int) (string, bool) {
	fake.installMutex.RLock()
	defer fake.installMutex.RUnlock()
	argsForCall := fake.installArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeFlux) InstallReturns(result1 []byte, result2 error) {
	fake.installMutex.Lock()
	defer fake.installMutex.Unlock()
	fake.InstallStub = nil
	fake.installReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeFlux) InstallReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.installMutex.Lock()
	defer fake.installMutex.Unlock()
	fake.InstallStub = nil
	if fake.installReturnsOnCall == nil {
		fake.installReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.installReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeFlux) Uninstall(arg1 string, arg2 bool) error {
	fake.uninstallMutex.Lock()
	ret, specificReturn := fake.uninstallReturnsOnCall[len(fake.uninstallArgsForCall)]
	fake.uninstallArgsForCall = append(fake.uninstallArgsForCall, struct {
		arg1 string
		arg2 bool
	}{arg1, arg2})
	stub := fake.UninstallStub
	fakeReturns := fake.uninstallReturns
	fake.recordInvocation("Uninstall", []interface{}{arg1, arg2})
	fake.uninstallMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeFlux) UninstallCallCount() int {
	fake.uninstallMutex.RLock()
	defer fake.uninstallMutex.RUnlock()
	return len(fake.uninstallArgsForCall)
}

func (fake *FakeFlux) UninstallCalls(stub func(string, bool) error) {
	fake.uninstallMutex.Lock()
	defer fake.uninstallMutex.Unlock()
	fake.UninstallStub = stub
}

func (fake *FakeFlux) UninstallArgsForCall(i int) (string, bool) {
	fake.uninstallMutex.RLock()
	defer fake.uninstallMutex.RUnlock()
	argsForCall := fake.uninstallArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeFlux) UninstallReturns(result1 error) {
	fake.uninstallMutex.Lock()
	defer fake.uninstallMutex.Unlock()
	fake.UninstallStub = nil
	fake.uninstallReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFlux) UninstallReturnsOnCall(i int, result1 error) {
	fake.uninstallMutex.Lock()
	defer fake.uninstallMutex.Unlock()
	fake.UninstallStub = nil
	if fake.uninstallReturnsOnCall == nil {
		fake.uninstallReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.uninstallReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFlux) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createHelmReleaseGitRepositoryMutex.RLock()
	defer fake.createHelmReleaseGitRepositoryMutex.RUnlock()
	fake.createHelmReleaseHelmRepositoryMutex.RLock()
	defer fake.createHelmReleaseHelmRepositoryMutex.RUnlock()
	fake.createKustomizationMutex.RLock()
	defer fake.createKustomizationMutex.RUnlock()
	fake.createSecretGitMutex.RLock()
	defer fake.createSecretGitMutex.RUnlock()
	fake.createSourceGitMutex.RLock()
	defer fake.createSourceGitMutex.RUnlock()
	fake.createSourceHelmMutex.RLock()
	defer fake.createSourceHelmMutex.RUnlock()
	fake.installMutex.RLock()
	defer fake.installMutex.RUnlock()
	fake.uninstallMutex.RLock()
	defer fake.uninstallMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFlux) recordInvocation(key string, args []interface{}) {
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

var _ flux.Flux = new(FakeFlux)
