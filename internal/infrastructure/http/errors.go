package http

import "fmt"

type InvalidJsonError struct {
	Data string
}

func (e InvalidJsonError) Error() string {
	if e.Data != "" {
		return fmt.Sprintf("Invaild JSON string: %s", e.Data)
	}

	return "Invalid JSON data"
}
