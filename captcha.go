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

func captcha(format, operator, leftOperand, rightOperand int) (string, error) {
	if format != 0 && format != 1 {
		return "", ErrFormatNotSupport
	}

	operatorStr, err := getOperator(operator)
	if err != nil {
		return "", err
	}

	switch format {
	case 0:
		rghOperandTxt, err := getNumberToText(rightOperand)
		if err != nil {
			return "", err
		}
		return displayCaptcha(strconv.Itoa(leftOperand), operatorStr, rghOperandTxt), nil
	case 1:
		lftOperandTxt, err := getNumberToText(leftOperand)
		if err != nil {
			return "", err
		}
		return displayCaptcha(lftOperandTxt, operatorStr, strconv.Itoa(rightOperand)), nil
	default:
		return "", ErrFormatNotSupport
	}
}

func getOperator(operator int) (string, error) {
	v, ok := operatorMap[operator]
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

func displayCaptcha(leftOperand, operator, rightOperand string) string {
	return leftOperand + " " + operator + " " + rightOperand
}
