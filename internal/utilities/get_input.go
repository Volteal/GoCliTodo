package utilities

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

func GetUserInput(r io.Reader, args ...string) (string, string, error) {
	if len(args) <= 0 {
		return "", "", errors.New("add string can not be empty")
	}

	if len(args) > 0 {
		values := strings.Join(args, " ")
		v := strings.Split(values, ",")
		return v[0], v[1], nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", "", err
	}

	fullTxt := scanner.Text()
	if len(fullTxt) == 0 {
		return "", "", errors.New("empty todo string is not allowed")
	}

	txt := strings.Split(fullTxt, ",")

	return txt[0], txt[1], nil
}
