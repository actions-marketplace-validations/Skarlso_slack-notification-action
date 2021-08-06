// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/Skarlso/slack-notification-action/pkg"
	"github.com/slack-go/slack"
)

type FakeSlackClient struct {
	PostMessageStub        func(string, ...slack.MsgOption) (string, string, error)
	postMessageMutex       sync.RWMutex
	postMessageArgsForCall []struct {
		arg1 string
		arg2 []slack.MsgOption
	}
	postMessageReturns struct {
		result1 string
		result2 string
		result3 error
	}
	postMessageReturnsOnCall map[int]struct {
		result1 string
		result2 string
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSlackClient) PostMessage(arg1 string, arg2 ...slack.MsgOption) (string, string, error) {
	fake.postMessageMutex.Lock()
	ret, specificReturn := fake.postMessageReturnsOnCall[len(fake.postMessageArgsForCall)]
	fake.postMessageArgsForCall = append(fake.postMessageArgsForCall, struct {
		arg1 string
		arg2 []slack.MsgOption
	}{arg1, arg2})
	stub := fake.PostMessageStub
	fakeReturns := fake.postMessageReturns
	fake.recordInvocation("PostMessage", []interface{}{arg1, arg2})
	fake.postMessageMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSlackClient) PostMessageCallCount() int {
	fake.postMessageMutex.RLock()
	defer fake.postMessageMutex.RUnlock()
	return len(fake.postMessageArgsForCall)
}

func (fake *FakeSlackClient) PostMessageCalls(stub func(string, ...slack.MsgOption) (string, string, error)) {
	fake.postMessageMutex.Lock()
	defer fake.postMessageMutex.Unlock()
	fake.PostMessageStub = stub
}

func (fake *FakeSlackClient) PostMessageArgsForCall(i int) (string, []slack.MsgOption) {
	fake.postMessageMutex.RLock()
	defer fake.postMessageMutex.RUnlock()
	argsForCall := fake.postMessageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSlackClient) PostMessageReturns(result1 string, result2 string, result3 error) {
	fake.postMessageMutex.Lock()
	defer fake.postMessageMutex.Unlock()
	fake.PostMessageStub = nil
	fake.postMessageReturns = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSlackClient) PostMessageReturnsOnCall(i int, result1 string, result2 string, result3 error) {
	fake.postMessageMutex.Lock()
	defer fake.postMessageMutex.Unlock()
	fake.PostMessageStub = nil
	if fake.postMessageReturnsOnCall == nil {
		fake.postMessageReturnsOnCall = make(map[int]struct {
			result1 string
			result2 string
			result3 error
		})
	}
	fake.postMessageReturnsOnCall[i] = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSlackClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.postMessageMutex.RLock()
	defer fake.postMessageMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSlackClient) recordInvocation(key string, args []interface{}) {
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

var _ pkg.SlackClient = new(FakeSlackClient)