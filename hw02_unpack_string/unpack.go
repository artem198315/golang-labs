package hw02unpackstring

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func validate(s string) (err error) {
	re := regexp.MustCompile(`^\d.+|.*\d{2,}.*`)

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err)
		}
	}()

	if re.MatchString(s) {
		return ErrInvalidString
	}

	return nil
}

func Unpack(s string) (string, error) {
	if err := validate(s); err != nil {
		return "", err
	}

	var o string

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && unicode.IsDigit(rune(s[i+1])) {
			continue
		} else if unicode.IsDigit(rune(s[i])) {
			num, _ := strconv.Atoi(string(s[i]))
			o += strings.Repeat(string(s[i-1]), num)
			continue
		}
		o += string(s[i])
	}
	return o, nil
}
