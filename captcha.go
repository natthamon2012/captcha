package captcha

import (
	"errors"
	"strconv"
)

var ErrFormatNotSupport = errors.New("format is not support")
var ErrOperatorNotSupport = errors.New("operator is not support")
var ErrNumberToText = errors.New("cannot convert number to text")

var operatorMap = map[int]string{
	0: "+",
	1: "-",
	2: "*",
}

var numberToText = map[int]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

type captcha struct {
	format       int
	leftOperand  int
	operator     int
	rightOperand int
}

func (c captcha) captcha() (string, error) {
	if c.format != 0 && c.format != 1 {
		return "", ErrFormatNotSupport
	}

	operator, err := c.getOperator()
	if err != nil {
		return "", err
	}

	leftOperand := strconv.Itoa(c.leftOperand)
	rightOperand := strconv.Itoa(c.rightOperand)

	if c.format == 0 {
		rightOperand, err = getNumberToText(c.rightOperand)
		if err != nil {
			return "", err
		}

	} else if c.format == 1 {
		leftOperand, err = getNumberToText(c.leftOperand)
		if err != nil {
			return "", err
		}
	}

	return display(leftOperand, operator, rightOperand), nil
}

func (c captcha) getOperator() (string, error) {
	v, ok := operatorMap[c.operator]
	if !ok {
		return "", ErrOperatorNotSupport
	}

	return v, nil
}

func getNumberToText(num int) (string, error) {
	v, ok := numberToText[num]
	if !ok {
		return "", ErrNumberToText
	}

	return v, nil
}

func display(left, operator, right string) string {
	return left + " " + operator + " " + right
}
