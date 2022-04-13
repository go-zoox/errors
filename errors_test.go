package errors

import (
	"errors"
	stderrors "errors"
	"testing"
)

func TestNew(t *testing.T) {
	err := New("test")
	if err == nil {
		t.Error("New() should return an error")
	}

	if err.Error() != "test" {
		t.Errorf("New(\"test\") = %v, want %v", err, "test")
	}
}

func TestWrap(t *testing.T) {
	originErr := stderrors.New("origin error")
	err := Wrap(originErr, "test")
	err2 := Wrap(err, "test2")
	err3 := Wrap(err2, "test3")

	if err == nil {
		t.Error("Wrap() should return an error")
	}

	if err.Error() != "origin error: test" {
		t.Errorf("Wrap(originErr, \"test\") = %v, want %v", err, "origin error: test")
	}

	if err2.Error() != "origin error: test: test2" {
		t.Errorf("Wrap(err, \"test2\") = %v, want %v", err2, "origin error: test: test2")
	}

	if err3.Error() != "origin error: test: test2: test3" {
		t.Errorf("Wrap(err2, \"test3\") = %v, want %v", err3, "origin error: test: test2: test3")
	}
}

func TestWrapf(t *testing.T) {
	originErr := stderrors.New("origin error")
	err := Wrapf(originErr, "test %d", 1)

	if err.Error() != "origin error: test 1" {
		t.Errorf("expected %v, got %v", "origin error: test 1", err.Error())
	}
}

func TestErrorf(t *testing.T) {
	err := Errorf("test %d", 1)

	if err.Error() != "test 1" {
		t.Errorf("expected %v, got %v", "test 1", err.Error())
	}
}

func TestUnwrap(t *testing.T) {
	originErr := stderrors.New("origin error")
	err := Wrap(originErr, "test")

	if err == nil {
		t.Error("Wrap() should return an error")
	}

	if err.Error() != "origin error: test" {
		t.Errorf("Wrap(originErr, \"test\") = %v, want %v", err, "origin error: test")
	}

	if stderrors.Unwrap(err) != originErr {
		t.Errorf("err.Unwrap() = %v, want %v", stderrors.Unwrap(err), originErr)
	}
}

func TestIs(t *testing.T) {
	originErr := stderrors.New("origin error")
	err := Wrap(originErr, "test")

	if stderrors.Is(err, originErr) != true {
		t.Errorf("stderrors.Is(err, originErr) = %v, want %v", stderrors.Is(err, originErr), true)
	}

	if stderrors.Is(err, err) != true {
		t.Errorf("stderrors.Is(err, err) = %v, want %v", stderrors.Is(err, err), true)
	}

	if stderrors.Is(err, errors.New("test")) != true {
		t.Errorf("stderrors.Is(err, errors.New(\"test\")) = %v, want %v", stderrors.Is(err, errors.New("test")), true)
	}

}
