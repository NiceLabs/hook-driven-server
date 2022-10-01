package main

import (
	"net/mail"
	"strings"
)

type Address mail.Address

func ParseAddress(address string) (*Address, error) {
	parsed, err := mail.ParseAddress(address)
	if err != nil {
		return nil, err
	}
	return &Address{Name: parsed.Name, Address: parsed.Address}, nil
}

func (a *Address) User() string {
	index := strings.LastIndexByte(a.Address, '@')
	return a.Address[:index]
}

func (a *Address) Domain() string {
	index := strings.LastIndexByte(a.Address, '@')
	return a.Address[index+1:]
}

func (a *Address) String() string {
	return (&mail.Address{Name: a.Name, Address: a.Address}).String()
}
