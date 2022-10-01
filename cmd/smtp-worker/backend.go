package main

import (
	"errors"
	"net/url"

	"github.com/emersion/go-smtp"
)

type Backend struct {
	HookProgram string
	Userinfo    *url.Userinfo
}

func (b *Backend) Login(state *smtp.ConnectionState, username, password string) (session smtp.Session, err error) {
	if url.UserPassword(username, password).String() != b.Userinfo.String() {
		err = errors.New("smtp-worker: user not found")
		return
	}
	session = &Session{
		HookProgram: b.HookProgram,
		state:       state,
	}
	return
}

func (b *Backend) AnonymousLogin(state *smtp.ConnectionState) (session smtp.Session, err error) {
	err = errors.New("smtp-worker: user not found")
	return
}
