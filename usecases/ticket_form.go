package usecases

var (
	REQUIRED_FIELD   = "Required field"
	FIELD_MAS_LENGTH = "Field has max length of 100"
)

type TicketFormError struct {
	Ok          bool
	Title       string
	Description string
}

func ValidateForm(title string, description string) TicketFormError {
	validations := TicketFormError{
		Ok: true,
	}

	if title == "" {
		validations.Ok = false
		validations.Title = "Required field"
	} else if len(title) > 100 {
		validations.Ok = false
		validations.Title = "Field has max length of 100"
	}

	if description == "" {
		validations.Ok = false
		validations.Description = "Required field"
	}

	return validations
}
