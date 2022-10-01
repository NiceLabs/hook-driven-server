package main

import (
	"errors"
	"io"
	"os"
	"os/exec"
	"strconv"

	"github.com/NiceLabs/hook-driven-server/cmd/smtp-worker/domains"
	"github.com/NiceLabs/hook-driven-server/utils"
	"github.com/emersion/go-smtp"
)

type Session struct {
	Workdir     string
	HookProgram string
	state       *smtp.ConnectionState
	from, to    *Address
	opts        smtp.MailOptions
}

func (s *Session) Mail(from string, opts smtp.MailOptions) (err error) {
	if s.from, err = ParseAddress(from); err != nil {
		return err
	}
	s.opts = opts
	return
}

func (s *Session) Rcpt(to string) (err error) {
	s.to, err = ParseAddress(to)
	return
}

func (s *Session) Data(r io.Reader) error {
	if s.HookProgram == "" {
		return errors.New("smtp-worker: handler is unset")
	}
	cmd := exec.Command(s.HookProgram, s.to.Address, s.from.Address)
	cmd.Dir = s.Workdir
	cmd.Stdin = r
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	utils.AddEnv(cmd, map[string]string{
		"SMTP_FROM":           s.from.Address,
		"SMTP_FROM_USERNAME":  s.from.User(),
		"SMTP_FROM_DOMAIN":    s.from.Domain(),
		"SMTP_TO":             s.to.Address,
		"SMTP_TO_USERNAME":    s.to.User(),
		"SMTP_TO_DOMAIN":      s.to.Domain(),
		"SMTP_TO_DOMAIN_TYPE": domains.Type(s.to.Domain()),
		"SMTP_HOSTNAME":       s.state.Hostname,
		"SMTP_LOCAL_ADDR":     s.state.LocalAddr.String(),
		"SMTP_REMOTE_ADDR":    s.state.RemoteAddr.String(),
		"SMTP_UTF8":           strconv.FormatBool(s.opts.UTF8),
		"SMTP_REQUIRE_TLS":    strconv.FormatBool(s.opts.RequireTLS),
		"SMTP_BODY_TYPE":      string(s.opts.Body),
		"SMTP_BODY_SIZE":      strconv.Itoa(s.opts.Size),
	})
	if auth := s.opts.Auth; auth != nil {
		utils.AddEnv(cmd, map[string]string{"SMTP_AUTH": *auth})
	}
	return cmd.Run()
}

func (s *Session) Logout() error {
	return nil
}

func (s *Session) Reset() {
}
