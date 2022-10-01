package main

import (
	"io/fs"
	"os"
	"strings"
	"time"
)

type DummyFileInfo string

func (d DummyFileInfo) Name() string       { return string(d) }
func (d DummyFileInfo) Size() int64        { return 0 }
func (d DummyFileInfo) Mode() fs.FileMode  { return os.ModeTemporary }
func (d DummyFileInfo) ModTime() time.Time { return time.UnixMilli(0) }
func (d DummyFileInfo) IsDir() bool        { return strings.HasSuffix(string(d), "/") }
func (d DummyFileInfo) Sys() any           { return nil }
func (d DummyFileInfo) Owner() string      { return "nobody" }
func (d DummyFileInfo) Group() string      { return "nobody" }
