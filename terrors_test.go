package terrors_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jalavosus/terrors"
)

func TestNotImplementedError_Is(t *testing.T) {
	type testCase struct {
		name string
		arg  error
		want assert.BoolAssertionFunc
	}

	tests := []testCase{
		{
			name: "NotImplementedError/true",
			arg: &terrors.NotImplementedError{
				Msg: "Hello!",
			},
			want: assert.True,
		},
		{
			name: "ErrNotImplemented/true",
			arg:  terrors.ErrNotImplemented,
			want: assert.True,
		},
		{
			name: "Random Error/false",
			arg:  errors.New("you're a dolphin"),
			want: assert.False,
		},
	}

	var check *terrors.NotImplementedError

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := errors.Is(check, tt.arg)
			tt.want(t, got)
		})
	}
}

func TestNotImplementedError_As(t *testing.T) {
	type testCase struct {
		name string
		arg  error
		want assert.BoolAssertionFunc
	}

	tests := []testCase{
		{
			name: "NotImplementedError/true",
			arg: &terrors.NotImplementedError{
				Msg: "Hello!",
			},
			want: assert.True,
		},
		{
			name: "ErrNotImplemented/false",
			arg:  terrors.ErrNotImplemented,
			want: assert.False,
		},
	}

	var check *terrors.NotImplementedError

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := errors.As(tt.arg, &check)
			tt.want(t, got)
		})
	}
}
