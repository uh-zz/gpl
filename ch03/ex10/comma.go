package main

import (
	"bytes"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func bufcomma(s string) string {

	var (
		buf   bytes.Buffer
		count int
	)
	strbytes := []byte(s)

	for i := len(strbytes) - 1; i >= 0; i-- {

		if count != 0 && count%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(strbytes[i])
		count++
	}
	return reverse(buf.String())
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func WithDefaultBuffer(s string) string {
	var buf bytes.Buffer // default Buffer
	utf8bytes := []byte(s)

	// when bc (byte counter) is zero, then insert a comma
	bc := len(utf8bytes) % 3
	if bc == 0 {
		bc = 3
	}

	for _, r := range utf8bytes {
		if bc == 0 {
			buf.WriteByte(',')
			bc = 3
		}
		buf.WriteByte(r)
		bc--
	}
	return buf.String()
}
