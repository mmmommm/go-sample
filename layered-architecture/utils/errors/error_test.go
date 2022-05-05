package errors

import (
	"errors"
	"testing"

	e "github.com/pkg/errors"
)

func TestErr(t *testing.T) {
	err1 := NewAppError(1, 0, "")
	var AppErr *AppError
	if !errors.As(err1, &appErr) {
		t.Errorf("not match")
	}
	err2 := e.Wrap(err1, "")
	if !errors.As(err2, &AppErr) {
		t.Errorf("not match")
	}
}