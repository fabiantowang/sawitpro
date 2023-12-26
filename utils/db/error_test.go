package db

import (
	"errors"
	"testing"
)

func TestErrorCode(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "new error", args: args{err: errors.New("")}, want: ""},
		{name: "duplicate", args: args{err: ErrUniqueViolation}, want: "23505"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrorCode(tt.args.err); got != tt.want {
				t.Errorf("ErrorCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
