package request

import (
	"fmt"
	"voice-out-be/internal/util"

	"github.com/go-playground/validator/v10"
)

type CreatePostRequest struct {
	From string `json:"from" validate:"required"`
	To   string `json:"to" validate:"required"`
	Body string `json:"body" validate:"required"`
}

func (r *CreatePostRequest) Validate() error {
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

			}
		}

		return fmt.Errorf("%s", validationErrorMessage)
	}

	return nil
}
