package helpers

import (
	"errors"
	"regexp"
	"task_5_pbi_btpns_RakaNugraha/models"
)

// IsValidEmail memeriksa apakah alamat email yang diberikan valid.
func IsValidEmail(email string) bool {
	// Pola regex untuk memvalidasi alamat email
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Memeriksa apakah alamat email cocok dengan pola regex
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}

// ValidateUserRegistration melakukan validasi terhadap data pengguna saat registrasi.
func ValidateUser(user models.User) error {
	if user.Username == "" {
		return errors.New("username must be provided")
	}

	if user.Email == "" {
		return errors.New("email must be provided")
	}

	if user.Password == "" {
		return errors.New("password must be provided")
	}

	if len(user.Password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	return nil
}
