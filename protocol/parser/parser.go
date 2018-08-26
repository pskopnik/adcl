// Package parser provides parsing functionality for ADC protocol messages as
// well as for reading messages from connections.
//
// Users will likely only interact with the Parser type.
//
// Overview
//
//     Parser (bufio.Reader)
//         |  \
//         |   +---------------+
//     ConnReader         ReadMessage() <- - - - - -
//         ¦                   |                   ¦
//         ¦                   |                   ¦
//         +- - - - - - -> MessageReader - - -> ParseMessage()
//                             |                  | | \
//                             |                  | |  \
//                           Lexer                .......
//
// The above diagram outlines the structure of this package in terms of
// encapsulation (solid lines) and data flow (dashed lines).
//
// The Parser type is initialised with a bufio.Reader. The ReadMessage()
// method uses a ConnReader instance contained in the Parser to read a message
// line from the reader. ReadMessage() then creates a MessageReader, which
// contains a Lexer instance for internal use. ReadMessage() passes the
// MessageReader to the ParseMessage() function, which in turn calls a bunch
// of further Parse... functions for extracting message header and content.
// The resulting Message is returned by ReadMessage().
package parser

import (
	"bufio"

	"github.com/seoester/adcl/protocol/message"
)

type Parser struct {
	connReader ConnReader
}

// New creates a new Parser reading from the passed in bufio.Reader.
//
// Equivalent to:
//     var parser Parser
//     parser.Reset(r)
func New(r *bufio.Reader) *Parser {
	return &Parser{
		connReader: *NewConnReader(r),
	}
}

// Reset sets r as the reader and resets the internal state.
func (p *Parser) Reset(r *bufio.Reader) {
	p.connReader.Reset(r)
}

// ReadMessage reads the next Message from the underlying reader.
// One line is consumed in any case.
func (p *Parser) ReadMessage() (message.Message, error) {
	var mes message.Message
	var m MessageReader

	line, err := p.connReader.ReadMessageLine()
	if err != nil {
		return mes, err
	}

	m.Reset(line)

	mes, err = ParseMessage(&m)
	if err != nil {
		return mes, err
	}

	return mes, nil
}
