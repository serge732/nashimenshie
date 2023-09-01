package merchandiser_utils

import (
	"bytes"
	"strconv"
)

func Progress(current int, end int) string {
	var digits []string

	r := end
	for {
		digits = append(digits, "0")

		r = r / 10

		if r == 0 {
			break
		}
	}

	r2 := current
	i := len(digits) - 1
	for {
		digits[i] = strconv.Itoa(r2 % 10)
		i--

		r2 = r2 / 10
		if r2 == 0 {
			break
		}
	}

	var buf bytes.Buffer
	for _, digit := range digits {
		buf.WriteString(digit)
	}

	return buf.String() + "/" + strconv.Itoa(end)
}
