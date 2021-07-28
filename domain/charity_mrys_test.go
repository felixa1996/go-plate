package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharityMrys(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		charityMrys CharityMrys
		args        CharityMrys
		expected    CharityMrys
		expectedErr error
	}{
		{
			name:        "Success creating charity mrys",
			charityMrys: CharityMrys{},
			args: CharityMrys{
				Name: "Test",
			},
			expected: CharityMrys{
				Name: "Test",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.charityMrys.Name = tt.args.Name

			assert.True(t, tt.charityMrys == tt.expected, "[TestCase '%s'] Result: '%v' | Expected: '%v'",
				tt.name,
				tt.charityMrys,
				tt.expected)
		})
	}

}
