package greetings

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Hello(t *testing.T) {
	var original_randomFormat_FN = randomFormat_FN

	defer func() {
		randomFormat_FN = original_randomFormat_FN
	}()

	testCase := []struct {
		name        string
		namePeople  string
		expectError string
		mock_fn     func()
	}{
		{
			name:        "Error name empty",
			namePeople:  "",
			expectError: errors.New("el campo name se encuentra vacio").Error(),
		},

		{
			name:        " hello name ok",
			namePeople:  "Lucho",
			expectError: "",
			mock_fn: func() {
				randomFormat_FN = func() string {
					return "Test Saludos %v"
				}
			},
		},
	}

	for _, tc := range testCase {
		if tc.mock_fn != nil {
			tc.mock_fn()
		}

		_, err := Hello(tc.name)
		if err != nil {
			assert.Equal(t, err.Error(), tc.expectError)
		}

	}
}
