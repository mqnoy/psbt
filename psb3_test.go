package main

import "testing"

func TestPsb3(t *testing.T) {
	testCases := []struct {
		name     string
		param    string
		expected string
	}{
		{
			name:     "example1",
			param:    "07:05:45PM",
			expected: "19:05:45",
		},
		{
			name:     "example2",
			param:    "01:00:45AM",
			expected: "01:00:45",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ConvertTo24HourFormat(tc.param)
			if err != nil {
				t.Error(err)
			}

			if result != tc.expected {
				t.Errorf("Expected: %s, Got: %s", tc.expected, result)
			}
		})
	}
}
