package validator

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrProfileMandatory = errors.New("phone number or fullname is mandatory")
	ErrPhoneLength      = errors.New("phone numbers must be at minimum 10 characters and maximum 13 characters")
	ErrPhoneContent     = errors.New("phone numbers must start with the Indonesia country code +62 and must be numeric")
	ErrNameLength       = errors.New("full name must be at minimum 3 characters and maximum 60 characters")
	ErrPasswordLength   = errors.New("passwords must be at minimum 6 characters and maximum 64 characters")
	ErrPasswordContent  = errors.New("passwords must contain at least 1 capital characters AND 1 number AND 1 special (non alpha-numeric) characters")
)

// New user validation
func ValidateNewUser(phone, fullname, password string) []error {
	var result []error

	err := validatePhoneLength(phone)
	if err != nil {
		result = append(result, err)
	}

	err = validatePhoneContent(phone)
	if err != nil {
		result = append(result, err)
	}

	err = validateNameLength(fullname)
	if err != nil {
		result = append(result, err)
	}

	err = validatePasswordLength(password)
	if err != nil {
		result = append(result, err)
	}

	err = validatePasswordContent(password)
	if err != nil {
		result = append(result, err)
	}

	return result
}

// Profile update validation
func ValidateProfileUpdate(phone, fullname *string) []error {
	var result []error

	if phone == nil && fullname == nil {
		result = append(result, ErrProfileMandatory)
		return result
	}

	if phone != nil {
		phoneString := *phone

		err := validatePhoneLength(phoneString)
		if err != nil {
			result = append(result, err)
		}

		err = validatePhoneContent(phoneString)
		if err != nil {
			result = append(result, err)
		}
	}
	if fullname != nil {
		fullnameString := *fullname

		err := validateNameLength(fullnameString)
		if err != nil {
			result = append(result, err)
		}
	}

	return result
}

// Phone length validation
func validatePhoneLength(phone string) error {
	lenPhone := len(phone)

	if lenPhone < 10 || lenPhone > 13 {
		return ErrPhoneLength
	}
	return nil
}

// Phone format validation
func validatePhoneContent(phone string) error {
	idCode := "+62"

	if !strings.HasPrefix(phone, idCode) {
		return ErrPhoneContent
	}

	if _, err := strconv.Atoi(strings.TrimPrefix(phone, idCode)); err != nil {
		return ErrPhoneContent
	}

	return nil
}

// Fullname length validation
func validateNameLength(fullname string) error {
	lenFullname := len(fullname)

	if lenFullname < 3 || lenFullname > 60 {
		return ErrNameLength
	}
	return nil
}

// Password length validation
func validatePasswordLength(password string) error {
	lenPassword := len(password)

	if lenPassword < 6 || lenPassword > 64 {
		return ErrPasswordLength
	}
	return nil
}

// Password format validation
func validatePasswordContent(password string) error {
	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
		if hasUpper && hasLower && hasNumber && hasSpecial {
			return nil
		}
	}

	return ErrPasswordContent
}
