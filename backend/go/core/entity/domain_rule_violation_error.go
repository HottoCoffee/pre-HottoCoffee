package entity

import "errors"

type DomainRuleViolationError error

func NewDomainRuleViolationError(text string) DomainRuleViolationError {
	return &errorString{text}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func IsDomainRuleViolationError(err error) bool {
	tmp := NewDomainRuleViolationError("")
	return errors.As(err, &tmp)
}
