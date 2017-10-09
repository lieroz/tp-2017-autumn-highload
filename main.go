package main

import (
	"os"
	"flag"
	"log"
	"bufio"
	"strings"
	"strconv"
	"runtime"
)

type Config struct {
	Port    string
	NumCpu  int64
	WebRoot string
}

func parseConfig(filePath string) *Config {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalln("error opening config file:", err)
	}
	defer f.Close()
	var conf Config
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		params := strings.Split(scanner.Text(), WordSeparator)
		if params[0] == "listen" {
			conf.Port = params[1]
		} else if params[0] == "cpu_limit" {
			conf.NumCpu, err = strconv.ParseInt(params[1], Base, 0)
			if err != nil {
				log.Fatalln("error parsing config file:", err)
			}
		} else if params[0] == "document_root" {
			conf.WebRoot = params[1]
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln("error parsing config file:", err)
	}
	return &conf
}

func main() {
	confFile := flag.String("c", "httpd.conf", "a string")
	flag.Parse()
	pwd, _ := os.Getwd()
	conf := parseConfig(pwd + "/" + *confFile)
	runtime.GOMAXPROCS(int(conf.NumCpu))
	serv := NewServer(conf.Port, conf.WebRoot)
	serv.ListenAndServe()
}
