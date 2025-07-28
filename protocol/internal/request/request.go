package request

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"me.httpfrom.tcp/internal/headers"
)

type parserState string

const (
	ParsingRequestLine parserState = "requestline"
	ParsingHeaders     parserState = "headers"
	ParsingBody        parserState = "body"
	ParsingComplete    parserState = "complete"
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
	Headers     headers.Headers
	Body        []byte

	state parserState
}

func newRequest() *Request {
	return &Request{
		state:   ParsingRequestLine,
		Headers: *headers.NewHeaders(),
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
		case ParsingRequestLine:
			rl, n, err := parseRequestLine(data[read:])
			if err != nil {
				return 0, err
			}

			if n == 0 {
				break outer
			}

			read += n
			r.RequestLine = *rl
			r.state = ParsingHeaders

		case ParsingHeaders:
			n, done, err := r.Headers.Parse(data[read:])
			if err != nil {
				return 0, err
			}
			if done {
				r.state = ParsingBody
			}
			if n == 0 {
				break outer
			}
			read += n

		case ParsingBody:
			contentLength := r.Headers.GetInt("content-length", 0)
			body := data[read:]

			remaining := contentLength - len(r.Body)
			if remaining == 0 {
				r.state = ParsingComplete
				break outer
			}

			if contentLength == 0 {
				r.state = ParsingComplete
				break outer
			}

			r.Body = append(r.Body, body...)
			read += len(body)
			break outer

		case ParsingComplete:
			break outer

		default:
			panic("skill issue")
		}
	}
	return read, nil
}

func (r *Request) done() bool {
	return r.state == ParsingComplete
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	req := newRequest()
	buf := make([]byte, 1024)
	bufLen := 0

	// WOAHHHHH,... ReadAll is automatically implemented for us!? - coz we use io.Reader as type!!? damm
	// data, err := io.ReadAll(reader)
outer:
	for !req.done() {
		// TODO: at the very end of exhausting buffer - bufLen could exceed the buffer size, causing bufLen + n to overflow
		if bufLen == len(buf) {
			return nil, fmt.Errorf("buffer full but request incomplete. possible overflow")
		}

		n, err := reader.Read(buf[bufLen:])
		bufLen += n

		readN, parseErr := req.parse(buf[:bufLen])
		if parseErr != nil {
			return nil, parseErr
		}

		// shift the unparsed data to the beginning of the buffer.
		if readN > 0 {
			copy(buf, buf[readN:bufLen])
			bufLen -= readN
		}

		// yea... like prime said, in irl we don't get EOF... after EOF, how can I even send the response back to a "closed" connection... Hilarious.
		if err == io.EOF {
			if !req.done() {
				return nil, fmt.Errorf("connection closed prematurely, or insufficient content")
			}
			break outer
		}

		if err != nil {
			return nil, errors.Join(
				fmt.Errorf("read failure %w", err))
		}
	}
	return req, nil
}
