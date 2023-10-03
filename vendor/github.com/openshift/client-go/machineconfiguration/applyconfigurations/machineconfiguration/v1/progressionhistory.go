// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/machineconfiguration/v1"
)

// ProgressionHistoryApplyConfiguration represents an declarative configuration of the ProgressionHistory type for use
// with apply.
type ProgressionHistoryApplyConfiguration struct {
	NameAndType *string           `json:"nameAndType,omitempty"`
	State       *v1.StateProgress `json:"state,omitempty"`
	Phase       *string           `json:"phase,omitempty"`
	Reason      *string           `json:"reason,omitempty"`
}

// ProgressionHistoryApplyConfiguration constructs an declarative configuration of the ProgressionHistory type for use with
// apply.
func ProgressionHistory() *ProgressionHistoryApplyConfiguration {
	return &ProgressionHistoryApplyConfiguration{}
}

// WithNameAndType sets the NameAndType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NameAndType field is set to the value of the last call.
func (b *ProgressionHistoryApplyConfiguration) WithNameAndType(value string) *ProgressionHistoryApplyConfiguration {
	b.NameAndType = &value
	return b
}

// WithState sets the State field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the State field is set to the value of the last call.
func (b *ProgressionHistoryApplyConfiguration) WithState(value v1.StateProgress) *ProgressionHistoryApplyConfiguration {
	b.State = &value
	return b
}

// WithPhase sets the Phase field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Phase field is set to the value of the last call.
func (b *ProgressionHistoryApplyConfiguration) WithPhase(value string) *ProgressionHistoryApplyConfiguration {
	b.Phase = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *ProgressionHistoryApplyConfiguration) WithReason(value string) *ProgressionHistoryApplyConfiguration {
	b.Reason = &value
	return b
}
