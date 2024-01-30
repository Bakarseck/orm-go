package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type User struct {
	Username string
	Password string `validate:"required,password,min=10"`
	Email    string `validate:"required,email"`
	Age      int    `validate:"required,min=13,max=99"`
}

// ValidateUser valide un utilisateur en fonction des balises de validation définies dans la structure.
func ValidateUser(user User) []string {
	var errors []string
	val := reflect.ValueOf(user)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		tag := val.Type().Field(i).Tag.Get("validate")
		if tag == "" {
			continue
		}

		errs := validateField(field.Interface(), tag)
		if len(errs) > 0 {
			errors = append(errors, fmt.Sprintf("%s: %s", val.Type().Field(i).Name, strings.Join(errs, ", ")))
		}
	}
	return errors
}

// validateField applique les règles de validation à une valeur de champ donnée.
func validateField(value interface{}, tag string) []string {
	rules := strings.Split(tag, ",")
	var errors []string

	for _, rule := range rules {
		switch {
		case rule == "required" && isZero(value):
			errors = append(errors, "is required")
		case strings.HasPrefix(rule, "min=") || strings.HasPrefix(rule, "max="):
			if err := checkMinMax(value, rule); err != "" {
				errors = append(errors, err)
			}
		case rule == "email" && !isValidEmail(fmt.Sprint(value)):
			errors = append(errors, "must be a valid email")
		case rule == "password" && !isValidPassword(fmt.Sprint(value)):
			errors = append(errors, "must contain at least one lowercase, one uppercase, one number, and one symbol")
		case rule == "username" && !isValidUsername(fmt.Sprint(value)):
			errors = append(errors, "must contain only letters, digits, dots, underscores, and dashes")
		}
	}
	return errors
}

// Fonctions auxiliaires pour la validation
func isZero(value interface{}) bool {
	return reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface())
}

func checkMinMax(value interface{}, rule string) string {
	parts := strings.Split(rule, "=")
	if len(parts) != 2 {
		return "invalid rule"
	}

	limit, err := strconv.Atoi(parts[1])
	if err != nil {
		return "invalid number"
	}

	switch v := value.(type) {
	case int:
		if parts[0] == "min" && v < limit {
			return fmt.Sprintf("must be greater than or equal to %d", limit)
		} else if parts[0] == "max" && v > limit {
			return fmt.Sprintf("must be less than or equal to %d", limit)
		}

	case string:
		if parts[0] == "min" && len(v) < limit {
			return fmt.Sprintf("length must be greater than or equal to %d", limit)
		} else if parts[0] == "max" && len(v) > limit {
			return fmt.Sprintf("length must be less than or equal to %d", limit)
		}
	}

	return ""
}

// Helper functions for specific validations
func isValidPassword(password string) bool {
	var (
		hasMin     = regexp.MustCompile(`[a-z]`).MatchString
		hasMaj     = regexp.MustCompile(`[A-Z]`).MatchString
		hasDigit   = regexp.MustCompile(`\d`).MatchString
		hasSpecial = regexp.MustCompile(`[\W_]`).MatchString
	)

	return hasMin(password) && hasMaj(password) && hasDigit(password) && hasSpecial(password)
}

func isValidEmail(s string) bool {
	var emailRegex = regexp.MustCompile(`^\S+@\S+\.\S+$`)
	return emailRegex.MatchString(s)
}

func isValidUsername(s string) bool {
	var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9._-]+$`)
	return usernameRegex.MatchString(s)
}

func main() {
	user := User{
		Username: "bakarseck",
		Password: "Pass12!",
		Email:    "john@example.com",
		Age:      25,
	}

	errors := ValidateUser(user)
	if len(errors) > 0 {
		fmt.Println("Validation errors:")
		for _, err := range errors {
			fmt.Println(err)
		}
	} else {
		fmt.Println("User is valid!")
	}
}
