package entities

import "errors"

var (
	ErrInvalidLenghtMsisdnInput = errors.New("invalid input. Length of an MSISDN should be from 8 to 15 digits")
	ErrInvalidDialingCode       = errors.New("invalid dialing code")
	ErrInvalidLengthNumber      = errors.New("invalid length number")
	ErrInvalidMnoCode           = errors.New("invalid mobile network operator code")
	ErrInputHasLetters          = errors.New("invalid MSISDN. Input should contain only digits")
)
