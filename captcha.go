package captcha

import (
	"errors"
)

var ErrFormatNotSupport = errors.New("format is not support")
var ErrOperatorNotSupport = errors.New("operator is not support")
var ErrNumberToText = errors.New("cannot convert number to text")

var operatorMap = map[string]string{
	"0": "+",
	"1": "-",
	"2": "*",
}

var numberToText = map[string]string{
	"1": "one",
	"2": "two",
	"3": "three",
	"4": "four",
	"5": "five",
	"6": "six",
	"7": "seven",
	"8": "eight",
	"9": "nine",
}

type captcha struct {
	format       string
	leftOperand  string
	operator     string
	rightOperand string
}

func (c captcha) captcha() (string, error) {
	if c.format != "0" && c.format != "1" {
		return "", ErrFormatNotSupport
	}

	var err error
	c.operator, err = c.getOperator()
	if err != nil {
		return "", err
	}

	if c.format == "0" {
		c.rightOperand, err = getNumberToText(c.rightOperand)
		if err != nil {
			return "", err
		}
	} else if c.format == "1" {
		c.leftOperand, err = getNumberToText(c.leftOperand)
		if err != nil {
			return "", err
		}
	}

	return c.displayCaptcha(), nil
}

func (c captcha) getOperator() (string, error) {
	v, ok := operatorMap[c.operator]
	if !ok {
		return "", ErrOperatorNotSupport
	}

	return v, nil
}

func getNumberToText(num string) (string, error) {
	v, ok := numberToText[num]
	if !ok {
		return "", ErrNumberToText
	}

	return v, nil
}

func (c captcha) displayCaptcha() string {
	return c.leftOperand + " " + c.operator + " " + c.rightOperand
}
