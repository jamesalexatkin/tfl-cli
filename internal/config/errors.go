package config

// FieldInvalidError is returned when a config field has an invalid value.
type FieldInvalidError struct {
	Field string
}

func (e FieldInvalidError) Error() string {
	return "config value invalid: " + e.Field
}
