package main

import (
	"testing"
)

func TestValidateUser(t *testing.T) {
	cases := []struct {
		user User
		want int // Nombre d'erreurs attendues
	}{
		{User{"john_doe", "Password123!", "john@example.com", 30}, 0},
		{User{"", "Password123!", "john@example.com", 30}, 1},              // Username manquant
		{User{"john_doe", "short", "john@example.com", 30}, 1},             // Mot de passe trop court
		{User{"john_doe", "Password123!", "invalid-email", 30}, 1},         // Email invalide
		{User{"john_doe", "Password123!", "john@example.com", 10}, 1},      // Âge trop jeune
		{User{"john_doe", "Password123!", "john@example.com", 100}, 1},     // Âge trop vieux
		{User{"", "", "", 0}, 4},                                            // Tout est manquant
	}

	for _, c := range cases {
		got := ValidateUser(c.user)
		if len(got) != c.want {
			t.Errorf("ValidateUser(%v) == %v, want %v", c.user, len(got), c.want)
		}
	}
}
