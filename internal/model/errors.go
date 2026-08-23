package model

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidID       = errors.New("graindry: invalid identifier")
	ErrNotFound        = errors.New("graindry: entity not found")
	ErrConflict        = errors.New("graindry: state conflict")
	ErrInterlock       = errors.New("graindry: interlock denied")
	ErrMoistureHold    = errors.New("graindry: moisture hold active")
	ErrAirflowSetpoint = errors.New("graindry: airflow setpoint violation")
	ErrFanFault        = errors.New("graindry: fan fault")
	ErrScheduleEmpty   = errors.New("graindry: schedule empty")
	ErrGradient        = errors.New("graindry: moisture gradient violation")
	ErrContextCanceled = errors.New("graindry: operation canceled")
)

type DomainError struct {
	Op   string
	Code string
	Err  error
}

func (e *DomainError) Error() string {
	if e == nil {
		return "<nil>"
	}
	if e.Err != nil {
		return fmt.Sprintf("graindry %s [%s]: %v", e.Op, e.Code, e.Err)
	}
	return fmt.Sprintf("graindry %s [%s]", e.Op, e.Code)
}

func (e *DomainError) Unwrap() error { return e.Err }

func Wrap(op, code string, err error) error {
	if err == nil {
		return nil
	}
	return &DomainError{Op: op, Code: code, Err: err}
}

func Is(err, target error) bool   { return errors.Is(err, target) }
func As(err error, target any) bool { return errors.As(err, target) }
