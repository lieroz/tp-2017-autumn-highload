package main

import (
	"net"
	tp "net/textproto"
	"strings"
	"log"
	"bufio"
)

type Request struct {
	Method  string
	AbsPath string
}

func (r *Request) Parse(conn net.Conn) error {
	reqLine, err := tp.NewReader(bufio.NewReader(conn)).ReadLine()
	if err != nil {
		log.Fatalln("error reading connection:", err)
	}
	reqParams := strings.Split(reqLine, " ")
	if !checkMethod(reqParams[0]) {
		return ErrMethodNotAllowed
	}
	if !checkUrl(reqParams[1]) {
		return ErrBadRequest
	}
	r.Method = reqParams[0]
	if strings.Contains(reqParams[1], "?") {
		r.AbsPath = reqParams[1][:strings.Index(reqParams[1], "?")]
	} else {
		r.AbsPath = reqParams[1]
	}
	return nil
}

func checkMethod(reqMethod string) bool {
	for _, method := range AllowedMethods {
		if reqMethod == method {
			return true
		}
	}
	return false
}

func checkUrl(reqUrl string) bool {
	if strings.Contains(reqUrl, "../") {
		return false
	}
	return true
}
