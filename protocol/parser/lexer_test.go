package parser_test

import (
	"io"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/seoester/adcl/protocol/parser"
)

var _ = Describe("Lexer", func() {
	Describe("Next() - Reading the next token", func() {

		It("should recognise an empty input", func() {
			lexer := NewLexer("")
			_, err := lexer.Next()
			Ω(err).Should(Equal(io.EOF))
			_, err = lexer.Next()
			Ω(err).Should(Equal(io.EOF))
		})

		It("should recognise a single token", func() {
			lexer := NewLexer("asdf")
			Ω(lexer.Next()).Should(Equal("asdf"))
			_, err := lexer.Next()
			Ω(err).Should(Equal(io.EOF))
		})

		It("should recognise two simple tokens", func() {
			lexer := NewLexer("asdf ghjk")
			Ω(lexer.Next()).Should(Equal("asdf"))
			Ω(lexer.Next()).Should(Equal("ghjk"))
			_, err := lexer.Next()
			Ω(err).Should(Equal(io.EOF))
		})

		It("should ignore trailing spaces", func() {
			lexer := NewLexer("asdf ")
			Ω(lexer.Next()).Should(Equal("asdf"))
			_, err := lexer.Next()
			Ω(err).Should(Equal(io.EOF))

			lexer = NewLexer("asdf  ")
			Ω(lexer.Next()).Should(Equal("asdf"))
			_, err = lexer.Next()
			Ω(err).Should(Equal(io.EOF))
		})

		It("should ignore leading spaces", func() {
			lexer := NewLexer(" asdf")
			Ω(lexer.Next()).Should(Equal("asdf"))
			_, err := lexer.Next()
			Ω(err).Should(Equal(io.EOF))

			lexer = NewLexer("  asdf")
			Ω(lexer.Next()).Should(Equal("asdf"))
			_, err = lexer.Next()
			Ω(err).Should(Equal(io.EOF))
		})

		It("should ignore multiple succeeding spaces", func() {
			lexer := NewLexer("asdf  ghjk")
			Ω(lexer.Next()).Should(Equal("asdf"))
			Ω(lexer.Next()).Should(Equal("ghjk"))
			_, err := lexer.Next()
			Ω(err).Should(Equal(io.EOF))
		})
	})

	Describe("Peek() - Peeking the next token", func() {
		It("should return the next token without advancing the offset", func() {
			lexer := NewLexer("asdf  ghjk")
			Ω(lexer.Peek()).Should(Equal("asdf"))
			Ω(lexer.Peek()).Should(Equal("asdf"))

			_, err := lexer.Next()
			Ω(err).ShouldNot(HaveOccurred())
			Ω(lexer.Peek()).Should(Equal("ghjk"))
			Ω(lexer.Peek()).Should(Equal("ghjk"))
		})
	})

	Describe("Put() - Return a token", func() {
		It("should return a single token", func() {
			lexer := NewLexer("asdf")
			tok, err := lexer.Next()
			Ω(err).ShouldNot(HaveOccurred())

			Ω(lexer.Put(tok)).ShouldNot(HaveOccurred())

			Ω(lexer.Next()).Should(Equal(tok))
		})

		It("should return tokens while ignoring separating spaces", func() {
			lexer := NewLexer("asdf    ghjk")
			tok, err := lexer.Next()
			Ω(err).ShouldNot(HaveOccurred())

			Ω(lexer.Put(tok)).ShouldNot(HaveOccurred())

			Ω(lexer.Next()).Should(Equal(tok))

			tok, err = lexer.Next()
			Ω(err).ShouldNot(HaveOccurred())

			Ω(lexer.Put(tok)).ShouldNot(HaveOccurred())

			Ω(lexer.Next()).Should(Equal(tok))
		})
	})
})
