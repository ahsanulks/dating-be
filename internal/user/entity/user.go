package entity

import (
	customerror "app/internal/custom_error"
	"app/internal/user/param/request"
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode"
)

const (
	phoneNumberMinLen = 10
	phoneNumberMaxLen = 15

	nameMinLen = 3
	nameMaxLen = 255

	passwordMinLen = 6
	passwordMaxLen = 64

	usernameMinLen = 5
	usernameMaxLen = 50
)

type Gender string

const (
	genderUnknown Gender = ""
	genderMale    Gender = "male"
	genderFemale  Gender = "female"
)

type User struct {
	ID          int64
	Name        string
	Username    string
	PhoneNumber string
	Gender      Gender
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewUser(param *request.CreateUser) (*User, error) {
	user := &User{
		Name:        param.Name,
		PhoneNumber: param.PhoneNumber,
		Password:    param.Password,
		Username:    strings.ToLower(param.Username),
		Gender:      genderUnknown.fromString(param.Gender),
	}

	validationError := customerror.NewValidationError()
	if err := user.validateUsername(); err != nil {
		validationError.Merge(err)
	}

	if user.Gender == genderUnknown {
		validationError.AddError("gender", "can only male and female")
	}

	if err := user.validatePhoneNumber(); err != nil {
		validationError.Merge(err)
	}

	if err := user.validateName(); err != nil {
		validationError.Merge(err)
	}

	if err := user.validatePassword(); err != nil {
		validationError.Merge(err)
	}

	if validationError.HasError() {
		return nil, validationError
	}

	return user, nil
}

func (user User) validateUsername() error {
	validationError := customerror.NewValidationError()

	if len(user.Username) < usernameMinLen || len(user.Username) > usernameMaxLen {
		validationError.AddError("username", fmt.Sprintf("must be between %d and %d characters in length", usernameMinLen, usernameMaxLen))
	}

	// Allowing usernames to contain letters, numbers, underscores, and dashes,
	usernamePattern := `^[a-z0-9_-]+$`
	regexpPattern := regexp.MustCompile(usernamePattern)
	if !regexpPattern.MatchString(user.Username) {
		validationError.AddError("username", "can only contain letter, numbers, underscores, and dashes")
	}
	if validationError.HasError() {
		return validationError
	}

	return nil
}

func (user User) validatePhoneNumber() error {
	validationError := customerror.NewValidationError()
	if len(user.PhoneNumber) < phoneNumberMinLen || len(user.PhoneNumber) > phoneNumberMaxLen {
		validationError.AddError("phoneNumber", fmt.Sprintf("must be between %d and %d characters in length", phoneNumberMinLen, phoneNumberMaxLen))
	}

	// check that have prefix +62 and only containt 0-9 after that
	phoneNumberRegex := regexp.MustCompile(`^\+62[0-9]`)
	match := phoneNumberRegex.MatchString(user.PhoneNumber)
	if !match {
		validationError.AddError("phoneNumber", "must start with '+62' and only containt number")
	}

	if validationError.HasError() {
		return validationError
	}

	return nil
}

func (user User) validateName() error {
	validationError := customerror.NewValidationError()
	if len(user.Name) < nameMinLen || len(user.Name) > nameMaxLen {
		validationError.AddError("name", fmt.Sprintf("must be between %d and %d characters in length", nameMinLen, nameMaxLen))
	}

	if validationError.HasError() {
		return validationError
	}

	return nil
}

func (user User) validatePassword() error {
	validationError := customerror.NewValidationError()
	if len(user.Password) < passwordMinLen || len(user.Password) > passwordMaxLen {
		validationError.AddError("password", fmt.Sprintf("must be between %d and %d characters in length", passwordMinLen, passwordMaxLen))
	}

	// Check for at least 1 capital letter, 1 number, and 1 special character
	if !strongPassword(user.Password) {
		validationError.AddError("password", "containing at least 1 capital characters AND 1 number AND 1 special (nonalpha-numeric) characters")
	}

	if validationError.HasError() {
		return validationError
	}

	return nil
}

func strongPassword(password string) bool {
	hasCapital := false
	hasLowercase := false
	hasNumber := false
	hasSpecialChar := false

	for _, char := range password {
		if unicode.IsUpper(char) {
			hasCapital = true
		} else if unicode.IsLower(char) {
			hasLowercase = true
		} else if unicode.IsDigit(char) {
			hasNumber = true
		} else if char >= 33 && char <= 47 || char >= 58 && char <= 64 || char >= 91 && char <= 96 || char >= 123 && char <= 126 {
			hasSpecialChar = true
		}
	}

	return hasCapital && hasLowercase && hasNumber && hasSpecialChar
}

func (g Gender) fromString(value string) Gender {
	lowerValue := strings.ToLower(value)
	if lowerValue == "male" {
		return genderMale
	} else if lowerValue == "female" {
		return genderFemale
	}
	return genderUnknown
}

func (g Gender) String() string {
	if g == genderMale {
		return "male"
	} else if g == genderFemale {
		return "female"
	}
	return ""
}
