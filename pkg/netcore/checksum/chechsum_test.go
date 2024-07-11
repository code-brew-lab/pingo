package checksum_test

import (
	"encoding/hex"
	"testing"

	"github.com/code-brew-lab/pingo/pkg/netcore/checksum"
)

type (
	CalculateTestCase struct {
		input    []byte
		expected uint16
	}

	VerifyTestCase struct {
		input   []byte
		isValid bool
	}
)

var (
	calculateCases = []CalculateTestCase{
		{
			input:    []byte{0x45, 0x00, 0x00, 0x28, 0x04, 0x59, 0x00, 0x00, 0x40, 0x06, 0x00, 0x00, 0xc0, 0xa8, 0x01, 0x6d, 0x0d, 0x45, 0x44, 0x40},
			expected: 25309,
		},
		{
			input:    []byte{0x45, 0x00, 0x05, 0x58, 0xe8, 0x82, 0x00, 0x00, 0x40, 0x01, 0x00, 0x00, 0xc0, 0xa8, 0x01, 0x6d, 0xc0, 0xa8, 0x01, 0x01},
			expected: 2404,
		},
		{
			input:    []byte{0x45, 0x00, 0x00, 0x54, 0xca, 0xf3, 0x00, 0x00, 0x38, 0x01, 0x08, 0xab, 0x68, 0x10, 0x00, 0x00, 0xc0, 0xa8, 0x01, 0x6d},
			expected: 34021,
		},
	}

	verifyCases = []VerifyTestCase{
		{
			input:   []byte{0x45, 0x00, 0x00, 0x28, 0x04, 0x59, 0x00, 0x00, 0x40, 0x06, 0x62, 0xdd, 0xc0, 0xa8, 0x01, 0x6d, 0x0d, 0x45, 0x44, 0x40},
			isValid: true,
		},
		{
			input:   []byte{0x45, 0x00, 0x00, 0x54, 0x28, 0x51, 0x00, 0x00, 0x40, 0x01, 0x90, 0xdb, 0xc0, 0xa8, 0x01, 0x6d, 0xd8, 0xef, 0x26, 0x78},
			isValid: true,
		},
		{
			input:   []byte{0x45, 0x00, 0x05, 0x58, 0x4e, 0xa0, 0x00, 0x00, 0x40, 0x01, 0x43, 0x2a, 0xc0, 0xa8, 0x01, 0x6d, 0x11, 0xfd, 0x0f, 0xc9},
			isValid: true,
		},
		{
			input:   []byte{0x45, 0x00, 0x05, 0x58, 0xe8, 0x82, 0x00, 0x00, 0x40, 0x01, 0x09, 0x64, 0xc0, 0xa8, 0x01, 0x6d, 0xc0, 0xa8, 0x01, 0x01},
			isValid: true,
		},
		{
			input:   []byte{0x45, 0x00, 0x00, 0x54, 0xca, 0xf3, 0x00, 0x00, 0x38, 0x01, 0x08, 0xab, 0x68, 0x10, 0x84, 0xe5, 0xc0, 0xa8, 0x01, 0x6d},
			isValid: true,
		},
		{
			input:   []byte{0x45, 0xdf, 0x05, 0x58, 0x4e, 0xa0, 0x00, 0x00, 0x40, 0x01, 0x43, 0x2a, 0xc0, 0xa8, 0x01, 0x6d, 0x11, 0xfd, 0x0f, 0xc9},
			isValid: false,
		},
		{
			input:   []byte{0x45, 0x00, 0x00, 0x28, 0x04, 0x59, 0x00, 0x00, 0x40, 0x06, 0x60, 0xdd, 0xc0, 0xa8, 0x01, 0x6d, 0x0d, 0x45, 0x44, 0x40},
			isValid: false,
		},
	}
)

func TestCalculate(t *testing.T) {
	for _, cc := range calculateCases {
		calculated := checksum.Calculate(cc.input)
		if calculated != cc.expected {
			t.Errorf("calculated: %d, expected: %d", calculated, cc.expected)
			t.Fail()
			continue
		}
	}
}

func TestVerify(t *testing.T) {
	for _, vc := range verifyCases {
		ok := checksum.Verify(vc.input)
		if ok != vc.isValid {
			t.Errorf("raw:%s, calculated: %v, expected: %v", hex.EncodeToString(vc.input), ok, vc.isValid)
			t.Fail()
			continue
		}
	}
}
