package validator

import (
	"reflect"
	"testing"
)

func Test_validatePhoneLength(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "too short", args: args{phone: "123"}, wantErr: true},
		{name: "too long", args: args{phone: "123456789012345"}, wantErr: true},
		{name: "ok", args: args{phone: "123456789012"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validatePhoneLength(tt.args.phone); (err != nil) != tt.wantErr {
				t.Errorf("validatePhoneLength() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validatePhoneContent(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "not id", args: args{phone: "+658767"}, wantErr: true},
		{name: "not number", args: args{phone: "+62ABC"}, wantErr: true},
		{name: "ok", args: args{phone: "+6286542527"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validatePhoneContent(tt.args.phone); (err != nil) != tt.wantErr {
				t.Errorf("validatePhoneContent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateNameLength(t *testing.T) {
	type args struct {
		fullname string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "too short", args: args{fullname: "a"}, wantErr: true},
		{name: "too long", args: args{fullname: "Lorem ipsum happy birthday merry christmas happy new year indonesia united states of america"}, wantErr: true},
		{name: "ok", args: args{fullname: "John Smith"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateNameLength(tt.args.fullname); (err != nil) != tt.wantErr {
				t.Errorf("validateNameLength() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validatePasswordLength(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "too short", args: args{password: "a"}, wantErr: true},
		{name: "too long", args: args{password: "Lorem ipsum happy birthday merry christmas happy new year indonesia united states of america yohoho forever"}, wantErr: true},
		{name: "ok", args: args{password: "John Smith"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validatePasswordLength(tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("validatePasswordLength() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validatePasswordContent(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "no upper", args: args{password: "abc12!"}, wantErr: true},
		{name: "no lower", args: args{password: "A1!"}, wantErr: true},
		{name: "no number", args: args{password: "Aa!"}, wantErr: true},
		{name: "no special", args: args{password: "Aa1"}, wantErr: true},
		{name: "ok", args: args{password: "Abc12!"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validatePasswordContent(tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("validatePasswordContent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateNewUser(t *testing.T) {
	type args struct {
		phone    string
		fullname string
		password string
	}
	tests := []struct {
		name string
		args args
		want []error
	}{
		{
			name: "not pass",
			args: args{phone: "123", fullname: "A", password: "abc"},
			want: []error{ErrPhoneLength, ErrPhoneContent, ErrNameLength, ErrPasswordLength, ErrPasswordContent},
		},
		{
			name: "ok",
			args: args{phone: "+6287541622", fullname: "John Smith", password: "Abc12!"},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateNewUser(tt.args.phone, tt.args.fullname, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateNewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateProfileUpdate(t *testing.T) {
	type args struct {
		phone    *string
		fullname *string
	}

	phone1 := "123"
	fullname1 := "a"

	phone2 := "+6281675452"
	fullname2 := "John Semith"

	tests := []struct {
		name string
		args args
		want []error
	}{
		{
			name: "empty",
			args: args{},
			want: []error{ErrProfileMandatory},
		},
		{
			name: "not pass",
			args: args{phone: &phone1, fullname: &fullname1},
			want: []error{ErrPhoneLength, ErrPhoneContent, ErrNameLength},
		},
		{
			name: "ok",
			args: args{phone: &phone2, fullname: &fullname2},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateProfileUpdate(tt.args.phone, tt.args.fullname); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateProfileUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}
