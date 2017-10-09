package main

import (
	"net"
	"strings"
	"strconv"
	"time"
	"bytes"
	"os"
	"path/filepath"
	"io"
)

const (
	Base = 10

	HttpVersion   = "HTTP/1.1"
	ServerName    = "tp-autumn-2017-highload"
	HttpSeparator = "\r\n"
	WordSeparator = " "
)

type Response struct {
	Code        int
	Description string
}

func (r *Response) WriteBody(conn net.Conn, f *os.File) {
	buf := r.writeFileInfo(f)
	io.Copy(buf, f)
	conn.Write(buf.Bytes())
}

func (r *Response) Write(conn net.Conn, f *os.File) {
	conn.Write(r.writeFileInfo(f).Bytes())
}

func (r *Response) writeFileInfo(f *os.File) *bytes.Buffer {
	fileInfo, _ := f.Stat()
	var contentHeaders = [][]string{
		{
			"Content-Length:", strconv.FormatInt(fileInfo.Size(), Base),
		}, {
			"Content-Type:", GetContentType(filepath.Ext(fileInfo.Name())[1:]),
		},
	}
	buf := bytes.NewBuffer(r.writeCommonHeaders().Bytes())
	for _, line := range contentHeaders {
		buf.WriteString(strings.Join(line, WordSeparator) + HttpSeparator)
	}
	buf.WriteString(HttpSeparator)
	return buf
}

func (r *Response) WriteCommonHeaders(conn net.Conn) {
	conn.Write(r.writeCommonHeaders().Bytes())
}

func (r *Response) writeCommonHeaders() *bytes.Buffer {
	var commonHeaders = [][]string{
		{
			HttpVersion, strconv.FormatInt(int64(r.Code), Base), r.Description,
		}, {
			"Date:", time.Now().String(),
		}, {
			"Server:", ServerName,
		}, {
			"Connection: Close",
		},
	}
	buf := bytes.NewBuffer(nil)
	for _, line := range commonHeaders {
		buf.WriteString(strings.Join(line, WordSeparator) + HttpSeparator)
	}
	return buf
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
