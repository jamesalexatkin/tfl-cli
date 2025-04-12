package config

import "fmt"

type ConfigInvalidError struct {
	Field string
}

func (e ConfigInvalidError) Error() string {
	return fmt.Sprintf("config value invalid: %s", e.Field)
}
