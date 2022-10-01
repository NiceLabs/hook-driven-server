package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/url"

	"github.com/emersion/go-smtp"
)

var (
	addr     string
	hook     string
	certFile string
	keyFile  string
	username string
	password string
)

func init() {
	flag.StringVar(&addr, "addr", "", "SMTP listen address")
	flag.StringVar(&certFile, "tls-cert-file", "", "TLS Certificate file")
	flag.StringVar(&keyFile, "tls-key-file", "", "TLS Key file")
	flag.StringVar(&hook, "on-request", "", "SMTP request handler")
	flag.StringVar(&username, "username", "user", "Username")
	flag.StringVar(&password, "password", "pass", "Password")
	flag.Parse()
}

func main() {
	var err error
	server := smtp.NewServer(&Backend{
		HookProgram: hook,
		Userinfo:    url.UserPassword(username, password),
	})
	server.Addr = addr
	server.EnableREQUIRETLS = true
	server.EnableBINARYMIME = true
	server.EnableSMTPUTF8 = true
	log.Println("Starting server at", server.Addr)
	if certFile != "" && keyFile != "" {
		if server.TLSConfig, err = simpleTLSConfig(certFile, keyFile); err != nil {
			log.Fatalln(err)
		}
		err = server.ListenAndServeTLS()
	} else {
		err = server.ListenAndServe()
	}
	if err != nil {
		log.Fatalln(err)
	}
}

func simpleTLSConfig(certFile, keyFile string) (config *tls.Config, err error) {
	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err == nil {
		config = &tls.Config{
			NextProtos:   []string{"smtp"},
			Certificates: []tls.Certificate{certificate},
		}
	}
	return
}
