package main

import (
	"flag"
	"log"
)

var (
	addr      string
	configure = new(TFTPConfigure)
)

func init() {
	flag.StringVar(&addr, "addr", ":tftp", "TFTP listen address")
	flag.StringVar(&configure.Workdir, "workdir", "", "Work directory")
	flag.StringVar(&configure.ReadHook, "on-read", "", "Read operation local hook program")
	flag.StringVar(&configure.WriteHook, "on-write", "", "Write operation local hook program")
	flag.Parse()
}

func main() {
	server := configure.NewServer()
	log.Println("Starting server at", addr)
	if err := server.ListenAndServe(addr); err != nil {
		log.Fatalln(err)
	}
}
