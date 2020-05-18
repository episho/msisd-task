package entities

import (
	"testing"
)

func TestCleanMsisdnInput(t *testing.T) {
	testCases := []struct {
		input         string
		expCleanInput string
		expError      error
	}{
		{
			"+389-71-833-789",
			"38971833789",
			nil,
		},
		{
			"00389 71 833 789",
			"38971833789",
			nil,
		},
		{
			"+389-71-833-789 test",
			"",
			ErrInputHasLetters,
		},
		{
			"00389 77 833 789 65486 68597967986 578974675",
			"",
			ErrInvalidLenghtMsisdnInput,
		},
		{
			"0 (389) 72 111 222",
			"38972111222",
			nil,
		},
	}


	for _, testCase := range testCases {
		res, err := CleanMsisdnInput(testCase.input)
		if err != testCase.expError {
			t.Errorf("CleanMsisdnInput: expected err: %v, got %v", testCase.expError, err)
		}

		if res != testCase.expCleanInput {
			t.Errorf("CleanMsisdnInput: expected cleanedInput: %s, got %s", testCase.expCleanInput, res)
		}
	}
}
