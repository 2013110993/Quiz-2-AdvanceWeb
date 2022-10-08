// Filename : internal/data/user.go

package data

import (
	"time"

	"fileuploading.federicorosado.net/internal/validator"
)

type User struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email,omitempty"`
	Address   string    `json:"address"`
	School    string    `json:"school"`
	Degree    string    `json:"degree"`
	Version   int32     `json:"version"`
}

func ValidateUser(v *validator.Validator, user *User) {

	v.Check(user.Name != "", "name", "must be provided")
	v.Check(len(user.Name) <= 200, "name", "must no more 200 characters")

	v.Check(user.Phone != "", "phone", "must be provided")
	v.Check(validator.Match(user.Phone, validator.PhoneRX), "phone", "must be a valid phone number")

	v.Check(user.Email != "", "email", "must be provided")
	v.Check(validator.Match(user.Email, validator.EmailRX), "email", "must be a valid email")

	v.Check(user.Address != "", "address", "must be provided")
	v.Check(len(user.Address) <= 500, "address", "must no more 500 characters")

	v.Check(user.School != "", "school", "must be provided")
	v.Check(len(user.School) <= 200, "school", "must no more 200 characters")

	v.Check(user.Degree != "", "degree", "must be provided")
	v.Check(len(user.Degree) <= 200, "degree", "must no more 200 characters")

}
