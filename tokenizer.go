package parsley

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// type token string

const (
	eof        = "EOF"
	identifier = "Identifier"
	number     = "Number"
)

type tokenizer struct {
	raw         string
	runes       []rune
	position    int
	currentRune rune

	Token      string
	Number     float64
	Identifier string
}

func newTokenizer(str string) (*tokenizer, error) {
	t := &tokenizer{raw: str, runes: []rune(str), position: 0}
	t.NextRune()
	err := t.NextToken()

	return t, err
}

func (t *tokenizer) NextToken() (err error) {
	// Skip whitespace
	for t.currentRune == ' ' {
		t.NextRune()
	}

	// Special characters
	switch t.currentRune {
	case '\000':
		t.Token = eof
		return

	case '+', '-', '*', '^', '/', '(', ')', '"', ',', '=', '>', '<', '&', '|':
		t.Token = string(t.currentRune)
		t.NextRune()

		switch t.currentRune {
		case '=', '&', '|':
			t.Token += string(t.currentRune)
			t.NextRune()
		}

		return
	}

	// Identifier - starts with letter or underscore
	if isPartOfVariable(t.currentRune) {
		sb := strings.Builder{}

		for isPartOfVariable(t.currentRune) {
			sb.WriteRune(t.currentRune)
			t.NextRune()
		}

		// Setup token
		t.Identifier = sb.String()
		t.Token = identifier
		return nil
	}

	// Number?
	if unicode.IsDigit(t.currentRune) || t.currentRune == '.' {
		// Capture digits/decimal point
		sb := strings.Builder{}
		haveDecimalPoint := false
		for unicode.IsDigit(t.currentRune) ||
			(!haveDecimalPoint && t.currentRune == '.') {
			sb.WriteRune(t.currentRune)

			haveDecimalPoint = t.currentRune == '.'
			t.NextRune()
		}

		// Parse it
		t.Number, err = strconv.ParseFloat(sb.String(), 64)
		if err != nil {
			return fmt.Errorf("error parsing float: %w", err)
		}

		t.Token = number
		return nil

	}

	return fmt.Errorf("unexpected character: %c", t.currentRune)
}

// Read the next character from the input stream
// and store it in currentRune, or load '\000' if EOF
func (t *tokenizer) NextRune() {
	if t.position < len(t.runes) {
		t.currentRune = t.runes[t.position]
	} else {
		t.currentRune = '\000'
	}
	t.position++
}

func isPartOfVariable(r rune) bool {
	return unicode.IsLetter(r) || r == '_' || r == '.'
}
