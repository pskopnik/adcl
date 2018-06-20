package parser

import (
	"errors"
	"io"
	"strings"
)

// Error variables related to Lexer.
var (
	ErrInvalidToken error = errors.New("invalid token")
)

// Lexer describes a string lexer suitable for tokenising ADC Messages.
//
// Space (0x20) is used as the only token separator. The lexer does not handle
// end-of-line characters (0x0A). When passing in ADC Messages, it expects the
// end-of-line chracter to already having been removed.
type Lexer struct {
	s string
	// offset is the index of the byte in s to be read next.
	// (strings are byte-indexed)
	offset int
}

// Reset sets string s as the input and resets the internal state. Afterwards,
// Next() will return the first token in s.
func (l *Lexer) Reset(s string) {
	l.s = s
	l.offset = 0
}

func (l *Lexer) readNext() (string, int, error) {
	start := l.offset

	for {
		if start == len(l.s) {
			return "", 0, io.EOF
		}

		ind := strings.IndexRune(l.s[start:], rune(space))
		if ind == 0 {
			start++
			continue
		} else if ind == -1 {
			n := len(l.s) - l.offset
			return l.s[start:], n, nil
		} else {
			n := (start - l.offset) + ind + 1
			return l.s[start : start+ind], n, nil
		}
	}
}

// Next reads the next token from the string and advances the offset.
// Additional spaces are skipped, the token is also returned without any
// surrounding spaces.
// If the end of the string is reached, io.EOF is returned as the error.
func (l *Lexer) Next() (string, error) {
	tok, n, err := l.readNext()
	if err != nil {
		return "", err
	}

	l.offset += n
	return tok, err
}

// Peek returns the next token in the string without advancing the offset.
// Additional spaces are skipped, the token is also returned without any
// surrounding spaces.
// If the end of the string is reached, io.EOF is returned as the error.
func (l *Lexer) Peek() (string, error) {
	tok, _, err := l.readNext()
	return tok, err
}

// Put returns the last emitted token back onto the string.
// The passed in token "tok" must be (match) the last emitted token, otherwise
// ErrInvalidToken is returned as the error.
func (l *Lexer) Put(tok string) error {
	end := l.offset

	for strings.LastIndexByte(l.s[:end], space) == end-1 {
		end--
	}

	if strings.Compare(l.s[end-len(tok):end], tok) == 0 {
		l.offset = end - len(tok)
		return nil
	} else {
		return ErrInvalidToken
	}
}
