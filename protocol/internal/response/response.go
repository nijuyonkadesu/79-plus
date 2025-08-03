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

type Writer struct {
	writer io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{
		writer: w,
	}
}

type Handler func(w *Writer, req *request.Request)

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

func GetDefaultHeaders(contentLength int) *headers.Headers {
	h := headers.NewHeaders()
	h.Set("content-length", fmt.Sprint(contentLength))
	h.Set("connection", "close")
	h.Set("content-type", "text/plain")

	return h
}

func (w *Writer) WriteStatusLine(status StatusCode) error {
	b := []byte{}
	reason := getStatusReason(status)
	b = fmt.Appendf(b, "HTTP/1.1 %d %s\r\n", status, reason)

	_, err := w.writer.Write(b)
	return err
}

func (w *Writer) WriteHeaders(header *headers.Headers) error {
	b := []byte{}
	header.ForEach(func(n, v string) {
		b = fmt.Appendf(b, "%s: %s\r\n", n, v)
	})

	b = fmt.Appendf(b, "\r\n")
	_, err := w.writer.Write(b)
	return err
}

func (w *Writer) WriteBody(p []byte) (int, error) {
	return w.writer.Write(p)
}

func (w *Writer) WriteChunkedBody(p []byte) (int, error) {
	_, err := w.writer.Write(fmt.Appendf(nil, "%x\r\n", len(p)))
	if err != nil {
		return 0, err
	}
	n, err := w.writer.Write(fmt.Appendf(p, "\r\n"))
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (w *Writer) WriteChunkedBodyDone() (int, error) {
	_, err := w.writer.Write(fmt.Appendf(nil, "%x\r\n\r\n", 0))
	return 0, err
}
