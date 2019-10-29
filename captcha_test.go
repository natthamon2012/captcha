package captcha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOperator(t *testing.T) {

	t.Run("valid operator should return operator", func(t *testing.T) {
		c := captcha{
			operator: "1",
		}
		actual, err := c.getOperator()
		expected := "-"
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("intvalid operator should return error", func(t *testing.T) {
		c := captcha{
			operator: "-1",
		}
		_, err := c.getOperator()
		expected := ErrOperatorNotSupport
		assert.Equal(t, expected, err)
	})
}

func TestGetNumberToText(t *testing.T) {
	t.Run("valid number should return number text", func(t *testing.T) {
		actual, err := getNumberToText("1")
		expected := "one"
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("intvalid number should return error", func(t *testing.T) {
		_, err := getNumberToText("-1")
		expected := ErrNumberToText
		assert.Equal(t, expected, err)
	})
}

func TestCaptcha(t *testing.T) {
	t.Run("format invalid should return error", func(t *testing.T) {
		c := captcha{
			format: "-1",
		}

		_, err := c.captcha()
		expected := ErrFormatNotSupport
		assert.Equal(t, expected, err)
	})

	t.Run("format type 0 should convert right operand number to text", func(t *testing.T) {
		c := captcha{
			format:       "0",
			leftOperand:  "0",
			operator:     "0",
			rightOperand: "1",
		}

		actual, err := c.captcha()
		expected := "0 + one"
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("format type 1 should convert left operand number to text", func(t *testing.T) {
		c := captcha{
			format:       "1",
			leftOperand:  "1",
			operator:     "0",
			rightOperand: "0",
		}
		actual, err := c.captcha()
		expected := "one + 0"
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}
