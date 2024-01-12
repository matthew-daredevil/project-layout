package error_sample

import (
    "fmt"
    "strings"

    "gopkg.in/go-playground/validator.v9"
)

type fieldError struct {
    err validator.FieldError
}

func (q fieldError) String() string {
    var sb strings.Builder

    sb.WriteString("validation failed on the field '" + q.err.Field() + "', condition: " + q.err.ActualTag())

    if q.err.Param() != "" {
        sb.WriteString(" { " + q.err.Param() + " }")
    }

    if q.err.Value() != nil && q.err.Value() != "" {
        sb.WriteString(fmt.Sprintf(", actual: %v", q.err.Value()))
    }

    return sb.String()
}