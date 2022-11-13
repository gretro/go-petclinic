package system

import (
	"fmt"
	"strings"
)

type EntityNotFoundError struct {
	EntityName  string
	ID          string
	OtherParams map[string]interface{}
}

func formatMapToString(params map[string]interface{}) string {
	var sBuilder strings.Builder

	for key, value := range params {
		sBuilder.WriteString(fmt.Sprintf("%q=%q;", key, value))
	}

	return sBuilder.String()
}

func (err *EntityNotFoundError) Error() string {
	var sBuilder strings.Builder

	if err.ID == "" {
		sBuilder.WriteString(fmt.Sprintf("Unable to find Entity of type %q", err.EntityName))
	} else {
		sBuilder.WriteString(fmt.Sprintf("Unable to find Entity of type %q with ID %q", err.EntityName, err.ID))
	}

	if err.OtherParams != nil && len(err.OtherParams) > 0 {
		sBuilder.WriteString(fmt.Sprintf(" (%s)", formatMapToString(err.OtherParams)))
	}

	return sBuilder.String()
}

type EntityConflictError struct {
	EntityName  string
	ID          string
	OtherParams map[string]interface{}
}

func (err *EntityConflictError) Error() string {
	var sBuilder strings.Builder

	if err.ID == "" {
		sBuilder.WriteString(fmt.Sprintf("Conflict with Entity of type %q", err.EntityName))
	} else {
		sBuilder.WriteString(fmt.Sprintf("Conflict Entity of type %q with ID %q", err.EntityName, err.ID))
	}

	if err.OtherParams != nil && len(err.OtherParams) > 0 {
		sBuilder.WriteString(fmt.Sprintf(" (%s)", formatMapToString(err.OtherParams)))
	}

	return sBuilder.String()
}
