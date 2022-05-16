package request

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"voice-out-be/internal/util"
)

type RegisterUserRequest struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

func (r *RegisterUserRequest) Validate() error {
	validate := util.GetNewValidator()

	if err := validate.Validate(r); err != nil {
		validationErrorMessage := validate.GetErrorMessage(err.(validator.ValidationErrors))

		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "alphanumeric":
				validationErrorMessage = fmt.Sprintf("%s has special characters",
					err.Field())

			case "required":
				validationErrorMessage = fmt.Sprintf("%s is required",
					err.Field())

			case "email":
				validationErrorMessage = fmt.Sprintf("%s is not valid email",
					err.Field())

			case "max":
				validationErrorMessage = "Exceeded maximum characters"
			}
		}

		return fmt.Errorf("%s", validationErrorMessage)
	}

	return nil
}
