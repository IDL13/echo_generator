package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewParser(t *testing.T) {
	type TestOptions struct {
		flags0 string `name:"help" value:"false" usage:"help"`
		flags1 string `name:"new" value:"echo" usage:"new"`
		flags2 string `name:"git" value:"false" usage:"git"`
	}

	tests := []struct {
		name           string
		input          TestOptions
		expectedBool   []bool
		expectedString []string
	}{
		// Testcase 1
		{
			name:           "Ok",
			input:          TestOptions{},
			expectedBool:   []bool{false, false},
			expectedString: []string{"echo"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			flagsB, flagS := NewParser(test.input)
			for i := range flagsB {
				assert.Equal(t, test.expectedBool[i], *flagsB[i])
			}
			for j := range flagS {
				assert.Equal(t, test.expectedString[j], *flagS[j])
			}
		})
	}
}
