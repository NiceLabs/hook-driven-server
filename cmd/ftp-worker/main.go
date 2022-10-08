package main

import (
	"flag"
	"log"

	. "goftp.io/server"
)

var (
	hostname     string
	port         int
	factory      = new(HookDriverFactory)
	auth         = new(SimpleAuth)
	certFile     string
	keyFile      string
	explicitFTPS bool
)

func init() {
	flag.StringVar(&hostname, "hostname", "", "FTP listen hostname")
	flag.IntVar(&port, "port", 0, "FTP listen port")
	flag.StringVar(&certFile, "tls-cert-file", "", "TLS Certificate file")
	flag.StringVar(&keyFile, "tls-key-file", "", "TLS Key file")
	flag.BoolVar(&explicitFTPS, "tls-explicit", false, "Explicit FTPS")
	flag.StringVar(&auth.Name, "username", "user", "Username")
	flag.StringVar(&auth.Password, "password", "pass", "Password")
	flag.StringVar(&factory.Workdir, "workdir", "", "Work directory")
	flag.StringVar(&factory.ReadHook, "on-read", "", "Read operation local hook program")
	flag.StringVar(&factory.WriteHook, "on-write", "", "Write operation local hook program")
	flag.Parse()
}

func main() {
	serve := NewServer(&ServerOpts{
		Name:         "FTP Hook Service",
		Auth:         auth,
		Hostname:     hostname,
		Port:         port,
		TLS:          certFile != "" && keyFile != "",
		CertFile:     certFile,
		KeyFile:      keyFile,
		ExplicitFTPS: explicitFTPS,
		Factory:      factory,
	})
	if err := serve.ListenAndServe(); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
