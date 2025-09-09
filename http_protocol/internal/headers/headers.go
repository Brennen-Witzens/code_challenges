package headers

import (
	"bytes"
	"fmt"
	"strings"
)

type Headers map[string]string

const crlf = "\r\n"

func NewHeaders() Headers {
	return make(map[string]string)
}

func (h Headers) Parse(data []byte) (n int, done bool, err error) {
	idx := bytes.Index(data, []byte(crlf))

	// Couldn't find the CRLF - \r\n
	if idx == -1 {
		return 0, false, nil
	} else if idx == 0 { // Found at start of data?
		return 0, true, nil
	}

	fieldLine := string(data[:idx])
	// "Host: localhost:42069\r\n\r\n"
	parts := strings.SplitN(strings.TrimSpace(fieldLine), ":", 2)

	if parts[0] == "Host " {
		return 0, false, fmt.Errorf("Invalid field name: %v", parts[0])
	}

	h[parts[0]] = strings.TrimSpace(parts[1])

	return idx + 2, false, nil
}
