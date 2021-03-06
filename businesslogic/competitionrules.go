package businesslogic

import "errors"

type IRule interface {
	Apply(registration EventRegistrationForm) error
}

type GenderRule struct {
	AllowSameSex bool
	IAccountRepository
	IPartnershipRepository
}

func (rule GenderRule) Apply(registration EventRegistrationForm) error {
	if registration.Couple.SameSex && (!rule.AllowSameSex) {
		return errors.New("same sex is not allowed")
	}
	return nil
}
