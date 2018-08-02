package pascal

import (
	"github.com/object88/writing-compilers-and-interpreters/frontend"
	"github.com/pkg/errors"
)

type SpecialSymbolToken struct {
	Token
}

func NewSpecialSymbolToken(s *frontend.Source) (*SpecialSymbolToken, error) {
	sst := &SpecialSymbolToken{
		Token: *NewToken(s),
	}

	err := sst.extract()

	return sst, err
}

func (sst *SpecialSymbolToken) extract() error {
	currentChar, err := sst.CurrentChar()
	if err != nil {
		return errors.Wrap(err, "pascal.SpecialSymbolToken::extract: failed to get current character from start")
	}

	var text string
	switch currentChar {
	case '+', '-', '*', '/', ',', ';', '\'', '=', '(', ')', '[', ']', '{', '}', '^':
		text = string(currentChar)
		_, err = sst.NextChar()
		if err != nil {
			return errors.Wrap(err, "pascal.SpecialSymbolToken::extract: failed to digest single rune special symbol token")
		}
	case ':':
		currentChar, err = sst.NextChar()
		if err != nil {
			return errors.Wrap(err, "pascal.SpecialSymbolToken::extract: failed to get next character after opening ':'")
		}
		if currentChar == '=' {
			text = ":="
			_, err = sst.NextChar()
			if err != nil {
				return errors.Wrap(err, "pascal.SpecialSymbolToken::extract: failed to get character after ':='")
			}
		} else {
			text = ":"
		}
	case '<':
		currentChar, err = sst.NextChar()
		if err != nil {
			return errors.Wrap(err, "pascal.SpecialSymbolToken::extract: failed to get next character after opening '<'")
		}
		if currentChar == '=' {
			text = "<="
			_, err = sst.NextChar()
			if err != nil {
				return errors.Wrap(err, "pascal.SpecialSymbolToken::extract: failed to get character after '<='")
			}
		} else if currentChar == '>' {
			text = "<>"
			_, err = sst.NextChar()
			if err != nil {
				return errors.Wrap(err, "pascal.SpecialSymbolToken::extract: failed to get character after '<>")
			}
		} else {
			text = "<"
		}
	case '>':
		currentChar, err = sst.NextChar()
		if err != nil {
			return errors.Wrap(err, "pascal.SpecialSymbolToken::extract: failed to get next character after opening '>'")
		}
		if currentChar == '=' {
			text = ">="
			_, err = sst.NextChar()
			if err != nil {
				return errors.Wrap(err, "pascal.SpecialSymbolToken::extract: failed to get character after '>='")
			}
		} else {
			text = ">"
		}
	case '.':
		currentChar, err = sst.NextChar()
		if err != nil {
			return errors.Wrap(err, "pascal.SpecialSymbolToken::extract: failed to get next character after opening '.'")
		}
		if currentChar == '.' {
			text = ".."
			_, err = sst.NextChar()
			if err != nil {
				return errors.Wrap(err, "pascal.SpecialSymbolToken::extract: failed to get character after '..'")
			}
		} else {
			text = "."
		}
	default:
		_, err = sst.NextChar()
		if err != nil {
			return errors.Wrapf(err, "pascal.SpecialSymbolToken::extract: failed to get next character after opening '%c' (invalid rune)", currentChar)
		}
	}

	typ, ok := SpecialSymbolTokenTypes[text]
	if !ok {
		typ = ErrorTokenType
	}
	sst.AssignTypeAndText(frontend.TokenType(typ), text)

	return nil
}
