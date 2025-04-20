package config

// FieldInvalidError is returned when a config field has an invalid value.
type FieldInvalidError struct {
	Field Key
}

func (e FieldInvalidError) Error() string {
	return "config value invalid: " + string(e.Field)
}
