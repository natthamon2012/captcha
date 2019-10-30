package captcha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOperator(t *testing.T) {

	t.Run("valid operator should return operator", func(t *testing.T) {
		c := captcha{
			operator: 1,
		}
		actual := c.getOperator()
		expected := "-"
		assert.Equal(t, expected, actual)
	})

	t.Run("intvalid operator should return error", func(t *testing.T) {
		c := captcha{
			operator: -1,
		}
		actual := c.getOperator()
		expected := ""
		assert.Equal(t, expected, actual)
	})
}

func TestGetNumberToText(t *testing.T) {
	t.Run("valid number should return number text", func(t *testing.T) {
		actual := getNumberToText(1)
		expected := "one"
		assert.Equal(t, expected, actual)
	})

	t.Run("intvalid number should return error", func(t *testing.T) {
		actual := getNumberToText(-1)
		expected := ""
		assert.Equal(t, expected, actual)
	})
}

func TestCaptcha(t *testing.T) {
	t.Run("format type 0 should convert right operand number to text", func(t *testing.T) {
		c := captcha{
			format:       0,
			leftOperand:  0,
			operator:     0,
			rightOperand: 1,
		}

		actual := c.captcha()
		expected := "0 + one"
		assert.Equal(t, expected, actual)
	})

	t.Run("format type 1 should convert left operand number to text", func(t *testing.T) {
		c := captcha{
			format:       1,
			leftOperand:  1,
			operator:     0,
			rightOperand: 0,
		}
		actual := c.captcha()
		expected := "one + 0"
		assert.Equal(t, expected, actual)
	})
}
