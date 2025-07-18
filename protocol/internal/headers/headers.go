package headers

import (
	"bytes"
	"fmt"
)

type Header map[string]string

func NewHeaders() Header {
	return map[string]string{}
}

var cr = []byte("\r\n")
var separator = ": "

var INCOMPLETE_VALUE = fmt.Errorf("TODO: is this necessary??")
var MALFORMED_FIELD_NAME = fmt.Errorf("multiple white space found in field name")
var MALFORMED_FIELD_VALUE = fmt.Errorf("header field value is missing")

// stage one: reading -> **wait till get the whole data** (I missed this) -> parse
// stage two: parsing -> get the data values -> update header -> return
// map & string are already a pointer (NEWS)

// Think: why would requestLine parsing needed state and not this?? it's just a design choice or do we have anything specific?

func parseHeader(data []byte) (string, string, int, error) {
	read := 0

	// SplitN because, there can be ':' in field value
	parts := bytes.SplitN(data, []byte(separator), 2)

	if len(parts) != 2 {
		// ahh, the point is , I'm expected to have the entire data, so okay to return err
		return "", "", 0, INCOMPLETE_VALUE
	}

	read += len(parts[0])
	parts[0] = bytes.TrimLeft(parts[0], " ")
	if bytes.HasSuffix(parts[0], []byte(" ")) {
		return "", "", 0, MALFORMED_FIELD_NAME
	}
	key := string(parts[0])

	read += len(parts[1]) + len(separator)
	value := string(bytes.TrimSpace(parts[1]))

	return key, value, read, nil
}

// Think: why Parse is exported in Headers but not in request? coz parsing headers is optional?
func (h Header) Parse(data []byte) (int, bool, error) {

	done := false
	read := 0

	for {
		idx := bytes.Index(data[read:], cr)

		if idx == -1 {
			break
		}

		// the empty line after all the field lines (end of headers) - yea 0 means, first character matchs registered nurse
		if idx == 0 {
			done = true
			break
		}

		// copy() is not needed because, we are not making using for buffers as we did in request line parsing
		key, value, n, err := parseHeader(data[read:read+idx])
		if err != nil {
			return 0, false, err
		}

		h[key] = value
		read += n + len(cr)
	}

	return read, done, nil
}
