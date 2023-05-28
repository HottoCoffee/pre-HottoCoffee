package entity

type DomainRuleViolationError struct {
	s string
}

func NewDomainRuleViolationError(text string) *DomainRuleViolationError {
	return &DomainRuleViolationError{text}
}

func (e *DomainRuleViolationError) Error() string {
	return e.s
}

func IsDomainRuleViolationError(err error) bool {
	_, ok := err.(*DomainRuleViolationError)
	return ok
}
