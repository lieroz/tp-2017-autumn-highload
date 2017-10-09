package main

import (
	"net"
	"strings"
	"strconv"
	"time"
	"bytes"
)

const (
	HttpVersion   = "HTTP/1.1"
	HttpSeparator = "\r\n"
	ServerName    = "tp-autumn-2017-highload"
)

type Response struct {
	Code        int
	Description string
}

func (r *Response) Write(conn net.Conn) {
}

func (r *Response) WriteCommonHeaders(conn net.Conn) {
	conn.Write(r.writeCommonHeaders())
}

func (r *Response) writeCommonHeaders() []byte {
	var commonHeaders = [][]string{
		{
			HttpVersion, strconv.FormatInt(int64(r.Code), 10), r.Description,
		}, {
			"Date:", time.Now().String(),
		}, {
			"Server:", ServerName,
		}, {
			"Connection: Close",
		},
	}
	buffer := bytes.NewBuffer(nil)
	for _, line := range commonHeaders {
		buffer.WriteString(strings.Join(line, " ") + HttpSeparator)
	}
	return buffer.Bytes()
}

func (r *Response) BuildErrResp(err error) {
	switch err {
	case ErrBadRequest:
		r.Code = StatusBadRequest
	case ErrForbidden:
		r.Code = StatusForbidden
	case ErrNotFound:
		r.Code = StatusNotFound
	case ErrMethodNotAllowed:
		r.Code = StatusMethodNotAllowed
	}
	r.Description = err.Error()
}
