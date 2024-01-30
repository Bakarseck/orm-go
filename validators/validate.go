package validators

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func ValidateStruct(v interface{}) []ValidationResult {
	var validationResults []ValidationResult

	valueOf := reflect.ValueOf(v)
	if valueOf.Kind() != reflect.Struct {
		return validationResults
	}

	typeOf := valueOf.Type()

	for i := 0; i < valueOf.NumField(); i++ {
		field := valueOf.Field(i)
		fieldName := typeOf.Field(i).Name
		fieldTag := typeOf.Field(i).Tag.Get("validate")

		tags := strings.Split(fieldTag, ",")

		for _, tag := range tags {
			tagParts := strings.Split(tag, "=")
			tagName := tagParts[0]
			tagValue := ""
			if len(tagParts) > 1 {
				tagValue = tagParts[1]
			}

			result := ValidationResult{Field: fieldName, Tag: tagName}

			switch tagName {
			case "required":
				if field.Kind() == reflect.String {
					if len(field.String()) == 0 || strings.TrimSpace(field.String()) == "" {
						result.Valid = false
						result.Reason = "String length must be more than zero and must not contain only spaces, tabs, or newlines"
					} else {
						result.Valid = true
					}
				} else if field.Kind() == reflect.Int {
					if field.Int() == 0 {
						result.Valid = false
						result.Reason = "Int value must not be zero"
					} else {
						result.Valid = true
					}
				}
			case "password":
				if field.Kind() == reflect.String {

					password := field.String()

					hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
					hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
					hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
					hasSymbol := regexp.MustCompile(`[^a-zA-Z0-9]`).MatchString(password)

					if !hasLower || !hasUpper || !hasDigit || !hasSymbol {
						result.Valid = false
						result.Reason = "String must contain at least one lowercase, one uppercase, one number, and one symbol"
					} else {
						result.Valid = true
					}
				}
			case "email":
				if field.Kind() == reflect.String {
					email := field.String()
					emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
					validEmail := regexp.MustCompile(emailRegex).MatchString(email)
					if !validEmail {
						result.Valid = false
						result.Reason = "String must be a valid email address"
					} else {
						result.Valid = true
					}
				}
			case "username":
				if field.Kind() == reflect.String {
					username := field.String()
					usernameRegex := `^[a-zA-Z0-9_.-]+$`
					validUsername := regexp.MustCompile(usernameRegex).MatchString(username)
					if !validUsername {
						result.Valid = false
						result.Reason = "String must contain only letters, digits, dots, underscores, and dashes"
					} else {
						result.Valid = true
					}
				}
			case "min":
				if field.Kind() == reflect.String {
					minLen, err := strconv.Atoi(tagValue)
					if err != nil {
						result.Valid = false
						result.Reason = "Invalid 'min' tag value"
						break
					}
					if len(field.String()) < minLen {
						result.Valid = false
						result.Reason = fmt.Sprintf("String length must be greater than or equal to %d", minLen)
					} else {
						result.Valid = true
					}
				} else if field.Kind() == reflect.Int {
					minValue, err := strconv.Atoi(tagValue)
					if err != nil {
						result.Valid = false
						result.Reason = "Invalid 'min' tag value"
						break
					}
					if int(field.Int()) < minValue {
						result.Valid = false
						result.Reason = fmt.Sprintf("Int value must be greater than or equal to %d", minValue)
					} else {
						result.Valid = true
					}
				}
			case "max":
				if field.Kind() == reflect.String {
					maxLen, err := strconv.Atoi(tagValue)
					if err != nil {
						result.Valid = false
						result.Reason = "Invalid 'max' tag value"
						break
					}
					if len(field.String()) > maxLen {
						result.Valid = false
						result.Reason = fmt.Sprintf("String length must be less than or equal to %d", maxLen)
					} else {
						result.Valid = true
					}
				} else if field.Kind() == reflect.Int {
					maxValue, err := strconv.Atoi(tagValue)
					if err != nil {
						result.Valid = false
						result.Reason = "Invalid 'max' tag value"
						break
					}
					if int(field.Int()) > maxValue {
						result.Valid = false
						result.Reason = fmt.Sprintf("Int value must be less than or equal to %d", maxValue)
					} else {
						result.Valid = true
					}
				}
			}

			validationResults = append(validationResults, result)
		}
	}

	return validationResults
}
