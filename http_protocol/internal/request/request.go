package request

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
	state       int // 0 = Done, 1 = Initalized
	// state RequestState
}

type RequestLine struct {
	Method        string
	RequestTarget string
	HttpVersion   string
}

// How Prime did it in solution files
//type RequestState int
//
//const (
//	requestStateInitialized RequestState = iota
//	requestStateDone
//)

const crlf = "\r\n"
const bufferSize = 8

func RequestFromReader(reader io.Reader) (*Request, error) {

	buf := make([]byte, bufferSize, bufferSize)
	readToIndex := 0
	request := Request{state: 1}

	for request.state != 0 {
		if readToIndex >= len(buf) {
			fullBuf := make([]byte, len(buf)*2)
			copy(fullBuf, buf)
			buf = fullBuf
		}

		// Read into buffer
		r, err := reader.Read(buf[readToIndex:])
		if err != nil {
			if errors.Is(err, io.EOF) {
				request.state = 0
				break
			}
			return nil, err
		}
		readToIndex += r

		// Parse the data
		n, err := request.parse(buf[:readToIndex])
		if err != nil {
			return nil, err
		}

		copy(buf, buf[n:])
		readToIndex -= n
	}

	return &request, nil
}

func (r *Request) parse(data []byte) (int, error) {

	// Check State for initalized
	switch r.state {
	case 1:
		parsedLine, bytesConsumed, err := parseRequestLine(data)

		// Something actually went wrong
		if err != nil {
			return 0, err
		}

		// We just need more data
		if bytesConsumed == 0 {
			return 0, nil
		}

		// Successful parsing
		r.RequestLine = *parsedLine
		r.state = 0
		return bytesConsumed, nil
	case 0:
		// If the state is 'done', return an error cause we shouldn't be trying to read more here
		return 0, fmt.Errorf("error: trying to read data in a done state")
	default:
		return 0, fmt.Errorf("error: unknown state")
	}

}

func parseRequestLine(rLine []byte) (*RequestLine, int, error) {

	idx := bytes.Index(rLine, []byte(crlf))

	// Couldn't find the CRLF - \r\n
	if idx == -1 {
		return nil, 0, nil
	}

	reqTextLine := string(rLine[:idx])
	requestLine, err := requestFromLineString(reqTextLine)

	if err != nil {
		return nil, idx, err
	}

	return requestLine, idx + 2, nil
}

func requestFromLineString(str string) (*RequestLine, error) {
	// Request line -> method SP request-target (route) SP HTTP-Version

	parts := strings.Split(str, " ")

	if len(parts) != 3 {
		return nil, fmt.Errorf("poorly formatted request line: %s", str)
	}

	method := parts[0]

	for _, c := range method {
		if c < 'A' || c > 'Z' {
			return nil, fmt.Errorf("invalid method: %s", method)
		}
	}

	requestRoute := parts[1]

	versionParts := strings.Split(parts[2], "/")
	if len(versionParts) != 2 {
		return nil, fmt.Errorf("malformed start line: %s", str)
	}

	httpPart := versionParts[0]

	if httpPart != "HTTP" {
		return nil, fmt.Errorf("unrecognized HTTP-version: %s", httpPart)
	}
	version := versionParts[1]
	if version != "1.1" {
		return nil, fmt.Errorf("unrecognized HTTP-version: %s", version)
	}

	return &RequestLine{
		Method:        method,
		RequestTarget: requestRoute,
		HttpVersion:   version,
	}, nil
}
