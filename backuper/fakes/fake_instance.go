// This file was generated by counterfeiter
package fakes

import (
	"io"
	"sync"

	"github.com/pivotal-cf/pcf-backup-and-restore/backuper"
)

type FakeInstance struct {
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	IDStub        func() string
	iDMutex       sync.RWMutex
	iDArgsForCall []struct{}
	iDReturns     struct {
		result1 string
	}
	IsBackupableStub        func() (bool, error)
	isBackupableMutex       sync.RWMutex
	isBackupableArgsForCall []struct{}
	isBackupableReturns     struct {
		result1 bool
		result2 error
	}
	IsRestorableStub        func() (bool, error)
	isRestorableMutex       sync.RWMutex
	isRestorableArgsForCall []struct{}
	isRestorableReturns     struct {
		result1 bool
		result2 error
	}
	BackupStub        func() error
	backupMutex       sync.RWMutex
	backupArgsForCall []struct{}
	backupReturns     struct {
		result1 error
	}
	RestoreStub        func() error
	restoreMutex       sync.RWMutex
	restoreArgsForCall []struct{}
	restoreReturns     struct {
		result1 error
	}
	CleanupStub        func() error
	cleanupMutex       sync.RWMutex
	cleanupArgsForCall []struct{}
	cleanupReturns     struct {
		result1 error
	}
	StreamBackupToStub        func(io.Writer) error
	streamBackupToMutex       sync.RWMutex
	streamBackupToArgsForCall []struct {
		arg1 io.Writer
	}
	streamBackupToReturns struct {
		result1 error
	}
	BackupChecksumStub        func() (string, error)
	backupChecksumMutex       sync.RWMutex
	backupChecksumArgsForCall []struct{}
	backupChecksumReturns     struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInstance) Name() string {
	fake.nameMutex.Lock()
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	} else {
		return fake.nameReturns.result1
	}
}

func (fake *FakeInstance) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeInstance) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInstance) ID() string {
	fake.iDMutex.Lock()
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct{}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	} else {
		return fake.iDReturns.result1
	}
}

func (fake *FakeInstance) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeInstance) IDReturns(result1 string) {
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInstance) IsBackupable() (bool, error) {
	fake.isBackupableMutex.Lock()
	fake.isBackupableArgsForCall = append(fake.isBackupableArgsForCall, struct{}{})
	fake.recordInvocation("IsBackupable", []interface{}{})
	fake.isBackupableMutex.Unlock()
	if fake.IsBackupableStub != nil {
		return fake.IsBackupableStub()
	} else {
		return fake.isBackupableReturns.result1, fake.isBackupableReturns.result2
	}
}

func (fake *FakeInstance) IsBackupableCallCount() int {
	fake.isBackupableMutex.RLock()
	defer fake.isBackupableMutex.RUnlock()
	return len(fake.isBackupableArgsForCall)
}

func (fake *FakeInstance) IsBackupableReturns(result1 bool, result2 error) {
	fake.IsBackupableStub = nil
	fake.isBackupableReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeInstance) IsRestorable() (bool, error) {
	fake.isRestorableMutex.Lock()
	fake.isRestorableArgsForCall = append(fake.isRestorableArgsForCall, struct{}{})
	fake.recordInvocation("IsRestorable", []interface{}{})
	fake.isRestorableMutex.Unlock()
	if fake.IsRestorableStub != nil {
		return fake.IsRestorableStub()
	} else {
		return fake.isRestorableReturns.result1, fake.isRestorableReturns.result2
	}
}

func (fake *FakeInstance) IsRestorableCallCount() int {
	fake.isRestorableMutex.RLock()
	defer fake.isRestorableMutex.RUnlock()
	return len(fake.isRestorableArgsForCall)
}

func (fake *FakeInstance) IsRestorableReturns(result1 bool, result2 error) {
	fake.IsRestorableStub = nil
	fake.isRestorableReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeInstance) Backup() error {
	fake.backupMutex.Lock()
	fake.backupArgsForCall = append(fake.backupArgsForCall, struct{}{})
	fake.recordInvocation("Backup", []interface{}{})
	fake.backupMutex.Unlock()
	if fake.BackupStub != nil {
		return fake.BackupStub()
	} else {
		return fake.backupReturns.result1
	}
}

func (fake *FakeInstance) BackupCallCount() int {
	fake.backupMutex.RLock()
	defer fake.backupMutex.RUnlock()
	return len(fake.backupArgsForCall)
}

func (fake *FakeInstance) BackupReturns(result1 error) {
	fake.BackupStub = nil
	fake.backupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) Restore() error {
	fake.restoreMutex.Lock()
	fake.restoreArgsForCall = append(fake.restoreArgsForCall, struct{}{})
	fake.recordInvocation("Restore", []interface{}{})
	fake.restoreMutex.Unlock()
	if fake.RestoreStub != nil {
		return fake.RestoreStub()
	} else {
		return fake.restoreReturns.result1
	}
}

func (fake *FakeInstance) RestoreCallCount() int {
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	return len(fake.restoreArgsForCall)
}

func (fake *FakeInstance) RestoreReturns(result1 error) {
	fake.RestoreStub = nil
	fake.restoreReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) Cleanup() error {
	fake.cleanupMutex.Lock()
	fake.cleanupArgsForCall = append(fake.cleanupArgsForCall, struct{}{})
	fake.recordInvocation("Cleanup", []interface{}{})
	fake.cleanupMutex.Unlock()
	if fake.CleanupStub != nil {
		return fake.CleanupStub()
	} else {
		return fake.cleanupReturns.result1
	}
}

func (fake *FakeInstance) CleanupCallCount() int {
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	return len(fake.cleanupArgsForCall)
}

func (fake *FakeInstance) CleanupReturns(result1 error) {
	fake.CleanupStub = nil
	fake.cleanupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) StreamBackupTo(arg1 io.Writer) error {
	fake.streamBackupToMutex.Lock()
	fake.streamBackupToArgsForCall = append(fake.streamBackupToArgsForCall, struct {
		arg1 io.Writer
	}{arg1})
	fake.recordInvocation("StreamBackupTo", []interface{}{arg1})
	fake.streamBackupToMutex.Unlock()
	if fake.StreamBackupToStub != nil {
		return fake.StreamBackupToStub(arg1)
	} else {
		return fake.streamBackupToReturns.result1
	}
}

func (fake *FakeInstance) StreamBackupToCallCount() int {
	fake.streamBackupToMutex.RLock()
	defer fake.streamBackupToMutex.RUnlock()
	return len(fake.streamBackupToArgsForCall)
}

func (fake *FakeInstance) StreamBackupToArgsForCall(i int) io.Writer {
	fake.streamBackupToMutex.RLock()
	defer fake.streamBackupToMutex.RUnlock()
	return fake.streamBackupToArgsForCall[i].arg1
}

func (fake *FakeInstance) StreamBackupToReturns(result1 error) {
	fake.StreamBackupToStub = nil
	fake.streamBackupToReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) BackupChecksum() (string, error) {
	fake.backupChecksumMutex.Lock()
	fake.backupChecksumArgsForCall = append(fake.backupChecksumArgsForCall, struct{}{})
	fake.recordInvocation("BackupChecksum", []interface{}{})
	fake.backupChecksumMutex.Unlock()
	if fake.BackupChecksumStub != nil {
		return fake.BackupChecksumStub()
	} else {
		return fake.backupChecksumReturns.result1, fake.backupChecksumReturns.result2
	}
}

func (fake *FakeInstance) BackupChecksumCallCount() int {
	fake.backupChecksumMutex.RLock()
	defer fake.backupChecksumMutex.RUnlock()
	return len(fake.backupChecksumArgsForCall)
}

func (fake *FakeInstance) BackupChecksumReturns(result1 string, result2 error) {
	fake.BackupChecksumStub = nil
	fake.backupChecksumReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeInstance) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.isBackupableMutex.RLock()
	defer fake.isBackupableMutex.RUnlock()
	fake.isRestorableMutex.RLock()
	defer fake.isRestorableMutex.RUnlock()
	fake.backupMutex.RLock()
	defer fake.backupMutex.RUnlock()
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	fake.streamBackupToMutex.RLock()
	defer fake.streamBackupToMutex.RUnlock()
	fake.backupChecksumMutex.RLock()
	defer fake.backupChecksumMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeInstance) recordInvocation(key string, args []interface{}) {
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

var _ backuper.Instance = new(FakeInstance)
