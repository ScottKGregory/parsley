package parsley

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type token string

const (
	eof         = "EOF"
	add         = "Add"
	subtract    = "Subtract"
	multiply    = "Multiply"
	divide      = "Divide"
	openParens  = "OpenParens"
	closeParens = "CloseParens"
	identifier  = "Identifier"
	number      = "Number"
	equal       = "Equal"
	greaterThan = "GreaterThan"
	lessThan    = "LessThan"
	power       = "Power"
	comma       = "Comma"
	quote       = "Quote"
	and         = "And"
	or          = "Or"
)

type tokenizer struct {
	raw         string
	runes       []rune
	position    int
	currentRune rune

	Token      token
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

	case '+':
		t.NextRune()
		t.Token = add
		return

	case '-':
		t.NextRune()
		t.Token = subtract
		return

	case '*':
		t.NextRune()
		t.Token = multiply
		return

	case '^':
		t.NextRune()
		t.Token = power
		return

	case '/':
		t.NextRune()
		t.Token = divide
		return

	case '(':
		t.NextRune()
		t.Token = openParens
		return

	case ')':
		t.NextRune()
		t.Token = closeParens
		return

	case '"':
		t.NextRune()
		t.Token = quote
		return

	case ',':
		t.NextRune()
		t.Token = comma
		return

	case '=':
		t.NextRune()
		t.Token = equal
		return

	case '>':
		t.NextRune()
		t.Token = greaterThan
		return

	case '<':
		t.NextRune()
		t.Token = lessThan
		return

	case '&':
		t.NextRune()
		t.Token = and
		return

	case '|':
		t.NextRune()
		t.Token = or
		return
	}

	// Identifier - starts with letter or underscore
	if isPartOfVariable(t.currentRune) {
		sb := strings.Builder{}

		for isPartOfVariable(t.currentRune) {
			_, err = sb.WriteRune(t.currentRune)
			if err != nil {
				return err
			}

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
			_, err = sb.WriteRune(t.currentRune)
			if err != nil {
				return err
			}

			haveDecimalPoint = t.currentRune == '.'
			t.NextRune()
		}

		// Parse it
		t.Number, err = strconv.ParseFloat(sb.String(), 64)
		if err != nil {
			return err
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
