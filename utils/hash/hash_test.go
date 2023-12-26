package hash

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	type args struct {
		password string
		salt     []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "default", args: args{password: "123", salt: []byte("")}, want: "+J+DQEe2S8B2wdJjqR7W2CyOhFRfeEhk7ycLYNIFMUs"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HashPassword(tt.args.password, tt.args.salt); got != tt.want {
				t.Errorf("HashPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPasswordsMatch(t *testing.T) {
	type args struct {
		hashedPassword string
		currPassword   string
		salt           []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "match", args: args{hashedPassword: "+J+DQEe2S8B2wdJjqR7W2CyOhFRfeEhk7ycLYNIFMUs", currPassword: "123", salt: []byte("")}, want: true},
		{name: "not-match", args: args{hashedPassword: "xyz", currPassword: "123", salt: []byte("")}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PasswordsMatch(tt.args.hashedPassword, tt.args.currPassword, tt.args.salt); got != tt.want {
				t.Errorf("PasswordsMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
