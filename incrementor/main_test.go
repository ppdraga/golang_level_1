package incrementor

import (
	"testing"
)

func TestIncrementor_SetMaximumValue(t *testing.T) {
	inc := NewIncrementor()
	err := inc.SetMaximumValue(-5)
	if err.Error() != "Нельзя установить число меньше нуля." {
		t.Error("Fail: expected error Нельзя установить число меньше нуля.")
	}
}

func TestIncrementor_IncrementNumber(t *testing.T) {
	inc := NewIncrementor()
	inc.SetMaximumValue(5)
	inc.IncrementNumber()
	got := inc.GetNumber()
	if got != 1 {
		t.Errorf("Fail: expected 1, got %d", got)
	}

	for i := 0; i < 7; i++ {
		inc.IncrementNumber()
	}

	got = inc.GetNumber()
	if got != 2 {
		t.Errorf("Fail: expected 2, got %d", got)
	}
}
