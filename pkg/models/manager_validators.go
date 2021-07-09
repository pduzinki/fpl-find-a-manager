package models

import "errors"

var ErrFullnameTooShort = errors.New("model: name is too short")

type managerValidatorFunc func(*Manager) error

func runManagerValidatorFuncs(manager *Manager, fns ...managerValidatorFunc) error {
	for _, f := range fns {
		err := f(manager)
		if err != nil {
			return err
		}
	}
	return nil
}

func (mv *managerValidator) FullNameLongerThanThreeRunes(manager *Manager) error {
	if len([]rune(manager.FullName)) <= 3 {
		return ErrFullnameTooShort
	}
	return nil
}
