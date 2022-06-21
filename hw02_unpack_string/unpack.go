package hw02unpackstring

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrInvalidString = errors.New("invalid string")
	Re               = regexp.MustCompile(`^\d.+|.*\d{2,}.*`)
)

func validate(s string) (err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err)
		}
	}()

	if Re.MatchString(s) {
		return ErrInvalidString
	}

	return nil
}

func Unpack(s string) (string, error) {
	if err := validate(s); err != nil {
		return "", err
	}

	var o string

	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
			continue
		} else if unicode.IsDigit(runes[i]) {
			num, _ := strconv.Atoi(string(runes[i]))
			o += strings.Repeat(string(runes[i-1]), num)
			continue
		}
		o += string(s[i])
	}
	return o, nil
}
