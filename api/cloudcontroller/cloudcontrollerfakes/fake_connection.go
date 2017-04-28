// This file was generated by counterfeiter
package cloudcontrollerfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/api/cloudcontroller"
)

type FakeConnection struct {
	MakeStub        func(request *cloudcontroller.Request, passedResponse *cloudcontroller.Response) error
	makeMutex       sync.RWMutex
	makeArgsForCall []struct {
		request        *cloudcontroller.Request
		passedResponse *cloudcontroller.Response
	}
	makeReturns struct {
		result1 error
	}
	makeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConnection) Make(request *cloudcontroller.Request, passedResponse *cloudcontroller.Response) error {
	fake.makeMutex.Lock()
	ret, specificReturn := fake.makeReturnsOnCall[len(fake.makeArgsForCall)]
	fake.makeArgsForCall = append(fake.makeArgsForCall, struct {
		request        *cloudcontroller.Request
		passedResponse *cloudcontroller.Response
	}{request, passedResponse})
	fake.recordInvocation("Make", []interface{}{request, passedResponse})
	fake.makeMutex.Unlock()
	if fake.MakeStub != nil {
		return fake.MakeStub(request, passedResponse)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.makeReturns.result1
}

func (fake *FakeConnection) MakeCallCount() int {
	fake.makeMutex.RLock()
	defer fake.makeMutex.RUnlock()
	return len(fake.makeArgsForCall)
}

func (fake *FakeConnection) MakeArgsForCall(i int) (*cloudcontroller.Request, *cloudcontroller.Response) {
	fake.makeMutex.RLock()
	defer fake.makeMutex.RUnlock()
	return fake.makeArgsForCall[i].request, fake.makeArgsForCall[i].passedResponse
}

func (fake *FakeConnection) MakeReturns(result1 error) {
	fake.MakeStub = nil
	fake.makeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConnection) MakeReturnsOnCall(i int, result1 error) {
	fake.MakeStub = nil
	if fake.makeReturnsOnCall == nil {
		fake.makeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.makeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConnection) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.makeMutex.RLock()
	defer fake.makeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeConnection) recordInvocation(key string, args []interface{}) {
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

var _ cloudcontroller.Connection = new(FakeConnection)
