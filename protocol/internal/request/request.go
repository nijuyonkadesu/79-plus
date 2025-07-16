package request

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

type parserState string

const (
	StateInit parserState = "init"
	StateDone parserState = "done"
)

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func (r *RequestLine) ValidHTTP() bool {
	return r.HttpVersion == "1.1"
}

type Request struct {
	RequestLine RequestLine
	state       parserState
}

func newRequest() *Request {
	return &Request{
		state: StateInit,
	}
}

var SEPARATOR = []byte("\r\n")
var ERROR_BAD_FORMAT = fmt.Errorf("content does not adhere to http 1.1 standard")
var UNNSUPPORTED_HTTP = fmt.Errorf("version is not HTTP/1.1")

func parseRequestLine(s []byte) (*RequestLine, int, error) {
	idx := bytes.Index(s, SEPARATOR)
	if idx == -1 {
		return nil, 0, nil
	}

	startLine := s[:idx]
	read := len(startLine) + len(SEPARATOR)

	parts := bytes.Split(startLine, []byte(" "))
	if len(parts) < 3 {
		return nil, 0, ERROR_BAD_FORMAT
	}

	method := parts[0]
	path := parts[1]
	httpVersion := bytes.Split(parts[2], []byte("/"))
	if len(httpVersion) < 2 {
		return nil, 0, UNNSUPPORTED_HTTP
	}
	version := httpVersion[1]

	rl := RequestLine{
		HttpVersion:   string(version),
		RequestTarget: string(path),
		Method:        string(method),
	}
	if rl.ValidHTTP() != true {
		return nil, 0, UNNSUPPORTED_HTTP
	}

	return &rl, read, nil
}

func (r *Request) parse(data []byte) (int, error) {

	read := 0
outer:
	for {
		switch r.state {
		case StateInit:
			rl, readTill, err := parseRequestLine(data[read:])
			if err != nil {
				return 0, err
			}

			if readTill == 0 {
				break outer
			}

			read += readTill
			if rl != nil {
				r.RequestLine = *rl
				r.state = StateDone
			}
		case StateDone:
			break outer
		}
	}
	return read, nil
}

func (r *Request) done() bool {
	return r.state == StateDone
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	req := newRequest()
	buf := make([]byte, 1024)
	bufLen := 0

	// WOAHHHHH,... ReadAll is automatically implemented for us!? - coz we use io.Reader as type!!? damm
	// data, err := io.ReadAll(reader)
	for !req.done() {
		// TODO: at the very end of exhausting buffer - bufLen could exceed the buffer size, causing bufLen + n to overflow
		n, err := reader.Read(buf[bufLen:])
		bufLen += n
		// TODO: handle io.EOF - what to do with errors in general ?
		if err != nil {
			return nil, errors.Join(
				fmt.Errorf("read failure %w", err))
		}

		readN, err := req.parse(buf[:bufLen])
		if err != nil {
			return nil, err
		}

		// TODO: but why...? why should we handle for a case where buffer might containe unwanted data? is it because buffer is small?
		copy(buf, buf[readN:bufLen])
		bufLen -= readN
	}
	return req, nil
}
