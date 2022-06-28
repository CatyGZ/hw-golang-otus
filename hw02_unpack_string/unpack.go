package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

const screening = `\`

type step struct {
	lastChar     string
	currentChar  string
	resultString strings.Builder
}

func Unpack(source string) (string, error) {
	// валидация строки
	isValidated := validateString(source)
	if !isValidated {
		return "", ErrInvalidString
	}

	var resultString strings.Builder
	step := step{
		lastChar:     "",
		currentChar:  "",
		resultString: resultString,
	}
	for _, r := range source {
		step.currentChar = string(r)
		// если текущий символ - символ экранирования
		if step.currentChar == screening {
			step.processScreening()
			continue
		}

		// если текущий символ - цифра
		if unicode.IsDigit(r) {
			err := step.processDigital()
			if err != nil {
				return "", err
			}
		} else {
			err := step.processChar()
			if err != nil {
				return "", err
			}
		}
	}

	return step.resultString.String(), nil
}

func validateString(str string) bool {
	if len(str) == 0 {
		return true
	}

	cutString := strings.TrimRight(str, screening)
	if (len(str)-len(cutString))%2 > 0 {
		return false
	}

	validateRegex := regexp.MustCompile(`^\D[\w\\]*$`)

	return validateRegex.Match([]byte(str))
}

func (s *step) processScreening() {
	if s.lastChar == screening {
		s.resultString.WriteString(s.currentChar)
		s.lastChar = ""
	} else {
		s.lastChar = s.currentChar
	}
}

func (s *step) processChar() error {
	if s.lastChar == screening {
		return ErrInvalidString
	}
	s.resultString.WriteString(s.currentChar)
	s.lastChar = s.currentChar

	return nil
}

func (s *step) processDigital() error {
	switch s.lastChar {
	case "":
		if s.resultString.Len() == 0 {
			return ErrInvalidString
		}
		resultRune := []rune(s.resultString.String())
		s.lastChar = string(resultRune[len(resultRune)-1])
		if s.lastChar == screening {
			cnt, err := strconv.Atoi(s.currentChar)
			if err != nil {
				return err
			}
			s.resultString.WriteString(strings.Repeat(s.lastChar, cnt-1))
			s.lastChar = ""
		} else {
			return ErrInvalidString
		}
	case screening:
		s.resultString.WriteString(s.currentChar)
		s.lastChar = s.currentChar
	default:
		cnt, err := strconv.Atoi(s.currentChar)
		if err != nil {
			return err
		}
		if cnt > 0 {
			s.resultString.WriteString(strings.Repeat(s.lastChar, cnt-1))
		} else {
			tempString := s.resultString.String()
			s.resultString.Reset()
			s.resultString.WriteString(removeLastChar(tempString))
		}
		s.lastChar = ""
	}
	return nil
}

func removeLastChar(source string) string {
	max := len(source)
	var resultString strings.Builder
	for i, r := range source {
		if i < max-1 {
			resultString.WriteRune(r)
		}
	}
	return resultString.String()
}
