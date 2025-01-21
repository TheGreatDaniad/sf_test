// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package api

import (
	"context"
	"sf_test/internal/models"
	"sf_test/internal/core"

	"sync"
)

// Ensure, that StepServiceMock does implement StepService.
// If this is not the case, regenerate this file with moq.
var _ core.StepService = &StepServiceMock{}

// StepServiceMock is a mock implementation of StepService.
//
//	func TestSomethingThatUsesStepService(t *testing.T) {
//
//		// make and configure a mocked StepService
//		mockedStepService := &StepServiceMock{
//			CreateStepFunc: func(ctx context.Context, step *models.Step) (int64, error) {
//				panic("mock out the CreateStep method")
//			},
//			DeleteStepFunc: func(ctx context.Context, id int64) error {
//				panic("mock out the DeleteStep method")
//			},
//			ListStepsFunc: func(ctx context.Context, sequenceID int64) ([]*models.Step, error) {
//				panic("mock out the ListSteps method")
//			},
//			UpdateStepFunc: func(ctx context.Context, step *models.Step) error {
//				panic("mock out the UpdateStep method")
//			},
//		}
//
//		// use mockedStepService in code that requires StepService
//		// and then make assertions.
//
//	}
type StepServiceMock struct {
	// CreateStepFunc mocks the CreateStep method.
	CreateStepFunc func(ctx context.Context, step *models.Step) (int64, error)

	// DeleteStepFunc mocks the DeleteStep method.
	DeleteStepFunc func(ctx context.Context, id int64) error

	// ListStepsFunc mocks the ListSteps method.
	ListStepsFunc func(ctx context.Context, sequenceID int64) ([]*models.Step, error)

	// UpdateStepFunc mocks the UpdateStep method.
	UpdateStepFunc func(ctx context.Context, step *models.Step) error

	// calls tracks calls to the methods.
	calls struct {
		// CreateStep holds details about calls to the CreateStep method.
		CreateStep []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Step is the step argument value.
			Step *models.Step
		}
		// DeleteStep holds details about calls to the DeleteStep method.
		DeleteStep []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID int64
		}
		// ListSteps holds details about calls to the ListSteps method.
		ListSteps []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// SequenceID is the sequenceID argument value.
			SequenceID int64
		}
		// UpdateStep holds details about calls to the UpdateStep method.
		UpdateStep []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Step is the step argument value.
			Step *models.Step
		}
	}
	lockCreateStep sync.RWMutex
	lockDeleteStep sync.RWMutex
	lockListSteps  sync.RWMutex
	lockUpdateStep sync.RWMutex
}

// CreateStep calls CreateStepFunc.
func (mock *StepServiceMock) CreateStep(ctx context.Context, step *models.Step) (int64, error) {
	if mock.CreateStepFunc == nil {
		panic("StepServiceMock.CreateStepFunc: method is nil but StepService.CreateStep was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Step *models.Step
	}{
		Ctx:  ctx,
		Step: step,
	}
	mock.lockCreateStep.Lock()
	mock.calls.CreateStep = append(mock.calls.CreateStep, callInfo)
	mock.lockCreateStep.Unlock()
	return mock.CreateStepFunc(ctx, step)
}

// CreateStepCalls gets all the calls that were made to CreateStep.
// Check the length with:
//
//	len(mockedStepService.CreateStepCalls())
func (mock *StepServiceMock) CreateStepCalls() []struct {
	Ctx  context.Context
	Step *models.Step
} {
	var calls []struct {
		Ctx  context.Context
		Step *models.Step
	}
	mock.lockCreateStep.RLock()
	calls = mock.calls.CreateStep
	mock.lockCreateStep.RUnlock()
	return calls
}

// DeleteStep calls DeleteStepFunc.
func (mock *StepServiceMock) DeleteStep(ctx context.Context, id int64) error {
	if mock.DeleteStepFunc == nil {
		panic("StepServiceMock.DeleteStepFunc: method is nil but StepService.DeleteStep was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  int64
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDeleteStep.Lock()
	mock.calls.DeleteStep = append(mock.calls.DeleteStep, callInfo)
	mock.lockDeleteStep.Unlock()
	return mock.DeleteStepFunc(ctx, id)
}

// DeleteStepCalls gets all the calls that were made to DeleteStep.
// Check the length with:
//
//	len(mockedStepService.DeleteStepCalls())
func (mock *StepServiceMock) DeleteStepCalls() []struct {
	Ctx context.Context
	ID  int64
} {
	var calls []struct {
		Ctx context.Context
		ID  int64
	}
	mock.lockDeleteStep.RLock()
	calls = mock.calls.DeleteStep
	mock.lockDeleteStep.RUnlock()
	return calls
}

// ListSteps calls ListStepsFunc.
func (mock *StepServiceMock) ListSteps(ctx context.Context, sequenceID int64) ([]*models.Step, error) {
	if mock.ListStepsFunc == nil {
		panic("StepServiceMock.ListStepsFunc: method is nil but StepService.ListSteps was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		SequenceID int64
	}{
		Ctx:        ctx,
		SequenceID: sequenceID,
	}
	mock.lockListSteps.Lock()
	mock.calls.ListSteps = append(mock.calls.ListSteps, callInfo)
	mock.lockListSteps.Unlock()
	return mock.ListStepsFunc(ctx, sequenceID)
}

// ListStepsCalls gets all the calls that were made to ListSteps.
// Check the length with:
//
//	len(mockedStepService.ListStepsCalls())
func (mock *StepServiceMock) ListStepsCalls() []struct {
	Ctx        context.Context
	SequenceID int64
} {
	var calls []struct {
		Ctx        context.Context
		SequenceID int64
	}
	mock.lockListSteps.RLock()
	calls = mock.calls.ListSteps
	mock.lockListSteps.RUnlock()
	return calls
}

// UpdateStep calls UpdateStepFunc.
func (mock *StepServiceMock) UpdateStep(ctx context.Context, step *models.Step) error {
	if mock.UpdateStepFunc == nil {
		panic("StepServiceMock.UpdateStepFunc: method is nil but StepService.UpdateStep was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Step *models.Step
	}{
		Ctx:  ctx,
		Step: step,
	}
	mock.lockUpdateStep.Lock()
	mock.calls.UpdateStep = append(mock.calls.UpdateStep, callInfo)
	mock.lockUpdateStep.Unlock()
	return mock.UpdateStepFunc(ctx, step)
}

// UpdateStepCalls gets all the calls that were made to UpdateStep.
// Check the length with:
//
//	len(mockedStepService.UpdateStepCalls())
func (mock *StepServiceMock) UpdateStepCalls() []struct {
	Ctx  context.Context
	Step *models.Step
} {
	var calls []struct {
		Ctx  context.Context
		Step *models.Step
	}
	mock.lockUpdateStep.RLock()
	calls = mock.calls.UpdateStep
	mock.lockUpdateStep.RUnlock()
	return calls
}
