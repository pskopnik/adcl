package parser

import (
	"bufio"
	"errors"
	"io"

	"github.com/seoester/adcl/protocol/encoding"
	"github.com/seoester/adcl/protocol/message"
)

// Constants related to ConnReader.
const (
	// MaxMessageLength is the maximum length a message may have.
	MaxMessageLength int = 16 << 10
)

// Error variables related to ConnReader.
var (
	ErrMessageTooLong = errors.New("message too long")
)

// ConnReader is a helper to read ADC messages from an bufio.Reader.
// Its primary method is ReadMessageLine, read its godoc for more information.
// An internal buffer is used, it may grow up to MaxMessageLength.
type ConnReader struct {
	r       *bufio.Reader
	lineBuf []byte
}

// NewConnReader creates a new ConnReader reading from the passed in
// bufio.Reader.
//
// Equivalent to:
//     var connReader ConnReader
//     connReader.Reset(r)
func NewConnReader(r *bufio.Reader) *ConnReader {
	return &ConnReader{
		r: r,
	}
}

// Reset sets r as the reader and resets the internal state.
func (c *ConnReader) Reset(r *bufio.Reader) {
	c.r = r
}

// ReadMessageLine reads the next message from the underlying reader.
// Exactly one line is read in all cases. The concluding end-of-line character
// is omitted from the returned line value.
// ReadMessageLine buffers at most MaxMessageLength bytes from the reader,
// if the line is longer, the message is discarded.
// During reading, the first 5 bytes of the message are validated using
// validateMessage().
func (c *ConnReader) ReadMessageLine() (line string, err error) {
	var buf, lineBuf []byte
	lineBuf = c.lineBuf

	// Note: ReadSlice includes the delimiter '\n'
	buf, err = c.r.ReadSlice(eol)

	if err != nil && err != bufio.ErrBufferFull {
		if err != io.EOF {
			// Try to read ahead until line ending
			_ = c.discardLine()
		}

		return
	}

	if len(buf) > MaxMessageLength {
		err = ErrMessageTooLong

		// Read ahead until line ending
		_ = c.discardLine()

		return
	}

	if len(buf) >= 5 {
		if err = c.validateMessage(buf); err != nil {
			_ = c.discardLine()

			return
		}
	}

	if err == nil {
		// Fast path: The whole message fit into the buffer.

		line = string(buf[:len(buf)-1])
		return
	}

	for {
		lineBuf = append(lineBuf, buf...)
		// Write reference to slice back into ConnReader, as lineBuf might have
		// been extended
		c.lineBuf = lineBuf[0:0]

		if err == nil {
			break
		}

		buf, err = c.r.ReadSlice(eol)

		if err != bufio.ErrBufferFull {
			if err != io.EOF {
				// Try to read ahead until line ending
				_ = c.discardLine()
			}

			return
		}

		if len(lineBuf)+len(buf) > MaxMessageLength {
			err = ErrMessageTooLong

			// Read ahead until line ending
			_ = c.discardLine()

			return
		}
	}

	line = string(lineBuf[:len(lineBuf)-1])
	return
}

// discardLine discards the current line, all bytes including the end-of-line
// character are read without allocating any buffers.
func (c *ConnReader) discardLine() error {
	var err error

	for {
		_, err = c.r.ReadSlice(eol)

		if err == nil {
			break
		} else if err != bufio.ErrBufferFull {
			return err
		}
	}

	return nil
}

// validateMessage validates the first 5 bytes of the message passed in.
func (c *ConnReader) validateMessage(buf []byte) error {
	if len(buf) < 5 {
		return ErrInvalidMessage
	}

	if buf[4] != space {
		return ErrInvalidMessage
	}

	if _, err := message.ParseType(buf[0]); err != nil {
		return err
	}

	// Do this manually instead of using message.ParseCommand to avoid the
	// allocating and copying done by string(buf)

	if !(encoding.IsUpperAlpha(buf[1]) &&
		encoding.IsUpperAlphaNum(buf[2]) &&
		encoding.IsUpperAlphaNum(buf[3])) {
		return message.ErrInvalidCommandName
	}

	return nil
}
