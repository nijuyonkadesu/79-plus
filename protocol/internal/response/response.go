package response

import (
	"fmt"
	"io"

	"me.httpfrom.tcp/internal/headers"
	"me.httpfrom.tcp/internal/request"
)

var INVALID_RESPONSE_CODE = fmt.Errorf("incorrect response code")

type StatusCode int

type Response struct {
	StatusLine string
	Headers    headers.Headers
	Body       []byte
}

var (
	OK         StatusCode = 200
	Failure    StatusCode = 500
	BadRequest StatusCode = 400
)

type HandlerError struct {
	Code    StatusCode
	Error   error
	Message string
}

type Handler func(w io.Writer, req *request.Request) *HandlerError

func (h *HandlerError) Write(w io.Writer) {
	WriteStatusLine(w, h.Code)
	WriteHeaders(w, GetDefaultHeaders(len(h.Message)))
	w.Write([]byte(h.Message))
}

func getStatusReason(status StatusCode) string {
	switch status {
	case OK:
		return "OK"
	case BadRequest:
		return "Bad Request"
	case Failure:
		return "Internal Server Error"
	default:
		return "" // it's still valid as per the HTTP protocol
	}
}

func WriteStatusLine(w io.Writer, status StatusCode) error {
	b := []byte{}
	reason := getStatusReason(status)
	b = fmt.Appendf(b, "HTTP/1.1 %d %s\r\n", status, reason)

	_, err := w.Write(b)
	return err
}

func GetDefaultHeaders(contentLength int) *headers.Headers {
	h := headers.NewHeaders()
	h.Set("content-length", fmt.Sprint(contentLength))
	h.Set("connection", "close")
	h.Set("content-type", "text/plain")

	return h
}

func WriteHeaders(w io.Writer, header *headers.Headers) error {
	b := []byte{}
	header.ForEach(func(n, v string) {
		b = fmt.Appendf(b, "%s: %s\r\n", n, v)
	})

	b = fmt.Appendf(b, "\r\n")
	_, err := w.Write(b)
	return err
}
