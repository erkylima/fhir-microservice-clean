package validation

import (
	"regexp"
	"time"

	"github.com/go-playground/validator"
	"github.com/qbem-repos/patient-service/internal/shared/util"
)

func validate(s interface{}) error {
	validate := validator.New()

	if err := validate.Struct(s); err != nil {
		return err.(validator.ValidationErrors)
	}
	return nil
}

func Validation(s interface{}) error {
	if err := validate(s); err != nil {
		return err
	}
	return nil
}

func IsValidBirthDate(birthDate string) bool {
	date, err := time.Parse("2006-01-02", birthDate)
	if err != nil {
		return false
	}
	age := util.CalculateAge(date, time.Now())
	return age >= 18
}

func IsISO8601Date(date string) bool {
	ISO8601DateRegexString := "^(\\d{4})-(\\d{2})-(\\d{2})$"
	ISO8601DateRegex := regexp.MustCompile(ISO8601DateRegexString)
	return ISO8601DateRegex.MatchString(date)
}
