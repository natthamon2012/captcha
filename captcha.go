package captcha

import (
	"fmt"
)

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

func (c captcha) captcha() string {
	operator := c.getOperator()

	var formatCaptcha = map[int]string{
		0: fmt.Sprintf("%d %s %s", c.leftOperand, operator, getNumberToText(c.rightOperand)),
		1: fmt.Sprintf("%s %s %d", getNumberToText(c.leftOperand), operator, c.rightOperand),
	}

	return formatCaptcha[c.format]
}

func (c captcha) getOperator() string {
	return operatorMap[c.operator]
}

func getNumberToText(num int) string {
	return numberToText[num]
}
