package main

import (
	"io"

	"goftp.io/server/core"
)

type HookDriverFactory struct {
	Workdir   string
	ReadHook  string
	WriteHook string
}

func (h *HookDriverFactory) NewDriver() (driver core.Driver, err error) {
	driver = &HookDriver{
		Workdir:   h.Workdir,
		ReadHook:  h.ReadHook,
		WriteHook: h.WriteHook,
		writeMaps: make(map[string]io.ReaderFrom),
	}
	return
}
