package message

import (
	"errors"
)

// Error variables related to status types.
var (
	ErrInvalidStatusCode   = errors.New("invalid status code")
	ErrInvalidSeverityCode = errors.New("invalid severity code")
	ErrInvalidErrorCode    = errors.New("invalid error code")

	errInvalidDigit = errors.New("invalid digit")
)

type Severity int

const (
	SeveritySuccess     Severity = 0
	SeverityRecoverable          = 1
	SeverityFatal                = 2
)

type ErrorCode int

type StatusCode struct {
	Severity Severity
	Error    ErrorCode
}

func (s StatusCode) Code() int {
	return int(s.Severity)*100 + int(s.Error)
}

func ParseStatusCode(s string) (status StatusCode, err error) {
	if len(s) != 3 {
		return status, ErrInvalidStatusCode
	}

	switch s[0] {
	case '0':
		status.Severity = SeveritySuccess
	case '1':
		status.Severity = SeverityRecoverable
	case '2':
		status.Severity = SeverityFatal
	default:
		return status, ErrInvalidSeverityCode
	}

	var tmp int

	tmp, err = intFromDigit(s[1])
	if err != nil {
		return status, ErrInvalidErrorCode
	}
	status.Error += ErrorCode(10 * tmp)

	tmp, err = intFromDigit(s[2])
	if err != nil {
		return status, ErrInvalidErrorCode
	}
	status.Error += ErrorCode(tmp)

	return
}

const (
	byte0 byte = '0'
)

func intFromDigit(b byte) (i int, err error) {
	i = int(b - byte0)
	if i > 9 || i < 0 {
		return 0, errInvalidDigit
	}

	return
}
