package model

import (
	"testing"
)

func TestInvitationCode(t *testing.T) {
	t.Run("generate given length int array", func(t *testing.T) {
		for length := 0; length < 100; length++ {
			codes := generateInvitationCode(length)
			if len(codes) != length {
				t.Errorf("Expected %d, Actual %d", length, len(codes))
			}
		}
	})

	t.Run("all number in code are 0 ~ 9", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			codes := generateInvitationCode(10)
			for _, c := range codes {
				if c < 0 || c > 9 {
					t.Errorf("Expected 0 ~ 9, Actual %d", c)
				}
			}
		}
	})
}
