// Package encoding provides utilties for working with encoding formats of the
// ADC protocol.
package encoding

import (
	"errors"
	"strings"
	"unicode/utf8"
)

// Error variables related to the encoding package.
var (
	ErrInvalidString = errors.New("invalid string, is not utf-8 encoded")
)

var encoder = strings.NewReplacer(
	" ", "\\s",
	"\n", "\\n",
	"\\", "\\\\",
)

var decoder = strings.NewReplacer(
	"\\s", " ",
	"\\n", "\n",
	"\\\\", "\\",
)

// EncodeToADCString encodes a string to the ADC string parameter format. An
// error is returned if the passed in string is not a valid UTF-8 string.
func EncodeToADCString(s string) (string, error) {
	if !utf8.ValidString(s) {
		return "", ErrInvalidString
	}

	return encoder.Replace(s), nil
}

// DecodeADCString decodes a ADC string parameter. An error is returned if the
// passed in string is not a valid ADC string parameter, i.e. is not a UTF-8
// string.
func DecodeADCString(s string) (string, error) {
	if !utf8.ValidString(s) {
		return "", ErrInvalidString
	}

	return decoder.Replace(s), nil
}

// Constants which are used in encodings and checks.
const (
	byteA byte = 'A'
	byteZ      = 'Z'
	byte0      = '0'
	byte9      = '9'
)

// IsUpperAlphaNum returns true if the byte b represents an upper letter or a
// number (in ASCII). In the ADC protocol specification this is referred to as a
// simple_alphanum.
func IsUpperAlphaNum(b byte) bool {
	return (b >= byteA && b <= byteZ) || (b >= byte0 && b <= byte9)
}

// IsUpperAlpha returns true if the byte b represents an upper letter (in
// ASCII). In the ADC protocol specification this is referred to as a
// simple_alpha.
func IsUpperAlpha(b byte) bool {
	return b >= byteA && b <= byteZ
}
