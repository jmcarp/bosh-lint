// This file was generated by counterfeiter
package releasedirfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/releasedir"
)

type FakeGitRepo struct {
	InitStub        func() error
	initMutex       sync.RWMutex
	initArgsForCall []struct{}
	initReturns     struct {
		result1 error
	}
	LastCommitSHAStub        func() (string, error)
	lastCommitSHAMutex       sync.RWMutex
	lastCommitSHAArgsForCall []struct{}
	lastCommitSHAReturns     struct {
		result1 string
		result2 error
	}
	MustNotBeDirtyStub        func(force bool) (dirty bool, err error)
	mustNotBeDirtyMutex       sync.RWMutex
	mustNotBeDirtyArgsForCall []struct {
		force bool
	}
	mustNotBeDirtyReturns struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGitRepo) Init() error {
	fake.initMutex.Lock()
	fake.initArgsForCall = append(fake.initArgsForCall, struct{}{})
	fake.recordInvocation("Init", []interface{}{})
	fake.initMutex.Unlock()
	if fake.InitStub != nil {
		return fake.InitStub()
	}
	return fake.initReturns.result1
}

func (fake *FakeGitRepo) InitCallCount() int {
	fake.initMutex.RLock()
	defer fake.initMutex.RUnlock()
	return len(fake.initArgsForCall)
}

func (fake *FakeGitRepo) InitReturns(result1 error) {
	fake.InitStub = nil
	fake.initReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGitRepo) LastCommitSHA() (string, error) {
	fake.lastCommitSHAMutex.Lock()
	fake.lastCommitSHAArgsForCall = append(fake.lastCommitSHAArgsForCall, struct{}{})
	fake.recordInvocation("LastCommitSHA", []interface{}{})
	fake.lastCommitSHAMutex.Unlock()
	if fake.LastCommitSHAStub != nil {
		return fake.LastCommitSHAStub()
	}
	return fake.lastCommitSHAReturns.result1, fake.lastCommitSHAReturns.result2
}

func (fake *FakeGitRepo) LastCommitSHACallCount() int {
	fake.lastCommitSHAMutex.RLock()
	defer fake.lastCommitSHAMutex.RUnlock()
	return len(fake.lastCommitSHAArgsForCall)
}

func (fake *FakeGitRepo) LastCommitSHAReturns(result1 string, result2 error) {
	fake.LastCommitSHAStub = nil
	fake.lastCommitSHAReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeGitRepo) MustNotBeDirty(force bool) (dirty bool, err error) {
	fake.mustNotBeDirtyMutex.Lock()
	fake.mustNotBeDirtyArgsForCall = append(fake.mustNotBeDirtyArgsForCall, struct {
		force bool
	}{force})
	fake.recordInvocation("MustNotBeDirty", []interface{}{force})
	fake.mustNotBeDirtyMutex.Unlock()
	if fake.MustNotBeDirtyStub != nil {
		return fake.MustNotBeDirtyStub(force)
	}
	return fake.mustNotBeDirtyReturns.result1, fake.mustNotBeDirtyReturns.result2
}

func (fake *FakeGitRepo) MustNotBeDirtyCallCount() int {
	fake.mustNotBeDirtyMutex.RLock()
	defer fake.mustNotBeDirtyMutex.RUnlock()
	return len(fake.mustNotBeDirtyArgsForCall)
}

func (fake *FakeGitRepo) MustNotBeDirtyArgsForCall(i int) bool {
	fake.mustNotBeDirtyMutex.RLock()
	defer fake.mustNotBeDirtyMutex.RUnlock()
	return fake.mustNotBeDirtyArgsForCall[i].force
}

func (fake *FakeGitRepo) MustNotBeDirtyReturns(result1 bool, result2 error) {
	fake.MustNotBeDirtyStub = nil
	fake.mustNotBeDirtyReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeGitRepo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.initMutex.RLock()
	defer fake.initMutex.RUnlock()
	fake.lastCommitSHAMutex.RLock()
	defer fake.lastCommitSHAMutex.RUnlock()
	fake.mustNotBeDirtyMutex.RLock()
	defer fake.mustNotBeDirtyMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeGitRepo) recordInvocation(key string, args []interface{}) {
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

var _ releasedir.GitRepo = new(FakeGitRepo)
